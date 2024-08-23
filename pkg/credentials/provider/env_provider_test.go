package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
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
	envURI := "TestEnvProvider_Credentials_URI"

	t.Run("no env", func(t *testing.T) {
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
			EnvCredentialsURI:  envURI,
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
			EnvCredentialsURI:  envURI,
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
		defer os.Unsetenv(envToken)
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
			EnvCredentialsURI:  envURI,
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

	t.Run("uri env", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
{
  "Code": "Success",
  "AccessKeyId": "<ak id>",
  "AccessKeySecret": "<ak secret>",
  "SecurityToken": "<security token>",
  "Expiration": "2006-01-02T15:04:05Z"
}
`))
		}))
		defer s.Close()

		os.Setenv(envURI, s.URL)
		defer os.Unsetenv(envURI)
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
			EnvCredentialsURI:  envURI,
		})
		defer p.Stop(context.TODO())
		cred, err := p.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if cred.AccessKeyId != "<ak id>" ||
			cred.AccessKeySecret != "<ak secret>" ||
			cred.SecurityToken != "<security token>" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})

	t.Run("oidc env", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
{
  "Credentials": {
     "AccessKeyId": "<oidc ak id>",
     "AccessKeySecret": "<oidc ak secret>",
     "SecurityToken": "<oidc security token>",
     "Expiration": "2006-01-02T15:04:05Z"
  }
}
`))
		}))
		defer s.Close()

		dir, _ := os.MkdirTemp("", "TestEnvProvider_Credentials")
		path := dir + "/" + "tt"
		os.WriteFile(path, []byte("test"), 0644)
		os.Setenv(envRoleArn, "foo")
		os.Setenv(envOidcP, "bar")
		os.Setenv(envOidcT, path)
		defer os.Unsetenv(envRoleArn)
		defer os.Unsetenv(envOidcP)
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
			EnvCredentialsURI:  envURI,
			stsEndpoint:        s.URL,
		})
		defer p.Stop(context.TODO())
		cred, err := p.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if cred.AccessKeyId != "<oidc ak id>" ||
			cred.AccessKeySecret != "<oidc ak secret>" ||
			cred.SecurityToken != "<oidc security token>" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})
}
