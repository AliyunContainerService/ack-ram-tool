package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	// github.com/aliyun/credentials-go >= v1.2.6
	"github.com/aliyun/credentials-go/credentials"
)

const (
	EnvRoleArn         = "ALIBABA_CLOUD_ROLE_ARN"
	EnvOidcProviderArn = "ALIBABA_CLOUD_OIDC_PROVIDER_ARN"
	EnvOidcTokenFile   = "ALIBABA_CLOUD_OIDC_TOKEN_FILE"
)

func testOSSSDK() {
	// 两种方法都可以
	cred, err := newCredential()
	// or
	// cred, err := newOidcCredential()
	if err != nil {
		panic(err)
	}
	if _, err := cred.GetAccessKeyId(); err != nil {
		log.Fatalf("get credentails failed: %+v", err)
	}

	provider := &CredentialsProvider{
		cred: cred,
	}
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", "", "",
		oss.SetCredentialsProvider(provider))
	if err != nil {
		panic(err)
	}
	ret, err := client.ListBuckets()
	if err != nil {
		panic(err)
	}

	fmt.Println("call oss.listBuckets via oidc token success:")
	for _, item := range ret.Buckets {
		fmt.Printf("- %s\n", item.Name)
	}
}

func newCredential() (credentials.Credential, error) {
	// https://www.alibabacloud.com/help/doc-detail/378661.html
	cred, err := credentials.NewCredential(nil)
	return cred, err
}

func newOidcCredential() (credentials.Credential, error) {
	// https://www.alibabacloud.com/help/doc-detail/378661.html
	config := new(credentials.Config).
		SetType("oidc_role_arn").
		SetRoleArn(os.Getenv(EnvRoleArn)).
		SetOIDCProviderArn(os.Getenv(EnvOidcProviderArn)).
		SetOIDCTokenFilePath(os.Getenv(EnvOidcTokenFile)).
		SetRoleSessionName("test-rrsa-oidc-token")
	// https://next.api.aliyun.com/product/Sts
	// config.SetSTSEndpoint("sts-vpc.cn-hangzhou.aliyuncs.com")

	oidcCredential, err := credentials.NewCredential(config)
	return oidcCredential, err
}

type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

type CredentialsProvider struct {
	cred credentials.Credential
}

func (c *Credentials) GetAccessKeyID() string {
	return c.AccessKeyId
}

func (c *Credentials) GetAccessKeySecret() string {
	return c.AccessKeySecret
}

func (c *Credentials) GetSecurityToken() string {
	return c.SecurityToken
}

func (p *CredentialsProvider) GetCredentials() oss.Credentials {
	cm, err := p.cred.GetCredential()
	if err != nil {
		log.Printf("get credentials failed: %+v", err)
		return &Credentials{}
	}

	return &Credentials{
		AccessKeyId:     tea.StringValue(cm.AccessKeyId),
		AccessKeySecret: tea.StringValue(cm.AccessKeySecret),
		SecurityToken:   tea.StringValue(cm.SecurityToken),
	}
}

func main() {
	// test oss sdk (https://github.com/aliyun/aliyun-oss-go-sdk) use rrsa oidc token
	log.Printf("test oss sdk using rrsa oidc token")
	testOSSSDK()
}
