package scanaddon

import (
	"context"
	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/spf13/cobra"
	"time"
)

type Option struct {
	clusterId         string
	privateIpAddress  bool
	temporaryDuration time.Duration
}

var opts = Option{
	temporaryDuration: time.Hour,
}

var cmd = &cobra.Command{
	Use:   "scan-addon",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := ctlcommon.SetupSignalHandler(context.Background())
		run(ctx)
	},
}

func run(ctx context.Context) {
	client := ctlcommon.GetClientOrDie()
	log.Logger.Infof("start to scan cluster %s", opts.clusterId)
	err := scan(ctx, client, opts)
	ctlcommon.ExitIfError(err)
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	cmd.Flags().StringVarP(&opts.clusterId, "cluster-id", "c", "", "cluster id")
	err := cmd.MarkFlagRequired("cluster-id")
	ctlcommon.ExitIfError(err)
}
