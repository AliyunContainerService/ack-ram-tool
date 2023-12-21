package scanuserpermissions

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac/binding"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/briandowns/spinner"
	"k8s.io/client-go/kubernetes"
	"time"
)

func scanAllClusters(ctx context.Context, openAPIClient openapi.ClientInterface) error {
	log.Logger.Info("start to scan users and bindings for all clusters")
	clusters, accounts, err := GetAllClustersAndAccountsWithSpin(ctx, openAPIClient)
	if err != nil {
		return err
	}
	for _, cluster := range clusters {
		clusterId := cluster.ClusterId
		log.Logger.Infof("---- %s (%s) ----", clusterId, cluster.Name)
		logger := log.Named(clusterId)
		if cluster.State.NoActiveApiServer() {
			logger.Errorf("invalid cluster state (%s), skip it", cluster.State)
			continue
		}

		clusterCtx := log.IntoContext(ctx, logger)
		if err := scanOneClusterWithAccounts(clusterCtx, openAPIClient, clusterId, accounts); err != nil {
			logger.Errorf("scan bindings for cluster %s failed: %s", clusterId, err)
		}
	}
	return nil
}

func scanOneClusterWithAccounts(ctx context.Context, openAPIClient openapi.ClientInterface,
	clusterId string, accounts map[int64]types.Account) error {
	logger := log.FromContext(ctx)
	logger.Infof("start to scan bindings for cluster %s", clusterId)
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Start()

	var bindings []binding.Binding
	var err error

	func() {
		defer spin.Stop()
		var kubeClient kubernetes.Interface
		kubeClient, err = getKubeClient(ctx, openAPIClient, clusterId)
		if err != nil {
			return
		}
		bindings, err = GetClusterBindings(ctx, kubeClient)
		if err != nil {
			return
		}
	}()
	if err != nil {
		return err
	}

	if opts.userId == 0 && !opts.allUsers {
		logger.Warn("by default, only deleted users are included. Use the --all-users flag to include all users")
	}
	fmt.Printf("ClusterId: %s\n", clusterId)
	outputTable(bindings, accounts)
	return nil
}

func GetAllClustersAndAccountsWithSpin(ctx context.Context,
	openAPIClient openapi.ClientInterface) ([]types.Cluster, map[int64]types.Account, error) {
	log.Logger.Info("start to get all clusters, users and roles")
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Start()
	defer spin.Stop()

	return getAllClustersAndAccounts(ctx, openAPIClient)
}

func getAllClustersAndAccounts(ctx context.Context,
	openAPIClient openapi.ClientInterface) ([]types.Cluster, map[int64]types.Account, error) {
	clusters, err := openAPIClient.ListClustersV1(ctx)
	if err != nil {
		return nil, nil, err
	}
	accounts, err := binding.ListAccounts(ctx, openAPIClient)
	if err != nil {
		return nil, nil, err
	}
	return clusters, accounts, nil
}
