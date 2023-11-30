package openapi

import (
	cs "github.com/alibabacloud-go/cs-20151215/v3/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ram "github.com/alibabacloud-go/ram-20150501/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
	// "github.com/aliyun/credentials-go/credentials"
)

var (
	defaultCsApiEndpoint  = "cs.aliyuncs.com"
	defaultRamApiEndpoint = "ram.aliyuncs.com"
	defaultStsApiEndpoint = "sts.aliyuncs.com"
)

type ClientInterface interface {
	RamClientInterface
	CSClientInterface
	StsClientInterface
}

type Client struct {
	ramClient *ram.Client
	stsClient *sts.Client
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
	stsClient, err := sts.NewClient(v1config)
	if err != nil {
		return nil, err
	}
	stsClient.Endpoint = tea.String(defaultStsApiEndpoint)

	return &Client{
		ramClient: ramClient,
		stsClient: stsClient,
		csClient:  csClient,
	}, nil
}

func (c *Client) Credential() credentials.Credential {
	return c.csClient.Credential
}
