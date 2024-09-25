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
	cred, err := c.cred.GetCredential()
	if err != nil {
		return nil, fmt.Errorf("get credentails failed: %w", err)
	}

	return &provider.Credentials{
		AccessKeyId:     tea.StringValue(cred.AccessKeyId),
		AccessKeySecret: tea.StringValue(cred.AccessKeySecret),
		SecurityToken:   tea.StringValue(cred.SecurityToken),
		Expiration:      time.Time{},
	}, nil
}
