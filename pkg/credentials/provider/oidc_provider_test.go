package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestOIDCProvider_Credentials_success(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak",
     "AccessKeySecret": "sk",
     "SecurityToken": "tt",
     "Expiration": "2206-01-02T15:04:05Z"
  }
}
`)
	})

	p := NewOIDCProvider(OIDCProviderOptions{
		STSEndpoint:     s.URL,
		RoleArn:         "role_arn",
		OIDCProviderArn: "oidc_arn",
		OIDCTokenFile:   os.Args[0],
	})

	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Log(err)
		t.Errorf("should no error: %+v", err)
	}

	if cred.AccessKeyId != "ak" ||
		cred.AccessKeySecret != "sk" ||
		cred.SecurityToken != "tt" ||
		cred.Expiration.IsZero() {
		t.Errorf("got unexpected cred")
	}
}

func TestOIDCProvider_Credentials_with_file_token(t *testing.T) {
	var gotToken string
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		gotToken = r.FormValue("OIDCToken")
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak",
     "AccessKeySecret": "sk",
     "SecurityToken": "tt",
     "Expiration": "2206-01-02T15:04:05Z"
  }
}
`)
	})

	dir := t.TempDir()
	tokenFile := fmt.Sprintf("%s/token", dir)
	tokenContent := "token_from_file"
	os.WriteFile(tokenFile, []byte(tokenContent), 0600)
	defer os.Remove(tokenFile)

	p := NewOIDCProvider(OIDCProviderOptions{
		STSEndpoint:     s.URL,
		RoleArn:         "role_arn",
		OIDCProviderArn: "oidc_arn",
		OIDCTokenFile:   tokenFile,
	})

	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Log(err)
		t.Errorf("should no error: %+v", err)
	}

	if cred.AccessKeyId != "ak" ||
		cred.AccessKeySecret != "sk" ||
		cred.SecurityToken != "tt" ||
		cred.Expiration.IsZero() {
		t.Errorf("got unexpected cred")
	}
	if gotToken != tokenContent {
		t.Errorf("got unexpected token: %s", gotToken)
	}
}

func TestOIDCProvider_Credentials_with_string_token(t *testing.T) {
	var gotToken string
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		gotToken = r.FormValue("OIDCToken")
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak",
     "AccessKeySecret": "sk",
     "SecurityToken": "tt",
     "Expiration": "2206-01-02T15:04:05Z"
  }
}
`)
	})

	tokenContent := "token_from_string"

	p := NewOIDCProvider(OIDCProviderOptions{
		STSEndpoint:     s.URL,
		RoleArn:         "role_arn",
		OIDCProviderArn: "oidc_arn",
		OIDCToken:       tokenContent,
	})

	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Log(err)
		t.Errorf("should no error: %+v", err)
	}

	if cred.AccessKeyId != "ak" ||
		cred.AccessKeySecret != "sk" ||
		cred.SecurityToken != "tt" ||
		cred.Expiration.IsZero() {
		t.Errorf("got unexpected cred")
	}
	if gotToken != tokenContent {
		t.Errorf("got unexpected token: %s", gotToken)
	}
}

func TestOIDCProvider_Credentials_error_no_token(t *testing.T) {
	p := NewOIDCProvider(OIDCProviderOptions{
		STSEndpoint:     "127.0.0.1",
		RoleArn:         "role_arn",
		OIDCProviderArn: "oidc_arn",
		OIDCToken:       "",
	})

	cred, err := p.Credentials(context.TODO())
	if err == nil {
		t.Error("should error")
	}
	t.Log(err)
	if !strings.Contains(err.Error(), "roleArn, oidcProviderArn or oidcTokenFile is empty") {
		t.Errorf("got unexpected error: %s", err)
	}
	if cred != nil {
		t.Errorf("got unexpected cred: %#v", cred)
	}
}

func TestOIDCProvider_Credentials_error_invalid_json(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak",
     "AccessKeySecret": "sk",
     "SecurityToken": "tt",
     "Expiration": "2206-01-02T15:04:05Z",
  }
}
`)
	})

	tokenContent := "token_from_string"

	p := NewOIDCProvider(OIDCProviderOptions{
		STSEndpoint:     s.URL,
		RoleArn:         "role_arn",
		OIDCProviderArn: "oidc_arn",
		OIDCToken:       tokenContent,
	})

	cred, err := p.Credentials(context.TODO())
	if err == nil {
		t.Error("should error")
	}
	t.Log(err)
	if !strings.Contains(err.Error(), "parse AssumeRoleWithOIDC body failed") {
		t.Errorf("got unexpected error: %s", err)
	}
	if cred != nil {
		t.Errorf("got unexpected cred: %#v", cred)
	}
}

func TestOIDCProvider_Credentials_error_invalid_data(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak",
     "AccessKeySecret": "sk"
  }
}
`)
	})

	tokenContent := "token_from_string"

	p := NewOIDCProvider(OIDCProviderOptions{
		STSEndpoint:     s.URL,
		RoleArn:         "role_arn",
		OIDCProviderArn: "oidc_arn",
		OIDCToken:       tokenContent,
	})

	cred, err := p.Credentials(context.TODO())
	if err == nil {
		t.Error("should error")
	}
	t.Log(err)
	if !strings.Contains(err.Error(), "call AssumeRoleWithOIDC failed") {
		t.Errorf("got unexpected error: %s", err)
	}
	if cred != nil {
		t.Errorf("got unexpected cred: %#v", cred)
	}
}

func TestOIDCProvider_Credentials_error_invalid_date(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak",
     "AccessKeySecret": "sk",
     "SecurityToken": "tt",
     "Expiration": "2206-01-02 15:04:05Z"
  }
}
`)
	})

	tokenContent := "token_from_string"

	p := NewOIDCProvider(OIDCProviderOptions{
		STSEndpoint:     s.URL,
		RoleArn:         "role_arn",
		OIDCProviderArn: "oidc_arn",
		OIDCToken:       tokenContent,
	})

	cred, err := p.Credentials(context.TODO())
	if err == nil {
		t.Error("should error")
	}
	t.Log(err)
	if !strings.Contains(err.Error(), `parse Expiration "2206-01-02 15:04:05Z" failed`) {
		t.Errorf("got unexpected error: %s", err)
	}
	if cred != nil {
		t.Errorf("got unexpected cred: %#v", cred)
	}
}
