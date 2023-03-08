package main

import (
	"log"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/credentialplugin"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/exportcredentials"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ack-ram-tool",
		Short: "A command line utility for using RAM in Alibaba Cloud Container Service For Kubernetes (ACK).",
		Long: `A command line utility for using RAM in Alibaba Cloud Container Service For Kubernetes (ACK).

More info: https://github.com/AliyunContainerService/ack-ram-tool`,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
		},
	}
)

func init() {
	rrsa.SetupRRSACmd(rootCmd)
	credentialplugin.SetupCredentialPluginCmd(rootCmd)
	version.SetupVersionCmd(rootCmd)
	rbac.SetupCmd(rootCmd)
	exportcredentials.SetupCmd(rootCmd)

	rootCmd.PersistentFlags().StringVar(&ctl.GlobalOption.Region, "region-id", "cn-hangzhou", "The region to use")
	rootCmd.PersistentFlags().BoolVarP(&ctl.GlobalOption.AssumeYes, "assume-yes", "y", false,
		"Automatic yes to prompts; assume \"yes\" as answer to all prompts and run non-interactively")
	rootCmd.PersistentFlags().StringVar(&ctl.GlobalOption.CredentialFilePath, "profile-file", "",
		"Path to credential file (default: ~/.alibabacloud/credentials or ~/.aliyun/config.json)")
	rootCmd.PersistentFlags().StringVar(&ctl.GlobalOption.ProfileName, "profile-name", "",
		"using this named profile when parse credentials from config.json of aliyun cli")
	rootCmd.PersistentFlags().BoolVar(&ctl.GlobalOption.IgnoreEnv,
		"ignore-env-credentials", false, "don't try to parse credentials from environment variables")
	rootCmd.PersistentFlags().BoolVar(&ctl.GlobalOption.IgnoreAliyuncliConfig,
		"ignore-aliyun-cli-credentials", false, "don't try to parse credentials from config.json of aliyun cli")
	//rootCmd.PersistentFlags().BoolVarP(&ctl.GlobalOption.InsecureSkipTLSVerify, "insecure-skip-tls-verify", "", false, "Skips the validity check for the server's certificate")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
