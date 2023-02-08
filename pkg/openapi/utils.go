package openapi

import (
	v1openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
)

func openapiConfigToV1(conf *openapi.Config) *v1openapi.Config {
	return &v1openapi.Config{
		AccessKeyId:          conf.AccessKeyId,
		AccessKeySecret:      conf.AccessKeySecret,
		SecurityToken:        conf.SecurityToken,
		Protocol:             conf.Protocol,
		Method:               conf.Method,
		RegionId:             conf.RegionId,
		ReadTimeout:          conf.ReadTimeout,
		ConnectTimeout:       conf.ConnectTimeout,
		HttpProxy:            conf.HttpsProxy,
		HttpsProxy:           conf.HttpsProxy,
		Credential:           conf.Credential,
		Endpoint:             conf.Endpoint,
		NoProxy:              conf.NoProxy,
		MaxIdleConns:         conf.MaxIdleConns,
		Network:              conf.Network,
		UserAgent:            conf.UserAgent,
		Suffix:               conf.Suffix,
		Socks5Proxy:          conf.Socks5Proxy,
		Socks5NetWork:        conf.Socks5NetWork,
		EndpointType:         conf.EndpointType,
		OpenPlatformEndpoint: conf.OpenPlatformEndpoint,
		Type:                 conf.Type,
		SignatureVersion:     conf.SignatureVersion,
		SignatureAlgorithm:   conf.SignatureAlgorithm,
	}
}
