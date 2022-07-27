package rrsa

import (
	"log"
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
	log.Println("error: " + msg)
	os.Exit(1)
}
