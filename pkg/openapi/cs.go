package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	cs "github.com/alibabacloud-go/cs-20151215/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"gopkg.in/yaml.v3"
)

type UpdateClusterOption struct {
	EnableRRSA *bool
}

type CSClientInterface interface {
	GetCluster(ctx context.Context, clusterId string) (*types.Cluster, error)
	GetRecentClusterLogs(ctx context.Context, clusterId string) ([]types.ClusterLog, error)
	UpdateCluster(ctx context.Context, clusterId string, opt UpdateClusterOption) (*types.ClusterTask, error)
	GetTask(ctx context.Context, taskId string) (*types.ClusterTask, error)
	GetUserKubeConfig(ctx context.Context, clusterId string, privateIpAddress bool, temporaryDuration time.Duration) (*types.KubeConfig, error)
	ListClusters(ctx context.Context) ([]types.Cluster, error)
	GetAddonMetaData(ctx context.Context, clusterId string, name string) (*types.ClusterAddon, error)
	GetAddonStatus(ctx context.Context, clusterId string, name string) (*types.ClusterAddon, error)
	InstallAddon(ctx context.Context, clusterId string, addon types.ClusterAddon) error
	ListAddons(ctx context.Context, clusterId string) ([]types.ClusterAddon, error)
}

func (c *Client) GetCluster(ctx context.Context, clusterId string) (*types.Cluster, error) {
	client := c.csClient
	resp, err := client.DescribeClusterDetail(&clusterId)
	if err != nil {
		return nil, err
	}
	cluster := &types.Cluster{}
	convertDescribeClusterDetailResponse(cluster, resp)
	return cluster, nil
}

func (c *Client) ListClusters(ctx context.Context) ([]types.Cluster, error) {
	client := c.csClient
	resp, err := client.DescribeClusters(&cs.DescribeClustersRequest{})
	if err != nil {
		return nil, err
	}

	return convertDescribeClustersResponse(resp), nil
}

func (c *Client) UpdateCluster(ctx context.Context, clusterId string, opt UpdateClusterOption) (*types.ClusterTask, error) {
	client := c.csClient
	req := &cs.ModifyClusterRequest{
		EnableRrsa: opt.EnableRRSA,
	}
	resp, err := client.ModifyCluster(tea.String(clusterId), req)
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, fmt.Errorf("parse body failed: %q", resp.String())
	}
	return &types.ClusterTask{TaskId: tea.StringValue(resp.Body.TaskId)}, nil
}

func (c *Client) GetTask(ctx context.Context, taskId string) (*types.ClusterTask, error) {
	client := c.csClient
	resp, err := client.DescribeTaskInfo(tea.String(taskId))
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, fmt.Errorf("parse body failed: %q", resp.String())
	}
	return &types.ClusterTask{
		TaskId: taskId,
		State:  types.ClusterTaskState(tea.StringValue(resp.Body.State)),
		// Error:  tea.StringValue(ret.Body.TaskError),
	}, nil
}

func convertDescribeClusterDetailResponse(c *types.Cluster, resp *cs.DescribeClusterDetailResponse) {
	body := resp.Body
	if body == nil {
		return
	}
	c.ClusterId = tea.StringValue(body.ClusterId)
	c.ClusterType = types.ClusterType(tea.StringValue(body.ClusterType))
	c.Name = tea.StringValue(body.Name)
	c.RegionId = tea.StringValue(body.RegionId)
	c.State = types.ClusterState(tea.StringValue(body.State))

	metadata := &types.ClusterMetaData{}
	_ = json.Unmarshal([]byte(tea.StringValue(body.MetaData)), metadata)
	c.MetaData = *metadata
}

func (c *Client) GetRecentClusterLogs(ctx context.Context, clusterId string) ([]types.ClusterLog, error) {
	client := c.csClient
	ret, err := client.DescribeClusterLogs(&clusterId)
	if err != nil {
		return nil, err
	}
	return convertDescribeClusterLogsResponse(ret), nil
}

