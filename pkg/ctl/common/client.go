package common

import (
	"fmt"
	"log"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper/env"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/aliyuncli"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	"github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

type ClientConfig struct {
	regionId                string
	credentialFilePath      string
	aliyuncliConfigFilePath string
	profileName             string
	ignoreEnv               bool
	ignoreAliyuncli         bool
}

func NewClient(config ClientConfig) (*openapi.Client, error) {
	crd, err := getCredential(getCredentialOption{
		credentialFilePath:      config.credentialFilePath,
		aliyuncliConfigFilePath: config.aliyuncliConfigFilePath,
		aliyuncliProfileName:    config.profileName,
		sessionName:             version.BinName(),
		ignoreEnv:               config.ignoreEnv,
		ignoreAliyuncli:         config.ignoreAliyuncli})
	if err != nil {
		return nil, err
	}
	return openapi.NewClient(&client.Config{
		RegionId:   tea.String(config.regionId),
		Credential: crd,
		UserAgent:  tea.String(version.UserAgent()),
	})
}

type getCredentialOption struct {
	credentialFilePath      string
	aliyuncliConfigFilePath string
	aliyuncliProfileName    string
	sessionName             string
	ignoreEnv               bool
	ignoreAliyuncli         bool
}

func getCredential(opt getCredentialOption) (credentials.Credential, error) {
	credentialFilePath := opt.credentialFilePath
	aliyuncliConfigFilePath := opt.aliyuncliConfigFilePath
	sessionName := opt.sessionName
	ignoreEnv := opt.ignoreEnv
	ignoreAliyuncli := opt.ignoreAliyuncli
	aliyuncliProfileName := opt.aliyuncliProfileName

	if credentialFilePath == "" && aliyuncliConfigFilePath == "" {
		if sessionName != "" {
			_ = os.Setenv(env.EnvRoleSessionName, sessionName)
		}
		if !ignoreEnv {
			if cred, err := env.NewCredential(); err == nil && cred != nil {
				log.Println("use credentials from environment variables")
				return cred, err
			}
		}
	}
	if aliyuncliConfigFilePath == "" {
		aliyuncliConfigFilePath, _ = utils.ExpandPath("~/.aliyun/config.json")
	}

	if !ignoreAliyuncli {
		acli, err := aliyuncli.NewCredentialHelper(aliyuncliConfigFilePath, aliyuncliProfileName)
		if err == nil && acli != nil {
			log.Printf("use credentials from aliyun cli (%s)", aliyuncliConfigFilePath)
			return acli.GetCredentials()
		}
	}

	if credentialFilePath != "" {
		if _, err := os.Stat(credentialFilePath); err == nil {
			_ = os.Setenv(credentials.ENVCredentialFile, credentialFilePath)
		}
	} else {
		path, err := utils.ExpandPath(credentials.PATHCredentialFile)
		if err == nil {
			credentialFilePath = path
		}
	}
	log.Printf("use default credentials from %s", credentialFilePath)
	return credentials.NewCredential(nil)
}

func GetClientOrDie() *openapi.Client {
	c, err := NewClient(
		ClientConfig{
			regionId:                ctl.GlobalOption.GetRegion(),
			credentialFilePath:      ctl.GlobalOption.GetCredentialFilePath(),
			aliyuncliConfigFilePath: ctl.GlobalOption.GetAliyuncliConfigFilePath(),
			profileName:             ctl.GlobalOption.GetProfileName(),
			ignoreEnv:               ctl.GlobalOption.GetIgnoreEnv(),
			ignoreAliyuncli:         ctl.GlobalOption.GetIgnoreAliyuncliConfig(),
		},
	)
	if err != nil {
		ExitByError(fmt.Sprintf("init client failed: %+v", err))
	}
	return c
}
