package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"
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
