package aliyuncli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudgo/env"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
)

type ProfileWrapper struct {
	cp   Profile
	conf *Configuration

	stsEndpoint string
	client      *http.Client
}

type CredentialHelper struct {
	profile *ProfileWrapper
}

func NewCredentialHelper(configPath, profileName, stsEndpoint string) (*CredentialHelper, error) {
	if configPath == "" {
		configPath = getDefaultConfigPath()
	}
	conf, profile, err := LoadProfile(configPath, profileName)
	if err != nil {
		return nil, fmt.Errorf("load profile: %w", err)
	}
	if err := profile.Validate(); err != nil {
		return nil, fmt.Errorf("validate profile: %w", err)
	}
	log.Logger.Debugf("use profile name: %s", profile.Name)
	c := &CredentialHelper{
		profile: &ProfileWrapper{
			cp:          profile,
			conf:        conf,
			stsEndpoint: stsEndpoint,
			client: &http.Client{
				Timeout: time.Second * 30,
			},
		},
	}
	return c, nil
}

func LoadProfile(path string, name string) (*Configuration, Profile, error) {
	var p Profile
	conf, err := LoadConfiguration(path)
	if err != nil {
		return nil, p, fmt.Errorf("init config: %w", err)
	}
	if name == "" {
		name = conf.CurrentProfile
	}
	p, ok := conf.GetProfile(name)
	if !ok {
		return nil, p, fmt.Errorf("unknown profile %s", name)
	}
	return conf, p, nil
}

func (c CredentialHelper) GetCredentials() (provider.CredentialsProvider, error) {
	return c.profile.GetProvider()
}

func (p *ProfileWrapper) GetProvider() (provider.CredentialsProvider, error) {
	cp := p.cp

	switch cp.Mode {
	case AK:
		log.Logger.Debugf("using %s mode", cp.Mode)
		return p.GetCredentialsByAK()
	case StsToken:
		log.Logger.Debugf("using %s mode", cp.Mode)
		return p.GetCredentialsBySts()
	case RamRoleArn:
		log.Logger.Debugf("using %s mode", cp.Mode)
		return p.GetCredentialsByRoleArn()
	case EcsRamRole:
		log.Logger.Debugf("using %s mode", cp.Mode)
		return p.GetCredentialsByEcsRamRole()
	//case config.RsaKeyPair:
	//	return p.GetCredentialsByPrivateKey()
	case RamRoleArnWithEcs:
		log.Logger.Debugf("using %s mode", cp.Mode)
		return p.GetCredentialsByRamRoleArnWithEcs()
	case ChainableRamRoleArn:
		log.Logger.Debugf("using %s mode", cp.Mode)
		return p.GetCredentialsByChainableRamRoleArn()
	case External:
		log.Logger.Debugf("using %s mode", cp.Mode)
		return p.GetCredentialsByExternal()
	case CredentialsURI:
		log.Logger.Debugf("using %s mode", cp.Mode)
		return p.GetCredentialsByCredentialsURI()
	default:
		return nil, fmt.Errorf("unexcepted certificate mode: %s", cp.Mode)
	}
}

func (p *ProfileWrapper) GetCredentialsByAK() (provider.CredentialsProvider, error) {
	cp := p.cp

	return provider.NewAccessKeyProvider(cp.AccessKeyId, cp.AccessKeySecret), nil
}

func (p *ProfileWrapper) GetCredentialsBySts() (provider.CredentialsProvider, error) {
	cp := p.cp

	return provider.NewSTSTokenProvider(cp.AccessKeyId, cp.AccessKeySecret, cp.StsToken), nil
}

func (p *ProfileWrapper) GetCredentialsByRoleArn() (provider.CredentialsProvider, error) {
	cp := p.cp

	preP := provider.NewAccessKeyProvider(cp.AccessKeyId, cp.AccessKeySecret)

	return p.GetCredentialsByRoleArnWithPro(preP)
}

func (p *ProfileWrapper) GetCredentialsByRoleArnWithPro(preP provider.CredentialsProvider) (provider.CredentialsProvider, error) {
	cp := p.cp

	credP := provider.NewRoleArnProvider(preP, cp.RamRoleArn, provider.RoleArnProviderOptions{
		STSEndpoint: p.stsEndpoint,
		SessionName: cp.RoleSessionName,
		Logger:      &log.ProviderLogWrapper{ZP: log.Logger},
	})
	return credP, nil
}

func (p *ProfileWrapper) GetCredentialsByEcsRamRole() (provider.CredentialsProvider, error) {
	cp := p.cp

	credP := provider.NewECSMetadataProvider(provider.ECSMetadataProviderOptions{
		RoleName: cp.RamRoleName,
		Logger:   &log.ProviderLogWrapper{ZP: log.Logger},
	})
	return credP, nil
}

//func (p *ProfileWrapper) GetCredentialsByPrivateKey() (credentials.Credential, error) {
//
//}

func (p *ProfileWrapper) GetCredentialsByRamRoleArnWithEcs() (provider.CredentialsProvider, error) {
	preP, err := p.GetCredentialsByEcsRamRole()
	if err != nil {
		return nil, err
	}
	return p.GetCredentialsByRoleArnWithPro(preP)
}

func (p *ProfileWrapper) GetCredentialsByChainableRamRoleArn() (provider.CredentialsProvider, error) {
	cp := p.cp
	profileName := cp.SourceProfile

	log.Logger.Debugf("get credentials from source profile %s", profileName)
	source, loaded := p.conf.GetProfile(profileName)
	if !loaded {
		return nil, fmt.Errorf("can not load the source profile: %s", profileName)
	}
	newP := &ProfileWrapper{
		cp:          source,
		conf:        p.conf,
		stsEndpoint: p.stsEndpoint,
		client:      p.client,
	}
	preP, err := newP.GetProvider()
	if err != nil {
		return nil, err
	}

	log.Logger.Debugf("using role arn by current profile %s", cp.Name)
	return p.GetCredentialsByRoleArnWithPro(preP)
}

func (p *ProfileWrapper) GetCredentialsByExternal() (provider.CredentialsProvider, error) {
	cp := p.cp
	args := strings.Fields(cp.ProcessCommand)
	cmd := exec.Command(args[0], args[1:]...) // #nosec G204
	log.Logger.Debugf("running external program: %s", cp.ProcessCommand)

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

	log.Logger.Debug("using profile from output of external program")
	newP := &ProfileWrapper{
		cp:   newCP,
		conf: p.conf,
	}
	return newP.GetProvider()
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

func (p *ProfileWrapper) GetCredentialsByCredentialsURI() (provider.CredentialsProvider, error) {
	cp := p.cp
	uri := cp.CredentialsURI
	if uri == "" {
		uri = os.Getenv(env.EnvCredentialsURI)
	}
	if uri == "" {
		return nil, fmt.Errorf("invalid credentials uri")
	}
	log.Logger.Debugf("get credentials from %s", uri)

	res, err := p.client.Get(uri) // #nosec G107
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // #nosec G307

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Get Credentials from %s failed, status code %d", uri, res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Code            string
		AccessKeyId     string
		AccessKeySecret string
		SecurityToken   string
		Expiration      string
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal credentials failed, the body %s", string(body))
	}

	if response.Code != "Success" {
		return nil, fmt.Errorf("get sts token err, Code is not Success")
	}

	return provider.NewSTSTokenProvider(
		response.AccessKeyId, response.AccessKeySecret,
		response.SecurityToken), nil
}

func (c CredentialHelper) ProfileName() string {
	return c.profile.cp.Name
}
