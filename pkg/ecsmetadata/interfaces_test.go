package ecsmetadata

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestGetMacs(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/', got '%s'", path)
							}
							return 200, "mac1\nmac2", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetMacs(ctx)
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		expected := []string{"mac1", "mac2"}
		if !equalStringSlices(result, expected) {
			t.Errorf("expected result '%v', got '%v'", expected, result)
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
							if path != "/latest/meta-data/network/interfaces/macs/" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/', got '%s'", path)
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
		result, err := client.GetMacs(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if len(result) != 0 {
			t.Errorf("expected empty result, got '%v'", result)
		}
	})
}

func TestGetInterfaceIdByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/network-interface-id" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/network-interface-id', got '%s'", path)
							}
							return 200, "interface-id-123", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetInterfaceIdByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "interface-id-123" {
			t.Errorf("expected result 'interface-id-123', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/network-interface-id" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/network-interface-id', got '%s'", path)
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
		result, err := client.GetInterfaceIdByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetNetMaskByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/netmask" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/netmask', got '%s'", path)
							}
							return 200, "255.255.255.0", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetNetMaskByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "255.255.255.0" {
			t.Errorf("expected result '255.255.255.0', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/netmask" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/netmask', got '%s'", path)
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
		result, err := client.GetNetMaskByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetVSwitchCidrBlockIdByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vswitch-cidr-block" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vswitch-cidr-block', got '%s'", path)
							}
							return 200, "192.168.1.0/24", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetVSwitchCidrBlockIdByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "192.168.1.0/24" {
			t.Errorf("expected result '192.168.1.0/24', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vswitch-cidr-block" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vswitch-cidr-block', got '%s'", path)
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
		result, err := client.GetVSwitchCidrBlockIdByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetPrivateIPV4sByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/private-ipv4s" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/private-ipv4s', got '%s'", path)
							}
							return 200, `["192.168.1.10", "192.168.1.11"]`, nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetPrivateIPV4sByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		expected := []string{"192.168.1.10", "192.168.1.11"}
		if !equalStringSlices(result, expected) {
			t.Errorf("expected result '%v', got '%v'", expected, result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/private-ipv4s" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/private-ipv4s', got '%s'", path)
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
		result, err := client.GetPrivateIPV4sByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if len(result) != 0 {
			t.Errorf("expected empty result, got '%v'", result)
		}
	})
}

func TestGetVpcIPV6CidrBlocksByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vpc-ipv6-cidr-blocks" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vpc-ipv6-cidr-blocks', got '%s'", path)
							}
							return 200, `["2001:db8::/64", "2001:db8:1::/64"]`, nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetVpcIPV6CidrBlocksByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		expected := []string{"2001:db8::/64", "2001:db8:1::/64"}
		if !equalStringSlices(result, expected) {
			t.Errorf("expected result '%v', got '%v'", expected, result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vpc-ipv6-cidr-blocks" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vpc-ipv6-cidr-blocks', got '%s'", path)
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
		result, err := client.GetVpcIPV6CidrBlocksByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if len(result) != 0 {
			t.Errorf("expected empty result, got '%v'", result)
		}
	})
}

func TestGetVSwitchIdByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vswitch-id" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vswitch-id', got '%s'", path)
							}
							return 200, "vsw-1234567890abcdef0", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetVSwitchIdByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "vsw-1234567890abcdef0" {
			t.Errorf("expected result 'vsw-1234567890abcdef0', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vswitch-id" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vswitch-id', got '%s'", path)
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
		result, err := client.GetVSwitchIdByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetVpcIdByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vpc-id" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vpc-id', got '%s'", path)
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
		result, err := client.GetVpcIdByMac(ctx, "mac1")
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
					return &MockWrapper{
						Mock: func(path string) (int, string, error) {
							if path == "/latest/api/token" {
								return 200, "token", nil
							}
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vpc-id" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vpc-id', got '%s'", path)
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
		result, err := client.GetVpcIdByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetPrimaryIPAddressByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/primary-ip-address" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/primary-ip-address', got '%s'", path)
							}
							return 200, "192.168.1.10", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetPrimaryIPAddressByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "192.168.1.10" {
			t.Errorf("expected result '192.168.1.10', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/primary-ip-address" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/primary-ip-address', got '%s'", path)
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
		result, err := client.GetPrimaryIPAddressByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetGatewayByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/gateway" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/gateway', got '%s'", path)
							}
							return 200, "192.168.1.1", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetGatewayByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "192.168.1.1" {
			t.Errorf("expected result '192.168.1.1', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/gateway" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/gateway', got '%s'", path)
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
		result, err := client.GetGatewayByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetIPV6sByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/ipv6s" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/ipv6s', got '%s'", path)
							}
							return 200, `["2001:db8::1", "2001:db8::2"]`, nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetIPV6sByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		expected := []string{"2001:db8::1", "2001:db8::2"}
		if !equalStringSlices(result, expected) {
			t.Errorf("expected result '%v', got '%v'", expected, result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/ipv6s" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/ipv6s', got '%s'", path)
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
		result, err := client.GetIPV6sByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if len(result) != 0 {
			t.Errorf("expected empty result, got '%v'", result)
		}
	})
}

func TestGetIPV6GatewayByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/ipv6-gateway" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/ipv6-gateway', got '%s'", path)
							}
							return 200, "2001:db8::1", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetIPV6GatewayByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "2001:db8::1" {
			t.Errorf("expected result '2001:db8::1', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/ipv6-gateway" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/ipv6-gateway', got '%s'", path)
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
		result, err := client.GetIPV6GatewayByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetVSwitchIPV6CidrBlockByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vswitch-ipv6-cidr-block" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vswitch-ipv6-cidr-block', got '%s'", path)
							}
							return 200, "2001:db8::/64", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetVSwitchIPV6CidrBlockByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "2001:db8::/64" {
			t.Errorf("expected result '2001:db8::/64', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/vswitch-ipv6-cidr-block" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/vswitch-ipv6-cidr-block', got '%s'", path)
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
		result, err := client.GetVSwitchIPV6CidrBlockByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}

func TestGetIPV4PrefixesByMac(t *testing.T) {
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/ipv4-prefixes" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/ipv4-prefixes', got '%s'", path)
							}
							return 200, "192.168.1.0/24", nil
						},
					}
				},
			},
		})
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		result, err := client.GetIPV4PrefixesByMac(ctx, "mac1")
		if err != nil {
			t.Errorf("expected no error, got '%v'", err)
		}
		if result != "192.168.1.0/24" {
			t.Errorf("expected result '192.168.1.0/24', got '%s'", result)
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
							if path != "/latest/meta-data/network/interfaces/macs/mac1/ipv4-prefixes" {
								t.Errorf("expected path '/latest/meta-data/network/interfaces/macs/mac1/ipv4-prefixes', got '%s'", path)
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
		result, err := client.GetIPV4PrefixesByMac(ctx, "mac1")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result, got '%s'", result)
		}
	})
}
