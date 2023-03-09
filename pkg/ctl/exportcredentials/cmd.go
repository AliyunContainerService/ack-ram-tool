package exportcredentials

import (
	"context"
	"fmt"
	"net/http"
	"strings"

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
		log.Logger.Infof("args: %#v", args)
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
			strings.Join(formats, " or ")))
	cmd.Flags().StringVarP(&opt.serve, "serve", "s", "",
		"start a server to export credentials")
}
