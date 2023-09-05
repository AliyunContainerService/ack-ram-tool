package common

import (
	"fmt"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudgo"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudgo/env"
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
	stsEndpoint             string
	sessionName             string
}

func NewClient(config ClientConfig) (*openapi.Client, error) {
	log.Logger.Debugf("use %s as sts endpoint", config.stsEndpoint)
	preP, err := getCredential(getCredentialOption{
		credentialFilePath:      config.credentialFilePath,
		aliyuncliConfigFilePath: config.aliyuncliConfigFilePath,
		aliyuncliProfileName:    config.profileName,
		sessionName:             config.SessionName(),
		ignoreEnv:               config.ignoreEnv,
		ignoreAliyuncli:         config.ignoreAliyuncli,
		stsEndpoint:             config.stsEndpoint,
	})
	if err != nil {
		return nil, fmt.Errorf("get credentials: %w", err)
	}

	p := preP
	if config.roleArn != "" {
		log.Logger.Debug("will use assume role to get credentails with pre credentials")
		p = provider.NewRoleArnProvider(preP, config.roleArn, provider.RoleArnProviderOptions{
			STSEndpoint: config.stsEndpoint,
			SessionName: config.SessionName(),
			Logger:      &log.ProviderLogWrapper{ZP: log.Logger},
		})
	}
	cred := provider.NewCredentialForV2SDK(p, provider.CredentialForV2SDKOptions{
		Logger: &log.ProviderLogWrapper{ZP: log.Logger},
	})
	regionId := config.regionId
	if regionId == "" {
		regionId = "cn-hangzhou"
	}

	return openapi.NewClient(&client.Config{
		RegionId:   tea.String(regionId),
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
	stsEndpoint             string
}

func getCredential(opt getCredentialOption) (provider.CredentialsProvider, error) {
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
			log.Logger.Debug("try to get credentials from environment variables")
			// TODO: support ecs ALIBABA_CLOUD_ECS_METADATA
			if cred, err := env.NewCredentialsProvider(env.CredentialsProviderOptions{
				STSEndpoint: opt.stsEndpoint,
			}); err == nil && cred != nil {
				log.Logger.Debugf("use credentials from environment variables")
				return cred, err
			}
			log.Logger.Debug("not found credentials from environment variables")
		}
	}
	if aliyuncliConfigFilePath == "" {
		aliyuncliConfigFilePath, _ = utils.ExpandPath("~/.aliyun/config.json")
		if path, err := checkFileExist(aliyuncliConfigFilePath); err != nil && os.IsNotExist(err) {
			log.Logger.Debugf("file %s not exist, ignore it", utils.ShortHomePath(aliyuncliConfigFilePath))
			ignoreAliyuncli = true
		} else {
			aliyuncliConfigFilePath = path
		}
	}

	if !ignoreAliyuncli {
		if path, err := checkFileExist(aliyuncliConfigFilePath); err != nil {
			return nil, fmt.Errorf("read file %s: %w", aliyuncliConfigFilePath, err)
		} else {
			aliyuncliConfigFilePath = path
		}
		log.Logger.Debugf("try to get credentials from aliyun cli config file: %s",
			utils.ShortHomePath(aliyuncliConfigFilePath))

		acli, err := aliyuncli.NewCredentialHelper(aliyuncliConfigFilePath, aliyuncliProfileName, opt.stsEndpoint)
		if err == nil && acli != nil {
			log.Logger.Debugf("get credentials from aliyun cli (%s) with profile name %s",
				utils.ShortHomePath(aliyuncliConfigFilePath), acli.ProfileName())
			return acli.GetCredentials()
		}
		if err != nil {
			return nil, fmt.Errorf("get credentials from aliyun cli (%s) failed: %w",
				utils.ShortHomePath(aliyuncliConfigFilePath), err)
		}
	}

	if credentialFilePath == "" {
		path, err := checkFileExist(credentials.PATHCredentialFile)
		if err != nil && os.IsNotExist(err) {
			doc := "https://aliyuncontainerservice.github.io/ack-ram-tool/#credentials"
			return nil, fmt.Errorf("not found any configs for credentials, please configure credentials: %s", doc)
		}
		credentialFilePath = path
	}

	log.Logger.Debugf("get default credentials from %s", utils.ShortHomePath(credentialFilePath))
	if path, err := checkFileExist(credentialFilePath); err != nil {
		return nil, fmt.Errorf("read file %s: %w", credentialFilePath, err)
	} else {
		credentialFilePath = path
	}

	_ = os.Setenv(credentials.ENVCredentialFile, credentialFilePath)
	cred, err := credentials.NewCredential(nil)
	if err != nil {
		return nil, err
	}
	return alibabacloudgo.NewCredentialsProviderWrapper(cred), nil
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
			stsEndpoint:             ctl.GlobalOption.GetSTSEndpoint(),
		},
	)
	if err != nil {
		ExitByError(fmt.Sprintf("init client failed: %+v", err))
	}
	return c
}

func (c ClientConfig) SessionName() string {
	if c.sessionName != "" {
		return c.sessionName
	}
	name := env.GetRoleSessionName()
	if name == "" {
		name = version.BinName()
	}
	return name
}

func checkFileExist(filepath string) (string, error) {
	path, err := utils.ExpandPath(filepath)
	if err == nil {
		filepath = path
	}
	_, err = os.ReadFile(filepath) // #nosec G304
	return filepath, err
}
