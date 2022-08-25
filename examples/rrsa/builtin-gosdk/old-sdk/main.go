package main

import (
	"fmt"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudgo/helper"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func main() {
	singer, err := helper.GetOidcSigner("test-old-sdk-use-odic-token")
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
	// get endpoint from https://www.alibabacloud.com/help/resource-access-management/latest/endpoints
	req.SetDomain("sts.aliyuncs.com")
	resp, err := client.GetCallerIdentity(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCallerIdentity with oidc token success:\n%s\n", resp.String())
}
