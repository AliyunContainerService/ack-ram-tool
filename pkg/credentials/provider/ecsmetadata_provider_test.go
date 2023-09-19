package provider

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestECSMetadataProvider_Credentials_success(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/latest/api/token":
			{
				fmt.Fprint(w, "token-xxx")
				return
			}
		default:
			{
				fmt.Fprint(w, `
{
     "AccessKeyId": "ak",
     "AccessKeySecret": "sk",
     "SecurityToken": "tt",
     "Expiration": "2206-01-02T15:04:05Z"
}
`)
			}
		}
	})

	p := NewECSMetadataProvider(ECSMetadataProviderOptions{
		Endpoint: s.URL,
		RoleName: "test",
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
