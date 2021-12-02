package main

import (
	"fmt"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func main() {
	roleArn := os.Getenv("ALIBABA_CLOUD_ROLE_ARN")
	oidcArn := os.Getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN")
	tokenFile := os.Getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE")
	singer, err := alibabacloudsdkgo.NewRAMRoleArnWithOIDCTokenSigner(oidcArn, roleArn, tokenFile, "", "", 0)
	if err != nil {
		panic(err)
	}
	client, err := sts.NewClientWithAccessKey("cn-hangzhou", "", "")
	if err != nil {
		panic(err)
	}
	client.SetSigner(singer)

	req := sts.CreateGetCallerIdentityRequest()
	req.Scheme = "https"
	req.SetDomain("sts.aliyuncs.com")
	resp, err := client.GetCallerIdentity(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCallerIdentity with oidc token success:\n%s\n", resp.String())
}
