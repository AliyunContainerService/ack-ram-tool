package associaterole

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"log"

	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	"github.com/spf13/cobra"
)

type Option struct {
	roleName             string
	namespace            string
	serviceAccount       string
	createRoleIfNotExist bool
	attachSystemPolicy   string
	attachCustomPolicy   string
}

var opts = Option{}

var cmd = &cobra.Command{
	Use:   "associate-role",
	Short: "Associate an RAM role to a Kubernetes Service Account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()
		clusterId := ctl.GlobalOption.ClusterId
		roleName := opts.roleName
		serviceAccount := opts.serviceAccount
		namespace := opts.namespace
		createRoleIfNotExist := opts.createRoleIfNotExist

		ctlcommon.YesOrExit(fmt.Sprintf(
			"Are you sure you want to associate RAM Role %q to service account %q (namespace: %q)?",
			roleName, serviceAccount, namespace))

		ctx := context.Background()
		c, err := common.GetRRSAStatus(ctx, clusterId, client)
		if err != nil {
			ctlcommon.ExitByError(fmt.Sprintf("get status failed: %+v", err))
		}
		if !c.MetaData.RRSAConfig.Enabled {
			ctlcommon.ExitByError("RRSA feature is not enabled!")
		}
		if err := AssociateRole(context.Background(), c, client,
			roleName, namespace, serviceAccount, createRoleIfNotExist); err != nil {
			ctlcommon.ExitByError(fmt.Sprintf("Associate RAM Role %q to service account %q (namespace: %q) failed: %+v",
				roleName, serviceAccount, namespace, err))
			return
		}
		log.Printf("Associate RAM Role %q to service account %q (namespace: %q) successfully",
			roleName, serviceAccount, namespace)

		attachPolices(ctx, client, roleName)
	},
}

func AssociateRole(ctx context.Context, c *types.Cluster, client *openapi.Client,
	roleName, namespace, serviceAccount string, createRoleIfNotExist bool) error {
	rrsac := c.MetaData.RRSAConfig
	role, err := client.GetRole(ctx, roleName)
	if err != nil {
		if openapi.IsRamRoleNotExistErr(err) && createRoleIfNotExist {
			return createRole(ctx, client, rrsac, roleName, namespace, serviceAccount)
		}
		return err
	}

	return updateRole(ctx, client, role, rrsac, namespace, serviceAccount)
}

func createRole(ctx context.Context, client *openapi.Client, rrsac types.RRSAConfig,
	roleName, namespace, serviceAccount string) error {
	rd := types.MakeRamPolicyDocument([]types.RamPolicyStatement{})
	assumeRolePolicyDocument := &rd
	role := types.RamRole{
		RoleName:                 roleName,
		Description:              "",
		AssumeRolePolicyDocument: assumeRolePolicyDocument,
	}
	policy := types.MakeAssumeRolePolicyStatementWithServiceAccount(
		rrsac.TokenIssuer(), rrsac.OIDCArn, namespace, serviceAccount)
	if err := assumeRolePolicyDocument.AppendPolicy(policy); err != nil {
		return err
	}

	log.Printf("will create RAM Role %q with blow AssumeRole Policy:\n%s\n",
		roleName, assumeRolePolicyDocument.JSON())
	ctlcommon.YesOrExit(fmt.Sprintf("Are you sure you want to create RAM Role %q?", roleName))

	_, err := client.CreateRole(ctx, role)
	return err
}

func updateRole(ctx context.Context, client *openapi.Client, role *types.RamRole,
	rrsac types.RRSAConfig, namespace, serviceAccount string) error {
	roleName := role.RoleName
	assumeRolePolicyDocument := role.AssumeRolePolicyDocument
	oldDocument := assumeRolePolicyDocument.JSON()
	policy := types.MakeAssumeRolePolicyStatementWithServiceAccount(
		rrsac.TokenIssuer(), rrsac.OIDCArn, namespace, serviceAccount)

	if exist, err := assumeRolePolicyDocument.IncludePolicy(policy); err != nil {
		return err
	} else if exist {
		log.Printf("Already associated RAM Role %q to service account %q (namespace: %q)",
			roleName, serviceAccount, namespace)
		return nil
	}

	if err := assumeRolePolicyDocument.AppendPolicy(policy); err != nil {
		return err
	}
	newDocument := assumeRolePolicyDocument.JSON()
	diff := utils.DiffPrettyText(oldDocument, newDocument)

	log.Printf("will change the AssumeRole Policy of RAM Role %q with blow content:\n%s\n",
		roleName, diff)
	ctlcommon.YesOrExit(fmt.Sprintf(
		"Are you sure you want to associate RAM Role %q to service account %q (namespace: %q)?",
		roleName, serviceAccount, namespace))

	_, err := client.UpdateRole(ctx, roleName, openapi.UpdateRamRoleOption{
		AssumeRolePolicyDocument: assumeRolePolicyDocument,
	})
	return err
}

func attachPolices(ctx context.Context, client *openapi.Client, roleName string) {
	if opts.attachSystemPolicy == "" && opts.attachCustomPolicy == "" {
		return
	}

	log.Println("Start to attach policies")
	if opts.attachSystemPolicy != "" {
		policyName := opts.attachSystemPolicy
		if err := attachPolicy(ctx, client, roleName, policyName, types.RamPolicyTypeSystem); err != nil {
			ctlcommon.ExitByError(fmt.Sprintf("Attach System policy %s failed: %+v", policyName, err))
			return
		}
	}
	if opts.attachCustomPolicy != "" {
		policyName := opts.attachCustomPolicy
		if err := attachPolicy(ctx, client, roleName, policyName, types.RamPolicyTypeCustom); err != nil {
			ctlcommon.ExitByError(fmt.Sprintf("Attach Custom policy %s failed: %+v", policyName, err))
			return
		}
	}
	log.Println("Attach policies successfully")
}

func attachPolicy(ctx context.Context, client *openapi.Client, roleName, policyName, policyType string) error {
	log.Printf("Start to attach the %s policy %s to the Role %s", policyType, policyName, roleName)
	ctlcommon.YesOrExit(fmt.Sprintf(
		"Are you sure you want to attach the %s policy %s to the Role %s?", policyType, policyName, roleName))

	err := client.AttachPolicyToRole(ctx, policyName, policyType, roleName)
	return err
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	ctlcommon.SetupClusterIdFlag(cmd)

	cmd.Flags().StringVarP(&opts.roleName, "role-name", "r", "", "The RAM Role name to use")
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "The Kubernetes namespace to use")
	cmd.Flags().StringVarP(&opts.serviceAccount, "service-account", "s", "", "The Kubernetes service account to use")
	cmd.Flags().BoolVar(&opts.createRoleIfNotExist, "create-role-if-not-exist", false, "Create the RAM Role if it does not exist")
	cmd.Flags().StringVar(&opts.attachSystemPolicy, "attach-system-policy", "", "Attach this system policy to the RAM Role")
	cmd.Flags().StringVar(&opts.attachCustomPolicy, "attach-custom-policy", "", "Attach this custom policy to the RAM Role")

	ctlcommon.ExitIfError(cmd.MarkFlagRequired("role-name"))
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("namespace"))
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("service-account"))
}