func (c *Client) GetUserKubeConfig(ctx context.Context, clusterId string,
	privateIpAddress bool, temporaryDuration time.Duration) (*types.KubeConfig, error) {
	client := c.csClient
	req := &cs.DescribeClusterUserKubeconfigRequest{
		PrivateIpAddress:         nil,
		TemporaryDurationMinutes: nil,
	}
	if temporaryDuration != 0 {
		dm := int64(temporaryDuration / time.Minute)
		if dm < 15 || dm > 4320 {
			return nil, fmt.Errorf("temporaryDuration should > 15 minutes and < 3 days")
		}
		req.TemporaryDurationMinutes = &dm
	}
	if privateIpAddress {
		req.PrivateIpAddress = &privateIpAddress
	}
	resp, err := client.DescribeClusterUserKubeconfig(&clusterId, req)
	if err != nil {
		return nil, err
	}

	ret := &types.KubeConfig{}
	if err := convertDescribeClusterUserKubeconfigResponse(ret, resp); err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) GetAddonMetaData(ctx context.Context, clusterId string, name string) (*types.ClusterAddon, error) {
	client := c.csClient
	resp, err := client.DescribeClusterAddonMetadata(
		tea.String(clusterId), tea.String(name), nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body
	return &types.ClusterAddon{
		Name:        tea.StringValue(body.Name),
		Version:     tea.StringValue(body.Version),
		NextVersion: "",
	}, nil
}

func (c *Client) GetAddonStatus(ctx context.Context, clusterId string, name string) (*types.ClusterAddon, error) {
	addons, err := c.ListAddons(ctx, clusterId)
	if err != nil {
		return nil, err
	}
	for _, addon := range addons {
		addon := addon
		if addon.Name == name {
			return &addon, nil
		}
	}
	return nil, nil
}

func (c *Client) InstallAddon(ctx context.Context, clusterId string, addon types.ClusterAddon) error {
	client := c.csClient
	req := &cs.InstallClusterAddonsRequest{
		Body: []*cs.InstallClusterAddonsRequestBody{
			{
				Config:  nil,
				Name:    tea.String(addon.Name),
				Version: tea.String(addon.NextVersion),
			},
		},
	}
	_, err := client.InstallClusterAddons(tea.String(clusterId), req)
	return err
}

func (c *Client) ListAddons(ctx context.Context, clusterId string) ([]types.ClusterAddon, error) {
	client := c.csClient
	resp, err := client.DescribeClusterAddonsVersion(tea.String(clusterId))
	if err != nil {
		return nil, err
	}

	addons := convertDescribeClusterAddonsVersionResponse(resp)
	return addons, nil
}

func convertDescribeClusterAddonsVersionResponse(resp *cs.DescribeClusterAddonsVersionResponse) []types.ClusterAddon {
	body := resp.Body
	if body == nil {
		return nil
	}
	var addons []types.ClusterAddon

	for _, value := range body {
		jsonV, err := json.Marshal(value)
		if err != nil {
			continue
		}
		var addon types.ClusterAddon
		if err := json.Unmarshal(jsonV, &addon); err == nil {
			addons = append(addons, addon)
		}
	}
	return addons
}

func convertDescribeClusterUserKubeconfigResponse(kubeconfig *types.KubeConfig, resp *cs.DescribeClusterUserKubeconfigResponse) error {
	body := resp.Body
	if body == nil {
		return nil
	}

	exp := tea.StringValue(body.Expiration)
	expT, err := time.Parse(time.RFC3339, exp)
	if err != nil {
		return err
	}

	rawConf := tea.StringValue(body.Config)
	kubeconfig.RawData = rawConf
	if err := yaml.Unmarshal([]byte(rawConf), kubeconfig); err != nil {
		return err
	}
	kubeconfig.Expiration = expT
	return nil
}

func convertDescribeClusterLogsResponse(resp *cs.DescribeClusterLogsResponse) []types.ClusterLog {
	body := resp.Body
	if body == nil {
		return nil
	}
	var ret []types.ClusterLog
	for _, item := range body {
		prefix := fmt.Sprintf("%s | ", tea.StringValue(item.ClusterId))
		t, _ := time.Parse("2006-01-02T15:04:05+07:00", tea.StringValue(item.Created))
		ret = append(ret, types.ClusterLog{
			Log:     strings.TrimLeft(tea.StringValue(item.ClusterLog), prefix),
			Created: t,
		})
	}
	return ret
}

func convertDescribeClustersResponse(resp *cs.DescribeClustersResponse) []types.Cluster {
	body := resp.Body
	if body == nil {
		return nil
	}
	var clusters []types.Cluster
	for _, item := range body {
		c := types.Cluster{}
		c.ClusterId = tea.StringValue(item.ClusterId)
		c.ClusterType = types.ClusterType(tea.StringValue(item.ClusterType))
		c.Name = tea.StringValue(item.Name)
		c.RegionId = tea.StringValue(item.RegionId)
		c.State = types.ClusterState(tea.StringValue(item.State))

		metadata := &types.ClusterMetaData{}
		_ = json.Unmarshal([]byte(tea.StringValue(item.MetaData)), metadata)
		c.MetaData = *metadata
		clusters = append(clusters, c)
	}

	return clusters
}
