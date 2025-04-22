package ecsmetadata

import (
	"context"
)

func (c *Client) GetVpcId(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/vpc-id")
}

func (c *Client) GetVpcCidrBlockId(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/vpc-cidr-block")
}

func (c *Client) GetVSwitchId(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/vswitch-id")
}

func (c *Client) GetVSwitchCidrBlockId(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/vswitch-cidr-block")
}

func (c *Client) GetPrivateIPV4(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/private-ipv4")
}

func (c *Client) GetPublicIPV4(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/public-ipv4")
}

func (c *Client) GetEIPV4(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/eipv4")
}

func (c *Client) GetNetworkType(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/network-type")
}

func (c *Client) GetMAC(ctx context.Context) (string, error) {
	return c.getTidyStringData(ctx, "/latest/meta-data/mac")
}
