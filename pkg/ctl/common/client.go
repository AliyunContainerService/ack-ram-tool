package common

import (
	"fmt"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	"github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

func NewClient(regionId, credentialFilePath, aliyuncliConfigFilePath, aliyuncliProfileName string) (*openapi.Client, error) {
	crd, err := getCredential(credentialFilePath, aliyuncliConfigFilePath, aliyuncliProfileName, version.BinName())
	if err != nil {
		return nil, err
	}
	return openapi.NewClient(&client.Config{
		RegionId:   tea.String(regionId),
		Credential: crd,
		UserAgent:  tea.String(version.UserAgent()),
	})
}

func getCredential(credentialFilePath, aliyuncliConfigFilePath, aliyuncliProfileName, sessionName string) (credentials.Credential, error) {
	cred, err := helper.NewCredential(credentialFilePath, aliyuncliConfigFilePath, aliyuncliProfileName, sessionName)
	if err == nil && cred != nil {
		return cred, err
	}
	if credentialFilePath != "" {
		if _, err := os.Stat(credentialFilePath); err == nil {
			os.Setenv(credentials.ENVCredentialFile, credentialFilePath)
		}
	}
	return credentials.NewCredential(nil)
}

func GetClientOrDie() *openapi.Client {
	c, err := NewClient(
		ctl.GlobalOption.Region,
		ctl.GlobalOption.CredentialFilePath,
		ctl.GlobalOption.AliyuncliConfigFilePath,
		ctl.GlobalOption.AliyuncliProfileName,
	)
	if err != nil {
		ExitByError(fmt.Sprintf("init client failed: %+v", err))
	}
	return c
}
