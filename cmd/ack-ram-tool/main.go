package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/credentialplugin"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/version"
	"github.com/aliyun/credentials-go/credentials"
	"github.com/spf13/cobra"
)

const (
	envAccessKeyId     = "ALIBABA_CLOUD_ACCESS_KEY_ID"
	envAccessKeySecret = "ALIBABA_CLOUD_ACCESS_KEY_SECRET"
)

var defaultProfilePath = filepath.Join("~", ".alibabacloud", "credentials")
var profilePath = ""

var (
	rootCmd = &cobra.Command{
		Use:   "ack-ram-tool",
		Short: "A command line utility for using RAM in Alibaba Cloud Container Service For Kubernetes (ACK).",
		Long: `A command line utility for using RAM in Alibaba Cloud Container Service For Kubernetes (ACK).

More info: https://github.com/AliyunContainerService/ack-ram-tool`,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if _, ok := os.LookupEnv(envAccessKeyId); ok {
				_ = os.Setenv(credentials.EnvVarAccessKeyId, os.Getenv(envAccessKeyId))
			}
			if _, ok := os.LookupEnv(envAccessKeySecret); ok {
				_ = os.Setenv(credentials.EnvVarAccessKeySecret, os.Getenv(envAccessKeySecret))
			}
			path, err := expandPath(profilePath)
			if err != nil {
				fmt.Printf("error: parse profile path %s failed: %+v", profilePath, err)
				os.Exit(1)
			}
			_ = os.Setenv(credentials.ENVCredentialFile, path)
		},
	}
)

func init() {
	rrsa.SetupRRSACmd(rootCmd)
	credentialplugin.SetupCredentialPluginCmd(rootCmd)
	version.SetupVersionCmd(rootCmd)

	rootCmd.PersistentFlags().StringVar(&ctl.GlobalOption.Region, "region-id", "cn-hangzhou", "The region to use")
	rootCmd.PersistentFlags().BoolVarP(&ctl.GlobalOption.AssumeYes, "assume-yes", "y", false,
		"Automatic yes to prompts; assume \"yes\" as answer to all prompts and run non-interactively")
	rootCmd.PersistentFlags().StringVar(&profilePath, "profile-file", defaultProfilePath, "Path to credential file")
	//rootCmd.PersistentFlags().BoolVarP(&ctl.GlobalOption.InsecureSkipTLSVerify, "insecure-skip-tls-verify", "", false, "Skips the validity check for the server's certificate")
}

func expandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[1:])
	}
	return path, nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
