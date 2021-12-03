package ctl

type globalOption struct {
	AssumeYes bool
	Region string
	UseVPCEndpoint bool
}

var GlobalOption = &globalOption{}
