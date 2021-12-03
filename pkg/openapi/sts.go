package openapi

import (
	"context"
	"fmt"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/oidctoken"
)

func GetStsEndpoint(region string, vpc bool) string {
	if region == "" {
		return defaultStsApiEndpoint
	}
	if !vpc {
		return fmt.Sprintf("sts.%s.aliyuncs.com", region)
	} else {
		return fmt.Sprintf("sts-vpc.%s.aliyuncs.com", region)
	}
}

func AssumeRoleWithOIDCToken(ctx context.Context, providerArn, roleArn string,
	sessionDuration time.Duration, token []byte, stsEndpoint string) (*oidctoken.Credential, error) {
	return oidctoken.AssumeRoleWithOIDCToken(ctx,
		providerArn, roleArn, string(token), stsEndpoint, "https", "",
		"", sessionDuration)
}
