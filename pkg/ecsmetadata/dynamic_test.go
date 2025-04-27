package ecsmetadata

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetDocument(t *testing.T) {
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
							if path != "/latest/dynamic/instance-identity/document" {
								t.Errorf("expected path '/latest/dynamic/instance-identity/document', got '%s'", path)
							}
							doc := map[string]string{
								"account-id":       "123456789012",
								"owner-account-id": "123456789012",
								"region-id":        "cn-hangzhou",
								"zone-id":          "cn-hangzhou-a",
							}
							data, _ := json.Marshal(doc)
							return 200, string(data), nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}

		doc, err := client.GetDocument(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if doc.RegionId != "cn-hangzhou" || doc.ZoneId != "cn-hangzhou-a" {
			t.Errorf("unexpected document content: %+v", doc)
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
							if path != "/latest/dynamic/instance-identity/document" {
								t.Errorf("expected path '/latest/dynamic/instance-identity/document', got '%s'", path)
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

		doc, err := client.GetDocument(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if doc != nil {
			t.Errorf("expected nil document, got '%+v'", doc)
		}
	})
}

func TestGetRawDocument(t *testing.T) {
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
							if path != "/latest/dynamic/instance-identity/document" {
								t.Errorf("expected path '/latest/dynamic/instance-identity/document', got '%s'", path)
							}
							return 200, `{"region-id":"cn-hangzhou","zone-id":"cn-hangzhou-a"}`, nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}

		rawDoc, err := client.GetRawDocument(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if !strings.Contains(rawDoc, `"region-id":"cn-hangzhou"`) {
			t.Errorf("unexpected raw document content: %s", rawDoc)
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
							if path != "/latest/dynamic/instance-identity/document" {
								t.Errorf("expected path '/latest/dynamic/instance-identity/document', got '%s'", path)
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

		rawDoc, err := client.GetRawDocument(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if rawDoc != "" {
			t.Errorf("expected empty raw document, got '%s'", rawDoc)
		}
	})
}

func TestNewDocumentPKCS7Signature(t *testing.T) {
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
							expectedPath := "/latest/dynamic/instance-identity/pkcs7?audience=test"
							if path != expectedPath {
								t.Errorf("expected path '%s', got '%s'", expectedPath, path)
							}
							return 200, "mock-signature", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}

		signature, err := client.NewDocumentPKCS7Signature(ctx, "test")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if signature != "mock-signature" {
			t.Errorf("expected signature 'mock-signature', got '%s'", signature)
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
							expectedPath := "/latest/dynamic/instance-identity/pkcs7?audience=test"
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

		signature, err := client.NewDocumentPKCS7Signature(ctx, "test")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if signature != "" {
			t.Errorf("expected empty signature, got '%s'", signature)
		}
	})
}
