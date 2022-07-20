package helper

import (
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"os"
)

const (
	EnvRoleArn         = "ALIBABA_CLOUD_ROLE_ARN"
	EnvOidcProviderArn = "ALIBABA_CLOUD_OIDC_PROVIDER_ARN"
	EnvOidcTokenFile   = "ALIBABA_CLOUD_OIDC_TOKEN_FILE"
)

func GetOidcSigner(sessionName string) (singer auth.Signer, err error) {
	roleArn := os.Getenv(EnvRoleArn)
	oidcArn := os.Getenv(EnvOidcProviderArn)
	tokenFile := os.Getenv(EnvOidcTokenFile)
	if roleArn == "" {
		return nil, fmt.Errorf("environment variable %q is missing", EnvRoleArn)
	}
	if oidcArn == "" {
		return nil, fmt.Errorf("environment variable %q is missing", EnvOidcProviderArn)
	}
	if tokenFile == "" {
		return nil, fmt.Errorf("environment variable %q is missing", EnvOidcTokenFile)
	}
	if _, err := os.Stat(tokenFile); err != nil {
		return nil, fmt.Errorf("unable to read file at %q: %s", tokenFile, err)
	}

	return alibabacloudsdkgo.NewRAMRoleArnWithOIDCTokenSigner(
		oidcArn, roleArn, tokenFile, "", sessionName, 0)
}
