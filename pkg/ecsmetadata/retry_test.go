package ecsmetadata

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

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
