package ramauthenticator

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
	"strings"
	"time"
)

const (
	tokenPrefixV1 = "k8s-ack-v2." // #nosec G101
)

var tokenExpiration = time.Minute * 15 // #nosec G101

var signParamsWhitelist = map[string]bool{
	"x-acs-action":          true,
	"x-acs-version":         true,
	"authorization":         true,
	"x-acs-signature-nonce": true,
	"x-acs-date":            true,
	"x-acs-content-sha256":  true,
	"x-acs-content-sm3":     true,
	"x-acs-security-token":  true,
	"ackclusterid":          true,
}

type Token struct {
	ClusterId string `json:"clusterId"`

	Method  string            `json:"method"`
	Path    string            `json:"path"`
	Query   map[string]string `json:"query"`
	Headers map[string]string `json:"headers"`

	Expiration time.Time `json:"-"`
}

func GenerateToken(clusterId string, cred credentials.Credential) (*Token, error) {
	q := &openapi.OpenApiRequest{
		Query: map[string]*string{
			"ACKClusterId": tea.String(clusterId),
		},
	}
	params := &openapi.Params{
		Action:      tea.String("GetCallerIdentity"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	req, err := generatePreSignedReq(q, params, cred)
	if err != nil {
		return nil, err
	}

	t := &Token{
		ClusterId: clusterId,
		Method:    tea.StringValue(params.Method),
		Path:      tea.StringValue(req.Pathname),
		Headers: map[string]string{
			"user-agent": version.UserAgent(),
		},
		Query: map[string]string{},
	}
	for k, v := range req.Headers {
		if !signParamsWhitelist[strings.ToLower(k)] {
			continue
		}
		t.Headers[k] = tea.StringValue(v)
	}
	for k, v := range req.Query {
		if !signParamsWhitelist[strings.ToLower(k)] {
			continue
		}
		t.Query[k] = tea.StringValue(v)
	}

	t.Expiration = time.Now().Add(tokenExpiration - 5*time.Minute).UTC()

	return t, nil
}

func (t *Token) String() string {
	req, _ := json.Marshal(t)
	return tokenPrefixV1 + base64.StdEncoding.EncodeToString(req)
}

func generatePreSignedReq(request *openapi.OpenApiRequest, params *openapi.Params, cred credentials.Credential) (*tea.Request, error) {
	newReq := tea.NewRequest()
	newReq.Protocol = util.DefaultString(nil, params.Protocol)
	newReq.Method = params.Method
	newReq.Pathname = params.Pathname

	newReq.Query = tea.Merge(request.Query)

	// endpoint is setted in product client
	newReq.Headers = tea.Merge(map[string]*string{
		//"host":                  client.Endpoint,
		"x-acs-version": params.Version,
		"x-acs-action":  params.Action,
		"user-agent":    tea.String(version.UserAgent()),
		"x-acs-date":    openapiutil.GetTimestamp(),
		//"x-acs-signature-nonce": util.GetNonce(),
		"accept": tea.String("application/json"),
	}, request.Headers)

	signatureAlgorithm := util.DefaultString(nil, tea.String("ACS3-HMAC-SHA256"))
	hashedRequestPayload := openapiutil.HexEncode(openapiutil.Hash(util.ToBytes(tea.String("")), signatureAlgorithm))

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		if tea.BoolValue(util.EqualString(params.ReqBodyType, tea.String("json"))) {
			jsonObj := util.ToJSONString(request.Body)
			hashedRequestPayload = openapiutil.HexEncode(openapiutil.Hash(util.ToBytes(jsonObj), signatureAlgorithm))
			newReq.Body = tea.ToReader(jsonObj)
			newReq.Headers["content-type"] = tea.String("application/json; charset=utf-8")
		} else {
			m := util.AssertAsMap(request.Body)
			formObj := openapiutil.ToForm(m)
			hashedRequestPayload = openapiutil.HexEncode(openapiutil.Hash(util.ToBytes(formObj), signatureAlgorithm))
			newReq.Body = tea.ToReader(formObj)
			newReq.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
		}
	}

	newReq.Headers["x-acs-content-sha256"] = hashedRequestPayload
	if !tea.BoolValue(util.EqualString(params.AuthType, tea.String("Anonymous"))) {
		authType := cred.GetType()
		if tea.BoolValue(util.EqualString(authType, tea.String("bearer"))) {
			bearerToken := cred.GetBearerToken()
			if tea.StringValue(bearerToken) == "" {
				return nil, errors.New("GetBearerToken failed")
			}

			newReq.Headers["x-acs-bearer-token"] = bearerToken
		} else {
			accessKeyId, _err := cred.GetAccessKeyId()
			if _err != nil {
				return nil, _err
			}
			accessKeySecret, _err := cred.GetAccessKeySecret()
			if _err != nil {
				return nil, _err
			}
			securityToken, _err := cred.GetSecurityToken()
			if _err != nil {
				return nil, _err
			}
			if !tea.BoolValue(util.Empty(securityToken)) {
				//newReq.Headers["x-acs-accesskey-id"] = accessKeyId
				newReq.Headers["x-acs-security-token"] = securityToken
			}

			newReq.Headers["Authorization"] = openapiutil.GetAuthorization(
				newReq, signatureAlgorithm, hashedRequestPayload, accessKeyId, accessKeySecret)
		}
	}

	return newReq, nil
}
