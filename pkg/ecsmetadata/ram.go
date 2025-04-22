package ecsmetadata

import (
	"context"
)

func (c *Client) GetRoleName(ctx context.Context) (string, error) {
	if c.roleName != "" {
		return c.roleName, nil
	}
	return c.getTidyStringData(ctx, "/latest/meta-data/ram/security-credentials/")
}
