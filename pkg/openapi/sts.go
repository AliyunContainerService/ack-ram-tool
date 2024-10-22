package openapi

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/alibabacloud-go/tea/tea"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/oidctoken"
)

type StsClientInterface interface {
	GetCallerIdentity(ctx context.Context) (*types.Account, error)
}

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

func (c *Client) GetCallerIdentity(ctx context.Context) (*types.Account, error) {
	client := c.stsClient
	resp, err := client.GetCallerIdentity()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, fmt.Errorf("unkown resp: %s", resp.String())
	}
	body := resp.Body
	arn := getRealArn(tea.StringValue(body.Arn))
	switch tea.StringValue(body.IdentityType) {
	case "Account":
		return &types.Account{
			PrincipalId: tea.StringValue(body.PrincipalId),
			Arn:         arn,
			Type:        types.AccountTypeRoot,
			RootUId:     tea.StringValue(body.AccountId),
			User: types.RamUser{
				Id: tea.StringValue(body.UserId),
			},
		}, nil
	case "RAMUser":
		parts := strings.Split(arn, "/")
		name := parts[len(parts)-1]
		return &types.Account{
			PrincipalId: tea.StringValue(body.PrincipalId),
			Arn:         arn,
			Type:        types.AccountTypeUser,
			RootUId:     tea.StringValue(body.AccountId),
			User: types.RamUser{
				Id:   tea.StringValue(body.UserId),
				Name: name,
			},
		}, nil
	case "AssumedRoleUser":
		parts := strings.Split(arn, "/")
		name := parts[len(parts)-1]
		return &types.Account{
			PrincipalId: tea.StringValue(body.PrincipalId),
			Arn:         arn,
			Type:        types.AccountTypeRole,
			RootUId:     tea.StringValue(body.AccountId),
			Role: types.RamRole{
				RoleId:   tea.StringValue(body.RoleId),
				Arn:      tea.StringValue(body.Arn),
				RoleName: name,
			},
		}, nil
	}

	return nil, fmt.Errorf("unkown resp: %s", resp.String())
}

func getRealArn(raw string) string {
	roleKey := ":assumed-role/"
	arn := raw
	if strings.Contains(raw, roleKey) {
		arn = strings.Replace(arn, roleKey, ":role/", 1)
		parts := strings.Split(arn, "/")
		parts = parts[:len(parts)-1]
		arn = strings.Join(parts, "/")
	}
	return arn
}
