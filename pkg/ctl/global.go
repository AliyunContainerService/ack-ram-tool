package ctl

type globalOption struct {
	AssumeYes             bool
	Region                string
	UseVPCEndpoint        bool
	InsecureSkipTLSVerify bool

	CredentialFilePath string

	AliyuncliConfigFilePath string
	ProfileName             string
	IgnoreEnv               bool
	IgnoreAliyuncliConfig   bool

	ClusterId string
}

var GlobalOption = &globalOption{}

func (g globalOption) GetRegion() string {
	return g.Region
}

func (g globalOption) GetCredentialFilePath() string {
	return g.CredentialFilePath
}

func (g globalOption) GetAliyuncliConfigFilePath() string {
	return g.AliyuncliConfigFilePath
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
