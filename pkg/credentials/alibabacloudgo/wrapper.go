package alibabacloudgo

import (
	"context"
	"fmt"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

type CredentialsProviderWrapper struct {
	cred credentials.Credential
}

func NewCredentialsProviderWrapper(cred credentials.Credential) *CredentialsProviderWrapper {
	return &CredentialsProviderWrapper{
		cred: cred,
	}
}

func (c CredentialsProviderWrapper) Credentials(ctx context.Context) (*provider.Credentials, error) {
	ak, err := c.cred.GetAccessKeyId()
	if err != nil {
		return nil, fmt.Errorf("get access key id failed: %w", err)
	}
	sk, err := c.cred.GetAccessKeySecret()
	if err != nil {
		return nil, fmt.Errorf("get access key secret failed: %w", err)
	}
	token, err := c.cred.GetSecurityToken()
	if err != nil {
		return nil, fmt.Errorf("get security token failed: %w", err)
	}

	return &provider.Credentials{
		AccessKeyId:     tea.StringValue(ak),
		AccessKeySecret: tea.StringValue(sk),
		SecurityToken:   tea.StringValue(token),
		Expiration:      time.Time{},
	}, nil
}
