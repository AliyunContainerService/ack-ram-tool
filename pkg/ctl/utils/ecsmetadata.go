package utils

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/spf13/cobra"
	"os"
	"time"
)

type ecsMetadataOpts struct {
	noLoop bool
}

var edOpts = ecsMetadataOpts{}

var testEcsMetadataCmd = &cobra.Command{
	Use: "test-ecs-metadata",
	Run: func(cmd *cobra.Command, args []string) {
		os.Setenv(ctl.EnvCredentialType, "imds") // #nosec G104
		client := ctlcommon.GetClientOrDie()
		ctx := ctlcommon.SetupSignalHandler(context.Background())
		sleep := time.Second * 30
		ticker := time.NewTicker(sleep)
		defer ticker.Stop()
		for {
			_, err := client.GetCallerIdentity(ctx)
			if err != nil {
				log.Logger.Error(err)
			} else {
				log.Logger.Info("success")
			}

			if edOpts.noLoop {
				return
			}
			select {
			case <-ctx.Done():
				fmt.Println("exit")
				return
			case <-ticker.C:
			}
		}
	},
}

func setupTestEdCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(testEcsMetadataCmd)
	testEcsMetadataCmd.Flags().BoolVar(&edOpts.noLoop, "no-loop", false, "")
}
