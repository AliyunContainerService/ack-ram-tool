package rrsa

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
)

var (
	defaultRegion = "cn-hangzhou"
	clusterId     string
)

var rrsaStatusTemplate = `
{{- if .Enabled }}
RRSA feature: enabled
OIDC Name:    {{ .OIDCName }}
OIDC Arn:     {{ .OIDCArn }}
Issuer:       {{ .Issuer }}
{{- else }}
RRSA feature: disabled
{{- end }}
`

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show RRSA feature status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		c, err := getRRSAStatus(context.Background(), clusterId, client)
		if err != nil {
			exitByError(fmt.Sprintf("fetch status failed: %+v", err))
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

func getRRSAStatus(ctx context.Context, clusterId string, client openapi.CSClientInterface) (*types.Cluster, error) {
	c, err := client.GetCluster(ctx, clusterId)
	return c, err
}

func setupStatusCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringVarP(&clusterId, "cluster-id", "c", "", "")
	err := statusCmd.MarkFlagRequired("cluster-id")
	exitIfError(err)
}
