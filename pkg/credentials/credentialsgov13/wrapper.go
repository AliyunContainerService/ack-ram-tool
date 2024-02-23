package credentialsgov13

import (
	"context"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/aliyun/credentials-go/credentials"
)

type CredentialsWrapper struct {
	*provider.CredentialForV2SDK

	p provider.CredentialsProvider
}

func NewCredentialsWrapper(p provider.CredentialsProvider,
	opts provider.CredentialForV2SDKOptions) *CredentialsWrapper {
	w := provider.NewCredentialForV2SDK(p, opts)
	return &CredentialsWrapper{
		CredentialForV2SDK: w,
		p:                  p,
	}
}

func (c *CredentialsWrapper) GetCredential() (*credentials.CredentialModel, error) {
	cred, err := c.p.Credentials(context.TODO())
	if err != nil {
		return nil, err
	}
	m := &credentials.CredentialModel{
		AccessKeyId:     &cred.AccessKeyId,
		AccessKeySecret: &cred.AccessKeySecret,
		SecurityToken:   nil,
		BearerToken:     nil,
		Type:            c.GetType(),
	}
	if cred.SecurityToken != "" {
		m.SecurityToken = &cred.SecurityToken
	}
	return m, nil
}
