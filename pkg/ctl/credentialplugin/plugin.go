package credentialplugin

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "credential-plugin",
	Short: "A kubectl/client-go credential plugin for authentication with an ACK cluster.",
	Long:  `A kubectl/client-go credential plugin for authentication with an ACK cluster.`,
}

func SetupCredentialPluginCmd(root *cobra.Command) {
	root.AddCommand(rootCmd)
}

func init() {
	setupGetKubeconfigCmd(rootCmd)
	setupGetCredentialCmdCmd(rootCmd)
	setupGetTokenCmdCmd(rootCmd)
}
