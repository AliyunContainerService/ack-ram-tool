package aliyuncli

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-cli/v3/cli"
	"github.com/aliyun/aliyun-cli/v3/config"
)

func GetCredentialsProvider(p provider.Profile) (provider.CredentialsProvider, error) {
	log.Logger.Debugf("try to get credential from aliyun cli profile %q", p.Name)
	newP := newProfile(p)
	if err := newP.Validate(); err != nil {
		return nil, fmt.Errorf("validate aliyun cli profile %q failed: %w", p.Name, err)
	}

	ctx := cli.NewCommandContext(os.Stdout, os.Stderr)

	cp := provider.NewFunctionProvider(func(c context.Context) (*provider.Credentials, error) {
		log.Logger.Debugf("try to get credential from aliyun cli profile %q with fallback mode", p.Name)
		ret, err := newP.GetCredential(ctx, nil)
		if err != nil {
			log.Logger.Debugf("get credential from aliyun cli profile %q failed: %s", p.Name, err)
			if newP.Mode == config.OAuth {
				if authErr := runOAuthFlow(c, newP); authErr != nil {
					log.Logger.Errorf("run OAuth Flow failed: %s", authErr)
					err = authErr
				} else {
					if cred, rerr := getCredentialsViaReload(c, os.Args[0], newP); rerr != nil {
						//
					} else {
						return cred, nil
					}
				}
			}
		}
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
			Expiration:      time.Now().Add(time.Minute * 5),
		}, nil
	})

	return cp, nil
}

func getCredentialsViaReload(ctx context.Context,
	bin string, profile config.Profile) (*provider.Credentials, error) {
	cmd := []string{
		bin,
		"export-credentials",
		"--profile-name",
		profile.Name,
		"-f",
		"aliyun-cli-config-json",
	}
	newP := provider.Profile{
		Name: "dummy", Mode: provider.External,
		ProcessCommand: strings.Join(cmd, " ")}
	cliProfile, err := provider.NewCLIConfigProvider(provider.CLIConfigProviderOptions{
		ConfigPath:                "",
		ProfileName:               newP.Name,
		STSEndpoint:               "",
		Logger:                    nil,
		GetProviderForUnknownMode: nil,
		Config: &provider.Configuration{
			CurrentProfile: newP.Name,
			Profiles:       []provider.Profile{newP},
			MetaPath:       "",
		},
	})
	if err != nil {
		return nil, err
	}
	return cliProfile.Credentials(ctx)
}

func runOAuthFlow(ctx context.Context, profile config.Profile) error {
	cmdArgs := buildOAuthCmd(profile)
	log.Logger.Infof("run OAuth Flow via %s", strings.Join(cmdArgs, " "))

	cmd := exec.CommandContext(ctx, cmdArgs[0], cmdArgs[1:]...) // #nosec G204
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func buildOAuthCmd(profile config.Profile) []string {
	var cmd []string
	cmd = append(cmd, "aliyun", "configure", "--mode", "OAuth")
	if profile.OAuthSiteType != "" {
		cmd = append(cmd, "--oauth-site-type", profile.OAuthSiteType)
	}
	if profile.Name != "" && profile.Name != "default" {
		cmd = append(cmd, "--profile", profile.Name)
	}
	return cmd
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
		OAuthAccessToken:          p.OAuthAccessToken,
		OAuthRefreshToken:         p.OAuthRefreshToken,
		OAuthAccessTokenExpire:    p.OAuthAccessTokenExpire,
		OAuthRefreshTokenExpire:   p.OAuthRefreshTokenExpire,
		OAuthSiteType:             p.OAuthSiteType,
		EndpointType:              p.EndpointType,
	}
}
