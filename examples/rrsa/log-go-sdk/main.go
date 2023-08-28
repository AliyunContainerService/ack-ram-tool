package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	sls "github.com/aliyun/aliyun-log-go-sdk"
)

type CredentialsProvider struct {
	p provider.CredentialsProvider
}

func NewSLSClient(c *CredentialsProvider, endpoint string) (sls.ClientInterface, error) {
	cred, err := c.Credentials(context.TODO())
	if err != nil {
		return nil, err
	}
	if cred.SecurityToken == "" {
		return sls.CreateNormalInterface(endpoint, cred.AccessKeyId, cred.AccessKeySecret, ""), nil
	}
	shutdown := make(chan struct{})
	return sls.CreateTokenAutoUpdateClient(endpoint, c.TokenUpdateFunc, shutdown)
}

func (c *CredentialsProvider) Credentials(ctx context.Context) (*provider.Credentials, error) {
	return c.p.Credentials(ctx)
}

func (c *CredentialsProvider) TokenUpdateFunc() (accessKeyID, accessKeySecret, securityToken string, expireTime time.Time, err error) {
	cred, err := c.p.Credentials(context.TODO())
	if err != nil {
		return "", "", "", time.Time{}, err
	}
	return cred.AccessKeyId, cred.AccessKeySecret, cred.SecurityToken, cred.Expiration, nil
}

func testLogSDK() {
	prov := provider.NewChainProvider(
		provider.NewEnvProvider(provider.EnvProviderOptions{}),
		provider.NewOIDCProvider(provider.OIDCProviderOptions{}),
	)
	p := &CredentialsProvider{
		p: prov,
	}

	endpoint := "cn-hangzhou.log.aliyuncs.com"
	client, err := NewSLSClient(p, endpoint)
	if err != nil {
		panic(err)
	}

	ret, err := client.ListProject()
	if err != nil {
		panic(err)
	}
	fmt.Println("call log.ListProject via oidc token success:")
	for _, name := range ret {
		fmt.Printf("- %s\n", name)
	}
}

func main() {
	// test log sdk (https://github.com/aliyun/aliyun-log-go-sdk) use rrsa oidc token
	log.Printf("test log sdk using rrsa oidc token")
	testLogSDK()
}
