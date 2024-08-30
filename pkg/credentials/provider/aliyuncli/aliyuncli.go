package aliyuncli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
)

type profileWrapper struct {
	cp   Profile
	conf *Configuration

	stsEndpoint string
	client      *http.Client
	logger      provider.Logger
}

type CLIProvider struct {
	profile *profileWrapper
	logger  provider.Logger
}

func NewCLIProvider(configPath, profileName, stsEndpoint string, logger provider.Logger) (*CLIProvider, error) {
	if configPath == "" {
		configPath = getDefaultConfigPath()
	}
	conf, profile, err := loadProfile(configPath, profileName)
	if err != nil {
		return nil, fmt.Errorf("load profile: %w", err)
	}
	if err := profile.validate(); err != nil {
		return nil, fmt.Errorf("validate profile: %w", err)
	}
	logger.Debug(fmt.Sprintf("use profile name: %s", profile.Name))
	c := &CLIProvider{
		profile: &profileWrapper{
			cp:          profile,
			conf:        conf,
			stsEndpoint: stsEndpoint,
			client: &http.Client{
				Timeout: time.Second * 30,
			},
			logger: logger,
		},
		logger: logger,
	}
	return c, nil
}

func loadProfile(path string, name string) (*Configuration, Profile, error) {
	var p Profile
	conf, err := loadConfiguration(path)
	if err != nil {
		return nil, p, fmt.Errorf("init config: %w", err)
	}
	if name == "" {
		name = conf.CurrentProfile
	}
	p, ok := conf.getProfile(name)
	if !ok {
		return nil, p, fmt.Errorf("unknown profile %s", name)
	}
	return conf, p, nil
}

func (c *CLIProvider) Credentials(ctx context.Context) (*provider.Credentials, error) {
	p, err := c.profile.getProvider()
	if err != nil {
		return nil, err
	}
	return p.Credentials(ctx)
}

func (p *profileWrapper) getProvider() (provider.CredentialsProvider, error) {
	cp := p.cp

	switch cp.Mode {
	case AK:
		p.logger.Debug(fmt.Sprintf("using %s mode", cp.Mode))
		return p.getCredentialsByAK()
	case StsToken:
		p.logger.Debug(fmt.Sprintf("using %s mode", cp.Mode))
		return p.getCredentialsBySts()
	case RamRoleArn:
		p.logger.Debug(fmt.Sprintf("using %s mode", cp.Mode))
		return p.getCredentialsByRoleArn()
	case EcsRamRole:
		p.logger.Debug(fmt.Sprintf("using %s mode", cp.Mode))
		return p.getCredentialsByEcsRamRole()
	//case config.RsaKeyPair:
	//	return p.getCredentialsByPrivateKey()
	case RamRoleArnWithEcs:
		p.logger.Debug(fmt.Sprintf("using %s mode", cp.Mode))
		return p.getCredentialsByRamRoleArnWithEcs()
	case ChainableRamRoleArn:
		p.logger.Debug(fmt.Sprintf("using %s mode", cp.Mode))
		return p.getCredentialsByChainableRamRoleArn()
	case External:
		p.logger.Debug(fmt.Sprintf("using %s mode", cp.Mode))
		return p.getCredentialsByExternal()
	case CredentialsURI:
		p.logger.Debug(fmt.Sprintf("using %s mode", cp.Mode))
		return p.getCredentialsByCredentialsURI()
	default:
		return nil, fmt.Errorf("unexcepted certificate mode: %s", cp.Mode)
	}
}

func (p *profileWrapper) getCredentialsByAK() (provider.CredentialsProvider, error) {
	cp := p.cp

	return provider.NewAccessKeyProvider(cp.AccessKeyId, cp.AccessKeySecret), nil
}

func (p *profileWrapper) getCredentialsBySts() (provider.CredentialsProvider, error) {
	cp := p.cp

	return provider.NewSTSTokenProvider(cp.AccessKeyId, cp.AccessKeySecret, cp.StsToken), nil
}

func (p *profileWrapper) getCredentialsByRoleArn() (provider.CredentialsProvider, error) {
	cp := p.cp

	preP := provider.NewAccessKeyProvider(cp.AccessKeyId, cp.AccessKeySecret)

	return p.getCredentialsByRoleArnWithPro(preP)
}

