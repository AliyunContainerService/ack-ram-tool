package ecsmetadata

import (
	"context"
	"errors"
	"fmt"
	"time"
)

const defaultMaxRetryTimes = 3

type RetryOptions struct {
	MaxRetryTimes  int
	RetryDelayFunc func(n int) time.Duration
}

func DefaultRetryOptions() *RetryOptions {
	return &RetryOptions{
		MaxRetryTimes: defaultMaxRetryTimes,
		RetryDelayFunc: func(n int) time.Duration {
			return time.Duration(n) * time.Second
		},
	}
}

func retryWithOptions(ctx context.Context, fn func(ctx context.Context) error, opts RetryOptions) error {
	if opts.MaxRetryTimes <= 0 {
		return fn(ctx)
	}

	var lastErr error
retry:
	for i := 0; i <= opts.MaxRetryTimes; i++ {
		lastErr = fn(ctx)
		if lastErr == nil {
			return nil
		}
		var nerr *noRetryError
		if errors.As(lastErr, &nerr) {
			return nerr.err
		}

		if opts.RetryDelayFunc != nil {
			delay := opts.RetryDelayFunc(i + 1)
			if delay > 0 {
				select {
				case <-ctx.Done():
					lastErr = ctx.Err()
					break retry
				case <-time.After(delay):
				}
			}
		}
	}

	return fmt.Errorf("retry failed after %d attempts: %w", opts.MaxRetryTimes, lastErr)
}
