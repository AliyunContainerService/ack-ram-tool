package common

import (
	"fmt"
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
	suffix := "get credential failed, more details about credential: https://github.com/AliyunContainerService/ack-ram-tool#credential"
	if strings.Contains(msg, "init client failed: ERROR: Can not open file open") {
		msg = fmt.Sprintf("%s. %s", msg, suffix)
	} else if strings.Contains(msg, "init client failed: No credential found") {
		msg = fmt.Sprintf("%s. %s", msg, suffix)
	}
	log.Println("error: " + msg)
	os.Exit(1)
}
