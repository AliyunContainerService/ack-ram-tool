package cleanupuserpermissions

import (
	"context"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac/binding"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac/scanuserpermissions"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/briandowns/spinner"
	"k8s.io/client-go/kubernetes"
	"time"
)

func cleanAllClusters(ctx context.Context, openAPIClient openapi.ClientInterface) error {
	log.Logger.Info("Start to scan users and bindings for all clusters")
	clusters, accounts, err := scanuserpermissions.GetAllClustersAndAccountsWithSpin(ctx, openAPIClient)
	if err != nil {
		return err
	}
	for _, cluster := range clusters {
		clusterId := cluster.ClusterId
		log.Logger.Infof("---- %s (%s) ----", clusterId, cluster.Name)
		logger := log.Named(clusterId)
		clusterCtx := log.IntoContext(ctx, logger)
		logger.Infof("start to cleanup bindings and permissions for cluster %s", clusterId)
		if err := cleanOneClusterWithAccounts(clusterCtx, openAPIClient, clusterId, accounts); err != nil {
			logger.Errorf("cleanup bindings and permissions for cluster %s failed: %s", clusterId, err)
		}
	}
	return nil
}

func cleanOneClusterWithAccounts(ctx context.Context, openAPIClient openapi.ClientInterface,
	clusterId string, accounts map[int64]types.Account) error {
	logger := log.FromContext(ctx)
	logger.Info("Start to scan users and bindings")
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Start()

	var kubeClient kubernetes.Interface
	var bindings []binding.Binding
	var err error

	func() {
		defer spin.Stop()
		kubeClient, err = getKubeClient(ctx, openAPIClient, clusterId)
		if err != nil {
			return
		}
		bindings, err = scanuserpermissions.GetClusterBindings(ctx, kubeClient)
		if err != nil {
			return
		}
	}()
	if err != nil {
		return err
	}

	return cleanupOneCluster(ctx, bindings, accounts, kubeClient, openAPIClient, clusterId)
}
