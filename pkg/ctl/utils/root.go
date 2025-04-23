package utils

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "utils",
}

func init() {
	setupTestEdCmd(rootCmd)
}

func SetupUtilsCmd(root *cobra.Command) {
	root.AddCommand(rootCmd)
}
