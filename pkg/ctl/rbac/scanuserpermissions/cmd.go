package scanuserpermissions

import (
	"context"
	"fmt"
	"time"

	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
)

// scan rbac permissions bound to alibaba cloud users
type OutputFormat string

type Option struct {
	userId uint64

	clusterId         string
	privateIpAddress  bool
	temporaryDuration time.Duration
	outputFormat      OutputFormat
}

var opts = Option{}

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
	kubeClient := getKubeClient(ctx, openAPIClient)

	rawBindings, err := listBindings(ctx, kubeClient)
	ctlcommon.ExitIfError(err)

	bindings := rawBindings.SortByUid()
	for _, b := range bindings {
		if b.AliUid == 0 {
			continue
		}
		fmt.Printf("UID: %-20d\t KIND: %s\t NAME: %s\n", b.AliUid, b.Kind, b.Name)
	}
}

func getKubeClient(ctx context.Context, openAPIClient *openapi.Client) kubernetes.Interface {
	clusterId := opts.clusterId
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
	err := cmd.MarkFlagRequired("cluster-id")
	ctlcommon.ExitIfError(err)
}
