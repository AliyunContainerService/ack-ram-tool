package ramauthenticator

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

const (
	tokenPrefixV1 = "k8s-ack-v1."
)

type Token struct {
	preSignedURLString string
}

func GenerateToken(clusterId, stsEndpoint string, cred credentials.Credential) (*Token, error) {
	getCallerIdentityURL, err := generateGetCallerIdentityURL(clusterId, stsEndpoint, cred)
	if err != nil {
		return nil, err
	}
	token := &Token{preSignedURLString: getCallerIdentityURL}
	return token, nil
}

func (t *Token) String() string {
	return tokenPrefixV1 + base64.RawURLEncoding.EncodeToString([]byte(t.preSignedURLString))
}

func generateGetCallerIdentityURL(clusterId, stsEndpoint string, cred credentials.Credential) (string, error) {
	if !strings.Contains(stsEndpoint, "://") {
		stsEndpoint = "https://" + stsEndpoint
	}
	rawReq := &openapi.OpenApiRequest{
		Query: map[string]*string{
			"ClusterId": tea.String(clusterId),
		},
	}
	teaReq := tea.NewRequest()
	teaReq.Method = tea.String("GET")
	teaReq.Pathname = tea.String("/")
	teaReq.Query = tea.Merge(map[string]*string{
		"Action":         tea.String("GetCallerIdentity"),
		"Format":         tea.String("json"),
		"Version":        tea.String("2015-04-01"),
		"Timestamp":      openapiutil.GetTimestamp(),
		"SignatureNonce": util.GetNonce(),
	}, rawReq.Query)

	accessKeyId, err := cred.GetAccessKeyId()
	if err != nil {
		return "", err
	}

	accessKeySecret, err := cred.GetAccessKeySecret()
	if err != nil {
		return "", err
	}

	securityToken, err := cred.GetSecurityToken()
	if err != nil {
		return "", err
	}

	if !tea.BoolValue(util.Empty(securityToken)) {
		teaReq.Query["SecurityToken"] = securityToken
	}

	teaReq.Query["SignatureMethod"] = tea.String("HMAC-SHA1")
	teaReq.Query["SignatureVersion"] = tea.String("1.0")
	teaReq.Query["AccessKeyId"] = accessKeyId

	signedParam := teaReq.Query
	teaReq.Query["Signature"] = openapiutil.GetRPCSignature(signedParam, teaReq.Method, accessKeySecret)

	requestURL := ""
	requestURL = fmt.Sprintf("%s%s", stsEndpoint, tea.StringValue(teaReq.Pathname))
	queryParams := teaReq.Query
	// sort QueryParams by key
	q := url.Values{}
	for key, value := range queryParams {
		q.Add(key, tea.StringValue(value))
	}
	querystring := q.Encode()
	if len(querystring) > 0 {
		if strings.Contains(requestURL, "?") {
			requestURL = fmt.Sprintf("%s&%s", requestURL, querystring)
		} else {
			requestURL = fmt.Sprintf("%s?%s", requestURL, querystring)
		}
	}

	return requestURL, nil
}
