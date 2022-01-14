package main

import (
	"fmt"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudgo"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
)

func main() {
	roleArn := os.Getenv("ALIBABA_CLOUD_ROLE_ARN")
	oidcArn := os.Getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN")
	tokenFile := os.Getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE")
	cred, err := alibabacloudgo.NewRAMRoleArnWithOIDCTokenCredential(oidcArn, roleArn, tokenFile, "", "", 0)
	if err != nil {
		panic(err)
	}
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
	fmt.Printf("GetCallerIdentity with oidc token success:\n%s\n", resp.String())
}
