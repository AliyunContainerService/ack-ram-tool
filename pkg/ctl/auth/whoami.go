package auth

import (
	"context"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Check who you are",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctl.GlobalOption.Verbose = true
		client := common.GetClientOrDie()

		acc, err := client.GetCallerIdentity(context.TODO())
		common.ExitIfError(err)
		table := tablewriter.NewWriter(os.Stdout)
		//table.SetAutoMergeCells(true)
		table.SetAutoWrapText(true)
		table.SetAutoFormatHeaders(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.EnableBorder(false)
		table.SetTablePadding("  ")
		table.SetNoWhiteSpace(true)

		table.Append([]string{"AccountId", acc.RootUId})
		table.Append([]string{"IdentityType", acc.IdentityType()})
		if acc.Role.RoleId != "" {
			table.Append([]string{"RoleId", acc.Role.RoleId})
			table.Append([]string{"RoleName", acc.Role.RoleName})
		} else if acc.User.Id != "" {
			table.Append([]string{"UserId", acc.User.Id})
			table.Append([]string{"UserName", acc.User.Name})
		}
		table.Append([]string{"Arn", acc.Arn})
		table.Append([]string{"PrincipalId", acc.PrincipalId})

		table.Render()
	},
}

func setupWhoamiCmdCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(whoamiCmd)

	whoamiCmd.Flags().StringVar(
		&ctl.GlobalOption.FinalAssumeRoleAnotherRoleArn, "role-arn", "",
		"Assume an RAM Role ARN when send request or sign token")
}
