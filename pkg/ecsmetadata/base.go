package ecsmetadata

import (
	"context"
)

func (c *Client) GetRegionId(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/region-id")
}

func (c *Client) GetZoneId(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/zone-id")
}

func (c *Client) GetOwnerAccountId(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/owner-account-id")
}

func (c *Client) GetHostname(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/hostname")
}

func (c *Client) GetSourceAddress(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/source-address")
}
