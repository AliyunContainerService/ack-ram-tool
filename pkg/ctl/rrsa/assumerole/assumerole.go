package assumerole

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/spf13/cobra"
)

type Option struct {
	roleArn       string
	oidcArn       string
	oidcTokenFile string
}

var opts = Option{}

var cmd = &cobra.Command{
	Use:   "assume-role",
	Short: "Retrieve a Security Token Service (STS) token to Assume a RAM Role via OIDC Token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var token []byte
		ctx := context.Background()
		roleArn := opts.roleArn
		oidcArn := opts.oidcArn
		oidcTokenFile := opts.oidcTokenFile
		if oidcTokenFile == "-" {
			token, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				ctlcommon.ExitByError(fmt.Sprintf("read token from stdin failed: %+v", err))
			}
		} else {
			token, err = ioutil.ReadFile(oidcTokenFile)
			if err != nil {
				ctlcommon.ExitByError(fmt.Sprintf("read token file failed: %+v", err))
			}
		}
		if len(token) < 4 || len(token) > 10000 {
			ctlcommon.ExitByError("invalid token: The length of OIDC Toke should be between 4 and 10000")
		}

		region := ctl.GlobalOption.Region
		if strings.HasPrefix(region, "cn-") && region != "cn-hongkong" && !ctl.GlobalOption.UseVPCEndpoint {
			// use default sts endpoint for Chinese mainland
			region = ""
		}
		cred, err := openapi.AssumeRoleWithOIDCToken(ctx, oidcArn, roleArn, time.Second*900, token,
			openapi.GetStsEndpoint(region, ctl.GlobalOption.UseVPCEndpoint))
		if err != nil {
			ctlcommon.ExitByError(fmt.Sprintf("Assume RAM Role failed: %+v", err))
		}
		log.Println("Retrieved a STS token:")
		fmt.Printf("AccessKeyId:       %s\n", cred.AccessKeyId)
		fmt.Printf("AccessKeySecret:   %s\n", cred.AccessKeySecret)
		fmt.Printf("SecurityToken:     %s\n", cred.SecurityToken)
		fmt.Printf("Expiration:        %s\n", cred.Expiration.UTC().Format(time.RFC3339))
	},
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)

	cmd.Flags().StringVarP(&opts.roleArn, "role-arn", "r", "", "The arn of RAM role")
	cmd.Flags().StringVarP(&opts.oidcArn, "oidc-provider-arn", "p", "", "The arn of OIDC provider")
	cmd.Flags().StringVarP(&opts.oidcTokenFile, "oidc-token-file", "t", "",
		"Path to OIDC token file. If value is '-', will read token from stdin")

	ctlcommon.ExitIfError(cmd.MarkFlagRequired("role-arn"))
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("oidc-provider-arn"))
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("oidc-token-file"))
}
