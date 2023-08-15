package provider

import (
	"context"
	"fmt"
	"os"
	"runtime"
)

var UserAgent = ""

type CredentialsProvider interface {
	Credentials(ctx context.Context) (*Credentials, error)
}

func init() {
	UserAgent = fmt.Sprintf("%s %s/%s %s", os.Args[0], runtime.GOOS, runtime.GOARCH, runtime.Version())
}

type NotEnableError struct {
	err error
}

func NewNotEnableError(err error) *NotEnableError {
	return &NotEnableError{err: err}
}
func (e NotEnableError) Error() string {
	return fmt.Sprintf("this provider is not enabled: %s", e.err.Error())
}
