package types

import "time"

type KubeConfig struct {
	Kind           string         `json:"kind,omitempty" yaml:"kind,omitempty"`
	APIVersion     string         `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	Clusters       []KubeCluster  `json:"clusters" yaml:"clusters"`
	Contexts       []KubeContext  `json:"contexts" yaml:"contexts"`
	CurrentContext string         `json:"current-context" yaml:"current-context"`
	Users          []KubeAuthUser `json:"users" yaml:"users"`
}

type KubeCluster struct {
	Name    string          `json:"name" yaml:"name"`
	Cluster KubeClusterInfo `json:"cluster" yaml:"cluster"`
}

type KubeClusterInfo struct {
	Server string `json:"server" yaml:"server"`
	//TLSServerName            string `json:"tls-server-name,omitempty" yaml:"tls-server-name,omitempty"`
	//InsecureSkipTLSVerify    bool   `json:"insecure-skip-tls-verify,omitempty" yaml:"insecure-skip-tls-verify,omitempty"`
	//CertificateAuthority     string `json:"certificate-authority,omitempty" yaml:"certificate-authority,omitempty"`
	CertificateAuthorityData string `json:"certificate-authority-data,omitempty" yaml:"certificate-authority-data,omitempty"`
	//ProxyURL                 string `json:"proxy-url,omitempty" yaml:"proxy-url,omitempty"`
}

type KubeAuthUser struct {
	Name string       `json:"name" yaml:"name"`
	User KubeAuthInfo `json:"user" yaml:"user"`
}

type KubeAuthInfo struct {
	//ClientCertificate     string `json:"client-certificate,omitempty" yaml:"client-certificate,omitempty"`
	ClientCertificateData string `json:"client-certificate-data,omitempty" yaml:"client-certificate-data,omitempty"`
	//ClientKey             string `json:"client-key,omitempty" yaml:"client-key,omitempty"`
	ClientKeyData string `json:"client-key-data,omitempty" yaml:"client-key-data,omitempty"`
	//Token                 string              `json:"token,omitempty" yaml:"token,omitempty"`
	//TokenFile             string              `json:"tokenFile,omitempty" yaml:"tokenFile,omitempty"`
	//Impersonate           string              `json:"act-as,omitempty" yaml:"act-as,omitempty"`
	//ImpersonateGroups     []string            `json:"act-as-groups,omitempty" yaml:"act-as-groups,omitempty"`
	//ImpersonateUserExtra  map[string][]string `json:"act-as-user-extra,omitempty" yaml:"act-as-user-extra,omitempty"`
	//Username              string              `json:"username,omitempty" yaml:"username,omitempty"`
	//Password              string              `json:"password,omitempty" yaml:"password,omitempty"`
	//AuthProvider          *KubeAuthProviderConfig `json:"auth-provider,omitempty" yaml:"auth-provider,omitempty"`
}

type KubeContext struct {
	Name    string          `json:"name" yaml:"name"`
	Context KubeContextInfo `json:"context" yaml:"context"`
}

type KubeContextInfo struct {
	Cluster string `json:"cluster" yaml:"cluster"`
	User    string `json:"user" yaml:"user"`
	//Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}

//type KubeAuthProviderConfig struct {
//	Name   string            `json:"name" yaml:"name"`
//	Config map[string]string `json:"config,omitempty" yaml:"config,omitempty"`
//}

type ExecCredential struct {
	KubeTypeMeta `json:",inline" yaml:",inline"`
	Spec         ExecCredentialSpec    `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status       *ExecCredentialStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type ExecCredentialSpec struct {
	Cluster     *ExecCluster `json:"cluster,omitempty" yaml:"cluster,omitempty"`
	Interactive bool         `json:"interactive" yaml:"interactive"`
}

type ExecCredentialStatus struct {
	ExpirationTimestamp   *KubeTime `json:"expirationTimestamp,omitempty" yaml:"expirationTimestamp,omitempty"`
	Token                 string    `json:"token,omitempty" yaml:"token,omitempty"`
	ClientCertificateData string    `json:"clientCertificateData,omitempty" yaml:"clientCertificateData,omitempty"`
	ClientKeyData         string    `json:"clientKeyData,omitempty" yaml:"clientKeyData,omitempty"`
}

type ExecCluster struct {
	Server                   string      `json:"server" yaml:"server"`
	TLSServerName            string      `json:"tls-server-name,omitempty" yaml:"tls-server-name,omitempty"`
	InsecureSkipTLSVerify    bool        `json:"insecure-skip-tls-verify,omitempty" yaml:"insecure-skip-tls-verify,omitempty"`
	CertificateAuthorityData string      `json:"certificate-authority-data,omitempty" yaml:"certificate-authority-data,omitempty"`
	ProxyURL                 string      `json:"proxy-url,omitempty" yaml:"proxy-url,omitempty"`
	DisableCompression       bool        `json:"disable-compression,omitempty" yaml:"disable-compression,omitempty"`
	Config                   interface{} `json:"config,omitempty" yaml:"config,omitempty"`
}

type KubeTypeMeta struct {
	Kind       string `json:"kind,omitempty" yaml:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
}

type KubeTime struct {
	time.Time
}

func (t KubeTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		// Encode unset/nil objects as JSON's "null".
		return []byte("null"), nil
	}
	buf := make([]byte, 0, len(time.RFC3339)+2)
	buf = append(buf, '"')
	// time cannot contain non escapable JSON characters
	buf = t.UTC().AppendFormat(buf, time.RFC3339)
	buf = append(buf, '"')
	return buf, nil
}

func (t KubeTime) MarshalYAML() (interface{}, error) {
	if t.IsZero() {
		return nil, nil
	}
	buf := make([]byte, 0, len(time.RFC3339)+2)
	// time cannot contain non escapable JSON characters
	buf = t.UTC().AppendFormat(buf, time.RFC3339)
	return string(buf), nil
}
