package types

import (
	"encoding/base64"
	"strings"
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

func (c RRSAConfig) TokenIssuer() string {
	issuers := strings.Split(c.Issuer, ",")
	if len(issuers) > 1 {
		return issuers[0]
	}
	return c.Issuer
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

type ClusterCredential struct {
	// server
	Server string
	// certificate-authority-data
	CertificateAuthorityData string
	// client-certificate-data
	ClientCertificateData string
	// client-key-data
	ClientKeyData string
	// expiration
	Expiration time.Time
}

func (c *ClusterCredential) LoadKubeConfig(conf *KubeConfig) error {
	c.Server = conf.Clusters[0].Cluster.Server

	ca, err := base64.StdEncoding.DecodeString(conf.Clusters[0].Cluster.CertificateAuthorityData)
	if err != nil {
		return err
	}
	c.CertificateAuthorityData = string(ca)

	cd, err := base64.StdEncoding.DecodeString(conf.Users[0].User.ClientCertificateData)
	if err != nil {
		return err
	}
	c.ClientCertificateData = string(cd)

	ck, err := base64.StdEncoding.DecodeString(conf.Users[0].User.ClientKeyData)
	if err != nil {
		return err
	}
	c.ClientKeyData = string(ck)

	c.Expiration = conf.Expiration
	return nil
}

type ClusterAddon struct {
	Name        string `json:"component_name"`
	Version     string `json:"version"`
	NextVersion string `json:"next_version"`
}

func (c ClusterAddon) Installed() bool {
	return c.Version != ""
}
