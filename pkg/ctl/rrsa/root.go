package rrsa

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "rrsa",
	Short: "Utils for using RAM Roles for Service Accounts(RRSA).",
	Long:  `Utils for using RAM Roles for Service Accounts(RRSA).`,
}

func init() {
	setupStatusCmd(rootCmd)
	setupEnableCmd(rootCmd)
	setupDisableCmd(rootCmd)
	setupAssociateRoleCmd(rootCmd)
	setupAssumeRoleCmd(rootCmd)
}

func SetupRRSACmd(root *cobra.Command) {
	root.AddCommand(rootCmd)
}
