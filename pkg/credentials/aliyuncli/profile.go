// Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aliyuncli

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"regexp"

	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

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

func NewProfile(name string) Profile {
	return Profile{
		Name:         name,
		Mode:         "",
		OutputFormat: "json",
		Language:     "en",
	}
}

func (cp *Profile) Validate() error {
	if cp.RegionId == "" {
		return fmt.Errorf("region can't be empty")
	}

	if !IsRegion(cp.RegionId) {
		return fmt.Errorf("invalid region %s", cp.RegionId)
	}

	if cp.Mode == "" {
		return fmt.Errorf("profile %s is not configure yet, run `aliyun configure --profile %s` first", cp.Name, cp.Name)
	}

	switch cp.Mode {
	case AK:
		return cp.ValidateAK()
	case StsToken:
		err := cp.ValidateAK()
		if err != nil {
			return err
		}
		if cp.StsToken == "" {
			return fmt.Errorf("invalid sts_token")
		}
	case RamRoleArn:
		err := cp.ValidateAK()
		if err != nil {
			return err
		}
		if cp.RamRoleArn == "" {
			return fmt.Errorf("invalid ram_role_arn")
		}
		if cp.RoleSessionName == "" {
			return fmt.Errorf("invalid role_session_name")
		}
	case EcsRamRole, RamRoleArnWithEcs:
	case RsaKeyPair:
		if cp.PrivateKey == "" {
			return fmt.Errorf("invalid private_key")
		}
		if cp.KeyPairName == "" {
			return fmt.Errorf("invalid key_pair_name")
		}
	case External:
		if cp.ProcessCommand == "" {
			return fmt.Errorf("invalid process_command")
		}
	case CredentialsURI:
		if cp.CredentialsURI == "" {
			return fmt.Errorf("invalid credentials_uri")
		}
	case ChainableRamRoleArn:
		if cp.SourceProfile == "" {
			return fmt.Errorf("invalid source_profile")
		}
		if cp.RamRoleArn == "" {
			return fmt.Errorf("invalid ram_role_arn")
		}
		if cp.RoleSessionName == "" {
			return fmt.Errorf("invalid role_session_name")
		}
	default:
		return fmt.Errorf("invalid mode: %s", cp.Mode)
	}
	return nil
}

func (cp *Profile) GetParent() *Configuration {
	return cp.parent
}

func (cp *Profile) ValidateAK() error {
	if len(cp.AccessKeyId) == 0 {
		return fmt.Errorf("invalid access_key_id: %s", cp.AccessKeyId)
	}
	if len(cp.AccessKeySecret) == 0 {
		return fmt.Errorf("invaild access_key_secret: %s", cp.AccessKeySecret)
	}
	return nil
}

func IsRegion(region string) bool {
	if match, _ := regexp.MatchString("^[-a-zA-Z0-9]+$", region); !match {
		return false
	}
	return true
}

func (cp *Profile) GetSessionCredential(client *sts.Client) (*sts.AssumeRoleResponseBodyCredentials, error) {
	req := &sts.AssumeRoleRequest{
		DurationSeconds: tea.Int64(int64(cp.ExpiredSeconds)),
		Policy:          nil,
		RoleArn:         tea.String(cp.RamRoleArn),
		RoleSessionName: tea.String(cp.RoleSessionName),
	}
	resp, err := client.AssumeRole(req)
	if err != nil {
		return nil, err
	}
	return resp.Body.Credentials, nil
}

func (cp *Profile) GetSTSClientByEcsRamRole() (*sts.Client, error) {
	conf := &credentials.Config{
		Type:     tea.String("ecs_ram_role"),
		RoleName: tea.String(cp.RamRoleName),
	}
	cred, err := credentials.NewCredential(conf)
	if err != nil {
		return nil, err
	}
	return sts.NewClient(&openapi.Config{
		Credential: cred,
	})
}
