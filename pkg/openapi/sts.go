package openapi

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/alibabacloud-go/tea/tea"
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
	switch tea.StringValue(body.IdentityType) {
	case "Account":
		return &types.Account{
			Type:    types.AccountTypeRoot,
			RootUId: tea.StringValue(body.AccountId),
			User: types.RamUser{
				Id: tea.StringValue(body.UserId),
			},
		}, nil
	case "RAMUser":
		return &types.Account{
			Type:    types.AccountTypeUser,
			RootUId: tea.StringValue(body.AccountId),
			User: types.RamUser{
				Id: tea.StringValue(body.UserId),
			},
		}, nil
	case "AssumedRoleUser":
		return &types.Account{
			Type:    types.AccountTypeRole,
			RootUId: tea.StringValue(body.AccountId),
			Role: types.RamRole{
				RoleId: tea.StringValue(body.RoleId),
				Arn:    tea.StringValue(body.Arn),
			},
		}, nil
	}

	return nil, fmt.Errorf("unkown resp: %s", resp.String())
}
