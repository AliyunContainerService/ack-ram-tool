package ecsmetadata

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
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
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
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
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
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
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
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
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
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
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.RequestURI()
			switch path {
			case "/latest/api/token":
				w.Write([]byte("token"))
			case "/latest/meta-data/owner-account-id":
				w.Write([]byte("123456789012"))
			default:
				t.Errorf("expected path '/latest/meta-data/owner-account-id', got '%s'", path)
			}
		}))
		defer s.Close()
		client, err := NewClient(ClientOptions{
			Endpoint: s.URL,
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
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.RequestURI()
			if path == "/latest/api/token" {
				w.Write([]byte("token"))
				return
			}
			if path != "/latest/meta-data/owner-account-id" {
				t.Errorf("expected path '/latest/meta-data/owner-account-id', got '%s'", path)
			}
			w.WriteHeader(400)
		}))
		defer s.Close()
		client, err := NewClient(ClientOptions{
			Endpoint: s.URL,
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

func TestGetHostname(t *testing.T) {
	ctx := context.Background()

	t.Run("normal case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/hostname" {
								t.Errorf("expected path '/latest/meta-data/hostname', got '%s'", path)
							}
							return 200, "example-hostname", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetHostname(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "example-hostname" {
			t.Errorf("expected result 'example-hostname', got '%s'", result)
		}
	})

	t.Run("error case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/hostname" {
								t.Errorf("expected path '/latest/meta-data/hostname', got '%s'", path)
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
		result, err := client.GetHostname(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetSourceAddress(t *testing.T) {
	ctx := context.Background()

	t.Run("normal case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/source-address" {
								t.Errorf("expected path '/latest/meta-data/source-address', got '%s'", path)
							}
							return 200, "http://mirrors.cloud.aliyuncs.com/", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetSourceAddress(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "http://mirrors.cloud.aliyuncs.com/" {
			t.Errorf("expected result 'http://mirrors.cloud.aliyuncs.com/', got '%s'", result)
		}
	})

	t.Run("error case", func(t *testing.T) {
		client, err := NewClient(ClientOptions{
			TransportWrappers: []TransportWrapper{
				func(rt http.RoundTripper) http.RoundTripper {
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/source-address" {
								t.Errorf("expected path '/latest/meta-data/source-address', got '%s'", path)
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
		result, err := client.GetSourceAddress(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetSourceAddressList(t *testing.T) {
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
							if path != "/latest/meta-data/source-address" {
								t.Errorf("expected path '/latest/meta-data/source-address', got '%s'", path)
							}
							return 200, "192.168.0.1\n10.0.0.1", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetSourceAddressList(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		expected := []string{"192.168.0.1", "10.0.0.1"}
		if len(result) != len(expected) {
			t.Errorf("expected result length %d, got %d", len(expected), len(result))
		}
		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("expected result[%d] '%s', got '%s'", i, expected[i], result[i])
			}
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
							if path != "/latest/meta-data/source-address" {
								t.Errorf("expected path '/latest/meta-data/source-address', got '%s'", path)
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
		result, err := client.GetSourceAddressList(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != nil {
			t.Errorf("expected nil result, got '%v'", result)
		}
	})
}
