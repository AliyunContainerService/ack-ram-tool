package provider

import (
	"context"
	"os"
	"path"
	"testing"
	"time"
)

func TestRemoteProvider_Credentials(t *testing.T) {
	d, err := os.MkdirTemp("", "TestRemoteProvider_Credentials")
	if err != nil {
		t.Errorf("should not error: %+v", err)
		return
	}
	fp := path.Join(d, "test.json")
	os.WriteFile(fp, []byte("abc"), 0600)
	getRawFunc := func(ctx context.Context) ([]byte, error) {
		return os.ReadFile(fp)
	}

	f := NewRemoteProvider(getRawFunc, func(ctx context.Context, data []byte) (*Credentials, error) {
		return &Credentials{
			AccessKeyId:     "ak_TestRemoteProvider_Credentials",
			AccessKeySecret: "sk_TestRemoteProvider_Credentials",
			SecurityToken:   "",
			Expiration:      time.Time{},
		}, nil
	}, RemoteProviderOptions{})

	cred, err := f.Credentials(context.TODO())
	if err != nil {
		t.Errorf("should not error: %+v", err)
		return
	}
	if cred.AccessKeyId != "ak_TestRemoteProvider_Credentials" ||
		cred.AccessKeySecret != "sk_TestRemoteProvider_Credentials" {
		t.Errorf("unexpected case found: %+v", *cred)
		return
	}
}
