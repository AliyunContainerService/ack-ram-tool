package rrsa

import (
	"context"
	"fmt"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	roleName             = ""
	namespace            = ""
	serviceAccount       = ""
	createRoleIfNotExist bool
)

var associateRoleCmd = &cobra.Command{
	Use:   "associate-role",
	Short: "Associate an RAM role to a Kubernetes Service Account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		yesOrExit(fmt.Sprintf(
			"Are you sure you want to associate RAM Role %s to service account %s (namespace: %s)?",
			roleName, serviceAccount, namespace))

		ctx := context.Background()
		c, err := getRRSAStatus(ctx, clusterId, client)
		if err != nil {
			exitByError(fmt.Sprintf("get status failed: %+v", err))
		}
		if !c.MetaData.RRSAConfig.Enabled {
			exitByError("RRSA feature is not enabled!")
		}
		if err := associateRole(context.Background(), c, client,
			roleName, namespace, serviceAccount, createRoleIfNotExist); err != nil {
			exitByError(fmt.Sprintf("Associate RAM Role %s to service account %s (namespace: %s) failed: %+v",
				roleName, serviceAccount, namespace, err))
			return
		}
		fmt.Printf("Associate RAM Role %s to service account %s (namespace: %s) successfully\n",
			roleName, serviceAccount, namespace)
	},
}

func associateRole(ctx context.Context, c *types.Cluster, client *openapi.Client,
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

	fmt.Printf("Will create RAM Role %q with blow AssumeRole Policy:\n%s\n",
		roleName, assumeRolePolicyDocument.JSON())
	yesOrExit(fmt.Sprintf("Are you sure you want to create RAM Role %s?", roleName))

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
		fmt.Printf("Already associated RAM Role %q to service account %q (namespace: %q). Skip to continue\n",
			roleName, serviceAccount, namespace)
		return nil
	}

	if err := assumeRolePolicyDocument.AppendPolicy(policy); err != nil {
		return err
	}
	newDocument := assumeRolePolicyDocument.JSON()
	diff := utils.DiffPrettyText(oldDocument, newDocument)

	fmt.Printf("Will change the AssumeRole Policy of RAM Role %q with blow content:\n%s\n",
		roleName, diff)
	yesOrExit(fmt.Sprintf(
		"Are you sure you want to associate RAM Role %q to service account %q (namespace: %q)?",
		roleName, serviceAccount, namespace))

	_, err := client.UpdateRole(ctx, roleName, openapi.UpdateRamRoleOption{
		AssumeRolePolicyDocument: assumeRolePolicyDocument,
	})
	return err
}

func setupAssociateRoleCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(associateRoleCmd)
	associateRoleCmd.Flags().StringVarP(&clusterId, "cluster-id", "c", "", "The cluster id to use")
	associateRoleCmd.Flags().StringVarP(&roleName, "role-name", "r", "", "The RAM role name to use")
	associateRoleCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "The Kubernetes namespace to use")
	associateRoleCmd.Flags().StringVarP(&serviceAccount, "service-account", "s", "", "The Kubernetes service account to use")
	associateRoleCmd.Flags().BoolVar(&createRoleIfNotExist, "create-role-if-not-exist", false, "Create the RAM role if it does not exist")
	err := associateRoleCmd.MarkFlagRequired("cluster-id")
	exitIfError(err)
	err = associateRoleCmd.MarkFlagRequired("role-name")
	exitIfError(err)
	err = associateRoleCmd.MarkFlagRequired("namespace")
	exitIfError(err)
	err = associateRoleCmd.MarkFlagRequired("service-account")
	exitIfError(err)
}
