package provider

import (
	"testing"
	"time"
)

func TestCredentialsExpired(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name        string
		credentials *Credentials
		expiryDelta time.Duration
		want        bool
	}{
		{
			name: "zero expiration time should not expire",
			credentials: &Credentials{
				Expiration: time.Time{}, // Zero time
			},
			expiryDelta: 0,
			want:        false,
		},
		{
			name: "expiration time in the future should not expire",
			credentials: &Credentials{
				Expiration: now.Add(time.Hour), // 1 hour in the future
			},
			expiryDelta: 0,
			want:        false,
		},
		{
			name: "expiration time in the past should expire",
			credentials: &Credentials{
				Expiration: now.Add(-time.Hour), // 1 hour in the past
			},
			expiryDelta: 0,
			want:        true,
		},
		{
			name: "expiration time equals now should expire",
			credentials: &Credentials{
				Expiration: now, // Now
			},
			expiryDelta: 0,
			want:        true,
		},
		{
			name: "expiration time in future but within expiryDelta should expire",
			credentials: &Credentials{
				Expiration: now.Add(30 * time.Minute), // 30 minutes in future
			},
			expiryDelta: 45 * time.Minute, // But we consider it expired 45 minutes before actual expiration
			want:        true,
		},
		{
			name: "expiration time in future and outside expiryDelta should not expire",
			credentials: &Credentials{
				Expiration: now.Add(60 * time.Minute), // 60 minutes in future
			},
			expiryDelta: 45 * time.Minute, // We consider it expired 45 minutes before actual expiration
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.credentials.expired(now, tt.expiryDelta)
			if got != tt.want {
				t.Errorf("Credentials.expired() = %v, want %v", got, tt.want)
			}
		})
	}
}
