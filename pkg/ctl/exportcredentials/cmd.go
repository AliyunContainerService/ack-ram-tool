package exportcredentials

import (
	"context"
	"fmt"
	"net/http"
	"net/netip"
	"strings"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/spf13/cobra"
)

type option struct {
	format string
	serve  string
}

var opt = option{}
var (
	formatAliyunCLIConfigJSON  = "aliyun-cli-config-json"
	formatAliyunCLIURIJSON     = "aliyun-cli-uri-json"
	formatECSMetadataJSON      = "ecs-metadata-json"
	formatCredentialFileIni    = "credential-file-ini" // #nosec G101
	formatEnvironmentVariables = "environment-variables"
)
var formats = []string{
	formatAliyunCLIConfigJSON,
	formatAliyunCLIURIJSON,
	formatECSMetadataJSON,
	formatCredentialFileIni,
	formatEnvironmentVariables,
}

var cmd = &cobra.Command{
	Use:   "export-credentials",
	Short: "Export credentials in various formats",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()

		if opt.serve == "" {
			cred, err := getCredentials(client)
			ctlcommon.ExitIfError(err)

			if opt.format == formatEnvironmentVariables && len(args) > 0 {
				err = runUserCommands(context.Background(), *cred, args, nil, nil)
				ctlcommon.ExitIfError(err)
				return
			}

			output := cred.Format(opt.format)
			fmt.Printf("%s\n", output)
			return
		}

		addr, err := netip.ParseAddrPort(opt.serve)
		if err != nil {
			ctlcommon.ExitByError(fmt.Sprintf("parse the --serve flag failed: %s", err))
		}
		if !addr.Addr().IsLoopback() {
			ctlcommon.ExitByError("the --serve flag only support loopback address")
		}
		log.Logger.Warnf("Serving HTTP on %s", opt.serve)
		if err := startCredServer(client); err != http.ErrServerClosed {
			ctlcommon.ExitIfError(err)
		}
	},
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)

	cmd.Flags().StringVarP(&opt.format, "format", "f", formatAliyunCLIConfigJSON,
		fmt.Sprintf("The output format to display credentials (%s)",
			strings.Join(formats, ", ")))
	cmd.Flags().StringVarP(&opt.serve, "serve", "s", "",
		"start a server to export credentials (e.g. 127.0.0.1:6666), the host part only support 127.0.0.1")

	cmd.Flags().StringVar(
		&ctl.GlobalOption.FinalAssumeRoleAnotherRoleArn, "role-arn", "",
		"Assume an RAM Role ARN when send request or sign token")
}
