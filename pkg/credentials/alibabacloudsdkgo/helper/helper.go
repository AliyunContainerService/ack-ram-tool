package helper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper/aliyuncli"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper/env"
	"github.com/aliyun/credentials-go/credentials"
)

const (
	EnvRoleArn         = "ALIBABA_CLOUD_ROLE_ARN"
	EnvOidcProviderArn = "ALIBABA_CLOUD_OIDC_PROVIDER_ARN"
	EnvOidcTokenFile   = "ALIBABA_CLOUD_OIDC_TOKEN_FILE"
)

// NewCredential return a Credential base on:
// * environment variables
// * credentialFilePath: credential file
// * aliyuncliConfigFilePath: aliyun cli config file
// * aliyuncliProfileName: profile name of aliyun cli
func NewCredential(credentialFilePath, aliyuncliConfigFilePath, aliyuncliProfileName, sessionName string) (credentials.Credential, error) {
	if credentialFilePath == "" {
		credentialFilePath = env.GetCredentialsFile()
	}
	if credentialFilePath != "" {
		credentialFilePath, _ = expandPath(credentialFilePath)
	}
	if credentialFilePath != "" {
		if _, err := os.Stat(credentialFilePath); err == nil {
			_ = os.Setenv(credentials.ENVCredentialFile, credentialFilePath)
		}
	}
	if aliyuncliProfileName == "" {
		aliyuncliProfileName = env.GetAliyuncliProfileName()
	}
	if sessionName != "" {
		_ = os.Setenv(env.EnvRoleSessionName, sessionName)
	}
	if rawP := env.GetAliyuncliProfilePath(); aliyuncliConfigFilePath == "" && rawP != "" {
		if path, err := expandPath(rawP); err == nil && path != "" {
			if _, err := os.Stat(path); err == nil {
				aliyuncliConfigFilePath = path
			}
		}
	}
	if aliyuncliConfigFilePath == "" || env.GetAliyuncliIgnoreProfile() == "TRUE" {
		if cred, err := env.NewCredential(); err == nil && cred != nil {
			return cred, err
		}
	}
	cred, err := aliyuncli.NewCredential(aliyuncliConfigFilePath, aliyuncliProfileName)
	return cred, err
}

func HaveOidcCredentialRequiredEnv() bool {
	return os.Getenv(EnvRoleArn) != "" &&
		os.Getenv(EnvOidcProviderArn) != "" &&
		os.Getenv(EnvOidcTokenFile) != ""
}

func NewOidcCredential(sessionName string) (credential credentials.Credential, err error) {
	return GetOidcCredential(sessionName)
}

// Deprecated: Use NewOidcCredential instead
func GetOidcCredential(sessionName string) (credential credentials.Credential, err error) {
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

	config := new(credentials.Config).
		SetType("oidc_role_arn").
		SetOIDCProviderArn(oidcArn).
		SetOIDCTokenFilePath(tokenFile).
		SetRoleArn(roleArn).
		SetRoleSessionName(sessionName)

	return credentials.NewCredential(config)
}

func expandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[1:])
	}
	return path, nil
}
