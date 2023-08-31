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
	logPrefix   string
}

func NewChainProvider(providers ...CredentialsProvider) *ChainProvider {
	if len(providers) == 0 {
		return DefaultChainProvider()
	}
	return &ChainProvider{
		providers: providers,
		logPrefix: "[ChainProvider]",
	}
}

func (c *ChainProvider) Credentials(ctx context.Context) (*Credentials, error) {
	var notEnableErrors []string

	for _, p := range c.providers {
		cred, err := p.Credentials(ctx)
		if err != nil {
			if _, ok := err.(*NotEnableError); ok {
				c.logger().Debug(fmt.Sprintf("%s provider %T not enabled will try to next: %s",
					c.logPrefix, p, err.Error()))
				notEnableErrors = append(notEnableErrors, fmt.Sprintf("provider %T not enabled: %s", p, err.Error()))
				continue
			}
		}
		pT := fmt.Sprintf("%T", p)
		if err == nil {
			if c.preProvider != pT {
				c.preProvider = pT
				c.logger().Info(fmt.Sprintf("%s switch to using provider %s", c.logPrefix, pT))
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

type DefaultChainProviderOptions struct {
	STSEndpoint string
	Logger      Logger
}

func NewDefaultChainProvider(opts DefaultChainProviderOptions) *ChainProvider {
	p := NewChainProvider(
		NewEnvProvider(EnvProviderOptions{}),
		NewOIDCProvider(OIDCProviderOptions{
			STSEndpoint:   opts.STSEndpoint,
			RefreshPeriod: time.Minute * 10,
			Logger:        opts.Logger,
		}),
		NewEncryptedFileProvider(EncryptedFileProviderOptions{
			RefreshPeriod: time.Minute * 10,
			Logger:        opts.Logger,
		}),
		NewECSMetadataProvider(ECSMetadataProviderOptions{
			RefreshPeriod: time.Minute * 10,
			Logger:        opts.Logger,
		}),
	)
	p.Logger = opts.Logger
	return p
}

// Deprecated: use NewDefaultChainProvider instead
func DefaultChainProvider() *ChainProvider {
	return NewDefaultChainProvider(DefaultChainProviderOptions{})
}

// Deprecated: use NewDefaultChainProvider instead
func DefaultChainProviderWithLogger(l Logger) *ChainProvider {
	return NewDefaultChainProvider(DefaultChainProviderOptions{
		Logger: l,
	})
}
