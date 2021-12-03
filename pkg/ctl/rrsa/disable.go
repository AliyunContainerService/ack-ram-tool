package rrsa

import (
	"context"
	"fmt"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable RRSA feature",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		yesOrExit("Are you sure you want to disable RRSA feature?")
		ctx := context.Background()
		c := allowRRSAFeatureOrDie(ctx, clusterId, client)
		if !c.MetaData.RRSAConfig.Enabled {
			fmt.Println("RRSA feature is already disabled. Skip to continue")
			return
		}

		var task *types.ClusterTask
		var err error
		spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		spin.Prefix = "Start to disable RRSA feature "
		spin.Start()
		if task, err = disableRRSA(ctx, clusterId, client); err != nil {
			spin.Stop()
			exitByError(fmt.Sprintf("Failed to disable RRSA feature for cluster %s: %+v", clusterId, err))
		}
		ctx, cancel := context.WithTimeout(ctx, time.Minute*15)
		defer cancel()
		if err := waitClusterUpdateFinished(ctx, clusterId, task.TaskId, client); err != nil {
			spin.Stop()
			exitByError(fmt.Sprintf("Failed to disable RRSA feature for cluster %s: %+v", clusterId, err))
		}
		spin.Stop()
		fmt.Printf("Disable RRSA feature for cluster %s successfully\n", clusterId)
	},
}

func disableRRSA(ctx context.Context, clusterId string, client openapi.CSClientInterface) (*types.ClusterTask, error) {
	boolValue := false
	return client.UpdateCluster(ctx, clusterId, openapi.UpdateClusterOption{EnableRRSA: &boolValue})
}

func setupDisableCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(disableCmd)
	disableCmd.Flags().StringVarP(&clusterId, "cluster-id", "c", "", "The cluster id to use")
	err := disableCmd.MarkFlagRequired("cluster-id")
	exitIfError(err)
}
