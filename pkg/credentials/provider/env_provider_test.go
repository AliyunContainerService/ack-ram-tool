package provider

import (
	"context"
	"os"
	"testing"
)

func TestEnvProvider_Credentials(t *testing.T) {
	envAk := "TestEnvProvider_Credentials_AK"
	envSK := "TestEnvProvider_Credentials_SK"
	envToken := "TestEnvProvider_Credentials_Token"
	envRoleArn := "TestEnvProvider_Credentials_Role_ARN"
	envOidcP := "TestEnvProvider_Credentials_OIDC_Pro"
	envOidcT := "TestEnvProvider_Credentials_OIDC_Token"

	t.Run("no env", func(t *testing.T) {
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
		})
		cred, err := p.Credentials(context.TODO())
		if err == nil {
			t.Errorf("should return error: %+v", err)
		}
		t.Log(err)
		if cred != nil {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})

	t.Run("only ak env", func(t *testing.T) {
		os.Setenv(envAk, "ak")
		os.Setenv(envSK, "sk")
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
		})
		cred, err := p.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if cred.AccessKeyId != "ak" ||
			cred.AccessKeySecret != "sk" ||
			cred.SecurityToken != "" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})

	t.Run("sts token env", func(t *testing.T) {
		os.Setenv(envAk, "ak")
		os.Setenv(envSK, "sk")
		os.Setenv(envToken, "sts-token")
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
		})
		cred, err := p.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if cred.AccessKeyId != "ak" ||
			cred.AccessKeySecret != "sk" ||
			cred.SecurityToken != "sts-token" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})
}
