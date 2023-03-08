package setupaddon

import (
	"context"
	"fmt"
	"strings"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/associaterole"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/setupaddon/addon"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
)

type Option struct {
	addonName string
}

var opts = Option{}

var cmd = &cobra.Command{
	Use:   "setup-addon",
	Short: "Setup RAM actions for cluster addon.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()
		addon := addon.GetAddon(opts.addonName)

		ctlcommon.YesOrExit(fmt.Sprintf(
			"Are you sure you want to setup RAM actions for this addon: %s?",
			opts.addonName))

		ctx := context.Background()
		clusterId := ctl.GlobalOption.ClusterId
		c, err := common.GetRRSAStatus(ctx, clusterId, client)
		if err != nil {
			ctlcommon.ExitByError(fmt.Sprintf("get status failed: %+v", err))
		}
		rrsaConfig := c.MetaData.RRSAConfig
		if !rrsaConfig.Enabled {
			ctlcommon.ExitByError("RRSA feature is not enabled!")
		}

		policy := addon.RamPolicy()
		log.Logger.Infof("start to ensure RAM policy(%q) is exist and be configured", policy.PolicyName)
		if err := ensureRamPolicy(ctx, addon, c, client); err != nil {
			ctlcommon.ExitIfError(err)
		}

		roleName := addon.RoleName(c.ClusterId)
		log.Logger.Infof("start to ensure RAM role(%q) is exist and be configured", roleName)
		if err := associaterole.AssociateRole(ctx, c, client, roleName,
			addon.NameSpace(), addon.ServiceAccountName(), true); err != nil {
			ctlcommon.ExitIfError(err)
		}

		ap := addon.RamPolicy()
		log.Logger.Infof("start to attach RAM polciy(%q) to role(%q)", ap.PolicyName, roleName)
		if err := ensureAttachPolicyToRamRole(ctx, addon, c, client, rrsaConfig, roleName); err != nil {
			ctlcommon.ExitIfError(err)
		}
		log.Logger.Infof("attach RAM polciy(%q) to role(%q) is successful", ap.PolicyName, roleName)
	},
}

func ensureRamPolicy(ctx context.Context, addon addon.Meta, c *types.Cluster, client *openapi.Client) error {
	ap := addon.RamPolicy()
	_, err := client.GetPolicy(ctx, ap.PolicyName, ap.PolicyType)
	if err != nil {
		if !openapi.IsRamPolicyNotExistErr(err) {
			return err
		}
		_, err = client.CreatePolicy(ctx, ap)
		if err != nil {
			return err
		}
	}
	return nil
}

func ensureAttachPolicyToRamRole(ctx context.Context, addon addon.Meta, c *types.Cluster,
	client *openapi.Client, rrsac types.RRSAConfig, roleName string) error {
	ap := addon.RamPolicy()

	policies, err := client.ListPoliciesForRole(ctx, roleName)
	if err != nil {
		return err
	}
	var attached bool
	for _, p := range policies {
		if p.PolicyName == ap.PolicyName {
			attached = true
		}
	}
	if attached {
		log.Logger.Infof("the policy(%q) has already been attached to the role(%q)",
			ap.PolicyName, roleName)
		return nil
	}

	return client.AttachPolicyToRole(ctx, ap.PolicyName, ap.PolicyType, roleName)
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	ctlcommon.SetupClusterIdFlag(cmd)

	cmd.Flags().StringVarP(&opts.addonName, "addon-name", "a", "The name of cluster addon",
		fmt.Sprintf("addon name: %s", strings.Join(addon.ListAddonNames(), ",")))
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("addon-name"))
}
