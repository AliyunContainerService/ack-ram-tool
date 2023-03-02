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
	"encoding/json"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	"os"
	"path"
)

const (
	defaultConfigDir         = "~/.aliyun"
	configFile               = "config.json"
	DefaultConfigProfileName = "default"
)

type Configuration struct {
	CurrentProfile string    `json:"current"`
	Profiles       []Profile `json:"profiles"`
	MetaPath       string    `json:"meta_path"`
	//Plugins 		[]Plugin `json:"plugin"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		CurrentProfile: DefaultConfigProfileName,
		Profiles: []Profile{
			NewProfile(DefaultConfigProfileName),
		},
	}
}

func (c *Configuration) GetProfile(pn string) (Profile, bool) {
	for _, p := range c.Profiles {
		if p.Name == pn {
			return p, true
		}
	}
	return Profile{Name: pn}, false
}

func getDefaultConfigPath() string {
	dir, err := utils.ExpandPath(defaultConfigDir)
	if err != nil {
		dir = defaultConfigDir
	}
	return path.Join(dir, configFile)
}

func LoadConfiguration(path string) (conf *Configuration, err error) {
	_, statErr := os.Stat(path)
	if os.IsNotExist(statErr) {
		conf = NewConfiguration()
		return
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("reading config from '%s' failed %v", path, err)
		return
	}

	conf, err = NewConfigFromBytes(bytes)
	return
}

func NewConfigFromBytes(bytes []byte) (conf *Configuration, err error) {
	conf = NewConfiguration()
	err = json.Unmarshal(bytes, conf)
	return
}
