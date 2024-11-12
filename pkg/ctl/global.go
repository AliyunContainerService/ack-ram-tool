package ctl

import (
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"os"
	"strconv"
	"strings"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
)

const (
	EnvAssumeYes                  = "ACK_RAM_TOOL_ASSUME_YES"
	EnvProfileFile                = "ACK_RAM_TOOL_PROFILE_FILE"
	EnvProfileNameOld             = "ACK_RAM_TOOL_PROFIL_ENAME"
	EnvProfileName                = "ACK_RAM_TOOL_PROFILE_NAME"
	EnvIgnoreEnvCredentials       = "ACK_RAM_TOOL_IGNORE_ENV_CREDENTIALS"        // #nosec G101
	EnvIgnoreAliyunCliCredentials = "ACK_RAM_TOOL_IGNORE_ALIYUN_CLI_CREDENTIALS" // #nosec G101
	EnvLogLevel                   = "ACK_RAM_TOOL_LOG_LEVEL"
	EnvRegionId                   = "ACK_RAM_TOOL_REGION_ID"
	EnvVerbose                    = "ACK_RAM_TOOL_VERBOSE"
	EnvCredentialType             = "ACK_RAM_TOOL_CREDENTIAL_TYPE" // #nosec G101

	DefaultRegion   = ""
	DefaultLogLevel = "info"
	debugLogLevel   = "debug"
)

type globalOption struct {
	AssumeYes             bool
	Region                string
	UseVPCEndpoint        bool
	InsecureSkipTLSVerify bool

	UseSpecifiedCredentialFile bool
	CredentialFilePath         string

	ProfileName                   string
	IgnoreEnv                     bool
	IgnoreAliyuncliConfig         bool
	FinalAssumeRoleAnotherRoleArn string

	LogLevel  string
	ClusterId string
	Verbose   bool
}

var GlobalOption = &globalOption{}

func (g *globalOption) GetRegion() string {
	return g.Region
}

func (g *globalOption) UpdateValues() {
	if v, err := strconv.ParseBool(os.Getenv(EnvAssumeYes)); err == nil && v {
		g.AssumeYes = true
	}
	if g.CredentialFilePath == "" {
		g.CredentialFilePath = os.Getenv(EnvProfileFile)
	}
	if g.ProfileName == "" {
		g.ProfileName = os.Getenv(EnvProfileName)
		if g.ProfileName == "" {
			g.ProfileName = os.Getenv(EnvProfileNameOld)
		}
	}
	if v, err := strconv.ParseBool(os.Getenv(EnvIgnoreEnvCredentials)); err == nil && v {
		g.IgnoreEnv = true
	}
	if v, err := strconv.ParseBool(os.Getenv(EnvIgnoreAliyunCliCredentials)); err == nil && v {
		g.IgnoreAliyuncliConfig = true
	}
	if v, err := strconv.ParseBool(os.Getenv(EnvVerbose)); err == nil && v {
		g.Verbose = true
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
	if g.Verbose {
		g.LogLevel = debugLogLevel
	}

	debugEnv := strings.Split(strings.ToLower(os.Getenv("DEBUG")), ",")
	for _, item := range debugEnv {
		if item == "sdk" || item == "tea" || item == "ack-ram-tool" {
			g.LogLevel = debugLogLevel
			break
		}
	}
}

func (g *globalOption) GetCredentialFilePath() string {
	if strings.HasSuffix(g.CredentialFilePath, ".json") {
		return ""
	}
	return g.CredentialFilePath
}

func (g *globalOption) GetAliyuncliConfigFilePath() string {
	if strings.HasSuffix(g.CredentialFilePath, ".json") {
		return g.CredentialFilePath
	}
	return ""
}

func (g *globalOption) GetProfileName() string {
	return g.ProfileName
}

func (g *globalOption) GetIgnoreEnv() bool {
	return g.IgnoreEnv
}

func (g *globalOption) GetIgnoreAliyuncliConfig() bool {
	return g.IgnoreAliyuncliConfig
}

func (g *globalOption) GetClusterId() string {
	return g.ClusterId
}

func (g *globalOption) GetLogLevel() string {
	return g.LogLevel
}

func (g *globalOption) GetRoleArn() string {
	return g.FinalAssumeRoleAnotherRoleArn
}

func (g *globalOption) GetEndpoints() openapi.Endpoints {
	region := g.GetRegion()
	endpoints := openapi.NewEndpoints(region, g.UseVPCEndpoint)
	endpoints.STS = g.GetSTSEndpoint()
	return endpoints
}

func (g *globalOption) GetCredentialType() string {
	return os.Getenv(EnvCredentialType)
}

func (g *globalOption) GetSTSEndpoint() string {
	region := g.GetRegion()
	return provider.GetSTSEndpoint(region, g.UseVPCEndpoint)
}
