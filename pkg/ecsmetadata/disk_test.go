package ecsmetadata

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestGetDisks(t *testing.T) {
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
							if path != "/latest/meta-data/disks/" {
								t.Errorf("expected path '/latest/meta-data/disks/', got '%s'", path)
							}
							return 200, "disk1\ndisk2", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetDisks(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		expected := []string{"disk1", "disk2"}
		if !equalStringSlices(result, expected) {
			t.Errorf("expected result '%v', got '%v'", expected, result)
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
							if path != "/latest/meta-data/disks/" {
								t.Errorf("expected path '/latest/meta-data/disks/', got '%s'", path)
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
		result, err := client.GetDisks(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if len(result) != 0 {
			t.Errorf("expected empty result, got '%v'", result)
		}
	})
}

func TestGetDiskId(t *testing.T) {
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
							if path != "/latest/meta-data/disks/diskSerial/id" {
								t.Errorf("expected path '/latest/meta-data/disks/diskSerial/id', got '%s'", path)
							}
							return 200, "disk-id-123", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetDiskId(ctx, "diskSerial")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "disk-id-123" {
			t.Errorf("expected result 'disk-id-123', got '%s'", result)
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
							if path != "/latest/meta-data/disks/diskSerial/id" {
								t.Errorf("expected path '/latest/meta-data/disks/diskSerial/id', got '%s'", path)
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
		result, err := client.GetDiskId(ctx, "diskSerial")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetDiskName(t *testing.T) {
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
							if path != "/latest/meta-data/disks/diskSerial/name" {
								t.Errorf("expected path '/latest/meta-data/disks/diskSerial/name', got '%s'", path)
							}
							return 200, "disk-name-123", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetDiskName(ctx, "diskSerial")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "disk-name-123" {
			t.Errorf("expected result 'disk-name-123', got '%s'", result)
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
							if path != "/latest/meta-data/disks/diskSerial/name" {
								t.Errorf("expected path '/latest/meta-data/disks/diskSerial/name', got '%s'", path)
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
		result, err := client.GetDiskName(ctx, "diskSerial")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}
