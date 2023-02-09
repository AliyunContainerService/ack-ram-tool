package openapi

import (
	cs "github.com/alibabacloud-go/cs-20151215/v3/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ram "github.com/alibabacloud-go/ram-20150501/client"
	"github.com/alibabacloud-go/tea/tea"
	// "github.com/aliyun/credentials-go/credentials"
)

var (
	defaultCsApiEndpoint  = "cs.aliyuncs.com"
	defaultRamApiEndpoint = "ram.aliyuncs.com"
	defaultStsApiEndpoint = "sts.aliyuncs.com"
)

type Client struct {
	ramClient *ram.Client
	csClient  *cs.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	csClient, err := cs.NewClient(config)
	if err != nil {
		return nil, err
	}
	csClient.Endpoint = tea.String(defaultCsApiEndpoint)

	v1config := openapiConfigToV1(config)
	ramClient, err := ram.NewClient(v1config)
	if err != nil {
		return nil, err
	}
	ramClient.Endpoint = tea.String(defaultRamApiEndpoint)
	return &Client{
		ramClient: ramClient,
		csClient:  csClient,
	}, nil
}
