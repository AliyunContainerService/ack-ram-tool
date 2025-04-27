package ecsmetadata

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestNewHTTPError(t *testing.T) {
	url := "http://example.com"
	statusCode := 500
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	body := []byte(`{"error": "internal server error"}`)
	err := errors.New("internal server error")

	httpErr := newHTTPError(err, url, &http.Response{
		StatusCode: statusCode,
		Header:     header,
	}, body)

	if httpErr.URL != url {
		t.Errorf("expected URL %s, got %s", url, httpErr.URL)
	}
	if httpErr.StatusCode != statusCode {
		t.Errorf("expected StatusCode %d, got %d", statusCode, httpErr.StatusCode)
	}
	if !strings.Contains(httpErr.Body, "internal server error") {
		t.Errorf("expected Body to contain 'internal server error', got %s", httpErr.Body)
	}
	if httpErr.Message != err.Error() {
		t.Errorf("expected Message %s, got %s", err.Error(), httpErr.Message)
	}
}

func TestNoRetryError(t *testing.T) {
	originalErr := errors.New("non-retryable error")
	noRetryErr := newNoRetryError(originalErr)

	if noRetryErr.Error() != originalErr.Error() {
		t.Errorf("expected Error() to return '%s', got '%s'", originalErr.Error(), noRetryErr.Error())
	}
	if noRetryErr.Unwrap() != originalErr {
		t.Errorf("expected Unwrap() to return original error, got %v", noRetryErr.Unwrap())
	}
}
