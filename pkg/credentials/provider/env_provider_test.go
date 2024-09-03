package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	os.Setenv(envINIConfigFile, "foo-bar-not-exist")
	os.Setenv(envProfileName, "foo-bar-test-not-exist")
}

func TestEnvProvider_Credentials_ak(t *testing.T) {
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
		os.Setenv(envAk, "ak_TestEnvProvider_Credentials_ak")
		os.Setenv(envSK, "sk")
		defer os.Unsetenv(envAk)
		defer os.Unsetenv(envSK)
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
		if cred.AccessKeyId != "ak_TestEnvProvider_Credentials_ak" ||
			cred.AccessKeySecret != "sk" ||
			cred.SecurityToken != "" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})
}

func TestEnvProvider_Credentials_sts(t *testing.T) {
	envAk := "TestEnvProvider_Credentials_AK_2"
	envSK := "TestEnvProvider_Credentials_SK_2"
	envToken := "TestEnvProvider_Credentials_Token_2"
	envRoleArn := "TestEnvProvider_Credentials_Role_ARN_2"
	envOidcP := "TestEnvProvider_Credentials_OIDC_Pro_2"
	envOidcT := "TestEnvProvider_Credentials_OIDC_Token_2"
	envURI := "TestEnvProvider_Credentials_URI_2"
	os.Setenv(envProfileName, "foo-bar-test")

	t.Run("sts token env", func(t *testing.T) {
		os.Setenv(envAk, "ak_TestEnvProvider_Credentials_sts")
		os.Setenv(envSK, "sk")
		os.Setenv(envToken, "sts-token")
		defer os.Unsetenv(envAk)
		defer os.Unsetenv(envSK)
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
		if cred.AccessKeyId != "ak_TestEnvProvider_Credentials_sts" ||
			cred.AccessKeySecret != "sk" ||
			cred.SecurityToken != "sts-token" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})
}

func TestEnvProvider_Credentials_ak_role(t *testing.T) {
	envAk := "TestEnvProvider_Credentials_AK_3"
	envSK := "TestEnvProvider_Credentials_SK_3"
	envToken := "TestEnvProvider_Credentials_Token_3"
	envRoleArn := "TestEnvProvider_Credentials_Role_ARN_3"
	envOidcP := "TestEnvProvider_Credentials_OIDC_Pro_3"
	envOidcT := "TestEnvProvider_Credentials_OIDC_Token_3"
	envURI := "TestEnvProvider_Credentials_URI_3"
	os.Setenv(envProfileName, "foo-bar-test")

	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak_TestEnvProvider_Credentials_ak_role",
     "AccessKeySecret": "sk",
     "SecurityToken": "sts-token",
     "Expiration": "2206-01-02T15:04:05Z"
  }
}
`)
	})
	defer s.Close()

	t.Run("ak with role env", func(t *testing.T) {
		os.Setenv(envAk, "ak")
		os.Setenv(envSK, "sk")
		os.Setenv(envRoleArn, "rol-arn")
		defer os.Unsetenv(envAk)
		defer os.Unsetenv(envSK)
		defer os.Unsetenv(envRoleArn)
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
			EnvCredentialsURI:  envURI,
			STSEndpoint:        s.URL,
		})
		cred, err := p.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if cred.AccessKeyId != "ak_TestEnvProvider_Credentials_ak_role" ||
			cred.AccessKeySecret != "sk" ||
			cred.SecurityToken != "sts-token" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}

		rp := p.cp.(*RoleArnProvider)
		rpCred, err := rp.cp.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if rpCred.SecurityToken != "" {
			t.Error("SecurityToken should is nil")
		}
	})
}

func TestEnvProvider_Credentials_sts_role(t *testing.T) {
	envAk := "TestEnvProvider_Credentials_AK_3_2"
	envSK := "TestEnvProvider_Credentials_SK_3_2"
	envToken := "TestEnvProvider_Credentials_Token_3_2"
	envRoleArn := "TestEnvProvider_Credentials_Role_ARN_3_2"
	envOidcP := "TestEnvProvider_Credentials_OIDC_Pro_3_2"
	envOidcT := "TestEnvProvider_Credentials_OIDC_Token_3_2"
	envURI := "TestEnvProvider_Credentials_URI_3_2"
	os.Setenv(envProfileName, "foo-bar-test")

	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak_TestEnvProvider_Credentials_sts_role",
     "AccessKeySecret": "sk",
     "SecurityToken": "sts-token",
     "Expiration": "2206-01-02T15:04:05Z"
  }
}
`)
	})
	defer s.Close()

	t.Run("ak with role env", func(t *testing.T) {
		os.Setenv(envAk, "ak")
		os.Setenv(envSK, "sk")
		os.Setenv(envToken, "sts-token-pre")
		os.Setenv(envRoleArn, "rol-arn")
		defer os.Unsetenv(envAk)
		defer os.Unsetenv(envSK)
		defer os.Unsetenv(envToken)
		defer os.Unsetenv(envRoleArn)
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:     envAk,
			EnvAccessKeySecret: envSK,
			EnvSecurityToken:   envToken,
			EnvRoleArn:         envRoleArn,
			EnvOIDCProviderArn: envOidcP,
			EnvOIDCTokenFile:   envOidcT,
			EnvCredentialsURI:  envURI,
			STSEndpoint:        s.URL,
		})
		cred, err := p.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if cred.AccessKeyId != "ak_TestEnvProvider_Credentials_sts_role" ||
			cred.AccessKeySecret != "sk" ||
			cred.SecurityToken != "sts-token" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}

		rp := p.cp.(*RoleArnProvider)
		rpCred, err := rp.cp.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if rpCred.SecurityToken != "sts-token-pre" {
			t.Error("SecurityToken should is sts-token-pre")
		}
	})
}

