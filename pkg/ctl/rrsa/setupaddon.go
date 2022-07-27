package rrsa

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/addon"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type SetupAddonOpts struct {
	addonName string
	clusterId string
}

var setupAddonOpts = SetupAddonOpts{}

var setupAddonCmd = &cobra.Command{
	Use:   "setup-addon",
	Short: "setup RAM actions for cluster addon.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		addon := addon.GetAddon(setupAddonOpts.addonName)

		yesOrExit(fmt.Sprintf(
			"Are you sure you want to setup RAM actions for this addon: %s?",
			setupAddonOpts.addonName))

		ctx := context.Background()
		c, err := getRRSAStatus(ctx, setupAddonOpts.clusterId, client)
		if err != nil {
			exitByError(fmt.Sprintf("get status failed: %+v", err))
		}
		rrsaConfig := c.MetaData.RRSAConfig
		if !rrsaConfig.Enabled {
			exitByError("RRSA feature is not enabled!")
		}

		policy := addon.RamPolicy()
		log.Printf("start to ensure RAM policy(%q) is exist and be configured", policy.PolicyName)
		if err := ensureRamPolicy(ctx, addon, c, client); err != nil {
			exitIfError(err)
		}

		roleName := addon.RoleName(c.ClusterId)
		log.Printf("start to ensure RAM role(%q) is exist and be configured", roleName)
		if err := associateRole(ctx, c, client, roleName,
			addon.NameSpace(), addon.ServiceAccountName(), true); err != nil {
			exitIfError(err)
		}

		ap := addon.RamPolicy()
		log.Printf("start to attach RAM polciy(%q) to role(%q)", ap.PolicyName, roleName)
		if err := ensureAttachPolicyToRamRole(ctx, addon, c, client, rrsaConfig, roleName); err != nil {
			exitIfError(err)
		}
		log.Printf("attach RAM polciy(%q) to role(%q) is successful", ap.PolicyName, roleName)
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
		log.Printf("the policy(%q) has already been attached to the role(%q)",
			ap.PolicyName, roleName)
		return nil
	}

	return client.AttachPolicyToRole(ctx, ap.PolicyName, ap.PolicyType, roleName)
}

func sSetupAddonCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(setupAddonCmd)
	setupAddonCmd.Flags().StringVarP(&setupAddonOpts.clusterId, "cluster-id", "c", "", "The cluster id to use")
	err := setupAddonCmd.MarkFlagRequired("cluster-id")
	exitIfError(err)

	setupAddonCmd.Flags().StringVarP(&setupAddonOpts.addonName, "addon-name", "a", "The name of cluster addon",
		fmt.Sprintf("addon name: %s", strings.Join(addon.ListAddonNames(), ",")))
	err = setupAddonCmd.MarkFlagRequired("addon-name")
	exitIfError(err)
}
