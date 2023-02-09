package common

import (
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
)

func YesOrExit(msg string) {
	if ctl.GlobalOption.AssumeYes {
		return
	}
	var promptRet bool
	prompt := &survey.Confirm{
		Message: msg,
	}
	_ = survey.AskOne(prompt, &promptRet)
	if !promptRet {
		log.Println("Canceled! Bye bye~")
		os.Exit(0)
	}
}

func SetupClusterIdFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&ctl.GlobalOption.ClusterId, "cluster-id", "c", "", "The cluster id to use")
	ExitIfError(cmd.MarkFlagRequired("cluster-id"))
}
