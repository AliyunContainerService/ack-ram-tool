package provider

import (
	"context"
	"testing"
)

func TestAccessKeyProvider_Credentials(t *testing.T) {
	p := NewAccessKeyProvider("ak", "sk")
	cred, err := p.Credentials(context.TODO())

	if err != nil {
		t.Log(err)
		t.Errorf("should not return error: %+v", err)
	}

	if cred.AccessKeyId != "ak" ||
		cred.AccessKeySecret != "sk" ||
		cred.SecurityToken != "" ||
		!cred.Expiration.IsZero() {
		t.Error("cred value is not expected")
	}
}
