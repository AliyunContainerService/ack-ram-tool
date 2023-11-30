package scanuserpermissions

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
	"strings"
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

const allClustersFlag = "all"

var opts = Option{
	temporaryDuration: time.Hour,
}

var cmd = &cobra.Command{
	Use:   "scan-user-permissions",
	Short: "scan RBAC permissions for one or all users",
	Long: `scan RBAC permissions for one or all users

Examples:
  # list all deleted users/roles for one cluster
  ack-ram-tool rbac scan-user-permissions -c <clusterId>

  # list all users/roles for one cluster
  ack-ram-tool rbac scan-user-permissions -c <clusterId> --all-users

  # list all deleted users/roles for all clusters
  ack-ram-tool rbac scan-user-permissions -c all

`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := ctlcommon.SetupSignalHandler(context.Background())
		run(ctx)
	},
}

func run(ctx context.Context) {
	openAPIClient := ctlcommon.GetClientOrDie()

	if opts.clusterId == allClustersFlag {
		err := scanAllClusters(ctx, openAPIClient)
		ctlcommon.ExitIfError(err)
	} else {
		err := scanOneCluster(ctx, openAPIClient, opts.clusterId)
		ctlcommon.ExitIfError(err)
	}
}

func scanOneCluster(ctx context.Context, openAPIClient openapi.ClientInterface, clusterId string) error {
	logger := log.FromContext(ctx)
	logger.Infof("Start to scan users and bindings for cluster %s", clusterId)
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
		bindings, err = GetClusterBindings(ctx, kubeClient)
		if err != nil {
			return
		}
	}()
	if err != nil {
		return err
	}

	outputTable(bindings, accounts)
	return nil
}

func GetClusterBindings(ctx context.Context, kubeClient kubernetes.Interface) ([]binding.Binding, error) {
	rawBindings, err := binding.ListBindings(ctx, kubeClient)
	if err != nil {
		return nil, err
	}
	bindings := rawBindings.SortByUid()
	return bindings, nil
}

func outputTable(bindings []binding.Binding, accounts map[int64]types.Account) {
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
		if !acc.Deleted() && !opts.allUsers {
			continue
		}
		newBindings = append(newBindings, b)
	}

	OutputBindingsTable(newBindings, accounts, true)
}

func OutputBindingsTable(bindings []binding.Binding, accounts map[int64]types.Account,
	highlightDeletedUsers bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"UID", "UserType", "UserName", "Binding"})
	//table.SetAutoMergeCells(true)
	table.SetAutoWrapText(true)
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.EnableBorder(false)
	table.SetTablePadding("  ")
	table.SetNoWhiteSpace(true)

	redColor := tablewriter.Colors{
		tablewriter.Bold,
		tablewriter.FgRedColor,
	}
	redColors := []tablewriter.Colors{
		redColor,
	}

	sort.Slice(bindings, func(i, j int) bool {
		ai := accounts[bindings[i].AliUid]
		bi := accounts[bindings[i].AliUid]
		if ai.Deleted() {
			return true
		}
		if bi.Deleted() {
			return false
		}
		return strings.Compare(ai.Id(), bi.Id()) == -1
	})

	for _, b := range bindings {
		acc := accounts[b.AliUid]

		var userComment string
		if acc.Deleted() {
			userComment = " (deleted)"
		}

		data := []string{
			fmt.Sprintf("%d%s", b.AliUid, userComment),
			string(acc.Type),
			acc.Name(),
			b.String(),
		}
		if acc.Deleted() && highlightDeletedUsers {
			table.Rich(data, redColors)
		} else {
			table.Append(data)
		}
	}

	table.Render()
}

func getKubeClient(ctx context.Context, openAPIClient openapi.ClientInterface,
	clusterId string) (kubernetes.Interface, error) {
	kubeconfig, err := openAPIClient.GetUserKubeConfig(ctx, clusterId,
		opts.privateIpAddress, opts.temporaryDuration)
	if err != nil {
		return nil, err
	}

	client, err := ctlcommon.NewKubeClient(kubeconfig.RawData)
	return client, err
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	cmd.Flags().Uint64VarP(&opts.userId, "user-id", "u", 0, "limit user id")
	cmd.Flags().StringVarP(&opts.clusterId, "cluster-id", "c", "", "cluster id")
	cmd.Flags().BoolVarP(&opts.allUsers, "all-users", "A", false, "list all users")
	err := cmd.MarkFlagRequired("cluster-id")
	ctlcommon.ExitIfError(err)
}
