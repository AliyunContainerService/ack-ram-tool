package rrsa

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
)

func getClientOrDie() *openapi.Client {
	client, err := NewClient(ctl.GlobalOption.Region)
	if err != nil {
		exitByError(fmt.Sprintf("init client failed: %+v", err))
	}
	return client
}

func yesOrExit(msg string) {
	if ctl.GlobalOption.AssumeYes {
		return
	}
	var promptRet bool
	prompt := &survey.Confirm{
		Message: msg,
	}
	_ = survey.AskOne(prompt, &promptRet)
	if !promptRet {
		fmt.Println("Canceled! Bye bye~")
		os.Exit(0)
	}
}

func allowRRSAFeatureOrDie(ctx context.Context, clusterId string, client *openapi.Client) *types.Cluster {
	c, err := getRRSAStatus(ctx, clusterId, client)
	if err != nil {
		exitByError(fmt.Sprintf("get status failed: %+v", err))
	}
	if c.State != types.ClusterStateRunning {
		exitByError(fmt.Sprintf("cluster state is not running: %s", c.State))
	}
	if c.ClusterType != types.ClusterTypeManagedKubernetes {
		exitByError("only support managed cluster")
	}
	return c
}

func waitClusterUpdateFinished(ctx context.Context, clusterId, taskId string, client openapi.CSClientInterface) error {
	n := int64(1)
	var taskSuccess bool
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		// check task status first
		if !taskSuccess {
			c, err := client.GetTask(ctx, taskId)
			if err != nil {
				return err
			}
			if c.State.IsNotSuccess() {
				return fmt.Errorf("task not success: %s, message: %s",
					c.State, getRRSAFailMessage(ctx, clusterId, client),
				)
			}
			if c.State == types.ClusterTaskStateSuccess {
				taskSuccess = true
				n = 1
			}
		}
		// if task successes, then check cluster status
		if taskSuccess {
			c, err := client.GetCluster(ctx, clusterId)
			if err != nil {
				return err
			}
			if c.State.IsRunning() {
				return nil
			}
		}
		jitter := time.Duration(rand.Int63n(int64(time.Second) * n))
		if jitter > time.Second*15 {
			jitter = time.Second*15 + time.Duration(rand.Int63n(int64(time.Second)*10))
		}
		time.Sleep(time.Second*20 + jitter)
		n++
	}
}

func getRRSAFailMessage(ctx context.Context, clusterId string, client openapi.CSClientInterface) string {
	logs, err := client.GetRecentClusterLogs(ctx, clusterId)
	if err != nil {
		// TODO: xxx
		return ""
	}
	max := 20
	n := 0
	for _, log := range logs {
		n++
		if n >= max {
			break
		}
		if !strings.Contains(log.Log, "Failed") {
			continue
		}
		if !strings.Contains(log.Log, "RRSA") {
			continue
		}
		return log.Log
	}
	return ""
}
