package openapi

import (
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	cs "github.com/alibabacloud-go/cs-20151215/v3/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ram "github.com/alibabacloud-go/ram-20150501/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
	"os"
	"strings"
)

var (
	defaultCsApiEndpoint  = "cs.aliyuncs.com"
	defaultRamApiEndpoint = "ram.aliyuncs.com"
	defaultStsApiEndpoint = "sts.aliyuncs.com"
)

var endpointsTpl = map[string][2]string{
	"cs":  {"cs.%s.aliyuncs.com", "cs-anony-vpc.%s.aliyuncs.com"},
	"ram": {"ram.aliyuncs.com", "ram.vpc-proxy.aliyuncs.com"},
	"sts": {"sts.%s.aliyuncs.com", "sts-vpc.%s.aliyuncs.com"},
}

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

type Endpoints struct {
	RAM string
	STS string
	CS  string
}

func NewClientWithEndpoints(config *openapi.Config, endpoints Endpoints) (*Client, error) {
	client, err := NewClient(config)
	if err != nil {
		return nil, err
	}
	if endpoints.STS != "" {
		client.stsClient.Endpoint = tea.String(endpoints.STS)
	}
	if endpoints.RAM != "" {
		client.ramClient.Endpoint = tea.String(endpoints.RAM)
	}
	if endpoints.CS != "" {
		client.csClient.Endpoint = tea.String(endpoints.CS)
	}

	return client, nil
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

func NewEndpoints(region string, vpc bool) Endpoints {
	endpoints := Endpoints{
		RAM: os.Getenv("RAM_ENDPOINT"),
		STS: provider.GetSTSEndpoint(region, vpc),
		CS:  os.Getenv("CS_ENDPOINT"),
	}
	index := 0
	if vpc {
		index = 1
	}
	if endpoints.RAM == "" {
		tpl := endpointsTpl["ram"][index]
		if strings.Contains(tpl, "%") {
			if region == "" {
				endpoints.RAM = defaultRamApiEndpoint
			} else {
				endpoints.RAM = fmt.Sprintf(tpl, region)
			}
		} else {
			endpoints.RAM = tpl
		}
	}
	if endpoints.STS == "" {
		tpl := endpointsTpl["sts"][index]
		if strings.Contains(tpl, "%") {
			if region == "" {
				endpoints.STS = defaultStsApiEndpoint
			} else {
				endpoints.STS = fmt.Sprintf(tpl, region)
			}
		} else {
			endpoints.STS = tpl
		}
	}
	if endpoints.CS == "" {
		tpl := endpointsTpl["cs"][index]
		if strings.Contains(tpl, "%") {
			if region == "" {
				endpoints.CS = defaultCsApiEndpoint
			} else {
				endpoints.CS = fmt.Sprintf(tpl, region)
			}
		} else {
			endpoints.CS = tpl
		}
	}

	return endpoints
}
