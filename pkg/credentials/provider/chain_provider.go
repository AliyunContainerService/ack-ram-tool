package provider

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type ChainProvider struct {
	providers []CredentialsProvider

	preProvider string
	Logger      Logger
}

func NewChainProvider(providers ...CredentialsProvider) *ChainProvider {
	if len(providers) == 0 {
		return DefaultChainProvider()
	}
	return &ChainProvider{
		providers: providers,
	}
}

func (c *ChainProvider) Credentials(ctx context.Context) (*Credentials, error) {
	var notEnableErrors []string

	for _, p := range c.providers {
		cred, err := p.Credentials(ctx)
		if err != nil {
			if _, ok := err.(*NotEnableError); ok {
				c.logger().Debug(fmt.Sprintf("provider %T not enabled will try to next: %s", p, err.Error()))
				notEnableErrors = append(notEnableErrors, fmt.Sprintf("provider %T not enabled: %s", p, err.Error()))
				continue
			}
		}
		pT := fmt.Sprintf("%T", p)
		if err == nil {
			if c.preProvider != pT {
				c.preProvider = pT
				c.logger().Info(fmt.Sprintf("switch to using provider %s", pT))
			}
			return cred, nil
		}
		return cred, fmt.Errorf("get credentials via %s failed: %w", pT, err)
	}
	return nil, fmt.Errorf("no available credentials providers: %s", strings.Join(notEnableErrors, ", "))
}

func (c *ChainProvider) logger() Logger {
	if c.Logger != nil {
		return c.Logger
	}
	return defaultLog
}

func DefaultChainProvider() *ChainProvider {
	return NewChainProvider(
		NewEnvProvider(EnvProviderOptions{}),
		NewOIDCProvider(OIDCProviderOptions{
			RefreshPeriod: time.Minute * 10,
		}),
		NewEncryptedFileProvider(EncryptedFileProviderOptions{
			RefreshPeriod: time.Minute * 10,
		}),
		NewECSMetadataProvider(ECSMetadataProviderOptions{
			RefreshPeriod: time.Minute * 10,
		}),
	)
}
