package ctl

import (
	"os"
	"strings"
)

const (
	EnvAssumeYes                  = "ACK_RAM_TOOL_ASSUME_YES"
	EnvProfileFile                = "ACK_RAM_TOOL_PROFILE_FILE"
	EnvProfileName                = "ACK_RAM_TOOL_PROFIL_ENAME"
	EnvIgnoreEnvCredentials       = "ACK_RAM_TOOL_IGNORE_ENV_CREDENTIALS"
	EnvIgnoreAliyunCliCredentials = "ACK_RAM_TOOL_IGNORE_ALIYUN_CLI_CREDENTIALS"
	EnvLogLevel                   = "ACK_RAM_TOOL_LOG_LEVEL"
	EnvRegionId                   = "ACK_RAM_TOOL_REGION_ID"

	DefaultRegion   = "cn-hangzhou"
	DefaultLogLevel = "info"
)

type globalOption struct {
	AssumeYes             bool
	Region                string
	UseVPCEndpoint        bool
	InsecureSkipTLSVerify bool

	UseSpecifiedCredentialFile bool
	CredentialFilePath         string

	ProfileName           string
	IgnoreEnv             bool
	IgnoreAliyuncliConfig bool

	LogLevel  string
	ClusterId string
}

var GlobalOption = &globalOption{}

func (g globalOption) GetRegion() string {
	return g.Region
}

func (g *globalOption) UpdateValues() {
	if os.Getenv(EnvAssumeYes) == "true" {
		g.AssumeYes = true
	}
	if g.CredentialFilePath == "" {
		g.CredentialFilePath = os.Getenv(EnvProfileFile)
	}
	if g.ProfileName == "" {
		g.ProfileName = os.Getenv(EnvProfileName)
	}
	if os.Getenv(EnvIgnoreEnvCredentials) == "true" {
		g.IgnoreEnv = true
	}
	if os.Getenv(EnvIgnoreAliyunCliCredentials) == "true" {
		g.IgnoreAliyuncliConfig = true
	}
	if g.LogLevel == "" {
		g.LogLevel = os.Getenv(EnvLogLevel)
	}
	if g.Region == "" {
		g.Region = os.Getenv(EnvRegionId)
	}

	if g.Region == "" {
		g.Region = DefaultRegion
	}
	if g.LogLevel == "" {
		g.LogLevel = DefaultLogLevel
	}
	if g.CredentialFilePath != "" {
		g.UseSpecifiedCredentialFile = true
	}
	if g.GetCredentialFilePath() != "" {
		g.IgnoreAliyuncliConfig = true
		g.IgnoreEnv = true
	}
}

func (g globalOption) GetCredentialFilePath() string {
	if strings.HasSuffix(g.CredentialFilePath, ".json") {
		return ""
	}
	return g.CredentialFilePath
}

func (g globalOption) GetAliyuncliConfigFilePath() string {
	if strings.HasSuffix(g.CredentialFilePath, ".json") {
		return g.CredentialFilePath
	}
	return ""
}

func (g globalOption) GetProfileName() string {
	return g.ProfileName
}

func (g globalOption) GetIgnoreEnv() bool {
	return g.IgnoreEnv
}

func (g globalOption) GetIgnoreAliyuncliConfig() bool {
	return g.IgnoreAliyuncliConfig
}

func (g globalOption) GetClusterId() string {
	return g.ClusterId
}

func (g globalOption) GetLogLevel() string {
	return g.LogLevel
}
