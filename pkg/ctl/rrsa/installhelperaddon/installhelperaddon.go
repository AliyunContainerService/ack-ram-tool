package installhelperaddon

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

const addonName = "ack-pod-identity-webhook"

var errAlreadyInstalled = errors.New("this addon is already installed")

var cmd = &cobra.Command{
	Use:   "install-helper-addon",
	Short: "Install ack-pod-identity-webhook to the cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()
		ctlcommon.YesOrExit("Are you sure you want to install ack-pod-identity-webhook?")
		ctx := context.Background()

		clusterId := ctl.GlobalOption.ClusterId
		c := common.AllowRRSAFeatureOrDie(ctx, clusterId, client)
		if !c.MetaData.RRSAConfig.Enabled {
			ctlcommon.ExitByError("RRSA feature is not enabled.")
			return
		}

		var err error
		log.Printf("Start to install %s", addonName)
		addon := types.ClusterAddon{Name: addonName}
		spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		spin.Start()

		if err = installAddon(ctx, clusterId, addon, client); err != nil {
			spin.Stop()
			if err == errAlreadyInstalled {
				log.Printf("Install %s for cluster %s successfully", addonName, clusterId)
				return
			}
			ctlcommon.ExitByError(
				fmt.Sprintf("Failed to install %s for cluster %s: %+v", addonName, clusterId, err))
		}

		ctx, cancel := context.WithTimeout(ctx, time.Minute*15)
		defer cancel()
		if err := common.WaitAddonActionFinished(ctx, clusterId, addon, client); err != nil {
			spin.Stop()
			ctlcommon.ExitByError(
				fmt.Sprintf("Failed to install %s for cluster %s: %+v", addonName, clusterId, err))
		}

		spin.Stop()
		log.Printf("Install %s for cluster %s successfully", addonName, clusterId)
	},
}

func installAddon(ctx context.Context, clusterId string, addon types.ClusterAddon,
	client openapi.CSClientInterface) error {
	err, installed := checkAddonInstalled(ctx, clusterId, addon, client)
	if err != nil {
		return err
	}
	if installed {
		return errAlreadyInstalled
	}

	err = client.InstallAddon(ctx, clusterId, addon)
	if err != nil {
		if strings.Contains(err.Error(), "AddonAlreadyInstalled") {
			return errAlreadyInstalled
		}
	}
	return nil
}

func checkAddonInstalled(ctx context.Context, clusterId string, addon types.ClusterAddon,
	client openapi.CSClientInterface) (error, bool) {
	ret, err := client.GetAddonStatus(ctx, clusterId, addon.Name)
	if err != nil {
		return err, false
	}
	if ret == nil {
		return nil, false
	}
	return nil, ret.Installed()
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	ctlcommon.SetupClusterIdFlag(cmd)
}
