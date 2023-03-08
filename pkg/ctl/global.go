package ctl

import "strings"

type globalOption struct {
	AssumeYes             bool
	Region                string
	UseVPCEndpoint        bool
	InsecureSkipTLSVerify bool

	CredentialFilePath string

	ProfileName           string
	IgnoreEnv             bool
	IgnoreAliyuncliConfig bool

	ClusterId string
}

var GlobalOption = &globalOption{}

func (g globalOption) GetRegion() string {
	return g.Region
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
