package scanaddon

import (
	"context"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
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
	kube, err := getKubeClient(ctx, openapiClient, opts.clusterId)
	if err != nil {
		return err
	}

	scanner := NewClusterScanner(openapiClient, kube, opts.clusterId)
	return scanner.Scan(ctx)
}
