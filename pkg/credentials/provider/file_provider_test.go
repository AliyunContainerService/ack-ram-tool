package provider

import (
	"context"
	"os"
	"path"
	"testing"
	"time"
)

func TestFileProvider_Credentials(t *testing.T) {
	d, err := os.MkdirTemp("", "TestFileProvider_Credentials")
	if err != nil {
		t.Errorf("should not error: %+v", err)
		return
	}
	fp := path.Join(d, "test.json")
	os.WriteFile(fp, []byte("abc"), 0600)

	f := NewFileProvider(fp, func(data []byte) (*Credentials, error) {
		return &Credentials{
			AccessKeyId:     "ak_TestFileProvider_Credentials",
			AccessKeySecret: "sk_TestFileProvider_Credentials",
			SecurityToken:   "",
			Expiration:      time.Time{},
		}, nil
	}, FileProviderOptions{})

	cred, err := f.Credentials(context.TODO())
	if err != nil {
		t.Errorf("should not error: %+v", err)
		return
	}
	if cred.AccessKeyId != "ak_TestFileProvider_Credentials" ||
		cred.AccessKeySecret != "sk_TestFileProvider_Credentials" {
		t.Errorf("unexpected case found: %+v", *cred)
		return
	}
}
