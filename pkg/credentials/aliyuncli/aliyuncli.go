package aliyuncli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper/env"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

type ProfileWrapper struct {
	cp   Profile
	conf *Configuration
}

type CredentialHelper struct {
	profile *ProfileWrapper
}

func NewCredentialHelper(configPath, profileName string) (*CredentialHelper, error) {
	if configPath == "" {
		configPath = getDefaultConfigPath()
	}
	conf, profile, err := LoadProfile(configPath, profileName)
	if err != nil {
		return nil, err
	}
	if err := profile.Validate(); err != nil {
		return nil, err
	}
	c := &CredentialHelper{
		profile: &ProfileWrapper{
			cp:   profile,
			conf: conf,
		},
	}
	return c, nil
}

func LoadProfile(path string, name string) (*Configuration, Profile, error) {
	var p Profile
	conf, err := LoadConfiguration(path)
	if err != nil {
		return nil, p, fmt.Errorf("init config failed %v", err)
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

func (c CredentialHelper) GetCredentials() (credentials.Credential, error) {
	return c.profile.GetCredentials()
}

func (p *ProfileWrapper) GetCredentials() (credentials.Credential, error) {
	cp := p.cp

	switch cp.Mode {
	case AK:
		return p.GetCredentialsByAK()
	case StsToken:
		return p.GetCredentialsBySts()
	case RamRoleArn:
		return p.GetCredentialsByRoleArn()
	case EcsRamRole:
		return p.GetCredentialsByEcsRamRole()
	//case config.RsaKeyPair:
	//	return p.GetCredentialsByPrivateKey()
	case RamRoleArnWithEcs:
		return p.GetCredentialsByRamRoleArnWithEcs()
	case ChainableRamRoleArn:
		return p.GetCredentialsByChainableRamRoleArn()
	case External:
		return p.GetCredentialsByExternal()
	case CredentialsURI:
		return p.GetCredentialsByCredentialsURI()
	default:
		return nil, fmt.Errorf("unexcepted certificate mode: %s", cp.Mode)
	}
}

func (p *ProfileWrapper) GetCredentialsByAK() (credentials.Credential, error) {
	cp := p.cp
	conf := &credentials.Config{
		Type:            tea.String("access_key"),
		AccessKeyId:     tea.String(cp.AccessKeyId),
		AccessKeySecret: tea.String(cp.AccessKeySecret),
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

func (p *ProfileWrapper) GetCredentialsBySts() (credentials.Credential, error) {
	cp := p.cp
	conf := &credentials.Config{
		Type:            tea.String("sts"),
		AccessKeyId:     tea.String(cp.AccessKeyId),
		AccessKeySecret: tea.String(cp.AccessKeySecret),
		SecurityToken:   tea.String(cp.StsToken),
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

func (p *ProfileWrapper) GetCredentialsByRoleArn() (credentials.Credential, error) {
	cp := p.cp
	conf := &credentials.Config{
		Type:            tea.String("ram_role_arn"),
		AccessKeyId:     tea.String(cp.AccessKeyId),
		AccessKeySecret: tea.String(cp.AccessKeySecret),
		RoleArn:         tea.String(cp.RamRoleArn),
		RoleSessionName: tea.String(cp.RoleSessionName),
	}
	if cp.ExpiredSeconds > 0 {
		conf.RoleSessionExpiration = tea.Int(cp.ExpiredSeconds)
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

func (p *ProfileWrapper) GetCredentialsByEcsRamRole() (credentials.Credential, error) {
	cp := p.cp
	conf := &credentials.Config{
		Type:     tea.String("ecs_ram_role"),
		RoleName: tea.String(cp.RamRoleName),
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

//func (p *ProfileWrapper) GetCredentialsByPrivateKey() (credentials.Credential, error) {
//
//}

func (p *ProfileWrapper) GetCredentialsByRamRoleArnWithEcs() (credentials.Credential, error) {
	cp := p.cp
	client, err := cp.GetSTSClientByEcsRamRole()
	if err != nil {
		return nil, err
	}
	resp, err := cp.GetSessionCredential(client)
	if err != nil {
		return nil, err
	}
	conf := &credentials.Config{
		Type:            tea.String("sts"),
		AccessKeyId:     resp.AccessKeyId,
		AccessKeySecret: resp.AccessKeySecret,
		SecurityToken:   resp.SecurityToken,
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

func (p *ProfileWrapper) GetCredentialsByChainableRamRoleArn() (credentials.Credential, error) {
	cp := p.cp
	profileName := cp.SourceProfile

	// 从 configuration 中重新获取 source profile
	source, loaded := p.conf.GetProfile(profileName)
	if !loaded {
		return nil, fmt.Errorf("can not load the source profile: " + profileName)
	}

	client, err := source.GetSTSClientByEcsRamRole()
	if err != nil {
		return nil, err
	}
	resp, err := cp.GetSessionCredential(client)
	if err != nil {
		return nil, err
	}
	conf := &credentials.Config{
		Type:            tea.String("sts"),
		AccessKeyId:     resp.AccessKeyId,
		AccessKeySecret: resp.AccessKeySecret,
		SecurityToken:   resp.SecurityToken,
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

func (p *ProfileWrapper) GetCredentialsByExternal() (credentials.Credential, error) {
	cp := p.cp
	args := strings.Fields(cp.ProcessCommand)
	cmd := exec.Command(args[0], args[1:]...) // #nosec G204
	buf, err := cmd.CombinedOutput()
	genmsg := func(err error) string {
		message := fmt.Sprintf(`run external program to get credentials faild:
  command: %s
  output: %s
  error: %s`,
			cp.ProcessCommand, string(buf), err.Error())
		return message
	}
	if err != nil {
		message := genmsg(err)
		return nil, errors.New(message)
	}
	var newCP Profile
	err = json.Unmarshal(buf, &newCP)
	if err != nil {
		message := genmsg(err)
		return nil, errors.New(message)
	}

	newP := &ProfileWrapper{
		cp:   newCP,
		conf: p.conf,
	}
	return newP.GetCredentials()
}

func (p *ProfileWrapper) GetCredentialsByCredentialsURI() (credentials.Credential, error) {
	cp := p.cp
	uri := cp.CredentialsURI
	if uri == "" {
		uri = os.Getenv(env.EnvCredentialsURI)
	}
	if uri == "" {
		return nil, fmt.Errorf("invalid credentials uri")
	}

	res, err := http.Get(uri) // #nosec G107
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
		return nil, fmt.Errorf("Get sts token err, Code is not Success")
	}
	conf := &credentials.Config{
		Type:            tea.String("sts"),
		AccessKeyId:     tea.String(response.AccessKeyId),
		AccessKeySecret: tea.String(response.AccessKeySecret),
		SecurityToken:   tea.String(response.SecurityToken),
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}
