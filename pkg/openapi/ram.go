package openapi

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	ram "github.com/alibabacloud-go/ram-20150501/client"
	"github.com/alibabacloud-go/tea/tea"
)

type RamClientInterface interface {
	GetRole(ctx context.Context, name string) (*types.RamRole, error)
	CreateRole(ctx context.Context, role types.RamRole) (*types.RamRole, error)
	UpdateRole(ctx context.Context, name string, opt UpdateRamRoleOption) (*types.RamRole, error)
	ListPoliciesForRole(ctx context.Context, name string) ([]types.RamRolePolicy, error)
	GetPolicy(ctx context.Context, name, policyType string) (*types.RamPolicy, error)
	CreatePolicy(ctx context.Context, policy types.RamPolicy) (*types.RamPolicy, error)
	AttachPolicyToRole(ctx context.Context, policyName, policyType, roleName string) error
	ListUsers(ctx context.Context) ([]types.RamUser, error)
	ListRoles(ctx context.Context) ([]types.RamRole, error)
}

func (c *Client) GetRole(ctx context.Context, name string) (*types.RamRole, error) {
	client := c.ramClient
	req := &ram.GetRoleRequest{
		RoleName: tea.String(name),
	}
	resp, err := client.GetRole(req)
	if err != nil {
		return nil, err
	}

	role := &types.RamRole{}
	convertGetRoleResponse(role, resp)
	return role, nil
}

func (c *Client) CreateRole(ctx context.Context, role types.RamRole) (*types.RamRole, error) {
	client := c.ramClient
	if role.AssumeRolePolicyDocument == nil || len(*role.AssumeRolePolicyDocument) == 0 {
		return nil, errors.New("AssumeRolePolicyDocument is required")
	}
	if role.RoleName == "" {
		return nil, errors.New("RoleName is required")
	}
	req := &ram.CreateRoleRequest{
		RoleName:                 tea.String(role.RoleName),
		Description:              tea.String(role.Description),
		AssumeRolePolicyDocument: tea.String(role.AssumeRolePolicyDocument.JSON()),
	}
	if role.MaxSessionDuration != 0 {
		req.MaxSessionDuration = tea.Int64(role.MaxSessionDuration)
	}
	resp, err := client.CreateRole(req)
	if err != nil {
		return nil, err
	}

	roleV2 := &types.RamRole{}
	convertCreateRoleResponse(roleV2, resp)
	return roleV2, nil
}

type UpdateRamRoleOption struct {
	AssumeRolePolicyDocument *types.RamPolicyDocument
}

func (c *Client) UpdateRole(ctx context.Context, name string, opt UpdateRamRoleOption) (*types.RamRole, error) {
	client := c.ramClient
	if opt.AssumeRolePolicyDocument == nil || len(*opt.AssumeRolePolicyDocument) == 0 {
		return nil, errors.New("AssumeRolePolicyDocument is required")
	}
	req := &ram.UpdateRoleRequest{
		RoleName:                    tea.String(name),
		NewAssumeRolePolicyDocument: tea.String(opt.AssumeRolePolicyDocument.JSON()),
	}
	resp, err := client.UpdateRole(req)
	if err != nil {
		return nil, err
	}

	role := &types.RamRole{}
	convertUpdateRoleResponse(role, resp)
	return role, nil
}

func (c *Client) ListPoliciesForRole(ctx context.Context, name string) ([]types.RamRolePolicy, error) {
	client := c.ramClient
	req := &ram.ListPoliciesForRoleRequest{
		RoleName: tea.String(name),
	}
	resp, err := client.ListPoliciesForRole(req)
	if err != nil {
		return nil, err
	}

	policies := convertListPoliciesForRoleResponse(resp)
	return policies, nil
}

func (c *Client) GetPolicy(ctx context.Context, name, policyType string) (*types.RamPolicy, error) {
	client := c.ramClient
	req := &ram.GetPolicyRequest{
		PolicyType: tea.String(policyType),
		PolicyName: tea.String(name),
	}
	resp, err := client.GetPolicy(req)
	if err != nil {
		return nil, err
	}

	policy := &types.RamPolicy{}
	convertGetRamPolicyResponse(policy, resp)
	return policy, nil
}

func (c *Client) CreatePolicy(ctx context.Context, policy types.RamPolicy) (*types.RamPolicy, error) {
	client := c.ramClient
	if policy.PolicyDocument == nil || len(*policy.PolicyDocument) == 0 {
		return nil, errors.New("PolicyDocument is required")
	}
	if policy.PolicyName == "" {
		return nil, errors.New("RoleName is required")
	}
	req := &ram.CreatePolicyRequest{
		PolicyName:     tea.String(policy.PolicyName),
		Description:    tea.String(policy.Description),
		PolicyDocument: tea.String(policy.PolicyDocument.JSON()),
	}
	resp, err := client.CreatePolicy(req)
	if err != nil {
		return nil, err
	}

	policyV2 := &types.RamPolicy{}
	convertCreateRamPolicyResponse(policyV2, resp)
	if policyV2.PolicyDocument == nil {
		policyV2.PolicyDocument = policy.PolicyDocument
	}
	return policyV2, nil
}

