package rrsa

import (
	"fmt"
	"os"
)

func exitIfError(err error) {
	if err == nil {
		return
	}
	if err != nil {
		exitByError(err.Error())
	}
}

func exitByError(msg string) {
	fmt.Println("error: " + msg)
	os.Exit(1)
}
