package provider

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
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
	var callCount int32
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&callCount, 1)
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

	cv := atomic.LoadInt32(&callCount)
	if cv < 1 {
		t.Errorf("callCount should >= 1: %v", cv)
	}

	time.Sleep(time.Second)
	cv = atomic.LoadInt32(&callCount)
	if cv <= 1 {
		t.Errorf("callCount should > 1: %v", cv)
	}

	cp.Stop(context.TODO())
	time.Sleep(time.Second)
	curr := atomic.LoadInt32(&callCount)
	time.Sleep(time.Second)

	cv = atomic.LoadInt32(&callCount)
	if cv != curr {
		t.Errorf("callCount should == %v: %v", curr, cv)
	}

	cp.Stop(context.TODO())
	cp.Stop(context.TODO())
	cp.Stop(context.TODO())
}
