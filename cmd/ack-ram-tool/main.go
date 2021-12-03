package main

import (
	"log"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ack-ram-tool",
		Short: "A command line utility for using RAM in ACK.",
		Long: `A command line utility for using RAM in ACK.

More info: https://github.com/AliyunContainerService/ack-ram-tool`,
	}
)

func init() {
	rrsa.SetupRRSACmd(rootCmd)
	version.SetupVersionCmd(rootCmd)
	rootCmd.PersistentFlags().StringVar(&ctl.GlobalOption.Region, "region-id", "cn-hangzhou", "The region to use")
	rootCmd.PersistentFlags().BoolVarP(&ctl.GlobalOption.AssumeYes, "assume-yes", "y", false,
		"Automatic yes to prompts; assume \"yes\" as answer to all prompts and run non-interactively")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
