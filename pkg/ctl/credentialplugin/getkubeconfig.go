package credentialplugin

import (
	"context"
	"fmt"
	"strings"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const (
	versionV1           = "v1"
	versionV1beta1      = "v1beta1"
	groupVersionV1      = "client.authentication.k8s.io/v1"
	groupVersionV1beta1 = "client.authentication.k8s.io/v1beta1"
	kindExecCredential  = "ExecCredential" // #nosec G101

	commandName = "ack-ram-tool"
)

type credentialMode string

var (
	modeRAMAuthenticatorToken credentialMode = "ram-authenticator-token"
	modeToken                 credentialMode = "token"
	modeCertificate           credentialMode = "certificate"
	modeCert                  credentialMode = "cert"
	selectedMode              string
)

var getKubeconfigCmd = &cobra.Command{
	Use:   "get-kubeconfig",
	Short: "Get a kubeconfig with exec credential plugin format.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := common.GetClientOrDie()
		ctx := context.Background()
		clusterId := ctl.GlobalOption.ClusterId

		kubeconfig, err := client.GetUserKubeConfig(ctx, clusterId,
			getCredentialOpts.privateIpAddress, getCredentialOpts.temporaryDuration)
		common.ExitIfError(err)
		newConf := generateExecKubeconfig(clusterId, kubeconfig, credentialMode(selectedMode))

		d, err := yaml.Marshal(newConf)
		common.ExitIfError(err)
		fmt.Println(string(d))
	},
}

func generateExecKubeconfig(clusterId string, config *types.KubeConfig, mode credentialMode) *types.KubeConfig {
	newConf := &types.KubeConfig{
		Kind:           config.Kind,
		APIVersion:     config.APIVersion,
		Clusters:       config.Clusters,
		Contexts:       config.Contexts,
		CurrentContext: config.CurrentContext,
		Users:          config.Users,
		Preferences:    config.Preferences,
	}
	var users []types.KubeAuthUser
	args := getExecArgs(clusterId, mode, getCredentialOpts)
	args = fillGlobalFlags(args)

	for _, u := range newConf.Users {
		newU := types.KubeAuthUser{
			Name: u.Name,
			User: types.KubeAuthInfo{
				Exec: &types.KubeExecConfig{
					Command:            commandName,
					Args:               args,
					APIVersion:         getApiVersion(getCredentialOpts.apiVersion),
					InstallHint:        "",
					ProvideClusterInfo: false,
					InteractiveMode:    types.NeverExecInteractiveMode,
				},
			},
		}
		users = append(users, newU)
	}
	newConf.Users = users
	return newConf
}

func fillGlobalFlags(args []string) []string {
	if ctl.GlobalOption.GetProfileName() != "" {
		args = append(args, "--profile-name", ctl.GlobalOption.GetProfileName())
	}
	if ctl.GlobalOption.GetIgnoreAliyuncliConfig() {
		args = append(args, "--ignore-aliyun-cli-credentials")
	}
	if ctl.GlobalOption.GetIgnoreEnv() {
		args = append(args, "--ignore-env-credentials")
	}
	if ctl.GlobalOption.GetRoleArn() != "" {
		args = append(args, "--role-arn", ctl.GlobalOption.GetRoleArn())
	}
	args = append(args, "--log-level", log.LogLevelError)
	return args
}

func getExecArgs(clusterId string, mode credentialMode, opt GetCredentialOpts) []string {
	switch mode {
	case modeRAMAuthenticatorToken, modeToken:
		return []string{
			"credential-plugin",
			"get-token",
			"--cluster-id",
			clusterId,
			"--api-version",
			opt.apiVersion,
		}
	default:
		return []string{
			"credential-plugin",
			"get-credential",
			"--cluster-id",
			clusterId,
			"--api-version",
			opt.apiVersion,
			"--expiration",
			fmt.Sprintf("%v", opt.temporaryDuration),
			"--credential-cache-dir",
			opt.cacheDir,
		}
	}
}

func setupGetKubeconfigCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(getKubeconfigCmd)
	common.SetupClusterIdFlag(getKubeconfigCmd)

	getKubeconfigCmd.Flags().DurationVar(&getCredentialOpts.temporaryDuration, "expiration", getCredentialOpts.temporaryDuration, "The certificate expiration")
	getKubeconfigCmd.Flags().BoolVar(&getCredentialOpts.privateIpAddress, "private-address", getCredentialOpts.privateIpAddress, "Use private ip as api-server address")
	getKubeconfigCmd.Flags().StringVarP(&selectedMode, "mode", "m", string(modeCert),
		fmt.Sprintf("credential mode: %s", strings.Join([]string{string(modeCert), string(modeToken)}, " or ")))
	getKubeconfigCmd.Flags().StringVar(&getCredentialOpts.apiVersion, "api-version", "v1beta1", "v1 or v1beta1")
	getKubeconfigCmd.Flags().StringVar(&getCredentialOpts.cacheDir, "credential-cache-dir", getCredentialOpts.cacheDir, "Directory to cache certificate")
	//getcredentialCmd.Flags().BoolVar(&getCredentialOpts.disableCache, "disable-credential-cache", false, "disable credential cache")
}
