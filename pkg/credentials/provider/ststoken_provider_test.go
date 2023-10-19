package provider

import (
	"context"
	"testing"
	"time"
)

func TestSTSTokenProvider_Credentials(t *testing.T) {
	p := NewSTSTokenProvider("ak", "sk", "tt")
	cred, err := p.Credentials(context.TODO())

	if err != nil {
		t.Log(err)
		t.Errorf("should not return error: %+v", err)
	}

	if cred.AccessKeyId != "ak" ||
		cred.AccessKeySecret != "sk" ||
		cred.SecurityToken != "tt" ||
		!cred.Expiration.IsZero() {
		t.Error("cred value is not expected")
	}
}

func TestSTSTokenProvider_SetExpiration(t *testing.T) {
	p := NewSTSTokenProvider("ak", "sk", "tt")
	tm := time.Now()
	p.SetExpiration(tm)
	cred, err := p.Credentials(context.TODO())

	if err != nil {
		t.Log(err)
		t.Errorf("should not return error: %+v", err)
	}

	if cred.AccessKeyId != "ak" ||
		cred.AccessKeySecret != "sk" ||
		cred.SecurityToken != "tt" ||
		cred.Expiration.IsZero() {
		t.Error("cred value is not expected")
	}
}
