package common

import (
	"errors"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdlatest "k8s.io/client-go/tools/clientcmd/api/latest"
)

func NewKubeClient(yamlConfig string) (kubernetes.Interface, error) {
	config, err := newKubConfig(yamlConfig)
	if err != nil {
		return nil, err
	}
	config.UserAgent = version.UserAgent()
	// use protobuf
	config.AcceptContentTypes = "application/vnd.kubernetes.protobuf,application/json"
	config.ContentType = "application/vnd.kubernetes.protobuf"
	return kubernetes.NewForConfig(config)
}

func newKubConfig(yamlConfig string) (*restclient.Config, error) {
	config := clientcmdapi.NewConfig()
	// if there's no data in a file, return the default object instead of failing (DecodeInto reject empty input)
	if len(yamlConfig) == 0 {
		return nil, errors.New("kubeconfig is empty")
	}
	decoded, _, err := clientcmdlatest.Codec.Decode([]byte(yamlConfig),
		&schema.GroupVersionKind{Version: clientcmdlatest.Version, Kind: "Config"},
		config)
	if err != nil {
		return nil, err
	}
	if cfg, ok := decoded.(*clientcmdapi.Config); ok {
		c := clientcmd.NewDefaultClientConfig(*cfg, nil)
		return c.ClientConfig()
	}
	return nil, errors.New("load kubeconfig error")
}
