package common

import (
	"fmt"
	"net/url"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/env"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	"github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

func NewClient(regionId string) (*openapi.Client, error) {
	crd, err := getCredential()
	if err != nil {
		return nil, err
	}
	return openapi.NewClient(&client.Config{
		RegionId:   tea.String(regionId),
		Credential: crd,
		UserAgent:  tea.String(version.UserAgent()),
	})
}

func getCredential() (credentials.Credential, error) {
	kid := env.GetAccessKeyId()
	ks := env.GetAccessKeySecret()
	st := env.GetSecurityToken()
	if kid != "" && ks != "" && st != "" {
		config := &credentials.Config{
			Type:            tea.String("sts"),
			AccessKeyId:     tea.String(kid),
			AccessKeySecret: tea.String(ks),
			SecurityToken:   tea.String(st),
		}
		return credentials.NewCredential(config)
	}

	if rawUri := env.GetCredentialsURI(); rawUri != "" {
		if _, err := url.Parse(rawUri); err == nil {
			config := &credentials.Config{
				Type: tea.String("credentials_uri"),
				Url:  tea.String(rawUri),
			}
			return credentials.NewCredential(config)
		}
	}

	if helper.HaveOidcCredentialRequiredEnv() {
		return helper.NewOidcCredential(version.BinName())
	}

	return credentials.NewCredential(nil)
}

func GetClientOrDie() *openapi.Client {
	c, err := NewClient(ctl.GlobalOption.Region)
	if err != nil {
		ExitByError(fmt.Sprintf("init client failed: %+v", err))
	}
	return c
}
