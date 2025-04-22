package ecsmetadata

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockTransportWrapper struct {
	rt       http.RoundTripper
	callback func(path string) (int, string, error)
}

func (m *MockTransportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	path := req.URL.RequestURI()
	code, body, err := m.callback(path)
	if err != nil {
		return nil, err
	}
	return &http.Response{
		StatusCode: code,
		Body:       io.NopCloser(strings.NewReader(body)),
	}, nil
}

func TestGetRegionId(t *testing.T) {
	ctx := context.Background()

	t.Run("normal case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockTransportWrapper{
						rt: rt,
						callback: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/region-id" {
								t.Errorf("expected path '/latest/meta-data/region-id', got '%s'", path)
							}
							return 200, "cn-hangzhou", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetRegionId(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "cn-hangzhou" {
			t.Errorf("expected result 'cn-hangzhou', got '%s'", result)
		}
	})

	t.Run("error case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockTransportWrapper{
						rt: rt,
						callback: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/region-id" {
								t.Errorf("expected path '/latest/meta-data/region-id', got '%s'", path)
							}
							return 400, "", errors.New("mock error")
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetRegionId(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetZoneId(t *testing.T) {
	ctx := context.Background()

	t.Run("normal case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockTransportWrapper{
						rt: rt,
						callback: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/zone-id" {
								t.Errorf("expected path '/latest/meta-data/zone-id', got '%s'", path)
							}
							return 200, "cn-hangzhou-a", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetZoneId(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "cn-hangzhou-a" {
			t.Errorf("expected result 'cn-hangzhou-a', got '%s'", result)
		}
	})

	t.Run("error case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockTransportWrapper{
						rt: rt,
						callback: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/zone-id" {
								t.Errorf("expected path '/latest/meta-data/zone-id', got '%s'", path)
							}
							return 200, "", errors.New("mock error")
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetZoneId(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetOwnerAccountId(t *testing.T) {
	ctx := context.Background()

	t.Run("normal case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockTransportWrapper{
						rt: rt,
						callback: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/owner-account-id" {
								t.Errorf("expected path '/latest/meta-data/owner-account-id', got '%s'", path)
							}
							return 200, "123456789012", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetOwnerAccountId(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "123456789012" {
			t.Errorf("expected result '123456789012', got '%s'", result)
		}
	})

	t.Run("error case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockTransportWrapper{
						rt: rt,
						callback: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/owner-account-id" {
								t.Errorf("expected path '/latest/meta-data/owner-account-id', got '%s'", path)
							}
							return 400, "", errors.New("mock error")
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetOwnerAccountId(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}
