package exportcredentials

import (
	"math/rand"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
)

type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	Expiration      string
}

// TODO: add cache
func getCredentials(client *openapi.Client) (*Credentials, error) {
	cc := client.Credential()
	ak, err := cc.GetAccessKeyId()
	if err != nil {
		return nil, err
	}
	as, err := cc.GetAccessKeySecret()
	if err != nil {
		return nil, err
	}
	st, err := cc.GetSecurityToken()
	if err != nil {
		return nil, err
	}
	exp := getExpirationWithJitter(time.Now())

	cred := Credentials{
		AccessKeyId:     *ak,
		AccessKeySecret: *as,
		SecurityToken:   *st,
		Expiration:      exp.UTC().Format("2006-01-02T15:04:05Z"),
	}

	return &cred, nil
}

func getExpirationWithJitter(t time.Time) time.Time {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))      // #nosec G404
	jitter := time.Duration(r.Int63n(int64(time.Minute) * 4)) // #nosec G404
	exp := t.Add(time.Minute*8 + jitter)                      // 8 + [0, 4) minutes
	return exp
}

func (c *Credentials) Format(format string) string {
	output := ""
	switch format {
	case formatCredentialFileIni, formatCredentialFileIniShort:
		output = toCredentialFileIni(*c)
	case formatAliyunCLIURIJSON, formatAliyunCLIURIJSONShort,
		formatECSMetadataJSON, formatECSMetadataJSONShort:
		output = toAliyunCLIURIBody(*c)
	case formatEnvironmentVariables, formatEnvironmentVariablesShort:
		output = toExportEnvironmentVariables(*c)
	default:
		output = toAliyunCLIConfigJSON(*c)
	}
	return output
}
