package aliyuncli

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-cli/v3/cli"
	"github.com/aliyun/aliyun-cli/v3/config"
	"os"
	"time"
)

func GetCredentialsProvider(p provider.Profile) (provider.CredentialsProvider, error) {
	log.Logger.Debugf("try to get credential from aliyun cli profile %q", p.Name)
	newP := newProfile(p)
	if err := newP.Validate(); err != nil {
		return nil, fmt.Errorf("validate aliyun cli profile %q failed: %w", p.Name, err)
	}

	ctx := cli.NewCommandContext(os.Stdout, os.Stderr)

	cp := provider.NewFunctionProvider(func(_ context.Context) (*provider.Credentials, error) {
		log.Logger.Debugf("try to get credential from aliyun cli profile %q with fallback mode", p.Name)
		ret, err := newP.GetCredential(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("can not get credential from aliyun cli profile %q: %w", p.Name, err)
		}
		cred, err := ret.GetCredential()
		if err != nil {
			return nil, fmt.Errorf("can not get credential from aliyun cli profile %q: %w", p.Name, err)
		}

		return &provider.Credentials{
			AccessKeyId:     tea.StringValue(cred.AccessKeyId),
			AccessKeySecret: tea.StringValue(cred.AccessKeySecret),
			SecurityToken:   tea.StringValue(cred.SecurityToken),
			Expiration:      time.Time{},
		}, nil
	})

	return cp, nil
}

func newProfile(p provider.Profile) config.Profile {
	return config.Profile{
		Name:                      p.Name,
		Mode:                      config.AuthenticateMode(string(p.Mode)),
		AccessKeyId:               p.AccessKeyId,
		AccessKeySecret:           p.AccessKeySecret,
		StsToken:                  p.StsToken,
		StsRegion:                 p.StsRegion,
		RamRoleName:               p.RamRoleName,
		RamRoleArn:                p.RamRoleArn,
		RoleSessionName:           p.RoleSessionName,
		ExternalId:                p.ExternalId,
		SourceProfile:             p.SourceProfile,
		PrivateKey:                p.PrivateKey,
		KeyPairName:               p.KeyPairName,
		ExpiredSeconds:            p.ExpiredSeconds,
		Verified:                  p.Verified,
		RegionId:                  p.RegionId,
		OutputFormat:              p.OutputFormat,
		Language:                  p.Language,
		Site:                      p.Site,
		ReadTimeout:               p.ReadTimeout,
		ConnectTimeout:            p.ConnectTimeout,
		RetryCount:                p.RetryCount,
		ProcessCommand:            p.ProcessCommand,
		CredentialsURI:            p.CredentialsURI,
		OIDCProviderARN:           p.OIDCProviderARN,
		OIDCTokenFile:             p.OIDCTokenFile,
		CloudSSOSignInUrl:         p.CloudSSOSignInUrl,
		AccessToken:               p.AccessToken,
		CloudSSOAccessTokenExpire: p.CloudSSOAccessTokenExpire,
		StsExpiration:             p.StsExpiration,
		CloudSSOAccessConfig:      p.CloudSSOAccessConfig,
		CloudSSOAccountId:         p.CloudSSOAccountId,
	}
}
