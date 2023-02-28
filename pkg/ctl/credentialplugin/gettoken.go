package credentialplugin

import (
	"encoding/json"
	"fmt"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ramauthenticator"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
)

type GetTokenOpts struct {
	//clusterId        string
	privateIpAddress bool
	stsEndpoint      string

	apiVersion string
	cacheDir   string
	//disableCache bool
}

var getTokenOpts = GetTokenOpts{
	privateIpAddress: false,
	apiVersion:       versionV1beta1,
	stsEndpoint:      "sts.aliyuncs.com",
	cacheDir:         defaultCacheDir,
}

var getTokenCmd = &cobra.Command{
	Use:   "get-token",
	Short: "Get token for ack-ram-authenticator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := common.GetClientOrDie()
		//ctx := context.Background()
		clusterId := ctl.GlobalOption.ClusterId
		getCredentialOpts.clusterId = clusterId

		token, err := ramauthenticator.GenerateToken(
			clusterId, getTokenOpts.stsEndpoint, client.Credential())
		common.ExitIfError(err)

		cred, err := newTokenExecCredential(token)
		common.ExitIfError(err)

		d, err := json.MarshalIndent(cred, "", " ")
		common.ExitIfError(err)
		fmt.Println(string(d))
	},
}

func newTokenExecCredential(token *ramauthenticator.Token) (*types.ExecCredential, error) {
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
			Token: token.String(),
		},
	}

	return cred, nil
}

func setupGetTokenCmdCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(getTokenCmd)
	common.SetupClusterIdFlag(getTokenCmd)

	getTokenCmd.Flags().StringVar(&getCredentialOpts.apiVersion, "api-version", getCredentialOpts.apiVersion, "v1 or v1beta1")
	//getCredentialCmd.Flags().StringVar(&getCredentialOpts.cacheDir, "credential-cache-dir", getCredentialOpts.cacheDir, "Directory to cache credential")
	//getcredentialCmd.Flags().BoolVar(&getCredentialOpts.disableCache, "disable-credential-cache", false, "disable credential cache")
}
