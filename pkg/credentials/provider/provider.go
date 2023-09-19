package provider

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
)

var UserAgent = ""

type CredentialsProvider interface {
	Credentials(ctx context.Context) (*Credentials, error)
}

type Stopper interface {
	Stop(ctx context.Context)
}

func init() {
	name := path.Base(os.Args[0])
	UserAgent = fmt.Sprintf("%s %s/%s ack-ram-tool/%s", name, runtime.GOOS, runtime.GOARCH, runtime.Version())
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
