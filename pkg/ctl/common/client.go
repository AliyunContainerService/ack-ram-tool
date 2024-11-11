package common

import (
	"context"
	"fmt"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudgo/env"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/credentialsgov13"
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

const defaultConnectTimeoutSeconds = 30

type credentialType string

const (
	credentialTypeImds    credentialType = "imds"
	credentialTypeEcsRole credentialType = "ecs-role"
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
	credentialType          string

	endpoints openapi.Endpoints
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
		credentialType:          credentialType(config.credentialType),
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
	cred := credentialsgov13.NewCredentialsWrapper(p, provider.CredentialForV2SDKOptions{
		Logger: &log.ProviderLogWrapper{ZP: log.Logger},
	})
	regionId := config.regionId
	if regionId == "" {
		regionId = "cn-hangzhou"
	}

	return openapi.NewClientWithEndpoints(&client.Config{
		RegionId:       tea.String(regionId),
		Credential:     cred,
		UserAgent:      tea.String(version.UserAgent()),
		ConnectTimeout: tea.Int(defaultConnectTimeoutSeconds),
	}, config.endpoints)
}

type getCredentialOption struct {
	credentialFilePath      string
	aliyuncliConfigFilePath string
	aliyuncliProfileName    string
	sessionName             string
	ignoreEnv               bool
	ignoreAliyuncli         bool
	stsEndpoint             string
	credentialType          credentialType
}

func getCredential(opt getCredentialOption) (provider.CredentialsProvider, error) {
	credentialFilePath := opt.credentialFilePath
	aliyuncliConfigFilePath := opt.aliyuncliConfigFilePath
	sessionName := opt.sessionName
	ignoreEnv := opt.ignoreEnv
	ignoreAliyuncli := opt.ignoreAliyuncli
	aliyuncliProfileName := opt.aliyuncliProfileName

	switch opt.credentialType {
	case credentialTypeImds, credentialTypeEcsRole:
		return provider.NewECSMetadataProvider(provider.ECSMetadataProviderOptions{
			Logger: log.ProviderLogger(),
		}), nil
	}

	// env
	if credentialFilePath == "" && aliyuncliConfigFilePath == "" {
		if sessionName != "" {
			_ = os.Setenv(env.EnvRoleSessionName, sessionName)
		}
		if !ignoreEnv {
			log.Logger.Debug("try to get credentials from environment variables")
			cred, _ := env.NewCredentialsProvider(env.CredentialsProviderOptions{
				STSEndpoint: opt.stsEndpoint,
			})
			if cred != nil {
				_, err := cred.Credentials(context.TODO())
				if err == nil || !provider.IsNotEnableError(err) {
					log.Logger.Debugf("use credentials from environment variables")
					return cred, nil
				} else {
					log.Logger.Debugf("get credentials from environment variables failed: %+v, try another method",
						err)
				}
			}
			log.Logger.Debug("not found credentials from environment variables")
		}
	}

	// aliyun cli config
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

		acli, err := provider.NewCLIConfigProvider(provider.CLIConfigProviderOptions{
			ConfigPath:  aliyuncliConfigFilePath,
			ProfileName: aliyuncliProfileName,
			STSEndpoint: opt.stsEndpoint,
			Logger:      log.ProviderLogger(),
		})
		if err == nil && acli != nil {
			log.Logger.Debugf("try to get credentials from aliyun cli (%s) with profile name %s",
				utils.ShortHomePath(aliyuncliConfigFilePath), acli.ProfileName())
			return acli, nil
		}
		if err != nil {
			return nil, fmt.Errorf("get credentials from aliyun cli (%s) failed: %w",
				utils.ShortHomePath(aliyuncliConfigFilePath), err)
		}
	}

	// ini config
	if credentialFilePath == "" {
		path, err := checkFileExist(credentials.PATHCredentialFile)
		if err != nil && os.IsNotExist(err) {
			doc := "https://aliyuncontainerservice.github.io/ack-ram-tool/#credentials"
			return nil, fmt.Errorf("not found any configs for credentials, please configure credentials: %s", doc)
		}
		credentialFilePath = path
	}

	log.Logger.Debugf("try to get default credentials from %s", utils.ShortHomePath(credentialFilePath))
	if path, err := checkFileExist(credentialFilePath); err != nil {
		return nil, fmt.Errorf("read file %s: %w", credentialFilePath, err)
	} else {
		credentialFilePath = path
	}

	cred, err := provider.NewIniConfigProvider(provider.INIConfigProviderOptions{
		ConfigPath:  credentialFilePath,
		STSEndpoint: opt.stsEndpoint,
		Logger:      log.ProviderLogger(),
	})
	if err != nil {
		return nil, err
	}
	return cred, nil
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
			endpoints:               ctl.GlobalOption.GetEndpoints(),
			credentialType:          ctl.GlobalOption.GetCredentialType(),
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
