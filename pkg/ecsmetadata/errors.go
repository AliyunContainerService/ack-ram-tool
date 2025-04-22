package ecsmetadata

import (
	"fmt"
	"net/http"
	"strings"
)

type HTTPError struct {
	URL        string
	StatusCode int
	Header     http.Header
	Body       string

	Err     error
	Message string
}

type noRetryError struct {
	err error
}

func newHTTPError(err error, url string, resp *http.Response, body []byte) *HTTPError {
	newBody := strings.ReplaceAll(string(body), "\n", " ")
	newBody = strings.ReplaceAll(newBody, "\r", " ")
	newBody = truncateStr(newBody, 80)
	return &HTTPError{
		URL:        url,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       newBody,
		Err:        err,
		Message:    err.Error(),
	}
}

func newNoRetryError(err error) *noRetryError {
	return &noRetryError{err: err}
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%s. send data to %s failed, status code: %d, body: %s",
		e.Message, e.URL, e.StatusCode, e.Body)
}

func (e HTTPError) Unwrap() error {
	return e.Err
}

func (e noRetryError) Error() string {
	return e.err.Error()
}

func (e noRetryError) Unwrap() error {
	return e.err
}

func truncateStr(raw string, maxLen int) string {
	currLen := len(raw)
	if currLen <= maxLen {
		return raw
	}
	return raw[:maxLen] + "..."
}
