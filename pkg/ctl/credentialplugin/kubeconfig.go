package credentialplugin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
)

const (
	versionV1          = "client.authentication.k8s.io/v1"
	versionV1beta1     = "client.authentication.k8s.io/v1beta1"
	kindExecCredential = "ExecCredential"
)

type GetKubeconfigOpts struct {
	clusterId         string
	temporaryDuration time.Duration
	privateIpAddress  bool

	apiVersion string
}

var getKubeconfigOpts = GetKubeconfigOpts{}

var getKubeconfigCmd = &cobra.Command{
	Use:   "get-credential",
	Short: "get ACK credential",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := common.GetClientOrDie()
		ctx := context.Background()

		cred, err := getKubeconfigExecCredential(ctx, client)
		common.ExitIfError(err)

		d, err := json.MarshalIndent(cred, "", " ")
		common.ExitIfError(err)
		fmt.Println(string(d))
	},
}

func getKubeconfigExecCredential(ctx context.Context, client *openapi.Client) (*types.ExecCredential, error) {
	config, err := client.GetUserKubeConfig(ctx, getKubeconfigOpts.clusterId,
		getKubeconfigOpts.privateIpAddress, getKubeconfigOpts.temporaryDuration)
	if err != nil {
		return nil, err
	}
	version := versionV1
	if getKubeconfigOpts.apiVersion != "v1" {
		version = versionV1beta1
	}
	cred := &types.ExecCredential{
		KubeTypeMeta: types.KubeTypeMeta{
			Kind:       kindExecCredential,
			APIVersion: version,
		},
		Spec: types.ExecCredentialSpec{
			//Cluster: &types.ExecCluster{
			//	Server:                   config.Server,
			//	CertificateAuthorityData: base64.StdEncoding.EncodeToString([]byte(config.CertificateAuthorityData)),
			//},
			Interactive: false,
		},
		Status: &types.ExecCredentialStatus{
			ExpirationTimestamp:   &types.KubeTime{Time: config.Expiration},
			ClientCertificateData: config.ClientCertificateData,
			ClientKeyData:         config.ClientKeyData,
		},
	}

	return cred, nil
}

func setupGetKubeconfigCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(getKubeconfigCmd)
	getKubeconfigCmd.Flags().StringVarP(&getKubeconfigOpts.clusterId, "cluster-id", "c", "", "The cluster id to use")
	err := getKubeconfigCmd.MarkFlagRequired("cluster-id")
	common.ExitIfError(err)

	getKubeconfigCmd.Flags().DurationVar(&getKubeconfigOpts.temporaryDuration, "expiration", time.Hour, "The credential expiration")
	getKubeconfigCmd.Flags().BoolVar(&getKubeconfigOpts.privateIpAddress, "private", false, "Use private ip as api-server address")
	getKubeconfigCmd.Flags().StringVar(&getKubeconfigOpts.apiVersion, "api-version", "v1beta1", "v1 or v1beta1")
}

func init() {
	setupGetKubeconfigCmd(rootCmd)
}
