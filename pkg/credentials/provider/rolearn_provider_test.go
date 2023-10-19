package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoleArnProvider_Credentials_success(t *testing.T) {
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

	p := NewRoleArnProvider(
		NewAccessKeyProvider("ak1", "sk1"),
		"role_arn",
		RoleArnProviderOptions{
			STSEndpoint: s.URL,
		},
	)

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

func TestRoleArnProvider_Credentials_stop_with_no_stop_method_cp(t *testing.T) {
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

	p := NewRoleArnProvider(
		NewAccessKeyProvider("ak1", "sk1"),
		"role_arn",
		RoleArnProviderOptions{
			STSEndpoint: s.URL,
			Logger:      TLogger{t},
		},
	)

	p.Stop(context.TODO())
	p.Stop(context.TODO())
}

func TestRoleArnProvider_Credentials_stop_with_stop_method_cp(t *testing.T) {
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

	p := NewRoleArnProvider(
		NewOIDCProvider(OIDCProviderOptions{
			Logger: TLogger{t},
		}),
		"role_arn",
		RoleArnProviderOptions{
			STSEndpoint: s.URL,
			Logger:      TLogger{t},
		},
	)

	p.Stop(context.TODO())
	p.Stop(context.TODO())
}

func setupHttpTestServer(t *testing.T, handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})
	s := httptest.NewServer(mux)
	return s
}
