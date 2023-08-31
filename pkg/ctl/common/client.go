package common

import (
	"fmt"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudgo"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper/env"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/aliyuncli"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
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
	roleArn                 string
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

	var p provider.CredentialsProvider
	p = alibabacloudgo.NewCredentialsProviderWrapper(crd)
	if config.roleArn != "" {
		p = provider.NewRoleArnProvider(p, config.roleArn, provider.RoleArnProviderOptions{
			SessionName: version.BinName(),
			Logger:      &log.ProviderLogWrapper{ZP: log.Logger},
		})
	}
	cred := provider.NewCredentialForV2SDK(p, provider.CredentialForV2SDKOptions{})

	return openapi.NewClient(&client.Config{
		RegionId:   tea.String(config.regionId),
		Credential: cred,
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
				log.Logger.Debugf("use credentials from environment variables")
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
			log.Logger.Debugf("use credentials from aliyun cli (%s) with profile name %s",
				utils.ShortHomePath(aliyuncliConfigFilePath), acli.ProfileName())
			return acli.GetCredentials()
		} else {
			log.Logger.Warnf("use credentials from aliyun cli (%s) failed: %+v",
				utils.ShortHomePath(aliyuncliConfigFilePath), err)
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
	log.Logger.Debugf("use default credentials from %s", utils.ShortHomePath(credentialFilePath))
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
			roleArn:                 ctl.GlobalOption.GetRoleArn(),
		},
	)
	if err != nil {
		ExitByError(fmt.Sprintf("init client failed: %+v", err))
	}
	return c
}
