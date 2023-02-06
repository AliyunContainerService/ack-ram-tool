package main

import (
	"fmt"
	"log"
	"os"

	cs20151215 "github.com/alibabacloud-go/cs-20151215/v3/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
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

func testOpenAPISDK() {
	// 两种方法都可以
	cred := newCredential()
	// or
	// cred := newOidcCredential()

	config := &openapi.Config{Credential: cred}
	config.Endpoint = tea.String("cs.cn-hangzhou.aliyuncs.com")
	client, err := cs20151215.NewClient(config)
	if err != nil {
		panic(err)
	}

	req := &cs20151215.DescribeClustersV1Request{}
	resp, err := client.DescribeClustersV1(req)
	if err != nil {
		panic(err)
	}
	for _, c := range resp.Body.Clusters {
		fmt.Printf("cluster id: %s, cluster name: %s\n", *c.ClusterId, *c.Name)
	}
}

func testOSSSDK() {
	// 两种方法都可以
	cred := newCredential()
	// or
	// cred := newOidcCredential()

	provider := &ossCredentialsProvider{cred: cred}
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
		fmt.Printf("-%s\n", item.Name)
	}
}

func newCredential() credentials.Credential {
	// https://www.alibabacloud.com/help/doc-detail/378661.html
	cred, err := credentials.NewCredential(nil)
	if err != nil {
		panic(err)
	}
	return cred
}

func newOidcCredential() credentials.Credential {
	// https://www.alibabacloud.com/help/doc-detail/378661.html
	config := new(credentials.Config).
		SetType("oidc_role_arn").
		SetRoleArn(os.Getenv(EnvRoleArn)).
		SetOIDCProviderArn(os.Getenv(EnvOidcProviderArn)).
		SetOIDCTokenFilePath(os.Getenv(EnvOidcTokenFile)).
		SetRoleSessionName("test-rrsa-oidc-token")

	oidcCredential, err := credentials.NewCredential(config)
	if err != nil {
		panic(err)
	}
	return oidcCredential
}

type ossCredentials struct {
	teaCred credentials.Credential
}

func (cred *ossCredentials) GetAccessKeyID() string {
	value, err := cred.teaCred.GetAccessKeyId()
	if err != nil {
		log.Printf("get access key id failed: %+v", err)
		return ""
	}
	return tea.StringValue(value)
}

func (cred *ossCredentials) GetAccessKeySecret() string {
	value, err := cred.teaCred.GetAccessKeySecret()
	if err != nil {
		log.Printf("get access key secret failed: %+v", err)
		return ""
	}
	return tea.StringValue(value)
}

func (cred *ossCredentials) GetSecurityToken() string {
	value, err := cred.teaCred.GetSecurityToken()
	if err != nil {
		log.Printf("get access security token failed: %+v", err)
		return ""
	}
	return tea.StringValue(value)
}

type ossCredentialsProvider struct {
	cred credentials.Credential
}

func (p *ossCredentialsProvider) GetCredentials() oss.Credentials {
	return &ossCredentials{teaCred: p.cred}
}

func main() {
	// test open api sdk (https://github.com/aliyun/alibabacloud-go-sdk) use rrsa oidc token
	log.Printf("test open api sdk use rrsa oidc token")
	testOpenAPISDK()

	// test oss sdk (https://github.com/aliyun/aliyun-oss-go-sdk) use rrsa oidc token
	if os.Getenv("TEST_OSS_SDK") == "true" {
		log.Printf("test oss sdk use rrsa oidc token")
		testOSSSDK()
	}
}
