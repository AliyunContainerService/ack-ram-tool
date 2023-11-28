package cleanupuserpermissions

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac/scanuserpermissions"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/briandowns/spinner"
	"time"

	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac/binding"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
)

type Option struct {
	userId uint64

	clusterId         string
	privateIpAddress  bool
	temporaryDuration time.Duration
	//outputFormat      OutputFormat
	allDeletedUsers bool
}

var opts = Option{
	temporaryDuration: time.Hour,
}

var cmd = &cobra.Command{
	Use:   "cleanup-user-permissions",
	Short: "cleanup RBAC permissions for one user or all deleted users",
	Long: `cleanup RBAC permissions for one user or all deleted users

Examples:
  # cleanup RBAC permissions for one cluster and one user
  ack-ram-tool rbac cleanup-user-permissions -c <clusterId> -u <uid>

  # cleanup RBAC permissions for one cluster and all deleted users
  ack-ram-tool rbac cleanup-user-permissions -c <clusterId> --all-deleted-users
`,
	Run: func(cmd *cobra.Command, args []string) {
		if !opts.allDeletedUsers && opts.userId == 0 {
			cmd.Help()
			log.Logger.Error("flag -u/--user-id/--all-deleted-users not set")
			return
		}

		ctx := ctlcommon.SetupSignalHandler(context.Background())
		run(ctx)
	},
}

func run(ctx context.Context) {
	openAPIClient := ctlcommon.GetClientOrDie()

	if err := cleanOneCluster(ctx, openAPIClient, opts.clusterId); err != nil {
		ctlcommon.ExitIfError(err)
	}
}

func cleanOneCluster(ctx context.Context, openAPIClient openapi.ClientInterface, clusterId string) error {
	log.Logger.Info("Start to scan users and bindings")
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Start()

	var kubeClient kubernetes.Interface
	var accounts map[int64]types.Account
	var bindings []binding.Binding
	var err error

	func() {
		defer spin.Stop()
		kubeClient, err = getKubeClient(ctx, openAPIClient, clusterId)
		if err != nil {
			return
		}
		accounts, err = binding.ListAccounts(ctx, openAPIClient)
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

	return cleanupOneCluster(ctx, bindings, accounts, kubeClient, clusterId)
}

func cleanupOneCluster(ctx context.Context, bindings []binding.Binding,
	accounts map[int64]types.Account, kube kubernetes.Interface, clusterId string) error {
	var newBindings []binding.Binding
	for _, b := range bindings {
		if b.AliUid == 0 {
			continue
		}
		if opts.userId != 0 && b.AliUid != int64(opts.userId) {
			continue
		}
		acc, ok := accounts[b.AliUid]
		if !ok {
			acc = types.NewFakeAccount(b.AliUid)
			acc.MarkDeleted()
			accounts[b.AliUid] = acc
		}
		if opts.allDeletedUsers && !acc.Deleted() {
			continue
		}
		newBindings = append(newBindings, b)
	}

	log.Logger.Info("will cleanup RBAC bindings as below:")
	scanuserpermissions.OutputBindingsTable(newBindings, accounts, false)

	ctlcommon.YesOrExit("Are you sure you want to cleanup these bindings?")

	for _, b := range newBindings {
		log.Logger.Infof("start to backup binding: %s", b.String())
		if p, err := binding.SaveBindingToFile(ctx, clusterId, b, kube); err != nil {
			return fmt.Errorf("backup binding %s: %w", b.String(), err)
		} else {
			log.Logger.Infof("the origin binding %s have been backed up to file %s", b.String(), p)
		}
		log.Logger.Infof("start to delete binding: %s", b.String())
		if err := binding.RemoveBinding(ctx, b, kube); err != nil {
			return fmt.Errorf("delete binding %s: %w", b.String(), err)
		}
		log.Logger.Infof("deleted binding: %s", b.String())
	}
	log.Logger.Info("all bindings have been cleanup")
	return nil
}

func getKubeClient(ctx context.Context, openAPIClient openapi.ClientInterface,
	clusterId string) (kubernetes.Interface, error) {
	kubeconfig, err := openAPIClient.GetUserKubeConfig(ctx, clusterId,
		opts.privateIpAddress, opts.temporaryDuration)
	if err != nil {
		return nil, fmt.Errorf("get kubeconfig: %w", err)
	}

	client, err := ctlcommon.NewKubeClient(kubeconfig.RawData)
	return client, err
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	cmd.Flags().Uint64VarP(&opts.userId, "user-id", "u", 0, "limit user id")
	cmd.Flags().StringVarP(&opts.clusterId, "cluster-id", "c", "", "cluster id")
	cmd.Flags().BoolVar(&opts.allDeletedUsers, "all-deleted-users", false, "cleanup all deleted users")
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("cluster-id"))
}
