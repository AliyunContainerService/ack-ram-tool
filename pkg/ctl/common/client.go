package common

import (
	"fmt"
	"log"
	"net/url"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/env"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	"github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

func NewClient(regionId string) (*openapi.Client, error) {
	config := getCredConfig()
	if config != nil {
		log.Printf("credential type: %s", tea.StringValue(config.Type))
	}
	crd, err := credentials.NewCredential(config)
	if err != nil {
		return nil, err
	}
	return openapi.NewClient(&client.Config{
		RegionId:   tea.String(regionId),
		Credential: crd,
		UserAgent:  tea.String(version.UserAgent()),
	})
}

func getCredConfig() *credentials.Config {
	var config *credentials.Config
	if rawUri := env.GetCredentialsURI(); rawUri != "" {
		if _, err := url.Parse(rawUri); err == nil {
			config = &credentials.Config{
				Type: tea.String("credentials_uri"),
				Url:  tea.String(rawUri),
			}
			return config
		}
	}

	kid := env.GetAccessKeyId()
	ks := env.GetAccessKeySecret()
	st := env.GetSecurityToken()
	if kid != "" && ks != "" && st != "" {
		config = &credentials.Config{
			Type:            tea.String("sts"),
			AccessKeyId:     tea.String(kid),
			AccessKeySecret: tea.String(ks),
			SecurityToken:   tea.String(st),
		}
		return config
	}
	return config
}

func GetClientOrDie() *openapi.Client {
	c, err := NewClient(ctl.GlobalOption.Region)
	if err != nil {
		ExitByError(fmt.Sprintf("init client failed: %+v", err))
	}
	return c
}
