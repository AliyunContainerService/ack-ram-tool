package ecsmetadata

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestGetInstanceType(t *testing.T) {
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
							if path != "/latest/meta-data/instance/instance-type" {
								t.Errorf("expected path '/latest/meta-data/instance/instance-type', got '%s'", path)
							}
							return 200, "ecs.g6.large", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetInstanceType(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "ecs.g6.large" {
			t.Errorf("expected result 'ecs.g6.large', got '%s'", result)
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
							if path != "/latest/meta-data/instance/instance-type" {
								t.Errorf("expected path '/latest/meta-data/instance/instance-type', got '%s'", path)
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
		result, err := client.GetInstanceType(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetInstanceName(t *testing.T) {
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
							if path != "/latest/meta-data/instance/instance-name" {
								t.Errorf("expected path '/latest/meta-data/instance/instance-name', got '%s'", path)
							}
							return 200, "test-instance", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetInstanceName(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "test-instance" {
			t.Errorf("expected result 'test-instance', got '%s'", result)
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
							if path != "/latest/meta-data/instance/instance-name" {
								t.Errorf("expected path '/latest/meta-data/instance/instance-name', got '%s'", path)
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
		result, err := client.GetInstanceName(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetInstanceId(t *testing.T) {
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
							if path != "/latest/meta-data/instance-id" {
								t.Errorf("expected path '/latest/meta-data/instance-id', got '%s'", path)
							}
							return 200, "i-1234567890abcdef0", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetInstanceId(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "i-1234567890abcdef0" {
			t.Errorf("expected result 'i-1234567890abcdef0', got '%s'", result)
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
							if path != "/latest/meta-data/instance-id" {
								t.Errorf("expected path '/latest/meta-data/instance-id', got '%s'", path)
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
		result, err := client.GetInstanceId(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetImageId(t *testing.T) {
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
							if path != "/latest/meta-data/image-id" {
								t.Errorf("expected path '/latest/meta-data/image-id', got '%s'", path)
							}
							return 200, "ubuntu_20_04_x64_20G_alibase_20230922.vhd", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetImageId(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "ubuntu_20_04_x64_20G_alibase_20230922.vhd" {
			t.Errorf("expected result 'ubuntu_20_04_x64_20G_alibase_20230922.vhd', got '%s'", result)
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
							if path != "/latest/meta-data/image-id" {
								t.Errorf("expected path '/latest/meta-data/image-id', got '%s'", path)
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
		result, err := client.GetImageId(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetSerialNumber(t *testing.T) {
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
							if path != "/latest/meta-data/serial-number" {
								t.Errorf("expected path '/latest/meta-data/serial-number', got '%s'", path)
							}
							return 200, "test-serial-number", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetSerialNumber(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "test-serial-number" {
			t.Errorf("expected result 'test-serial-number', got '%s'", result)
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
							if path != "/latest/meta-data/serial-number" {
								t.Errorf("expected path '/latest/meta-data/serial-number', got '%s'", path)
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
		result, err := client.GetSerialNumber(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetSpotTerminationTime(t *testing.T) {
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
							if path != "/latest/meta-data/instance/spot/termination-time" {
								t.Errorf("expected path '/latest/meta-data/instance/spot/termination-time', got '%s'", path)
							}
							return 200, "2023-10-01T12:00:00Z", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetSpotTerminationTime(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		expectedTime, _ := time.Parse(time.RFC3339, "2023-10-01T12:00:00Z")
		if !result.Equal(expectedTime) {
			t.Errorf("expected result '%v', got '%v'", expectedTime, result)
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
							if path != "/latest/meta-data/instance/spot/termination-time" {
								t.Errorf("expected path '/latest/meta-data/instance/spot/termination-time', got '%s'", path)
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
		result, err := client.GetSpotTerminationTime(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if !result.IsZero() {
			t.Errorf("expected zero time, got '%v'", result)
		}
	})
}
