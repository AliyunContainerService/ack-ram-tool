package provider

import (
	"context"
	"fmt"
	"os"
)

const (
	envAccessKeyId     = "ALIBABA_CLOUD_ACCESS_KEY_ID"
	envAccessKeySecret = "ALIBABA_CLOUD_ACCESS_KEY_SECRET"
	envSecurityToken   = "ALIBABA_CLOUD_SECURITY_TOKEN"
)

type EnvProvider struct {
	cred *Credentials

	envAccessKeyId     string
	envAccessKeySecret string
	envSecurityToken   string
}

type EnvProviderOptions struct {
	EnvAccessKeyId     string
	EnvAccessKeySecret string
	EnvSecurityToken   string
}

func NewEnvProvider(opts EnvProviderOptions) *EnvProvider {
	opts.applyDefaults()

	return &EnvProvider{
		cred: &Credentials{
			AccessKeyId:     os.Getenv(opts.EnvAccessKeyId),
			AccessKeySecret: os.Getenv(opts.EnvAccessKeySecret),
			SecurityToken:   os.Getenv(opts.EnvSecurityToken),
		},
		envAccessKeyId:     opts.EnvAccessKeyId,
		envAccessKeySecret: opts.EnvAccessKeySecret,
		envSecurityToken:   opts.EnvSecurityToken,
	}
}

func (e *EnvProvider) Credentials(ctx context.Context) (*Credentials, error) {
	if e.cred.AccessKeyId == "" || e.cred.AccessKeySecret == "" {
		return nil, NewNotEnableError(fmt.Errorf("env %s or %s is empty",
			e.envAccessKeyId, e.envAccessKeySecret))
	}

	return e.cred.DeepCopy(), nil
}

func (o *EnvProviderOptions) applyDefaults() {
	if o.EnvAccessKeyId == "" {
		o.EnvAccessKeyId = envAccessKeyId
	}
	if o.EnvAccessKeySecret == "" {
		o.EnvAccessKeySecret = envAccessKeySecret
	}
	if o.EnvSecurityToken == "" {
		o.EnvSecurityToken = envSecurityToken
	}
}
