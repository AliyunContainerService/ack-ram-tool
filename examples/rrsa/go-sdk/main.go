package main

import (
	"fmt"
	"log"
	"os"

	cs20151215 "github.com/alibabacloud-go/cs-20151215/v5/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
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
	cred, err := newCredential()
	// or
	// cred, err := newOidcCredential()
	if err != nil {
		panic(err)
	}
	if _, err := cred.GetCredential(); err != nil {
		log.Fatalf("get credentails failed: %+v", err)
	}

	config := &openapi.Config{
		Credential: cred,
	}
	config.Endpoint = tea.String("cs.cn-hangzhou.aliyuncs.com")
	client, err := cs20151215.NewClient(config)
	if err != nil {
		panic(err)
	}

	req := &cs20151215.DescribeClustersForRegionRequest{}
	region := tea.String("cn-hangzhou")
	resp, err := client.DescribeClustersForRegion(region, req)
	if err != nil {
		panic(err)
	}
	for _, c := range resp.Body.Clusters {
		fmt.Printf("cluster id: %s, cluster name: %s\n", *c.ClusterId, *c.Name)
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

func main() {
	// test open api sdk (https://github.com/aliyun/alibabacloud-go-sdk) using rrsa oidc token
	log.Printf("test open api sdk use rrsa oidc token")
	testOpenAPISDK()
}
