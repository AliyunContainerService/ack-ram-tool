package ecsmetadata

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"
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

func TestGetRoleCredentials(t *testing.T) {
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
							expectedPath := "/latest/meta-data/ram/security-credentials/test-role"
							if path != expectedPath {
								t.Errorf("expected path '%s', got '%s'", expectedPath, path)
							}
							rawData := map[string]string{
								"AccessKeyId":     "mock-access-key-id",
								"AccessKeySecret": "mock-access-key-secret",
								"SecurityToken":   "mock-security-token",
								"Expiration":      "2023-10-01T12:00:00Z",
								"LastUpdated":     "2023-09-30T12:00:00Z",
								"Code":            "Success",
							}
							data, _ := json.Marshal(rawData)
							return 200, string(data), nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}

		result, err := client.GetRoleCredentials(ctx, "test-role")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result.AccessKeyId != "mock-access-key-id" || result.AccessKeySecret != "mock-access-key-secret" {
			t.Errorf("unexpected RoleCredentials: %+v", result)
		}
		expectedExp, _ := time.Parse(time.RFC3339, "2023-10-01T12:00:00Z")
		if !result.Expiration.Equal(expectedExp) {
			t.Errorf("expected Expiration %v, got %v", expectedExp, result.Expiration)
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
							expectedPath := "/latest/meta-data/ram/security-credentials/test-role"
							if path != expectedPath {
								t.Errorf("expected path '%s', got '%s'", expectedPath, path)
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

		result, err := client.GetRoleCredentials(ctx, "test-role")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != nil {
			t.Errorf("expected nil result, got '%+v'", result)
		}
	})
}
