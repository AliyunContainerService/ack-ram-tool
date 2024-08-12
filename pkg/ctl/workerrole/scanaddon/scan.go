package scanaddon

import (
	"context"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/briandowns/spinner"
	"go.uber.org/zap/zapcore"
	"time"
)

var scanNamespaces = []string{
	"kube-system",
	"mse-ingress-controller",
	"arms-prom",
	"ack-onepilot",
	"csdr",
	"fluid-system",
	"kruise-daemon-config",
	"kruise-system",
	"kube-ai",
	"kube-queue",
	"aliyun-acr-acceleration",
	"arena-system",
	"argo",
}

func scan(ctx context.Context,
	openapiClient openapi.ClientInterface, opts Option) error {
	kube, err := getKubeClient(ctx, openapiClient, opts.clusterId, opts.kubeconfigPath)
	if err != nil {
		return err
	}

	scanner := NewClusterScanner(openapiClient, kube, opts.clusterId)
	return scanner.Scan(ctx)
}

func newSpinner() func() {
	if log.Logger.Level() > zapcore.DebugLevel {
		spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		spin.Start()
		return spin.Stop
	}
	return func() {}
}
