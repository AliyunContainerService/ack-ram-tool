package cleanupuserpermissions

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rbac/scanuserpermissions"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"k8s.io/apimachinery/pkg/api/errors"
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

const allClustersFlag = "all"

var opts = Option{
	temporaryDuration: time.Hour,
}

var cmd = &cobra.Command{
	Use:   "cleanup-user-permissions",
	Short: "clean up RBAC permissions for one user or all deleted users",
	Long: `clean up RBAC permissions for one user or all deleted users

Examples:
  # clean up RBAC permissions for one cluster and one user
  ack-ram-tool rbac cleanup-user-permissions -c <clusterId> -u <uid>

  # clean up RBAC permissions for all cluster and one user
  ack-ram-tool rbac cleanup-user-permissions -c all -u <uid>
`,
	Run: func(cmd *cobra.Command, args []string) {
		//if !opts.allDeletedUsers && opts.userId == 0 {
		//	cmd.Help()
		//	log.Logger.Error("flag -u/--user-id not set")
		//	return
		//}

		ctx := ctlcommon.SetupSignalHandler(context.Background())
		run(ctx)
	},
}

func run(ctx context.Context) {
	openAPIClient := ctlcommon.GetClientOrDie()

	if opts.clusterId == allClustersFlag {
		if err := cleanAllClusters(ctx, openAPIClient); err != nil {
			ctlcommon.ExitIfError(err)
		}
	} else {
		if err := cleanOneCluster(ctx, openAPIClient, opts.clusterId); err != nil {
			ctlcommon.ExitIfError(err)
		}
	}
}

func cleanOneCluster(ctx context.Context, openAPIClient openapi.ClientInterface, clusterId string) error {
	logger := log.FromContext(ctx)
	logger.Info("start to scan users and bindings")
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

	return cleanupOneCluster(ctx, bindings, accounts, kubeClient, openAPIClient, clusterId, false)
}

func cleanupOneCluster(ctx context.Context, bindings []binding.Binding,
	accounts map[int64]types.Account, kube kubernetes.Interface,
	openAPIClient openapi.ClientInterface,
	clusterId string, allowSkip bool) error {
	var newBindings []binding.Binding
	var toCleanupUids []int64
	logger := log.FromContext(ctx)
	toCleanupUidsDup := make(map[int64]bool)

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
		if _, ok := toCleanupUidsDup[b.AliUid]; !ok {
			toCleanupUids = append(toCleanupUids, b.AliUid)
			toCleanupUidsDup[b.AliUid] = true
		}
	}
	for id, acc := range accounts {
		if opts.userId != 0 && id != int64(opts.userId) {
			continue
		}
		if opts.allDeletedUsers && !acc.Deleted() {
			continue
		}
		if _, ok := toCleanupUidsDup[id]; !ok {
			toCleanupUids = append(toCleanupUids, id)
			toCleanupUidsDup[id] = true
		}
	}

	logger.Warn("we will clean up RBAC bindings as follows:")
	scanuserpermissions.OutputBindingsTable(newBindings, accounts, false)

	logger.Warn("we will clean up kubeconfig permissions for users as follows:")
	for _, uid := range toCleanupUids {
		fmt.Printf("UID: %d\n", uid)
	}
	for _, uid := range toCleanupUids {
		logger.Infof("start to check cluster audit log for user %d", uid)
		resp, err := openAPIClient.DescribeUserClusterActivityState(ctx, clusterId, uid)
		if err != nil {
			logger.Errorf("check cluster audit log failed: %s", err)
		} else if resp.Active {
			warn := color.RedString("this user has been active in the past 7 days, and the last activity time was: %s", resp.LastLocalActivity())
			logger.Warnf("%s. You will find the relevant audit log details below:\nsls project: %s\nsls logstore: %s\nlast activity: %s (auditID: %s)",
				warn, resp.LogProjectName, resp.LogStoreName, resp.LastLocalActivity(), resp.LastAuditId)
		} else if !resp.Active {
			logger.Info("no activity has been found in the cluster audit log for this user in the past 7 days")
		}
	}

	confirm := "Are you sure you want to clean up these bindings and permissions?"
	if !allowSkip {
		ctlcommon.YesOrExit(confirm)
	} else if !ctlcommon.Yes(confirm) {
		logger.Warn("we will skip this cluster!")
		return nil
	}

	for _, b := range newBindings {
		if err := backupRBACBinding(ctx, b, kube, clusterId); err != nil {
			return err
		}
	}

	//ctlcommon.YesOrExit("Are you sure you want to clean up these permissions?")

	if err := removePermissions(ctx, openAPIClient, clusterId, toCleanupUids); err != nil {
		return err
	}

	logger.Info("all bindings and permissions have been cleaned up")
	return nil
}

func removePermissions(ctx context.Context, openAPIClient openapi.ClientInterface,
	clusterId string, uids []int64) error {
	logger := log.FromContext(ctx)
	for _, uid := range uids {
		logger.Infof("start to clean up kubeconfig permissions for uid %d", uid)
		if err := openAPIClient.CleanClusterUserPermissions(ctx, clusterId, uid, true); err != nil {
			return fmt.Errorf("clean up kubeconfig permissions for uid %d: %w", uid, err)
		}
		logger.Infof("finished clean up kubeconfig permissions for uid %d", uid)
	}
	return nil
}

func backupRBACBinding(ctx context.Context, b binding.Binding, kube kubernetes.Interface, clusterId string) error {
	logger := log.FromContext(ctx)
	logger.Infof("start to backup binding %s", b.String())
	if p, err := binding.SaveBindingToFile(ctx, clusterId, b, kube); err != nil {
		if errors.IsNotFound(err) {
			logger.Infof("skip binding %s which is not founds", b.String())
			return nil
		}
		return fmt.Errorf("backup binding %s: %w", b.String(), err)
	} else {
		logger.Infof("the origin binding %s have been backed up to file %s", b.String(), p)
	}
	return nil
}

func removeRBACBinding(ctx context.Context, b binding.Binding, kube kubernetes.Interface, clusterId string) error {
	logger := log.FromContext(ctx)
	logger.Infof("start to backup binding %s", b.String())
	if p, err := binding.SaveBindingToFile(ctx, clusterId, b, kube); err != nil {
		if errors.IsNotFound(err) {
			logger.Infof("skip binding %s which is not founds", b.String())
			return nil
		}
		return fmt.Errorf("backup binding %s: %w", b.String(), err)
	} else {
		logger.Infof("the origin binding %s have been backed up to file %s", b.String(), p)
	}

	logger.Infof("start to delete binding %s", b.String())
	if err := binding.RemoveBinding(ctx, b, kube); err != nil {
		return fmt.Errorf("delete binding %s: %w", b.String(), err)
	}
	logger.Infof("deleted binding %s", b.String())
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
	//cmd.Flags().BoolVar(&opts.allDeletedUsers, "all-deleted-users", false, "clean up all deleted users")
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("cluster-id"))
	ctlcommon.ExitIfError(cmd.MarkFlagRequired("user-id"))
}
