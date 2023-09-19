package provider

import (
	"context"
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
