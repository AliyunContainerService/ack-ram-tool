package credentialplugin

import (
	"encoding/json"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ramauthenticator"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const envTokenExtraQueryKeyPrefix = "ACK_RAM_TOOL_TOKEN_EXTRA_KEY_PREFIX"

type GetTokenOpts struct {
	//clusterId        string
	privateIpAddress bool
	stsEndpoint      string

	apiVersion string
	cacheDir   string
	//disableCache bool
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

		generator := ramauthenticator.NewTokenGenerator(clusterId, client.Credential())
		generator.SetExtraQuery(getExtraTokenQuery())
		token, err := generator.NewToken()
		common.ExitIfError(err)

		cred, err := newTokenExecCredential(token)
		common.ExitIfError(err)

		d, err := json.MarshalIndent(cred, "", " ")
		common.ExitIfError(err)
		fmt.Println(string(d))
	},
}

func getExtraTokenQuery() map[string]string {
	query := make(map[string]string)
	prefix := os.Getenv(envTokenExtraQueryKeyPrefix)
	if prefix == "" {
		return query
	}
	for _, item := range os.Environ() {
		before, after, found := strings.Cut(item, "=")
		if !found {
			continue
		}
		if after == "" {
			continue
		}
		if !strings.HasPrefix(before, prefix) {
			continue
		}
		k := strings.ToLower(strings.TrimPrefix(before, prefix))
		query[k] = after
	}
	return query
}

func newTokenExecCredential(token *ramauthenticator.Token) (*types.ExecCredential, error) {
	version := getApiVersion(getCredentialOpts.apiVersion)
	var exp *types.KubeTime
	if !token.Expiration.IsZero() {
		t := types.NewKubeTime(token.Expiration)
		exp = &t
	}

	cred := &types.ExecCredential{
		KubeTypeMeta: types.KubeTypeMeta{
			Kind:       kindExecCredential,
			APIVersion: version,
		},
		Spec: types.ExecCredentialSpec{
			Interactive: false,
		},
		Status: &types.ExecCredentialStatus{
			ExpirationTimestamp: exp,
			Token:               token.String(),
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
