package ecsmetadata

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestRetryWithOptions(t *testing.T) {
	ctx := context.Background()
	delayCount := 0
	opts := RetryOptions{
		MaxRetryTimes: 3,
		RetryDelayFunc: func(n int) time.Duration {
			delayCount++
			return 0
		},
	}

	callCount := 0
	fn := func(ctx context.Context) error {
		callCount++
		return fmt.Errorf("retryable error")
	}

	err := retryWithOptions(ctx, fn, opts)
	t.Log(err)
	t.Log(callCount)
	t.Log(delayCount)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if callCount != 4 {
		t.Fatalf("expected function to be called 4 times, got %d calls", callCount)
	}
	if delayCount != 3 {
		t.Fatalf("expected delay to be called 3 times, got %d calls", delayCount)
	}
}

func TestRetryWithOptions_SuccessOnFirstAttempt(t *testing.T) {
	ctx := context.Background()
	opts := DefaultRetryOptions()

	callCount := 0
	fn := func(ctx context.Context) error {
		callCount++
		return nil // Success on first attempt
	}

	err := retryWithOptions(ctx, fn, *opts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if callCount != 1 {
		t.Fatalf("expected function to be called once, got %d calls", callCount)
	}
}

func TestRetryWithOptions_SuccessAfterRetries(t *testing.T) {
	ctx := context.Background()
	opts := DefaultRetryOptions()

	callCount := 0
	fn := func(ctx context.Context) error {
		callCount++
		if callCount < opts.MaxRetryTimes {
			return fmt.Errorf("retryable error")
		}
		return nil // Success after retries
	}

	err := retryWithOptions(ctx, fn, *opts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if callCount != opts.MaxRetryTimes {
		t.Fatalf("expected function to be called %d times, got %d calls", opts.MaxRetryTimes, callCount)
	}
}

func TestRetryWithOptions_FailureAfterMaxRetries(t *testing.T) {
	ctx := context.Background()
	opts := DefaultRetryOptions()

	fn := func(ctx context.Context) error {
		return fmt.Errorf("retryable error") // Always fail
	}

	err := retryWithOptions(ctx, fn, *opts)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	expectedErrMsg := fmt.Sprintf("retry failed after %d attempts", opts.MaxRetryTimes)
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Fatalf("expected error message '%s', got '%s'", expectedErrMsg, err.Error())
	}
}

func TestRetryWithOptions_ContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	opts := DefaultRetryOptions()

	fn := func(ctx context.Context) error {
		cancel() // Cancel context on first attempt
		return fmt.Errorf("retryable error")
	}

	err := retryWithOptions(ctx, fn, *opts)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled error, got %v", err)
	}
}

func Test_isRetryable(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil error should not be retryable",
			args: args{err: nil},
			want: false,
		},
		{
			name: "generic error should be retryable",
			args: args{err: errors.New("generic error")},
			want: true,
		},
		{
			name: "HTTPError with 404 status code should not be retryable",
			args: args{err: &HTTPError{StatusCode: http.StatusNotFound}},
			want: false,
		},
		{
			name: "HTTPError with 400 status code should not be retryable",
			args: args{err: &HTTPError{StatusCode: http.StatusBadRequest}},
			want: false,
		},
		{
			name: "HTTPError with 500 status code should be retryable",
			args: args{err: &HTTPError{StatusCode: http.StatusInternalServerError}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRetryable(tt.args.err); got != tt.want {
				t.Errorf("isRetryable() = %v, want %v", got, tt.want)
			}
		})
	}
}
