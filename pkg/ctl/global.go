package ctl

type globalOption struct {
	AssumeYes             bool
	Region                string
	UseVPCEndpoint        bool
	InsecureSkipTLSVerify bool

	CredentialFilePath      string
	AliyuncliConfigFilePath string
	AliyuncliProfileName    string

	ClusterId string
}

var GlobalOption = &globalOption{}
