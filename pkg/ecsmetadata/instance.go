package ecsmetadata

import (
	"context"
)

func (c *Client) GetInstanceType(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/instance/instance-type")
}

func (c *Client) GetInstanceName(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/instance/instance-name")
}

func (c *Client) GetInstanceId(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/instance-id")
}
