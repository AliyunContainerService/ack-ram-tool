package rrsa

import (
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	"github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

func NewClient(regionId string) (*openapi.Client, error) {
	crd, err := credentials.NewCredential(nil)
	if err != nil {
		return nil, err
	}
	return openapi.NewClient(&client.Config{
		RegionId:   tea.String(regionId),
		Credential: crd,
		UserAgent:  tea.String(version.UserAgent()),
	})
}
