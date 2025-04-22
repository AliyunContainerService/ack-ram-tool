package ecsmetadata

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestGetRoleName(t *testing.T) {
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
							if path != "/latest/meta-data/ram/security-credentials/" {
								t.Errorf("expected path '/latest/meta-data/ram/security-credentials/', got '%s'", path)
							}
							return 200, "test-role", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetRoleName(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "test-role" {
			t.Errorf("expected result 'test-role', got '%s'", result)
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
							if path != "/latest/meta-data/ram/security-credentials/" {
								t.Errorf("expected path '/latest/meta-data/ram/security-credentials/', got '%s'", path)
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
		result, err := client.GetRoleName(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}