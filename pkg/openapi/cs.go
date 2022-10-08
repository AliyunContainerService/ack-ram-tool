package openapi

import (
	"context"
	"encoding/base64"
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
	GetUserKubeConfig(ctx context.Context, clusterId string, privateIpAddress bool, temporaryDuration time.Duration) (*types.ClusterKubeConfig, error)
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

func (c *Client) GetUserKubeConfig(ctx context.Context, clusterId string,
	privateIpAddress bool, temporaryDuration time.Duration) (*types.ClusterKubeConfig, error) {
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

	ret := &types.ClusterKubeConfig{}
	if err := convertDescribeClusterUserKubeconfigResponse(ret, resp); err != nil {
		return nil, err
	}
	return ret, nil
}

func convertDescribeClusterUserKubeconfigResponse(c *types.ClusterKubeConfig, resp *cs.DescribeClusterUserKubeconfigResponse) error {
	body := resp.Body
	if body == nil {
		return nil
	}

	exp := tea.StringValue(body.Expiration)
	expT, err := time.Parse(time.RFC3339, exp)
	if err != nil {
		return err
	}
	c.Expiration = expT

	rawConf := tea.StringValue(body.Config)
	kubeconfig := types.KubeConfig{}
	if err := yaml.Unmarshal([]byte(rawConf), &kubeconfig); err != nil {
		return err
	}

	c.Server = kubeconfig.Clusters[0].Cluster.Server
	ca, err := base64.StdEncoding.DecodeString(kubeconfig.Clusters[0].Cluster.CertificateAuthorityData)
	if err != nil {
		return err
	}
	c.CertificateAuthorityData = string(ca)

	cd, err := base64.StdEncoding.DecodeString(kubeconfig.Users[0].User.ClientCertificateData)
	if err != nil {
		return err
	}
	c.ClientCertificateData = string(cd)

	ck, err := base64.StdEncoding.DecodeString(kubeconfig.Users[0].User.ClientKeyData)
	if err != nil {
		return err
	}
	c.ClientKeyData = string(ck)
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
