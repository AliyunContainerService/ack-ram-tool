package ecsmetadata

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestGetVpcId(t *testing.T) {
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
							if path != "/latest/meta-data/vpc-id" {
								t.Errorf("expected path '/latest/meta-data/vpc-id', got '%s'", path)
							}
							return 200, "vpc-1234567890abcdef0", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetVpcId(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "vpc-1234567890abcdef0" {
			t.Errorf("expected result 'vpc-1234567890abcdef0', got '%s'", result)
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
							if path != "/latest/meta-data/vpc-id" {
								t.Errorf("expected path '/latest/meta-data/vpc-id', got '%s'", path)
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
		result, err := client.GetVpcId(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetVpcCidrBlockId(t *testing.T) {
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
							if path != "/latest/meta-data/vpc-cidr-block" {
								t.Errorf("expected path '/latest/meta-data/vpc-cidr-block', got '%s'", path)
							}
							return 200, "192.168.0.0/16", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetVpcCidrBlockId(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "192.168.0.0/16" {
			t.Errorf("expected result '192.168.0.0/16', got '%s'", result)
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
							if path != "/latest/meta-data/vpc-cidr-block" {
								t.Errorf("expected path '/latest/meta-data/vpc-cidr-block', got '%s'", path)
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
		result, err := client.GetVpcCidrBlockId(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetMAC(t *testing.T) {
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
							if path != "/latest/meta-data/mac" {
								t.Errorf("expected path '/latest/meta-data/mac', got '%s'", path)
							}
							return 200, "00:16:3e:12:34:56", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetMAC(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "00:16:3e:12:34:56" {
			t.Errorf("expected result '00:16:3e:12:34:56', got '%s'", result)
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
							if path != "/latest/meta-data/mac" {
								t.Errorf("expected path '/latest/meta-data/mac', got '%s'", path)
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
		result, err := client.GetMAC(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}
