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
)

type UpdateClusterOption struct {
	EnableRRSA *bool
}

type CSClientInterface interface {
	GetCluster(ctx context.Context, clusterId string) (*types.Cluster, error)
	GetRecentClusterLogs(ctx context.Context, clusterId string) ([]types.ClusterLog, error)
	UpdateCluster(ctx context.Context, clusterId string, opt UpdateClusterOption) (*types.ClusterTask, error)
	GetTask(ctx context.Context, taskId string) (*types.ClusterTask, error)
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
