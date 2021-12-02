package rrsa

import (
	"fmt"
	"os"
)

func exitByError(msg string) {
	fmt.Println("error: " + msg)
	os.Exit(1)
}
