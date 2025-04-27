package ecsmetadata

import (
	"context"
	"fmt"
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

func TestGetTokenWith404Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut && r.URL.Path == "/latest/api/token" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Write([]byte("data"))
	}))
	defer server.Close()

	opts := ClientOptions{Endpoint: server.URL}
	client, _ := NewClient(opts)

	t.Run("token request returns 404", func(t *testing.T) {
		token, err := client.getToken(context.Background())
		if err == nil || !strings.Contains(err.Error(), "not found") {
			t.Fatalf("expected not found error, got %v", err)
		}
		if token != "" {
			t.Errorf("expected empty token, got %s", token)
		}
	})
}

func TestDefaultClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/latest/api/token" {
			w.Write([]byte("test-token"))
			return
		}
		if v := r.Header.Get("X-aliyun-ecs-metadata-token"); v != "test-token" {
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

func TestGetMetaDataWithRetry(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/latest/api/token" {
			// Simulate 404 error for token request
			if strings.Contains(r.URL.Path, "test-token-500-error") {
				http.Error(w, "internal server", http.StatusInternalServerError)
				return
			}
			w.Write([]byte("test-token"))
			return
		}
		if v := r.Header.Get("X-aliyun-ecs-metadata-token"); v != "test-token" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		// Simulate retryable error
		if strings.Contains(r.URL.Path, "test-retryable-error") {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		// Simulate non-retryable error
		if strings.Contains(r.URL.Path, "non-retryable-error") {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		w.Write([]byte("metadata-value"))
	}))
	defer server.Close()

	opts := ClientOptions{Endpoint: server.URL}
	client, _ := NewClient(opts)
	client.metadataToken = "test-token"
	client.metadataTokenExp = time.Now().Add(time.Hour)

	t.Run("disable retry", func(t *testing.T) {
		client.disableRetry = true
		data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path/test-retryable-error")
		if err == nil || !strings.Contains(err.Error(), "internal server error") {
			t.Fatalf("expected internal server error, got %v", err)
		}
		if data != nil {
			t.Errorf("expected no data, got %s", string(data))
		}
	})

	t.Run("retryable error", func(t *testing.T) {
		client.disableRetry = false
		client.retryOptions.MaxRetryTimes = 3
		data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path/test-retryable-error")
		if err == nil || !strings.Contains(err.Error(), "internal server error") {
			t.Fatalf("expected internal server error after retries, got %v", err)
		}
		if data != nil {
			t.Errorf("expected no data, got %s", string(data))
		}
	})

	t.Run("non-retryable error", func(t *testing.T) {
		client.disableRetry = false
		data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path/non-retryable-error")
		if err == nil || !strings.Contains(err.Error(), "bad request") {
			t.Fatalf("expected bad request error without retries, got %v", err)
		}
		if data != nil {
			t.Errorf("expected no data, got %s", string(data))
		}
	})
}

func TestGetMetaDataWithToken404Ignored(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/latest/api/token" {
			// Simulate 404 error for token request
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		// Ensure the metadata request is processed without a token
		w.Write([]byte("metadata-value"))
	}))
	defer server.Close()

	opts := ClientOptions{Endpoint: server.URL}
	client, _ := NewClient(opts)

	t.Run("ignore 404 error from get token and return metadata", func(t *testing.T) {
		data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != "metadata-value" {
			t.Errorf("expected metadata-value, got %s", string(data))
		}
	})
}

func TestGetMetaDataWithToken500NotIgnored(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/latest/api/token" {
			// Simulate 500 error for token request
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		// Ensure the metadata request is processed without a token
		w.Write([]byte("metadata-value"))
	}))
	defer server.Close()

	opts := ClientOptions{Endpoint: server.URL}
	client, _ := NewClient(opts)

	data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path")
	if err == nil || !strings.Contains(err.Error(), "internal server error") {
		t.Fatalf("expected internal server error, got %v", err)
	}
	if data != nil {
		t.Errorf("expected no data, got %s", string(data))
	}
}

func TestReuseToken(t *testing.T) {
	accessTokenCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/latest/api/token" {
			accessTokenCount++
			w.Write([]byte(fmt.Sprintf("test-token-%d", accessTokenCount)))
			return
		}
		if v := r.Header.Get("X-aliyun-ecs-metadata-token"); v != "test-token-1" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		w.Write([]byte("metadata-value"))
	}))
	defer server.Close()
	client, err := NewClient(ClientOptions{Endpoint: server.URL})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for i := 0; i < 5; i++ {
		data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != "metadata-value" {
			t.Errorf("expected metadata-value, got %s", string(data))
		}
		if accessTokenCount != 1 {
			t.Errorf("expected 1 access token, got %d", accessTokenCount)
		}
	}
}

func TestRefreshToken(t *testing.T) {
	accessTokenCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/latest/api/token" {
			accessTokenCount++
			w.Write([]byte(fmt.Sprintf("test-token-%d", accessTokenCount)))
			return
		}
		w.Write([]byte("metadata-value"))
	}))
	defer server.Close()
	client, err := NewClient(ClientOptions{
		Endpoint:        server.URL,
		TokenTTLSeconds: 75 + 5,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for i := 0; i < 5; i++ {
		data, err := client.GetMetaData(context.Background(), http.MethodGet, "/test-path")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != "metadata-value" {
			t.Errorf("expected metadata-value, got %s", string(data))
		}
		time.Sleep(time.Second)
	}
	if accessTokenCount != 2 {
		t.Errorf("expected 2 access token, got %d", accessTokenCount)
	}
	if client.metadataToken != "test-token-2" {
		t.Errorf("expected test-token-2, got %s", client.metadataToken)
	}
	t.Log(time.Now())
	t.Log(client.metadataTokenExp)
}
