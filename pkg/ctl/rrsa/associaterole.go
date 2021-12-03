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
	roleName       = ""
	namespace      = ""
	serviceAccount = ""
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
		if err := associateRole(context.Background(), c, client); err != nil {
			exitByError(fmt.Sprintf("Associate RAM Role %s to service account %s (namespace: %s) failed: %+v",
				roleName, serviceAccount, namespace, err))
			return
		}
		fmt.Printf("Associate RAM Role %s to service account %s (namespace: %s) successfully\n",
			roleName, serviceAccount, namespace)
	},
}

func associateRole(ctx context.Context, c *types.Cluster, client *openapi.Client) error {
	rrsac := c.MetaData.RRSAConfig
	role, err := client.GetRole(ctx, roleName)
	if err != nil {
		return err
	}
	assumeRolePolicyDocument := role.AssumeRolePolicyDocument
	oldDocument := assumeRolePolicyDocument.JSON()
	policy := types.MakeAssumeRolePolicyStatementWithServiceAccount(
		rrsac.Issuer, rrsac.OIDCArn, namespace, serviceAccount)
	if exist, err := assumeRolePolicyDocument.IncludePolicy(policy); err != nil {
		return err
	} else if exist {
		fmt.Printf("Already associated RAM Role %s to service account %s (namespace: %s). Skip to continue\n",
			roleName, serviceAccount, namespace)
		return nil
	}
	if err := assumeRolePolicyDocument.AppendPolicy(policy); err != nil {
		return err
	}
	newDocument := assumeRolePolicyDocument.JSON()
	diff := utils.DiffPrettyText(oldDocument, newDocument)
	fmt.Printf("Will change the assumeRolePolicyDocument of RAM Role %s with blow content:\n%s\n",
		roleName, diff)
	yesOrExit(fmt.Sprintf(
		"Are you sure you want to associate RAM Role %s to service account %s (namespace: %s)?",
		roleName, serviceAccount, namespace))

	_, err = client.UpdateRole(ctx, roleName, openapi.UpdateRamRoleOption{
		AssumeRolePolicyDocument: assumeRolePolicyDocument,
	})
	return err
}

func setupAssociateRoleCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(associateRoleCmd)
	associateRoleCmd.Flags().StringVarP(&clusterId, "cluster-id", "c", "", "The cluster id to use")
	associateRoleCmd.Flags().StringVarP(&roleName, "role-name", "r", "", "The RAM role name to use")
	associateRoleCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "The Kubernetes namespace to use")
	associateRoleCmd.Flags().StringVarP(&serviceAccount, "service-account", "s", "The Kubernetes service account to use", "")
	err := associateRoleCmd.MarkFlagRequired("cluster-id")
	exitIfError(err)
	err = associateRoleCmd.MarkFlagRequired("role-name")
	exitIfError(err)
	err = associateRoleCmd.MarkFlagRequired("namespace")
	exitIfError(err)
	err = associateRoleCmd.MarkFlagRequired("service-account")
	exitIfError(err)
}
