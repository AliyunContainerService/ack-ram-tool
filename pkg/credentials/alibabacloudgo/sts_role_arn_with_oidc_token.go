package alibabacloudgo

import (
	"context"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/oidctoken"
	"github.com/alibabacloud-go/tea/tea"
)

type RAMRoleArnWithOIDCTokenCredential struct {
	provider *oidctoken.RoleProvider
}

func NewRAMRoleArnWithOIDCTokenCredential(providerArn, roleArn, tokenFile, policy, roleSessionName string, sessionDuration int) (*RAMRoleArnWithOIDCTokenCredential, error) {
	r, err := oidctoken.NewRoleProvider(
		providerArn, roleArn, tokenFile, policy, roleSessionName, time.Second*time.Duration(sessionDuration))
	if err != nil {
		return nil, err
	}
	return &RAMRoleArnWithOIDCTokenCredential{provider: r}, nil
}

func (r *RAMRoleArnWithOIDCTokenCredential) GetAccessKeyId() (*string, error) {
	cred, err := r.provider.GetCredential(context.TODO())
	if err != nil {
		return nil, err
	}
	return tea.String(cred.AccessKeyId), nil
}

func (r *RAMRoleArnWithOIDCTokenCredential) GetAccessKeySecret() (*string, error) {
	cred, err := r.provider.GetCredential(context.TODO())
	if err != nil {
		return nil, err
	}
	return tea.String(cred.AccessKeySecret), nil
}

func (r *RAMRoleArnWithOIDCTokenCredential) GetSecurityToken() (*string, error) {
	cred, err := r.provider.GetCredential(context.TODO())
	if err != nil {
		return nil, err
	}
	return tea.String(cred.SecurityToken), nil
}

func (r *RAMRoleArnWithOIDCTokenCredential) GetBearerToken() *string {
	return tea.String("")
}

func (r *RAMRoleArnWithOIDCTokenCredential) GetType() *string {
	return tea.String("ram_role_arn_with_oidc_token")
}
