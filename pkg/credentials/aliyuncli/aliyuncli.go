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
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/aliyun-cli/cli"
	"github.com/aliyun/aliyun-cli/config"
	"github.com/aliyun/credentials-go/credentials"
)

const configFile = "config.json"

type Profile struct {
	cp   config.Profile
	conf *config.Configuration
}

type CredentialHelper struct {
	profile *Profile
	cred    credentials.Credential
}

func NewCredentialHelper(configPath, profileName string) (*CredentialHelper, error) {
	//profile, err := config.LoadCurrentProfile()
	//clictx := cli.NewCommandContext(bytes.NewBuffer(nil), bytes.NewBuffer(nil))
	//profile := cfg.GetCurrentProfile(clictx)
	if configPath == "" {
		configPath = config.GetConfigPath() + "/" + configFile
	}
	conf, profile, err := LoadProfile(configPath, profileName)
	if err != nil {
		return nil, err
	}
	if err := profile.Validate(); err != nil {
		return nil, err
	}
	c := &CredentialHelper{
		profile: &Profile{
			cp:   profile,
			conf: conf,
		},
	}
	return c, nil
}

func LoadProfile(path string, name string) (*config.Configuration, config.Profile, error) {
	var p config.Profile
	conf, err := config.LoadConfiguration(path)
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

func (p *Profile) GetCredentials() (credentials.Credential, error) {
	cp := p.cp

	switch cp.Mode {
	case config.AK:
		return p.GetCredentialsByAK()
	case config.StsToken:
		return p.GetCredentialsBySts()
	case config.RamRoleArn:
		return p.GetCredentialsByRoleArn()
	case config.EcsRamRole:
		return p.GetCredentialsByEcsRamRole()
	//case config.RsaKeyPair:
	//	return p.GetCredentialsByPrivateKey()
	case config.RamRoleArnWithEcs:
		return p.GetCredentialsByRamRoleArnWithEcs()
	case config.ChainableRamRoleArn:
		return p.GetCredentialsByChainableRamRoleArn()
	case config.External:
		return p.GetCredentialsByExternal()
	case config.CredentialsURI:
		return p.GetCredentialsByCredentialsURI()
	default:
		return nil, fmt.Errorf("unexcepted certificate mode: %s", cp.Mode)
	}
}

func (p *Profile) GetCredentialsByAK() (credentials.Credential, error) {
	cp := p.cp
	conf := &credentials.Config{
		Type:            tea.String("access_key"),
		AccessKeyId:     tea.String(cp.AccessKeyId),
		AccessKeySecret: tea.String(cp.AccessKeySecret),
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

func (p *Profile) GetCredentialsBySts() (credentials.Credential, error) {
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

func (p *Profile) GetCredentialsByRoleArn() (credentials.Credential, error) {
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

func (p *Profile) GetCredentialsByEcsRamRole() (credentials.Credential, error) {
	cp := p.cp
	conf := &credentials.Config{
		Type:     tea.String("ecs_ram_role"),
		RoleName: tea.String(cp.RamRoleName),
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

//func (p *Profile) GetCredentialsByPrivateKey() (credentials.Credential, error) {
//
//}

func (p *Profile) GetCredentialsByRamRoleArnWithEcs() (credentials.Credential, error) {
	cp := p.cp
	client, err := cp.GetClientByEcsRamRole(sdk.NewConfig())
	if err != nil {
		return nil, err
	}
	accessKeyID, accessKeySecret, StsToken, err := cp.GetSessionCredential(client)
	if err != nil {
		return nil, err
	}
	conf := &credentials.Config{
		Type:            tea.String("sts"),
		AccessKeyId:     tea.String(accessKeyID),
		AccessKeySecret: tea.String(accessKeySecret),
		SecurityToken:   tea.String(StsToken),
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

func (p *Profile) GetCredentialsByChainableRamRoleArn() (credentials.Credential, error) {
	cp := p.cp
	profileName := cp.SourceProfile

	// 从 configuration 中重新获取 source profile
	source, loaded := p.conf.GetProfile(profileName)
	if !loaded {
		return nil, fmt.Errorf("can not load the source profile: " + profileName)
	}

	client, err := source.GetClient(cli.NewCommandContext(bytes.NewBuffer(nil), bytes.NewBuffer(nil)))
	if err != nil {
		return nil, err
	}
	accessKeyID, accessKeySecret, StsToken, err := cp.GetSessionCredential(client)
	if err != nil {
		return nil, err
	}
	conf := &credentials.Config{
		Type:            tea.String("sts"),
		AccessKeyId:     tea.String(accessKeyID),
		AccessKeySecret: tea.String(accessKeySecret),
		SecurityToken:   tea.String(StsToken),
	}
	cred, err := credentials.NewCredential(conf)
	return cred, err
}

func (p *Profile) GetCredentialsByExternal() (credentials.Credential, error) {
	cp := p.cp
	args := strings.Fields(cp.ProcessCommand)
	cmd := exec.Command(args[0], args[1:]...)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	var newCP config.Profile
	err = json.Unmarshal(buf, &newCP)
	if err != nil {
		message := fmt.Sprintf("%s\n%s\n%s", cp.ProcessCommand, string(buf), err.Error())
		return nil, errors.New(message)
	}

	newP := &Profile{
		cp:   newCP,
		conf: p.conf,
	}
	return newP.GetCredentials()
}

func (p *Profile) GetCredentialsByCredentialsURI() (credentials.Credential, error) {
	cp := p.cp
	uri := cp.CredentialsURI
	if uri == "" {
		uri = os.Getenv("ALIBABA_CLOUD_CREDENTIALS_URI")
	}
	if uri == "" {
		return nil, fmt.Errorf("invalid credentials uri")
	}

	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

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
