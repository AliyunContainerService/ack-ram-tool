package env

import (
	"errors"
	"os"

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

// NewCredential return a Credential base on environment variables
func NewCredential() (credentials.Credential, error) {
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
		config.Type = stringPoint("sts")
	} else if credURI != "" {
		config.Type = stringPoint("credentials_uri")
	} else if keyId != "" && keySecret != "" {
		config.Type = stringPoint("access_key")
	} else if roleArn != "" && oidcProviderArn != "" && oidcTokenFile != "" {
		config.Type = stringPoint("oidc_role_arn")
	} else {
		return nil, errors.New("not found credentials related environment variables")
	}

	return credentials.NewCredential(config)
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
