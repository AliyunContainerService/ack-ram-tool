package rrsa

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/spf13/cobra"
)

var (
	roleArn       = ""
	oidcArn       = ""
	oidcTokenFile = ""
)

var assumeRoleCmd = &cobra.Command{
	Use:   "assume-role",
	Short: "Retrieve a Security Token Service (STS) token to Assume a RAM Role via OIDC Token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var token []byte
		ctx := context.Background()
		if oidcTokenFile == "-" {
			token, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				exitByError(fmt.Sprintf("read token from stdin failed: %+v", err))
			}
		} else {
			token, err = ioutil.ReadFile(oidcTokenFile)
			if err != nil {
				exitByError(fmt.Sprintf("read token file failed: %+v", err))
			}
		}
		if len(token) < 4 || len(token) > 10000 {
			exitByError("invalid token: The length of OIDCToke should be between 4 and 10000")
		}

		cred, err := openapi.AssumeRoleWithOIDCToken(ctx, oidcArn, roleArn, time.Second*900, token,
			openapi.GetStsEndpoint(ctl.GlobalOption.Region, ctl.GlobalOption.UseVPCEndpoint))
		if err != nil {
			exitByError(fmt.Sprintf("Assume RAM Role failed: %+v", err))
		}
		fmt.Println("Retrieved a STS token:")
		fmt.Printf("AccessKeyId:       %s\n", cred.AccessKeyId)
		fmt.Printf("AccessKeySecret:   %s\n", cred.AccessKeySecret)
		fmt.Printf("SecurityToken:     %s\n", cred.SecurityToken)
		fmt.Printf("Expiration:        %s\n", cred.Expiration.UTC().Format(time.RFC3339))
	},
}

func setupAssumeRoleCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(assumeRoleCmd)
	assumeRoleCmd.Flags().StringVarP(&roleArn, "role-arn", "r", "", "")
	assumeRoleCmd.Flags().StringVarP(&oidcArn, "oidc-provider-arn", "p", "", "")
	assumeRoleCmd.Flags().StringVarP(&oidcTokenFile, "oidc-token-file", "t", "", "")
	assumeRoleCmd.MarkFlagRequired("role-arn")
	assumeRoleCmd.MarkFlagRequired("oidc-provider-arn")
	assumeRoleCmd.MarkFlagRequired("oidc-token-file")
}
