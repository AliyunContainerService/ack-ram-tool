package ecsmetadata

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		opts := ClientOptions{}
		client, err := NewClient(opts)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if client.endpoint != DefaultEndpoint {
			t.Errorf("expected endpoint %s, got %s", DefaultEndpoint, client.endpoint)
		}
		if client.tokenTTLSeconds != defaultTokenTTLSeconds {
			t.Errorf("expected tokenTTLSeconds %d, got %d", defaultTokenTTLSeconds, client.tokenTTLSeconds)
		}
	})

	t.Run("invalid TokenTTLSeconds", func(t *testing.T) {
		opts := ClientOptions{TokenTTLSeconds: maxTokenTTLSeconds + 1}
		_, err := NewClient(opts)
		if err == nil || !strings.Contains(err.Error(), "invalid TokenTTLSeconds") {
			t.Fatalf("expected invalid TokenTTLSeconds error, got %v", err)
		}
	})
}

func TestGetToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/latest/api/token" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Write([]byte("test-token"))
	}))
	defer server.Close()

	t.Run("IMDSv2 disabled", func(t *testing.T) {
		opts := ClientOptions{DisableIMDSV2: true, Endpoint: server.URL}
		client, _ := NewClient(opts)
		token, err := client.getToken(context.Background())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if token != "" {
			t.Errorf("expected empty token, got %s", token)
		}
	})

	t.Run("token not expired", func(t *testing.T) {
		opts := ClientOptions{Endpoint: server.URL}
		client, _ := NewClient(opts)
		client.metadataToken = "cached-token"
		client.metadataTokenExp = time.Now().Add(time.Minute)
		token, err := client.getToken(context.Background())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if token != "cached-token" {
			t.Errorf("expected cached token, got %s", token)
		}
	})

	t.Run("fetch new token", func(t *testing.T) {
		opts := ClientOptions{Endpoint: server.URL}
		client, _ := NewClient(opts)
		token, err := client.getToken(context.Background())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if token != "test-token" {
			t.Errorf("expected test-token, got %s", token)
		}
	})
}

func TestDefaultClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/latest/api/token" {
			w.Write([]byte("test-token"))
			return
		}
		if v := r.Header.Get("X-aliyun-ecs-metadata-token"); v != "" && v != "test-token" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		w.Write([]byte("metadata-value"))
	}))
	defer server.Close()
	DefaultClient.endpoint = server.URL

	t.Run("successful request", func(t *testing.T) {
		data, err := DefaultClient.GetMetaData(context.Background(), http.MethodGet, "/test-path")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != "metadata-value" {
			t.Errorf("expected metadata-value, got %s", string(data))
		}
	})
}

func TestGetMetaData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/latest/api/token" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		if v := r.Header.Get("X-aliyun-ecs-metadata-token"); v != "" && v != "test-token" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		w.Write([]byte("metadata-value"))
	}))
	defer server.Close()

	opts := ClientOptions{Endpoint: server.URL}
	client, _ := NewClient(opts)
	client.metadataToken = "test-token"
	client.metadataTokenExp = time.Now().Add(time.Hour)

	t.Run("successful request", func(t *testing.T) {
		data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != "metadata-value" {
			t.Errorf("expected metadata-value, got %s", string(data))
		}
	})

	t.Run("forbidden error from get token was ignored", func(t *testing.T) {
		client.metadataToken = ""
		data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != "metadata-value" {
			t.Errorf("expected metadata-value, got %s", string(data))
		}
	})
}

func TestSendWithRetry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "retry") {
			http.Error(w, "retryable error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("success"))
	}))
	defer server.Close()

	opts := ClientOptions{Endpoint: server.URL}
	client, _ := NewClient(opts)

	t.Run("success without retry", func(t *testing.T) {
		data, err := client.sendWithRetry(context.Background(), http.MethodGet, "/success", nil)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != "success" {
			t.Errorf("expected success, got %s", string(data))
		}
	})

	t.Run("retry logic", func(t *testing.T) {
		_, err := client.sendWithRetry(context.Background(), http.MethodGet, "/retry", nil)
		if err == nil || !strings.Contains(err.Error(), "retry failed") {
			t.Fatalf("expected retry error, got %v", err)
		}
	})
}
