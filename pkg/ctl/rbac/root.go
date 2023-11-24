package rbac

import (
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac/cleanupuserpermissions"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac/scanuserpermissions"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rbac",
	Short: "Utils for RBAC permissions.",
	Long:  `Utils for RBAC permissions.`,
}

func init() {
	scanuserpermissions.SetupCmd(rootCmd)
	cleanupuserpermissions.SetupCmd(rootCmd)
}

func SetupCmd(root *cobra.Command) {
	root.AddCommand(rootCmd)
}
