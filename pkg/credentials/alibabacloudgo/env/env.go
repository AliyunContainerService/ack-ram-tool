package env

import (
	"errors"
	"fmt"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudgo"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/aliyun/credentials-go/credentials"
)

var (
	accessKeyIdEnvs = []string{
		envNewSdkAccessKeyId,
		envOldSdkAccessKeyID,
		envAliyuncliAccessKeyId1,
		envAliyuncliAccessKeyId2,
		envAccAlibabaCloudAccessKeyId,
		envAliyuncliAccessKeyId3,
	}

	accessKeySecretEnvs = []string{
		envNewSdkAccessKeySecret,
		envOldSdkAccessKeySecret,
		envAliyuncliAccessKeySecret1,
		envAliyuncliAccessKeySecret2,
		envAccAlibabaCloudAccessKeySecret,
		envAliyuncliAccessKeySecret3,
	}

	securityTokenEnvs = []string{
		envNewSdkSecurityToken,
		envOldSdkAccessKeyStsToken,
		envAliyuncliStsToken1,
		envAliyuncliStsToken2,
		envAccAlibabaCloudSecurityToken,
		envAliyuncliStsToken3,
	}

	roleArnEnvs = []string{
		envRoleArn,
	}
	oidcProviderArnEnvs = []string{
		envOidcProviderArn,
	}
	oidcTokenFileEnvs = []string{
		envOidcTokenFile,
	}
	roleSessionNameEnvs = []string{
		envNewSdkRoleSessionName,
		envOldSdkRoleSessionName,
	}

	credentialsURIEnvs = []string{
		envNewSdkCredentialsURI,
	}

	credentialFileEnvs = []string{
		envNewSdkCredentialFile,
	}

	aliyuncliProfileNameEnvs = []string{
		envAliyuncliProfileName1,
		envAliyuncliProfileName2,
		envAliyuncliProfileName3,
	}
	aliyuncliIgnoreProfileEnvs = []string{
		envAliyuncliIgnoreProfile,
	}
	aliyuncliProfilePathEnvs = []string{
		envAliyuncliProfilePath,
	}
)

type CredentialsProviderOptions struct {
	STSEndpoint string
}

// NewCredentialsProvider return a CredentialsProvider base on environment variables
func NewCredentialsProvider(opts CredentialsProviderOptions) (provider.CredentialsProvider, error) {
	keyId := GetAccessKeyId()
	keySecret := GetAccessKeySecret()
	stsToken := GetSecurityToken()
	credURI := GetCredentialsURI()
	roleArn := GetRoleArn()
	oidcProviderArn := GetOIDCProviderArn()
	oidcTokenFile := GetOIDCTokenFile()
	sessionName := GetRoleSessionName()

	config := &credentials.Config{
		AccessKeyId:       stringPoint(keyId),
		AccessKeySecret:   stringPoint(keySecret),
		SecurityToken:     stringPoint(stsToken),
		Url:               stringPoint(credURI),
		RoleArn:           stringPoint(roleArn),
		OIDCProviderArn:   stringPoint(oidcProviderArn),
		OIDCTokenFilePath: stringPoint(oidcTokenFile),
		RoleSessionName:   stringPoint(sessionName),
	}
	if keyId != "" && keySecret != "" && stsToken != "" {
		return provider.NewSTSTokenProvider(keyId, keySecret, stsToken), nil
	}
	if roleArn != "" && oidcProviderArn != "" && oidcTokenFile != "" {
		return provider.NewOIDCProvider(provider.OIDCProviderOptions{
			STSEndpoint:     opts.STSEndpoint,
			SessionName:     sessionName,
			RoleArn:         roleArn,
			OIDCProviderArn: oidcProviderArn,
			OIDCTokenFile:   oidcTokenFile,
			Logger:          &log.ProviderLogWrapper{ZP: log.Logger},
		}), nil
	}
	if keyId != "" && keySecret != "" {
		return provider.NewAccessKeyProvider(keyId, keySecret), nil
	}
	if credURI != "" {
		config.Type = stringPoint("credentials_uri")
	} else {
		return nil, errors.New("not found credentials related environment variables")
	}

	cred, err := credentials.NewCredential(config)
	if err != nil {
		return nil, fmt.Errorf("init credential failed: %w", err)
	}
	return alibabacloudgo.NewCredentialsProviderWrapper(cred), nil
}

func GetAccessKeyId() string {
	return getEnvsValue(accessKeyIdEnvs)
}

func GetAccessKeySecret() string {
	return getEnvsValue(accessKeySecretEnvs)
}

func GetSecurityToken() string {
	return getEnvsValue(securityTokenEnvs)
}

func GetCredentialsURI() string {
	return getEnvsValue(credentialsURIEnvs)
}

func GetCredentialsFile() string {
	return getEnvsValue(credentialFileEnvs)
}

func GetAliyuncliProfileName() string {
	return getEnvsValue(aliyuncliProfileNameEnvs)
}

func GetAliyuncliIgnoreProfile() string {
	return getEnvsValue(aliyuncliIgnoreProfileEnvs)
}

func GetAliyuncliProfilePath() string {
	return getEnvsValue(aliyuncliProfilePathEnvs)
}

func GetRoleArn() string {
	return getEnvsValue(roleArnEnvs)
}

func GetOIDCProviderArn() string {
	return getEnvsValue(oidcProviderArnEnvs)
}

func GetOIDCTokenFile() string {
	return getEnvsValue(oidcTokenFileEnvs)
}

func GetRoleSessionName() string {
	return getEnvsValue(roleSessionNameEnvs)
}

func getEnvsValue(keys []string) string {
	for _, key := range keys {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return ""
}

func stringPoint(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
