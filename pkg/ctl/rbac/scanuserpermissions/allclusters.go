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
	log.Logger.Info("Start to scan users and bindings for all clusters")
	clusters, accounts, err := getAllClustersAndAccountsWithSpin(ctx, openAPIClient)
	if err != nil {
		return err
	}
	for _, cluster := range clusters {
		clusterId := cluster.ClusterId
		log.Logger.Infof("---- %s (%s) ----", clusterId, cluster.Name)
		if err := scanOneClusterWithAccounts(ctx, openAPIClient, clusterId, accounts); err != nil {
			log.Logger.Errorf("scan bindings for cluster %s failed: %s", clusterId, err)
		}
	}
	return nil
}

func scanOneClusterWithAccounts(ctx context.Context, openAPIClient openapi.ClientInterface,
	clusterId string, accounts map[int64]types.Account) error {
	log.Logger.Infof("Start to scan bindings for cluster %s", clusterId)
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

	fmt.Printf("ClusterId: %s\n", clusterId)
	outputTable(bindings, accounts)
	return nil
}

func getAllClustersAndAccountsWithSpin(ctx context.Context,
	openAPIClient openapi.ClientInterface) ([]types.Cluster, map[int64]types.Account, error) {
	log.Logger.Info("Start to get all clusters, users and roles")
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
