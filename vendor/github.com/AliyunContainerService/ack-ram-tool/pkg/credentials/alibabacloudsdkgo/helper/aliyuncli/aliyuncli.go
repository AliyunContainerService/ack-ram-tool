package aliyuncli

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/credentials-go/credentials"
	"io/ioutil"
	"os"
	"path/filepath"
)

// parse aliyun cli configuration

var (
	errUnsupportedAuthenticateMode = errors.New("unsupported mode")
	errProfileNameNotFound         = errors.New("not found such profile name")
)

const (
	defaultProfileFile = "~/.aliyun/config.json"
	defaultProfile
)

type Configuration struct {
	CurrentProfile string    `json:"current"`
	Profiles       []Profile `json:"profiles"`
	MetaPath       string    `json:"meta_path"`
	//Plugins 		[]Plugin `json:"plugin"`
}

type AuthenticateMode string

const (
	AK                  = AuthenticateMode("AK")
	StsToken            = AuthenticateMode("StsToken")
	RamRoleArn          = AuthenticateMode("RamRoleArn")
	EcsRamRole          = AuthenticateMode("EcsRamRole")
	RsaKeyPair          = AuthenticateMode("RsaKeyPair")
	RamRoleArnWithEcs   = AuthenticateMode("RamRoleArnWithRoleName")
	ChainableRamRoleArn = AuthenticateMode("ChainableRamRoleArn")
	External            = AuthenticateMode("External")
	CredentialsURI      = AuthenticateMode("CredentialsURI")
)

type Profile struct {
	Name            string           `json:"name"`
	Mode            AuthenticateMode `json:"mode"`
	AccessKeyId     string           `json:"access_key_id"`
	AccessKeySecret string           `json:"access_key_secret"`
	StsToken        string           `json:"sts_token"`
	StsRegion       string           `json:"sts_region"`
	RamRoleName     string           `json:"ram_role_name"`
	RamRoleArn      string           `json:"ram_role_arn"`
	RoleSessionName string           `json:"ram_session_name"`
	SourceProfile   string           `json:"source_profile"`
	PrivateKey      string           `json:"private_key"`
	KeyPairName     string           `json:"key_pair_name"`
	ExpiredSeconds  int              `json:"expired_seconds"`
	Verified        string           `json:"verified"`
	RegionId        string           `json:"region_id"`
	OutputFormat    string           `json:"output_format"`
	Language        string           `json:"language"`
	Site            string           `json:"site"`
	ReadTimeout     int              `json:"retry_timeout"`
	ConnectTimeout  int              `json:"connect_timeout"`
	RetryCount      int              `json:"retry_count"`
	ProcessCommand  string           `json:"process_command"`
	CredentialsURI  string           `json:"credentials_uri"`
	parent          *Configuration   //`json:"-"`
}

// NewCredential return Credential base on config file from aliyun cli
// when configFilePath is empty, it'll use ~/.aliyun/config.json
// when profileName is empty, it'll use value of 'current' field from configFilePath
func NewCredential(configFilePath, profileName string) (credentials.Credential, error) {
	c, err := LoadConfiguration(configFilePath)
	if err != nil {
		return nil, err
	}
	if profileName == "" {
		profileName = c.CurrentProfile
	}
	p, err := c.GetProfile(profileName)
	if err != nil {
		return nil, err
	}
	return p.Credential()
}

func LoadConfiguration(path string) (*Configuration, error) {
	if path == "" {
		path = defaultProfileFile
	}
	path, err := expandPath(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Configuration
	if err := json.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("parse configuration from %s: %w", path, err)
	}
	return &conf, nil
}

func (c *Configuration) GetProfile(name string) (*Profile, error) {
	for _, p := range c.Profiles {
		p := p
		if p.Name == name {
			return &p, nil
		}
	}
	return nil, errProfileNameNotFound
}

func (p *Profile) Credential() (credentials.Credential, error) {
	config := &credentials.Config{
		AccessKeyId:     stringPoint(p.AccessKeyId),
		AccessKeySecret: stringPoint(p.AccessKeySecret),
		SecurityToken:   stringPoint(p.StsToken),
		Url:             stringPoint(p.CredentialsURI),
		RoleName:        stringPoint(p.RamRoleName),
		RoleArn:         stringPoint(p.RamRoleArn),
		RoleSessionName: stringPoint(p.RoleSessionName),
		PrivateKeyFile:  stringPoint(p.PrivateKey),
		PublicKeyId:     stringPoint(p.KeyPairName),
	}
	switch p.Mode {
	case AK:
		config.Type = stringPoint("access_key")
		break
	case StsToken:
		config.Type = stringPoint("sts")
		break
	case CredentialsURI:
		config.Type = stringPoint("credentials_uri")
		break
	case EcsRamRole:
		config.Type = stringPoint("ecs_ram_role")
		break
	case RamRoleArn:
		config.Type = stringPoint("ram_role_arn")
		if config.RoleSessionName == nil || *config.RoleSessionName == "" {
			config.RoleSessionName = stringPoint("session-name")
		}
		break
	//case RsaKeyPair:
	//	config.Type = stringPoint("rsa_key_pair")
	default:
		return nil, errUnsupportedAuthenticateMode
	}

	return credentials.NewCredential(config)
}

func stringPoint(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

func expandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[1:])
	}
	return path, nil
}
