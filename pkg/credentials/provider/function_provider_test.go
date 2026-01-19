package provider

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestFunctionProvider_Credentials(t *testing.T) {
	f := NewFunctionProvider(func(ctx context.Context) (*Credentials, error) {
		return &Credentials{
			AccessKeyId:     "ak_TestFunctionProvider_Credentials",
			AccessKeySecret: "sk_TestFunctionProvider_Credentials",
			SecurityToken:   "",
			Expiration:      time.Time{},
		}, nil
	})

	cred, err := f.Credentials(context.TODO())
	if err != nil {
		t.Errorf("should not error: %+v", err)
		return
	}
	if cred.AccessKeyId != "ak_TestFunctionProvider_Credentials" ||
		cred.AccessKeySecret != "sk_TestFunctionProvider_Credentials" {
		t.Errorf("unexpected case found: %+v", *cred)
		return
	}
}

func TestFunctionProvider_Credentials_cache(t *testing.T) {
	var n int

	f := NewFunctionProvider(func(ctx context.Context) (*Credentials, error) {
		n++
		return &Credentials{
			AccessKeyId:     "ak_TestFunctionProvider_Credentials",
			AccessKeySecret: "sk_TestFunctionProvider_Credentials",
			SecurityToken:   fmt.Sprintf("token-%d", n),
			Expiration:      time.Now().Add(time.Hour),
		}, nil
	})

	cred, err := f.Credentials(context.TODO())
	if err != nil {
		t.Errorf("should not error: %+v", err)
		return
	}
	if cred.AccessKeyId != "ak_TestFunctionProvider_Credentials" ||
		cred.AccessKeySecret != "sk_TestFunctionProvider_Credentials" {
		t.Errorf("unexpected case found: %+v", *cred)
		return
	}

	token := cred.SecurityToken
	t.Log(token)
	for i := 0; i < 10; i++ {
		cred, _ = f.Credentials(context.TODO())
		t.Log(cred.SecurityToken)
		if cred.SecurityToken != token {
			t.Errorf("unexpected case found: %+v", *cred)
			return
		}
	}

	f.setNowFunc(func() time.Time {
		return time.Now().Add(time.Hour)
	})
	cred, _ = f.Credentials(context.TODO())
	t.Log(cred.SecurityToken)
	if cred.SecurityToken == token {
		t.Errorf("unexpected case found: %+v", *cred)
	}
}
