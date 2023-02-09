package credentialplugin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
)

type GetCredentialOpts struct {
	clusterId         string
	temporaryDuration time.Duration
	privateIpAddress  bool

	apiVersion string
	cacheDir   string
	//disableCache bool
}

var getCredentialOpts = GetCredentialOpts{
	temporaryDuration: time.Hour,
	privateIpAddress:  false,
	apiVersion:        versionV1beta1,
	cacheDir:          defaultCacheDir,
}

var getCredentialCmd = &cobra.Command{
	Use:   "get-credential",
	Short: "Get ACK credential",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := common.GetClientOrDie()
		ctx := context.Background()
		clusterId := ctl.GlobalOption.ClusterId
		getCredentialOpts.clusterId = clusterId

		cacheDir, err := common.EnsureDir(getCredentialOpts.cacheDir)
		common.ExitIfError(err)
		cache := NewCredentialCache(cacheDir, getCredentialOpts)
		cred, err := cache.GetCredential()
		if err != nil && err != errNoValidCache && err != errNeedRefreshCache {
			common.ExitIfError(err)
		}

		if cred == nil {
			cred, err = getKubeconfigExecCredential(ctx, clusterId, client)
			common.ExitIfError(err)
			err = cache.SaveCredential(cred)
			common.ExitIfError(err)
		}

		d, err := json.MarshalIndent(cred, "", " ")
		common.ExitIfError(err)
		fmt.Println(string(d))
	},
}

func getKubeconfigExecCredential(ctx context.Context, clusterId string, client *openapi.Client) (*types.ExecCredential, error) {
	kubeconfig, err := client.GetUserKubeConfig(ctx, clusterId,
		true, getCredentialOpts.temporaryDuration)
	if err != nil {
		return nil, err
	}
	config := &types.ClusterCredential{}
	if err := config.LoadKubeConfig(kubeconfig); err != nil {
		return nil, err
	}
	version := getApiVersion(getCredentialOpts.apiVersion)
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

func getApiVersion(version string) string {
	if version != versionV1 {
		return groupVersionV1beta1
	}
	return groupVersionV1
}

func setupGetCredentialCmdCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(getCredentialCmd)
	common.SetupClusterIdFlag(getCredentialCmd)

	getCredentialCmd.Flags().DurationVar(&getCredentialOpts.temporaryDuration, "expiration", time.Hour, "The credential expiration")
	//getCredentialCmd.Flags().BoolVar(&getCredentialOpts.privateIpAddress, "private-address", getCredentialOpts.privateIpAddress, "Use private ip as api-server address")
	getCredentialCmd.Flags().StringVar(&getCredentialOpts.apiVersion, "api-version", getCredentialOpts.apiVersion, "v1 or v1beta1")
	getCredentialCmd.Flags().StringVar(&getCredentialOpts.cacheDir, "credential-cache-dir", getCredentialOpts.cacheDir, "Directory to cache credential")
	//getcredentialCmd.Flags().BoolVar(&getCredentialOpts.disableCache, "disable-credential-cache", false, "disable credential cache")
}
