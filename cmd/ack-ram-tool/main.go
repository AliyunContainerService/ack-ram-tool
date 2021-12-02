package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/version"
	"github.com/aliyun/credentials-go/credentials"
	"github.com/spf13/cobra"
)

var defaultProfilePath = filepath.Join("~", ".alibabacloud", "credentials")
var profilePath = ""

var (
	rootCmd = &cobra.Command{
		Use:   "ack-ram-tool",
		Short: "A command line utility for using RAM in ACK.",
		Long: `A command line utility for using RAM in ACK.

More info: https://github.com/AliyunContainerService/ack-ram-tool`,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
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
	version.SetupVersionCmd(rootCmd)
	rootCmd.PersistentFlags().BoolVarP(&ctl.GlobalOption.AssumeYes, "assume-yes", "y", false,
		"Automatic yes to prompts; assume \"yes\" as answer to all prompts and run non-interactively")
	rootCmd.PersistentFlags().StringVar(&profilePath, "profile-file", defaultProfilePath, "Path to credential file")
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
