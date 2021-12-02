package openapi

import (
	"context"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/oidctoken"
)

func AssumeRoleWithOIDCToken(ctx context.Context, providerArn, roleArn string, sessionDuration time.Duration, token []byte) (*oidctoken.Credential, error) {
	return oidctoken.AssumeRoleWithOIDCToken(ctx,
		providerArn, roleArn, string(token), StsApiEndpoint, "https", "",
		"", sessionDuration)
}
