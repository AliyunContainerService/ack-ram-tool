package rrsa

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/spf13/cobra"
)

type DemoOpts struct {
	noLoop bool
	region string
}

var demoOpts = DemoOpts{
	region: "cn-hangzhou",
}

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "A demo for using RRSA Token in ACK Cluster when running it as pod container",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if demoOpts.region != "" {
			ctl.GlobalOption.Region = demoOpts.region
		}
		sleep := time.Second * 30
		for {
			log.Logger.Info("======= [begin] list ACK clusters with RRSA =======")
			client := common.GetClientOrDie()
			cs, err := client.ListClustersForRegion(context.Background(), demoOpts.region)
			if err != nil {
				if demoOpts.noLoop {
					common.ExitByError(err.Error())
				} else {
					log.Logger.Error(err)
				}
			} else {
				fmt.Println("clusters:")
				for _, c := range cs {
					fmt.Printf("cluster id: %s, cluster name: %s\n", c.ClusterId, c.Name)
				}
				log.Logger.Info("======= [end]   list ACK clusters with RRSA =======")
				if demoOpts.noLoop {
					break
				}
			}
			log.Logger.Info("will try again after 30 seconds")
			time.Sleep(sleep)
		}
	},
}

func setupDemoCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demoCmd)
	demoCmd.Flags().BoolVar(&demoOpts.noLoop, "no-loop", false, "")
	demoCmd.Flags().StringVar(&demoOpts.region, "region", demoOpts.region, "")
}