func (c *Client) AttachPolicyToRole(ctx context.Context, policyName, policyType, roleName string) error {
	client := c.ramClient
	req := &ram.AttachPolicyToRoleRequest{
		PolicyType: tea.String(policyType),
		PolicyName: tea.String(policyName),
		RoleName:   tea.String(roleName),
	}
	_, err := client.AttachPolicyToRole(req)
	return err
}

func convertGetRoleResponse(r *types.RamRole, resp *ram.GetRoleResponse) {
	body := resp.Body
	if body == nil {
		return
	}
	role := body.Role
	if role == nil {
		return
	}

	r.RoleId = tea.StringValue(role.RoleId)
	r.RoleName = tea.StringValue(role.RoleName)
	r.Arn = tea.StringValue(role.Arn)
	r.Description = tea.StringValue(role.Description)
	r.MaxSessionDuration = tea.Int64Value(role.MaxSessionDuration)
	if role.AssumeRolePolicyDocument != nil {
		policy := &types.RamPolicyDocument{}
		if err := json.Unmarshal([]byte(*role.AssumeRolePolicyDocument), policy); err != nil {
			log.Logger.Errorf("unmarshal AssumeRolePolicyDocument failed: %+v: \n%s\n", *role.AssumeRolePolicyDocument, err)
		}
		r.AssumeRolePolicyDocument = policy
	}
}

func convertUpdateRoleResponse(r *types.RamRole, resp *ram.UpdateRoleResponse) {
	body := resp.Body
	if body == nil {
		return
	}
	role := body.Role
	if role == nil {
		return
	}

	r.RoleId = tea.StringValue(role.RoleId)
	r.RoleName = tea.StringValue(role.RoleName)
	r.Arn = tea.StringValue(role.Arn)
	r.Description = tea.StringValue(role.Description)
	r.MaxSessionDuration = tea.Int64Value(role.MaxSessionDuration)
	if role.AssumeRolePolicyDocument != nil {
		policy := &types.RamPolicyDocument{}
		if err := json.Unmarshal([]byte(*role.AssumeRolePolicyDocument), policy); err != nil {
			log.Logger.Errorf("unmarshal AssumeRolePolicyDocument failed: %+v: \n%s\n", *role.AssumeRolePolicyDocument, err)
		}
		r.AssumeRolePolicyDocument = policy
	}
}

func convertCreateRoleResponse(r *types.RamRole, resp *ram.CreateRoleResponse) {
	body := resp.Body
	if body == nil {
		return
	}
	role := body.Role
	if role == nil {
		return
	}

	r.RoleId = tea.StringValue(role.RoleId)
	r.RoleName = tea.StringValue(role.RoleName)
	r.Arn = tea.StringValue(role.Arn)
	r.Description = tea.StringValue(role.Description)
	r.MaxSessionDuration = tea.Int64Value(role.MaxSessionDuration)
	if role.AssumeRolePolicyDocument != nil {
		policy := &types.RamPolicyDocument{}
		if err := json.Unmarshal([]byte(*role.AssumeRolePolicyDocument), policy); err != nil {
			log.Logger.Errorf("unmarshal AssumeRolePolicyDocument failed: %+v: \n%s\n", *role.AssumeRolePolicyDocument, err)
		}
		r.AssumeRolePolicyDocument = policy
	}
}

func convertListPoliciesForRoleResponse(resp *ram.ListPoliciesForRoleResponse) []types.RamRolePolicy {
	body := resp.Body
	if body == nil {
		return nil
	}
	if body.Policies == nil {
		return nil
	}
	policies := body.Policies.Policy
	if policies == nil {
		return nil
	}

	var rs []types.RamRolePolicy
	for _, p := range policies {
		r := types.RamRolePolicy{
			DefaultVersion: tea.StringValue(p.DefaultVersion),
			Description:    tea.StringValue(p.Description),
			PolicyName:     tea.StringValue(p.PolicyName),
			//AttachDate:     tea.StringValue(p.AttachDate),
			PolicyType: tea.StringValue(p.PolicyType),
		}
		rs = append(rs, r)
	}

	return rs
}

