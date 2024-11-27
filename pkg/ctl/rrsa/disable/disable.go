package disable

import (
	"context"
	"fmt"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable RRSA feature",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()
		ctlcommon.YesOrExit("Are you sure you want to disable RRSA feature?")
		ctx := context.Background()

		clusterId := ctl.GlobalOption.ClusterId
		c := common.AllowRRSAFeatureOrDie(ctx, clusterId, client)
		if !c.MetaData.RRSAConfig.Enabled {
			log.Logger.Info("RRSA feature is already disabled")
			return
		}

		var task *types.ClusterTask
		var err error
		log.Logger.Info("Start to disable RRSA feature")
		spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		spin.Start()
		if task, err = disableRRSA(ctx, clusterId, client); err != nil {
			spin.Stop()
			ctlcommon.ExitByError(fmt.Sprintf("Failed to disable RRSA feature for cluster %s: %+v", clusterId, err))
		}
		time.Sleep(time.Second * 30)
		ctx, cancel := context.WithTimeout(ctx, time.Minute*15)
		defer cancel()
		if err := common.WaitClusterUpdateFinished(ctx, clusterId, task.TaskId, client); err != nil {
			spin.Stop()
			ctlcommon.ExitByError(fmt.Sprintf("Failed to disable RRSA feature for cluster %s: %+v", clusterId, err))
		}
		spin.Stop()
		log.Logger.Infof("Disable RRSA feature for cluster %s successfully\n", clusterId)
	},
}

func disableRRSA(ctx context.Context, clusterId string, client openapi.CSClientInterface) (*types.ClusterTask, error) {
	boolValue := false
	return client.UpdateCluster(ctx, clusterId, openapi.UpdateClusterOption{EnableRRSA: &boolValue})
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	ctlcommon.SetupClusterIdFlag(cmd)
}
