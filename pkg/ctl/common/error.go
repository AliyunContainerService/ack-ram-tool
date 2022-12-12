package common

import (
	"log"
	"os"
	"strings"
)

func ExitIfError(err error) {
	if err == nil {
		return
	}
	if err != nil {
		ExitByError(err.Error())
	}
}

func ExitByError(msg string) {
	if strings.Contains(msg, "init client failed: ERROR: Can not open file open") {
		msg = "get credential info failed. " + msg
		msg += ". more details about credential: https://github.com/AliyunContainerService/ack-ram-tool#credential"
	}
	log.Println("error: " + msg)
	os.Exit(1)
}
