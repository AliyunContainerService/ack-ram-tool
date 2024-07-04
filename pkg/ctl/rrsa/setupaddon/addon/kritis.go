package addon

import (
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
)

func init() {
	registryAddon(&Kritis{})
}

type Kritis struct{}

func (k *Kritis) AddonName() string {
	return "kritis-validation-hook"
}

func (k *Kritis) RoleName(clusterId string) string {
	return fmt.Sprintf("%s-%s", k.AddonName(), clusterId)
}

func (k *Kritis) RamPolicy() types.RamPolicy {
	policy := types.MakeRamPolicyDocument([]types.RamPolicyStatement{
		{
			"Effect": "Allow",
			"Action": []string{
				"cr:ListInstance",
				"cr:ListMetadataOccurrences",
			},
			"Resource": "*",
		},
	})
	return types.RamPolicy{
		Description:    fmt.Sprintf("policy for ack cluster addon %s", k.AddonName()),
		PolicyDocument: &policy,
		PolicyName:     fmt.Sprintf("ack-addon-policy-%s", k.AddonName()),
		PolicyType:     types.RamPolicyTypeCustom,
	}
}

func (k *Kritis) NameSpace() string {
	return "kube-system"
}

func (k *Kritis) ServiceAccountName() string {
	return "kritis"
}