func (p *profileWrapper) getCredentialsByRoleArnWithPro(preP provider.CredentialsProvider) (provider.CredentialsProvider, error) {
	cp := p.cp

	credP := provider.NewRoleArnProvider(preP, cp.RamRoleArn, provider.RoleArnProviderOptions{
		STSEndpoint: p.stsEndpoint,
		SessionName: cp.RoleSessionName,
		Logger:      p.logger,
	})
	return credP, nil
}

func (p *profileWrapper) getCredentialsByEcsRamRole() (provider.CredentialsProvider, error) {
	cp := p.cp

	credP := provider.NewECSMetadataProvider(provider.ECSMetadataProviderOptions{
		RoleName: cp.RamRoleName,
		Logger:   p.logger,
	})
	return credP, nil
}

//func (p *profileWrapper) getCredentialsByPrivateKey() (credentials.Credential, error) {
//
//}

func (p *profileWrapper) getCredentialsByRamRoleArnWithEcs() (provider.CredentialsProvider, error) {
	preP, err := p.getCredentialsByEcsRamRole()
	if err != nil {
		return nil, err
	}
	return p.getCredentialsByRoleArnWithPro(preP)
}

func (p *profileWrapper) getCredentialsByChainableRamRoleArn() (provider.CredentialsProvider, error) {
	cp := p.cp
	profileName := cp.SourceProfile

	p.logger.Debug(fmt.Sprintf("get credentials from source profile %s", profileName))
	source, loaded := p.conf.getProfile(profileName)
	if !loaded {
		return nil, fmt.Errorf("can not load the source profile: " + profileName)
	}
	newP := &profileWrapper{
		cp:          source,
		conf:        p.conf,
		stsEndpoint: p.stsEndpoint,
		client:      p.client,
	}
	preP, err := newP.getProvider()
	if err != nil {
		return nil, err
	}

	p.logger.Debug(fmt.Sprintf("using role arn by current profile %s", cp.Name))
	return p.getCredentialsByRoleArnWithPro(preP)
}

func (p *profileWrapper) getCredentialsByExternal() (provider.CredentialsProvider, error) {
	cp := p.cp
	args := strings.Fields(cp.ProcessCommand)
	cmd := exec.Command(args[0], args[1:]...) // #nosec G204
	p.logger.Debug(fmt.Sprintf("running external program: %s", cp.ProcessCommand))

	genmsg := func(buf []byte, err error) string {
		message := fmt.Sprintf(`run external program to get credentials faild:
  command: %s
  output: %s
  error: %s`,
			cp.ProcessCommand, string(buf), err.Error())
		return message
	}

	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	allOutput := stderr.String() + "\n" + stdout.String()
	if err != nil {
		message := genmsg([]byte(allOutput), err)
		return nil, errors.New(message)
	}

	buf := stdout.Bytes()
	var newCP Profile
	err = json.Unmarshal(buf, &newCP)
	if err != nil {
		if pp := tryToParseProfileFromOutput(string(buf)); pp != nil {
			newCP = *pp
		} else {
			message := genmsg([]byte(allOutput), err)
			return nil, errors.New(message)
		}
	}

	p.logger.Debug(fmt.Sprintf("using profile from output of external program"))
	newP := &profileWrapper{
		cp:   newCP,
		conf: p.conf,
	}
	return newP.getProvider()
}

var regexpCredJSON = regexp.MustCompile(`{[^}]+"mode":[^}]+}`)

func tryToParseProfileFromOutput(output string) *Profile {
	ret := regexpCredJSON.FindAllString(output, 1)
	if len(ret) < 1 {
		return nil
	}
	credJSON := ret[0]
	var p Profile
	if err := json.Unmarshal([]byte(credJSON), &p); err == nil && p.Mode != "" {
		return &p
	}
	return nil
}

func (p *profileWrapper) getCredentialsByCredentialsURI() (provider.CredentialsProvider, error) {
	cp := p.cp
	uri := cp.CredentialsURI
	if uri == "" {
		uri = os.Getenv(provider.EnvCredentialsURI)
	}
	if uri == "" {
		return nil, fmt.Errorf("invalid credentials uri")
	}
	p.logger.Debug(fmt.Sprintf("get credentials from uri %s", uri))

	newPr := provider.NewURIProvider(uri, provider.URIProviderOptions{})
	return newPr, nil
}

func (c *CLIProvider) ProfileName() string {
	return c.profile.cp.Name
}
