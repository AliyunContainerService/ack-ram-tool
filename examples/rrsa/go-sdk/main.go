package main

import (
	"fmt"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

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

	client, err := sts.NewClient(&openapi.Config{
		Endpoint:   tea.String("sts.aliyuncs.com"),
		Credential: oidcCredential,
	})
	if err != nil {
		panic(err)
	}
	resp, err := client.GetCallerIdentity()
	if err != nil {
		panic(err)
	}
	fmt.Printf("call GetCallerIdentity via oidc token success:\n%s\n", resp.String())
}