func TestEnvProvider_Credentials_uri(t *testing.T) {
	envAk := "TestEnvProvider_Credentials_AK_4"
	envSK := "TestEnvProvider_Credentials_SK_4"
	envToken := "TestEnvProvider_Credentials_Token_4"
	envRoleArn := "TestEnvProvider_Credentials_Role_ARN_4"
	envOidcP := "TestEnvProvider_Credentials_OIDC_Pro_4"
	envOidcT := "TestEnvProvider_Credentials_OIDC_Token_4"
	envURI := "TestEnvProvider_Credentials_URI_4"
	os.Setenv(envProfileName, "foo-bar-test")

	t.Run("uri env", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
{
  "Code": "Success",
  "AccessKeyId": "ak_TestEnvProvider_Credentials_uri",
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
		if cred.AccessKeyId != "ak_TestEnvProvider_Credentials_uri" ||
			cred.AccessKeySecret != "<ak secret>" ||
			cred.SecurityToken != "<security token>" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})
}

func TestEnvProvider_Credentials_oidc(t *testing.T) {
	envAk := "TestEnvProvider_Credentials_AK_5"
	envSK := "TestEnvProvider_Credentials_SK_5"
	envToken := "TestEnvProvider_Credentials_Token_5"
	envRoleArn := "TestEnvProvider_Credentials_Role_ARN_5"
	envOidcP := "TestEnvProvider_Credentials_OIDC_Pro_5"
	envOidcT := "TestEnvProvider_Credentials_OIDC_Token_5"
	envURI := "TestEnvProvider_Credentials_URI_5"
	os.Setenv(envProfileName, "foo-bar-test")
	t.Run("oidc env", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
{
  "Credentials": {
     "AccessKeyId": "ak_TestEnvProvider_Credentials_oidc",
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
			STSEndpoint:        s.URL,
		})
		defer p.Stop(context.TODO())
		cred, err := p.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if cred.AccessKeyId != "ak_TestEnvProvider_Credentials_oidc" ||
			cred.AccessKeySecret != "<oidc ak secret>" ||
			cred.SecurityToken != "<oidc security token>" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})
}

func TestEnvProvider_Credentials_ini(t *testing.T) {
	envAk := "TestEnvProvider_Credentials_AK_2"
	envSK := "TestEnvProvider_Credentials_SK_2"
	envToken := "TestEnvProvider_Credentials_Token_2"
	envRoleArn := "TestEnvProvider_Credentials_Role_ARN_2"
	envOidcP := "TestEnvProvider_Credentials_OIDC_Pro_2"
	envOidcT := "TestEnvProvider_Credentials_OIDC_Token_2"
	envURI := "TestEnvProvider_Credentials_URI_2"
	envConfig := "TestEnvProvider_Credentials_ini_env"
	envProfile := "TestEnvProvider_Credentials_ini_name"
	os.Setenv(envProfileName, "foo-bar-test")

	t.Run("sts token env", func(t *testing.T) {
		os.Setenv(envConfig, "testdata/ak.ini")
		os.Setenv(envProfile, "default")
		defer os.Unsetenv(envConfig)
		defer os.Unsetenv(envProfile)
		p := NewEnvProvider(EnvProviderOptions{
			EnvAccessKeyId:       envAk,
			EnvAccessKeySecret:   envSK,
			EnvSecurityToken:     envToken,
			EnvRoleArn:           envRoleArn,
			EnvOIDCProviderArn:   envOidcP,
			EnvOIDCTokenFile:     envOidcT,
			EnvCredentialsURI:    envURI,
			EnvConfigFile:        envConfig,
			EnvConfigSectionName: envProfile,
		})
		cred, err := p.Credentials(context.TODO())
		if err != nil {
			t.Errorf("should no error: %+v", err)
		}
		if cred.AccessKeyId != "foo_from_ini" ||
			cred.AccessKeySecret != "bar_from_ini" ||
			cred.SecurityToken != "" {
			t.Errorf("got unexpected cred: %+v", *cred)
		}
	})
}
