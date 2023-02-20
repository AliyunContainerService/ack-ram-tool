package status

import (
	"bytes"
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"strings"
	"text/template"

	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
)

var rrsaStatusTemplate = `
{{- if .Enabled }}
RRSA feature:          enabled
OIDC Provider Name:    {{ .OIDCName }}
OIDC Provider Arn:     {{ .OIDCArn }}
OIDC Token Issuer:     {{ .TokenIssuer }}
{{- else }}
RRSA feature: disabled
{{- end }}
`

var cmd = &cobra.Command{
	Use:   "status",
	Short: "Show RRSA feature status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()
		clusterId := ctl.GlobalOption.ClusterId
		c, err := common.GetRRSAStatus(context.Background(), clusterId, client)
		if err != nil {
			ctlcommon.ExitByError(fmt.Sprintf("fetch status failed: %+v", err))
		}
		displayRRSAStatus(c)
	},
}

func displayRRSAStatus(c *types.Cluster) {
	rrsac := c.MetaData.RRSAConfig
	buf := bytes.NewBuffer(nil)
	t, _ := template.New("rrsa").Parse(rrsaStatusTemplate)
	_ = t.Execute(buf, rrsac)
	fmt.Printf("%s\n", strings.TrimSpace(buf.String()))
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	ctlcommon.SetupClusterIdFlag(cmd)
}
