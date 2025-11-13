package provider

import (
	"testing"
	"time"
)

func TestCredentialsShouldRefresh(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name        string
		credentials *Credentials
		want        bool
	}{
		{
			name: "expired credentials should refresh",
			credentials: &Credentials{
				AccessKeyId:     "test-id",
				AccessKeySecret: "test-secret",
				SecurityToken:   "test-token",
				Expiration:      now.Add(-time.Minute),
				nextRefresh:     now.Add(time.Hour),
			},
			want: true,
		},
		{
			name: "not expired credentials with zero nextRefresh should not refresh",
			credentials: &Credentials{
				AccessKeyId:     "test-id",
				AccessKeySecret: "test-secret",
				SecurityToken:   "test-token",
				Expiration:      now.Add(time.Hour),
				nextRefresh:     time.Time{},
			},
			want: false,
		},
		{
			name: "not expired credentials with future nextRefresh should not refresh",
			credentials: &Credentials{
				AccessKeyId:     "test-id",
				AccessKeySecret: "test-secret",
				SecurityToken:   "test-token",
				Expiration:      now.Add(time.Hour),
				nextRefresh:     now.Add(time.Minute),
			},
			want: false,
		},
		{
			name: "not expired credentials with past nextRefresh should refresh",
			credentials: &Credentials{
				AccessKeyId:     "test-id",
				AccessKeySecret: "test-secret",
				SecurityToken:   "test-token",
				Expiration:      now.Add(time.Hour),
				nextRefresh:     now.Add(-time.Minute),
			},
			want: true,
		},
		{
			name: "not expired credentials with zero expiration and past nextRefresh should refresh",
			credentials: &Credentials{
				AccessKeyId:     "test-id",
				AccessKeySecret: "test-secret",
				SecurityToken:   "test-token",
				Expiration:      time.Time{},
				nextRefresh:     now.Add(-time.Minute),
			},
			want: true,
		},
		{
			name: "not expired credentials with zero expiration and zero nextRefresh should not refresh",
			credentials: &Credentials{
				AccessKeyId:     "test-id",
				AccessKeySecret: "test-secret",
				SecurityToken:   "test-token",
				Expiration:      time.Time{},
				nextRefresh:     time.Time{},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.credentials.shouldRefresh(now)
			if got != tt.want {
				t.Errorf("Credentials.shouldRefresh() = %v, want %v", got, tt.want)
			}
		})
	}
}
