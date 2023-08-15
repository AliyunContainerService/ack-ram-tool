package provider

import "log"

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(err error, msg string)
}

var defaultLog Logger = defaultLogger{}

type defaultLogger struct {
}

func (d defaultLogger) Info(msg string) {
	log.Print(msg)
}

func (d defaultLogger) Debug(msg string) {
	// log.Print(msg)
}

func (d defaultLogger) Error(err error, msg string) {
	log.Print(msg)
}
