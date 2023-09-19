package provider

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestChainProvider_Credentials_success(t *testing.T) {
	p1 := NewAccessKeyProvider("", "")
	p2 := NewAccessKeyProvider("", "")
	p3 := NewSTSTokenProvider("ak3", "sk3", "sts3")
	cp := NewChainProvider(p1, p2, p3)

	cred, err := cp.Credentials(context.TODO())
	if err != nil {
		t.Errorf("should no error: %+v", err)
	}
	if cred.AccessKeyId != "ak3" ||
		cred.AccessKeySecret != "sk3" ||
		cred.SecurityToken != "sts3" {
		t.Errorf("unexpect ret: %+v", *cred)
	}
}

func TestChainProvider_Credentials_no_provider(t *testing.T) {
	cp := NewChainProvider(NewAccessKeyProvider("", ""))
	cred, err := cp.Credentials(context.TODO())
	if err == nil {
		t.Errorf("should return error: %+v", err)
	}
	t.Log(err)
	if cred != nil {
		t.Errorf("should return nil: %+v", *cred)
	}
}

func TestChainProvider_Stop(t *testing.T) {
	callCount := 0
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		callCount++
		fmt.Fprint(w, `
{}
`)
	})

	p := NewRoleArnProvider(
		NewAccessKeyProvider("ak1", "sk1"),
		"role_arn",
		RoleArnProviderOptions{
			STSEndpoint:   s.URL,
			RefreshPeriod: time.Millisecond * 100,
			Logger:        TLogger{t: t},
		},
	)

	cp := NewChainProvider(p)
	cp.Logger = TLogger{t: t}
	cp.Credentials(context.TODO())

	if callCount < 1 {
		t.Errorf("callCount should >= 1: %v", callCount)
	}

	time.Sleep(time.Second)
	if callCount <= 1 {
		t.Errorf("callCount should > 1: %v", callCount)
	}

	cp.Stop(context.TODO())
	time.Sleep(time.Second)
	curr := callCount
	time.Sleep(time.Second)

	if callCount != curr {
		t.Errorf("callCount should == %v: %v", curr, callCount)
	}

	cp.Stop(context.TODO())
	cp.Stop(context.TODO())
	cp.Stop(context.TODO())
}