func convertGetRamPolicyResponse(r *types.RamPolicy, resp *ram.GetPolicyResponse) {
	body := resp.Body
	if body == nil {
		return
	}
	p := body.Policy
	if p == nil {
		return
	}

	r.PolicyType = tea.StringValue(p.PolicyType)
	r.PolicyName = tea.StringValue(p.PolicyName)
	r.DefaultVersion = tea.StringValue(p.DefaultVersion)
	r.Description = tea.StringValue(p.Description)
	r.AttachmentCount = tea.Int32Value(p.AttachmentCount)

	if p.PolicyDocument != nil {
		policy := &types.RamPolicyDocument{}
		if err := json.Unmarshal([]byte(*p.PolicyDocument), policy); err != nil {
			log.Logger.Errorf("unmarshal PolicyDocument failed: %+v: \n%s\n", *p.PolicyDocument, err)
		}
		r.PolicyDocument = policy
	}
}

func convertCreateRamPolicyResponse(r *types.RamPolicy, resp *ram.CreatePolicyResponse) {
	body := resp.Body
	if body == nil {
		return
	}
	p := body.Policy
	if p == nil {
		return
	}

	r.PolicyType = tea.StringValue(p.PolicyType)
	r.PolicyName = tea.StringValue(p.PolicyName)
	r.DefaultVersion = tea.StringValue(p.DefaultVersion)
	r.Description = tea.StringValue(p.Description)
}

func (c *Client) ListUsers(ctx context.Context) ([]types.RamUser, error) {
	var users []types.RamUser
	var marker string
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		req := &ram.ListUsersRequest{
			Marker:   nil,
			MaxItems: tea.Int32(1000),
		}
		if marker != "" {
			req.Marker = tea.String(marker)
		}
		resp, err := c.ramClient.ListUsers(req)
		if err != nil {
			return nil, err
		}
		items := convertListUsersResponse(resp)
		users = append(users, items...)
		if resp.Body != nil {
			if !tea.BoolValue(resp.Body.IsTruncated) {
				break
			}
			marker = tea.StringValue(resp.Body.Marker)
		}
	}
	return users, nil
}

func (c *Client) ListRoles(ctx context.Context) ([]types.RamRole, error) {
	var roles []types.RamRole
	var marker string
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		req := &ram.ListRolesRequest{
			Marker:   nil,
			MaxItems: tea.Int32(1000),
		}
		if marker != "" {
			req.Marker = tea.String(marker)
		}
		resp, err := c.ramClient.ListRoles(req)
		if err != nil {
			return nil, err
		}
		items := convertListRolesResponse(resp)
		roles = append(roles, items...)
		if resp.Body != nil {
			if !tea.BoolValue(resp.Body.IsTruncated) {
				break
			}
			marker = tea.StringValue(resp.Body.Marker)
		}
	}
	return roles, nil
}

func convertListUsersResponse(resp *ram.ListUsersResponse) []types.RamUser {
	body := resp.Body
	if body == nil {
		return nil
	}
	p := body.Users
	if p == nil {
		return nil
	}
	us := p.User
	if us == nil {
		return nil
	}

	var ret []types.RamUser
	for _, u := range us {
		ret = append(ret, types.RamUser{
			Id:          tea.StringValue(u.UserId),
			Name:        tea.StringValue(u.UserName),
			DisplayName: tea.StringValue(u.DisplayName),
			Deleted:     false,
		})
	}

	return ret
}

func convertListRolesResponse(resp *ram.ListRolesResponse) []types.RamRole {
	body := resp.Body
	if body == nil {
		return nil
	}
	p := body.Roles
	if p == nil {
		return nil
	}
	us := p.Role
	if us == nil {
		return nil
	}

	var ret []types.RamRole
	for _, u := range us {
		ret = append(ret, types.RamRole{
			RoleId:                   tea.StringValue(u.RoleId),
			RoleName:                 tea.StringValue(u.RoleName),
			Arn:                      tea.StringValue(u.Arn),
			Description:              tea.StringValue(u.Description),
			AssumeRolePolicyDocument: nil,
			MaxSessionDuration:       tea.Int64Value(u.MaxSessionDuration),
			Deleted:                  false,
		})
	}

	return ret
}

func IsRamRoleNotExistErr(err error) bool {
	return isSdkErrWithCode(err, "EntityNotExist.Role")
}

func IsRamUserNotExistErr(err error) bool {
	return isSdkErrWithCode(err, "EntityNotExist.User")
}

func IsRamPolicyNotExistErr(err error) bool {
	return isSdkErrWithCode(err, "EntityNotExist.Policy")
}

func IsRamPolicyAttachedToRoleErr(err error) bool {
	return isSdkErrWithCode(err, "EntityAlreadyExists.Role.Policy")
}

func isSdkErrWithCode(err error, code string) bool {
	if err == nil {
		return false
	}
	sdkErr, ok := err.(*tea.SDKError)
	if !ok {
		return false
	}
	return tea.StringValue(sdkErr.Code) == code
}
