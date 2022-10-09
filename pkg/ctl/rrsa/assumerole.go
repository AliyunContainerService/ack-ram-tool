package rrsa

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/spf13/cobra"
)

type AssumeRoleOpts struct {
	roleArn       string
	oidcArn       string
	oidcTokenFile string
}

var assumeRoleOpts = AssumeRoleOpts{}

var assumeRoleCmd = &cobra.Command{
	Use:   "assume-role",
	Short: "Retrieve a Security Token Service (STS) token to Assume a RAM Role via OIDC Token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var token []byte
		ctx := context.Background()
		roleArn := assumeRoleOpts.roleArn
		oidcArn := assumeRoleOpts.oidcArn
		oidcTokenFile := assumeRoleOpts.oidcTokenFile
		if oidcTokenFile == "-" {
			token, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				common.ExitByError(fmt.Sprintf("read token from stdin failed: %+v", err))
			}
		} else {
			token, err = ioutil.ReadFile(oidcTokenFile)
			if err != nil {
				common.ExitByError(fmt.Sprintf("read token file failed: %+v", err))
			}
		}
		if len(token) < 4 || len(token) > 10000 {
			common.ExitByError("invalid token: The length of OIDC Toke should be between 4 and 10000")
		}

		region := ctl.GlobalOption.Region
		if strings.HasPrefix(region, "cn-") && region != "cn-hongkong" && !ctl.GlobalOption.UseVPCEndpoint {
			// use default sts endpoint for Chinese mainland
			region = ""
		}
		cred, err := openapi.AssumeRoleWithOIDCToken(ctx, oidcArn, roleArn, time.Second*900, token,
			openapi.GetStsEndpoint(region, ctl.GlobalOption.UseVPCEndpoint))
		if err != nil {
			common.ExitByError(fmt.Sprintf("Assume RAM Role failed: %+v", err))
		}
		log.Println("Retrieved a STS token:")
		fmt.Printf("AccessKeyId:       %s\n", cred.AccessKeyId)
		fmt.Printf("AccessKeySecret:   %s\n", cred.AccessKeySecret)
		fmt.Printf("SecurityToken:     %s\n", cred.SecurityToken)
		fmt.Printf("Expiration:        %s\n", cred.Expiration.UTC().Format(time.RFC3339))
	},
}

func setupAssumeRoleCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(assumeRoleCmd)
	assumeRoleCmd.Flags().StringVarP(&assumeRoleOpts.roleArn, "role-arn", "r", "", "The arn of RAM role")
	assumeRoleCmd.Flags().StringVarP(&assumeRoleOpts.oidcArn, "oidc-provider-arn", "p", "", "The arn of OIDC provider")
	assumeRoleCmd.Flags().StringVarP(&assumeRoleOpts.oidcTokenFile, "oidc-token-file", "t", "",
		"Path to OIDC token file. If value is '-', will read token from stdin")
	err := assumeRoleCmd.MarkFlagRequired("role-arn")
	common.ExitIfError(err)
	err = assumeRoleCmd.MarkFlagRequired("oidc-provider-arn")
	common.ExitIfError(err)
	err = assumeRoleCmd.MarkFlagRequired("oidc-token-file")
	common.ExitIfError(err)
}
