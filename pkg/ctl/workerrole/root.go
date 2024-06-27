package workerrole

import (
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/workerrole/scanaddon"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "worker-role",
	Short: "Utils for worker role.",
	Long:  `Utils for worker role.`,
}

func init() {
	scanaddon.SetupCmd(rootCmd)
}

func SetupCmd(root *cobra.Command) {
	root.AddCommand(rootCmd)
}
