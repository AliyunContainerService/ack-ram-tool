package ecsmetadata

import (
	"context"
)

func (c *Client) GetUserData(ctx context.Context) (string, error) {
	data, err := c.getRawStringData(ctx, "/latest/user-data")
	if err != nil {
		return "", err
	}
	return data, nil
}
