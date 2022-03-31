package main

import (
	"fmt"
	"log"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	// github.com/aliyun/credentials-go >= v1.2.1
	"github.com/aliyun/credentials-go/credentials"
)

func testOpenAPISDK(cred credentials.Credential) {
	client, err := sts.NewClient(&openapi.Config{
		Endpoint:   tea.String("sts.aliyuncs.com"),
		Credential: cred,
	})
	if err != nil {
		panic(err)
	}
	resp, err := client.GetCallerIdentity()
	if err != nil {
		panic(err)
	}
	fmt.Printf("call sts.GetCallerIdentity via oidc token success:\n%s\n", resp.String())
}

func testOSSSDK(cred credentials.Credential) {
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
	roleArn := os.Getenv("ALIBABA_CLOUD_ROLE_ARN")
	oidcArn := os.Getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN")
	tokenFile := os.Getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE")

	config := new(credentials.Config).
		SetType("oidc_role_arn").
		SetOIDCProviderArn(oidcArn).
		SetOIDCTokenFilePath(tokenFile).
		SetRoleArn(roleArn).
		SetRoleSessionName("test-rrsa-oidc-token")

	oidcCredential, err := credentials.NewCredential(config)
	if err != nil {
		panic(err)
	}

	// test open api sdk (https://github.com/aliyun/alibabacloud-go-sdk) use rrsa oidc token
	fmt.Println("\ntest open api sdk use rrsa oidc token")
	testOpenAPISDK(oidcCredential)

	// test oss sdk (https://github.com/aliyun/aliyun-oss-go-sdk) use rrsa oidc token
	if os.Getenv("TEST_OSS_SDK") == "true" {
		fmt.Println("\ntest oss sdk use rrsa oidc token")
		testOSSSDK(oidcCredential)
	}
}
