package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURIProvider_Credentials_success(t *testing.T) {
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

	e := NewURIProvider(s.URL, URIProviderOptions{})
	cred, err := e.Credentials(context.TODO())
	if err != nil {
		t.Error(err.Error())
	}
	if cred.AccessKeyId != "<ak id>" {
		t.Errorf("AccessKeyId is wrong: %s", cred.AccessKeyId)
	}
	if cred.AccessKeySecret != "<ak secret>" {
		t.Errorf("AccessKeySecret is wrong: %s", cred.AccessKeySecret)
	}
	if cred.SecurityToken != "<security token>" {
		t.Errorf("SecurityToken is wrong: %s", cred.SecurityToken)
	}
	if cred.Expiration.Format("2006-01-02T15:04:05Z") != "2006-01-02T15:04:05Z" {
		t.Errorf("Expiration is wrong: %s", cred.Expiration.Format("2006-01-02T15:04:05Z"))
	}

	e.Stop(context.TODO())
}

func TestURIProvider_Credentials_404(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer s.Close()

	e := NewURIProvider(s.URL, URIProviderOptions{})
	defer e.Stop(context.TODO())
	cred, err := e.Credentials(context.TODO())
	if err == nil {
		t.Errorf("err should not nil")
	}
	t.Log(err)
	if cred != nil {
		t.Errorf("cred should nil")
	}
}

func TestURIProvider_Credentials_wrong_resp(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
  "Code": "Success",
  "AccessKeyId": "<ak id>",
  "SecurityToken": "<security token>",
  "Expiration": "2006-01-02T15:04:05Z"
}
`))
	}))
	defer s.Close()

	e := NewURIProvider(s.URL, URIProviderOptions{})
	defer e.Stop(context.TODO())

	cred, err := e.Credentials(context.TODO())
	if err == nil {
		t.Errorf("err should not nil")
	}
	t.Log(err)
	if cred != nil {
		t.Errorf("cred should nil")
	}
}
