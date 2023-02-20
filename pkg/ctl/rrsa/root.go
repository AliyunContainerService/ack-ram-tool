package rrsa

import (
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/associaterole"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/assumerole"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/disable"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/enable"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/installhelperaddon"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/setupaddon"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/status"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rrsa",
	Short: "Utils for using RAM Roles for Service Accounts(RRSA).",
	Long:  `Utils for using RAM Roles for Service Accounts(RRSA).`,
}

func init() {
	status.SetupCmd(rootCmd)
	enable.SetupCmd(rootCmd)
	disable.SetupCmd(rootCmd)
	associaterole.SetupCmd(rootCmd)
	assumerole.SetupCmd(rootCmd)
	setupaddon.SetupCmd(rootCmd)
	installhelperaddon.SetupCmd(rootCmd)
	setupDemoCmd(rootCmd)
}

func SetupRRSACmd(root *cobra.Command) {
	root.AddCommand(rootCmd)
}
