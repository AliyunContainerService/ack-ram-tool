package utils

import "github.com/fatih/color"

func Bold(format string, a ...interface{}) string {
	return color.New(color.Bold).Sprintf(format, a...)
}

func Underline(format string, a ...interface{}) string {
	return color.New(color.Underline).Sprintf(format, a...)
}
