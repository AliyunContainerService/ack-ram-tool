package auth

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auth",
	Short: "Utils for auth.",
	Long:  `Utils for auth.`,
}

func init() {
	setupWhoamiCmdCmd(rootCmd)
}

func SetupCmd(root *cobra.Command) {
	root.AddCommand(rootCmd)
}
