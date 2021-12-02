package alibabacloudsdkgo

import (
	"context"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/oidctoken"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/signers"
)

type RAMRoleArnWithOIDCTokenSigner struct {
	provider *oidctoken.RoleProvider
}

func NewRAMRoleArnWithOIDCTokenSigner(providerArn, roleArn, tokenFile, policy, roleSessionName string, sessionDuration int) (*RAMRoleArnWithOIDCTokenSigner, error) {
	r, err := oidctoken.NewRoleProvider(
		providerArn, roleArn, tokenFile, policy, roleSessionName, time.Second*time.Duration(sessionDuration))
	if err != nil {
		return nil, err
	}
	return &RAMRoleArnWithOIDCTokenSigner{provider: r}, nil
}

func (r *RAMRoleArnWithOIDCTokenSigner) GetName() string {
	return "HMAC-SHA1"
}

func (r *RAMRoleArnWithOIDCTokenSigner) GetType() string {
	return ""
}

func (r *RAMRoleArnWithOIDCTokenSigner) GetVersion() string {
	return "1.0"
}

func (r *RAMRoleArnWithOIDCTokenSigner) GetAccessKeyId() (string, error) {
	s, err := r.getStsSinger()
	if err != nil {
		return "", err
	}
	return s.GetAccessKeyId()
}

func (r *RAMRoleArnWithOIDCTokenSigner) GetExtraParam() map[string]string {
	s, err := r.getStsSinger()
	if err != nil {
		return nil
	}
	return s.GetExtraParam()
}

func (r *RAMRoleArnWithOIDCTokenSigner) Sign(stringToSign, secretSuffix string) string {
	s, err := r.getStsSinger()
	if err != nil {
		return ""
	}
	return s.Sign(stringToSign, secretSuffix)
}

func (r *RAMRoleArnWithOIDCTokenSigner) getStsSinger() (*signers.StsTokenSigner, error) {
	cred, err := r.provider.GetCredential(context.TODO())
	if err != nil {
		return nil, err
	}
	c := credentials.NewStsTokenCredential(cred.AccessKeyId, cred.AccessKeySecret, cred.SecurityToken)
	s := signers.NewStsTokenSigner(c)
	return s, nil
}
