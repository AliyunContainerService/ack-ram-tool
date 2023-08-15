package provider

import (
	"context"
	"fmt"
	"os"
	"time"
)

type FileProvider struct {
	u *Updater

	filepath string
	decoder  func(data []byte) (*Credentials, error)
}

type FileProviderOptions struct {
	RefreshPeriod time.Duration
	ExpiryWindow  time.Duration
}

func NewFileProvider(filepath string, decoder func(data []byte) (*Credentials, error), opts FileProviderOptions) *FileProvider {
	e := &FileProvider{
		filepath: filepath,
		decoder:  decoder,
	}
	e.u = NewUpdater(e.getCredentials, UpdaterOptions{
		ExpiryWindow:  opts.ExpiryWindow,
		RefreshPeriod: opts.RefreshPeriod,
	})
	e.u.Start(context.TODO())

	return e
}

func (f *FileProvider) Credentials(ctx context.Context) (*Credentials, error) {
	return f.u.Credentials(ctx)
}

func (f *FileProvider) getCredentials(ctx context.Context) (*Credentials, error) {
	data, err := os.ReadFile(f.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, NewNotEnableError(fmt.Errorf("read file %s failed: %w", f.filepath, err))
		}
		return nil, fmt.Errorf("read file %s failed: %w", f.filepath, err)
	}

	cred, err := f.decoder(data)
	if err != nil {
		return nil, fmt.Errorf("decode data from %s failed: %w", f.filepath, err)
	}
	return cred, nil
}
