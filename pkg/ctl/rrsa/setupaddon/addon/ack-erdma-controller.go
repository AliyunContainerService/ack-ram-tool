package addon

import (
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
)

func init() {
	registryAddon(&AckErdmaController{})
}

const (
	addonName = "ack-erdma-controller"
)

type AckErdmaController struct{}

func (a *AckErdmaController) AddonName() string {
	return addonName
}

func (a *AckErdmaController) RoleName(clusterId string) string {
	return addonName + "-" + clusterId
}

func (a *AckErdmaController) RamPolicy() types.RamPolicy {
	policy := types.MakeRamPolicyDocument([]types.RamPolicyStatement{
		{
			"Effect": "Allow",
			"Action": []string{
				"ecs:DescribeInstances",
				"ecs:DescribeInstanceTypes",
				"ecs:DescribeNetworkInterfaces",
				"ecs:ModifyNetworkInterfaceAttribute",
				"ecs:CreateNetworkInterface",
				"ecs:AttachNetworkInterface",
			},
			"Resource": "*",
		},
	})
	return types.RamPolicy{
		Description:    fmt.Sprintf("policy for ack cluster addon %s", a.AddonName()),
		PolicyDocument: &policy,
		PolicyName:     fmt.Sprintf("ack-addon-policy-%s", a.AddonName()),
		PolicyType:     types.RamPolicyTypeCustom,
	}
}

func (a *AckErdmaController) NameSpace() string {
	return addonName
}

func (a *AckErdmaController) ServiceAccountName() string {
	return addonName
}
