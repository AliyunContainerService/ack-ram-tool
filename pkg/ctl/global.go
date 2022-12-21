package ctl

type globalOption struct {
	AssumeYes             bool
	Region                string
	UseVPCEndpoint        bool
	InsecureSkipTLSVerify bool

	CredentialFilePath      string
	AliyuncliConfigFilePath string
	AliyuncliProfileName    string
}

var GlobalOption = &globalOption{}
