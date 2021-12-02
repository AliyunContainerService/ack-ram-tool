package types

import (
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
)

type ClusterType string
type ClusterState string
type ClusterTaskState string

var ClusterStateRunning ClusterState = "running"
var ClusterTypeManagedKubernetes ClusterType = "ManagedKubernetes"
var (
	ClusterTaskStateSuccess  ClusterTaskState = "success"
	ClusterTaskStateFail     ClusterTaskState = "fail"
	ClusterTaskStateTimeout  ClusterTaskState = "timeout"
	ClusterTaskStateCanceled ClusterTaskState = "canceled"
)

type Cluster struct {
	ClusterId   string
	ClusterType ClusterType
	MetaData    ClusterMetaData
	Name        string
	RegionId    string
	State       ClusterState
}

type ClusterMetaData struct {
	RRSAConfig RRSAConfig `json:"RRSAConfig"`
}

type RRSAConfig struct {
	Enabled bool `json:"enabled"`

	Issuer   string `json:"issuer"`
	Audience string `json:"audience"`

	OIDCName string `json:"oidc_name"`
	OIDCArn  string `json:"oidc_arn"`
}

type ClusterTask struct {
	TaskId string
	State  ClusterTaskState
	Error  interface{}
	Result interface{}
}

func (t ClusterTask) Err() string {
	return utils.ReplaceNewLine(string(utils.JSONValue(t.Error)))
}

func (s ClusterTaskState) IsNotSuccess() bool {
	return s == ClusterTaskStateFail || s == ClusterTaskStateTimeout || s == ClusterTaskStateCanceled
}

func (s ClusterState) IsRunning() bool {
	return s == ClusterStateRunning
}

type ClusterLog struct {
	Log     string
	Created time.Time
}
