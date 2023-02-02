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
		client := common.GetClientOrDie()
		for {
			log.Println("======= [begin] list ACK clusters with RRSA =======")
			cs, err := client.ListClusters(context.Background())
			if demoOpts.noLoop {
				common.ExitIfError(err)
			} else if err != nil {
				log.Println(err)
			}

			fmt.Println("clusters:")
			for _, c := range cs {
				fmt.Printf("cluster id: %s, cluster name: %s\n", c.ClusterId, c.Name)
			}
			log.Println("======= [end]   list ACK clusters with RRSA =======")
			if demoOpts.noLoop {
				break
			}
			log.Println("will try again after 30 seconds")
			time.Sleep(time.Second * 30)
		}
	},
}

func setupDemoCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demoCmd)
	demoCmd.Flags().BoolVar(&demoOpts.noLoop, "no-loop", false, "")
}
