package rrsa

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/spf13/cobra"
)

type DemoOpts struct {
	noLoop bool
}

var demoOpts = DemoOpts{}

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "A demo for using RRSA Token in ACK Cluster when running it as pod container",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sleep := time.Second * 30
		for {
			log.Println("======= [begin] list ACK clusters with RRSA =======")
			client := common.GetClientOrDie()
			cs, err := client.ListClusters(context.Background())
			if err != nil {
				if demoOpts.noLoop {
					common.ExitByError(err.Error())
				} else {
					log.Println(err)
				}
			} else {
				fmt.Println("clusters:")
				for _, c := range cs {
					fmt.Printf("cluster id: %s, cluster name: %s\n", c.ClusterId, c.Name)
				}
				log.Println("======= [end]   list ACK clusters with RRSA =======")
				if demoOpts.noLoop {
					break
				}
			}
			log.Println("will try again after 30 seconds")
			time.Sleep(sleep)
		}
	},
}

func setupDemoCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demoCmd)
	demoCmd.Flags().BoolVar(&demoOpts.noLoop, "no-loop", false, "")
}
