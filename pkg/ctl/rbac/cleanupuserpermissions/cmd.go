package cleanupuserpermissions

import (
	"context"
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
	allUsers bool
}

var opts = Option{
	temporaryDuration: time.Hour,
}

var cmd = &cobra.Command{
	Use:   "cleanup-user-permissions",
	Short: "cleanup RBAC permissions for one user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	ctx := context.Background()
	openAPIClient := ctlcommon.GetClientOrDie()

	oneCluster(ctx, openAPIClient, opts.clusterId)
}

func oneCluster(ctx context.Context, openAPIClient openapi.ClientInterface, clusterId string) {
	log.Logger.Info("Start to scan users and bindings")
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Start()

	kubeClient := getKubeClient(ctx, openAPIClient, clusterId)
	rawBindings, err := binding.ListBindings(ctx, kubeClient)
	ctlcommon.ExitIfError(err)
	accounts, err := binding.ListAccounts(ctx, openAPIClient)
	ctlcommon.ExitIfError(err)
	spin.Stop()

	bindings := rawBindings.SortByUid()
	cleanup(ctx, bindings, accounts, kubeClient, clusterId)
}

func cleanup(ctx context.Context, bindings []binding.Binding,
	accounts map[int64]types.Account, kube kubernetes.Interface, clusterId string) {
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
		newBindings = append(newBindings, b)
	}

	log.Logger.Info("will cleanup RBAC bindings as blow:")
	scanuserpermissions.OutputBindingsTable(newBindings, accounts, false)

	ctlcommon.YesOrExit("Are you sure you want to cleanup these bindings?")
	for _, b := range newBindings {
		log.Logger.Infof("start to backup binding: %s", b.String())
		if p, err := binding.SaveBindingToFile(ctx, clusterId, b, kube); err != nil {
			ctlcommon.ExitIfError(err)
		} else {
			log.Logger.Infof("the origin binding %s have been backed up to file %s", b.String(), p)
		}
		log.Logger.Infof("start to cleanup binding: %s", b.String())
		if err := binding.RemoveBinding(ctx, b, kube); err != nil {
			ctlcommon.ExitIfError(err)
		}
		log.Logger.Infof("finished cleanup binding: %s", b.String())
	}
	log.Logger.Info("all bindings have been cleanup")
}

func getKubeClient(ctx context.Context, openAPIClient openapi.ClientInterface, clusterId string) kubernetes.Interface {
	kubeconfig, err := openAPIClient.GetUserKubeConfig(ctx, clusterId,
		opts.privateIpAddress, opts.temporaryDuration)
	ctlcommon.ExitIfError(err)

	client, err := ctlcommon.NewKubeClient(kubeconfig.RawData)
	ctlcommon.ExitIfError(err)
	return client
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	cmd.Flags().Uint64VarP(&opts.userId, "user-id", "u", 0, "limit user id")
	cmd.Flags().StringVarP(&opts.clusterId, "cluster-id", "c", "", "cluster id")
	//cmd.Flags().BoolVarP(&opts.allUsers, "all-users", "A", false, "list all users")
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("cluster-id"))
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("user-id"))
}
