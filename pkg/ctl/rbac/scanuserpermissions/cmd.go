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

var opts = Option{
	temporaryDuration: time.Hour,
}

var cmd = &cobra.Command{
	Use:   "scan-user-permissions",
	Short: "scan RBAC permissions for one or all users",
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
	outputTable(bindings, accounts)
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
		redColor, redColor, redColor, redColor, redColor, redColor,
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
	cmd.Flags().BoolVarP(&opts.allUsers, "all-users", "A", false, "list all users")
	err := cmd.MarkFlagRequired("cluster-id")
	ctlcommon.ExitIfError(err)
}
