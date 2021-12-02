package version

import (
	"fmt"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version information",
	Long:  `Show the version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:     %s\n", version.Version)
		fmt.Printf("GitCommit:   %s\n", version.GitCommit)
	},
}

func SetupVersionCmd(root *cobra.Command) {
	root.AddCommand(rootCmd)
}
