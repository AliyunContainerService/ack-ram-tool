package common

import (
	"log"
	"os"
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
	log.Println("error: " + msg)
	os.Exit(1)
}
