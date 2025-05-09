package exportcredentials

import (
	"context"
	"fmt"
	"strings"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/spf13/cobra"
)

type option struct {
	format string
}

var opt = option{}
var (
	formatAliyunCLIConfigJSON       = "aliyun-cli-config-json"
	formatAliyunCLIConfigJSONShort  = "aliyuncli-json"
	formatAliyunCLIURIJSON          = "aliyun-cli-uri-json"
	formatAliyunCLIURIJSONShort     = "aliyuncli-uri"
	formatECSMetadataJSON           = "ecs-metadata-json"
	formatECSMetadataJSONShort      = "ecs"
	formatCredentialFileIni         = "credential-file-ini" // #nosec G101
	formatCredentialFileIniShort    = "ini"                 // #nosec G101
	formatEnvironmentVariables      = "environment-variables"
	formatEnvironmentVariablesShort = "env"
)
var formats = []string{
	formatAliyunCLIConfigJSON,
	formatAliyunCLIConfigJSONShort,
	formatAliyunCLIURIJSON,
	formatAliyunCLIURIJSONShort,
	formatECSMetadataJSON,
	formatECSMetadataJSONShort,
	formatCredentialFileIni,
	formatCredentialFileIniShort,
	formatEnvironmentVariables,
	formatEnvironmentVariablesShort,
}

var cmd = &cobra.Command{
	Use:   "export-credentials",
	Short: "Export credentials in various formats",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()
		cred, err := getCredentials(client)
		ctlcommon.ExitIfError(err)

		if (opt.format == formatEnvironmentVariables ||
			opt.format == formatEnvironmentVariablesShort) &&
			len(args) > 0 {
			err = runUserCommands(context.Background(), *cred, args, nil, nil, nil)
			ctlcommon.ExitIfError(err)
			return
		}

		output := cred.Format(opt.format)
		fmt.Printf("%s\n", output)
	},
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)

	cmd.Flags().StringVarP(&opt.format, "format", "f", formatAliyunCLIConfigJSON,
		fmt.Sprintf("The output format to display credentials (%s)",
			strings.Join(formats, ", ")))

	cmd.Flags().StringVar(
		&ctl.GlobalOption.FinalAssumeRoleAnotherRoleArn, "role-arn", "",
		"Assume an RAM Role ARN when send request or sign token")
}
