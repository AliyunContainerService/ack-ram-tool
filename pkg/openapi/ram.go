package openapi

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	ram "github.com/alibabacloud-go/ram-20150501/client"
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateRamRoleOption struct {
	AssumeRolePolicyDocument *types.AssumeRolePolicyDocument
}

type RamClientInterface interface {
	GetRole(ctx context.Context, name string) (*types.RamRole, error)
	UpdateRole(ctx context.Context, name string, opt UpdateRamRoleOption) (*types.RamRole, error)
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
		policy := &types.AssumeRolePolicyDocument{}
		if err := json.Unmarshal([]byte(*role.AssumeRolePolicyDocument), policy); err != nil {
			log.Printf("unmarshal AssumeRolePolicyDocument failed: %+v: \n%s\n", *role.AssumeRolePolicyDocument, err)
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
		policy := &types.AssumeRolePolicyDocument{}
		if err := json.Unmarshal([]byte(*role.AssumeRolePolicyDocument), policy); err != nil {
			log.Printf("unmarshal AssumeRolePolicyDocument failed: %+v: \n%s\n", *role.AssumeRolePolicyDocument, err)
		}
		r.AssumeRolePolicyDocument = policy
	}
}
