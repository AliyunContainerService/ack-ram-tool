package provider

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestFileCacheProvider_Credentials_use_cache(t *testing.T) {
	cacheDir := ensureTmpDir(t)
	defer os.RemoveAll(cacheDir)

	exp := time.Now().Add(time.Hour * 4).Truncate(0)
	callCount := 0
	cp := NewFunctionProvider(func(ctx context.Context) (*Credentials, error) {
		callCount++
		return &Credentials{
			AccessKeyId:     "test",
			AccessKeySecret: "test",
			Expiration:      exp,
			SecurityToken:   "test",
		}, nil
	})

	fp := NewFileCacheProvider(cacheDir, cp, FileCacheProviderOptions{})
	cred1, err := fp.Credentials(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// use cache
	cred2, err := fp.Credentials(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if callCount != 1 {
		t.Fatal("callCount should be 1")
	}
	if !cred1.Expiration.Equal(exp) {
		t.Fatal("cred1 should be equal to exp")
	}
	if !cred1.Expiration.Equal(cred2.Expiration) {
		t.Fatal("cred1 should be equal to cred2")
	}
}

func TestFileCacheProvider_Credentials_cache_expired(t *testing.T) {
	cacheDir := ensureTmpDir(t)
	defer os.RemoveAll(cacheDir)

	exp := time.Now().Add(time.Minute * 8).Truncate(0)
	exp2 := time.Now().Add(time.Minute * 10).Truncate(0)
	callCount := 0
	cp := NewFunctionProvider(func(ctx context.Context) (*Credentials, error) {
		callCount++
		cred := &Credentials{
			AccessKeyId:     "test",
			AccessKeySecret: "test",
			Expiration:      exp,
			SecurityToken:   "test",
		}
		if callCount > 1 {
			cred.Expiration = exp2
		}
		return cred, nil
	})

	fp := NewFileCacheProvider(cacheDir, cp, FileCacheProviderOptions{})
	cred1, err := fp.Credentials(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// use cache but cache is expired
	cred2, err := fp.Credentials(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if callCount != 2 {
		t.Fatal("callCount should be 2")
	}
	if !cred1.Expiration.Equal(exp) {
		t.Fatal("cred1 should be equal to exp")
	}
	if !cred2.Expiration.Equal(exp2) {
		t.Fatal("cred2 should be equal to exp2")
	}
}

func TestFileCacheProvider_Credentials_cache_broken(t *testing.T) {
	cacheConents := []string{
		"invalid",
		"aW52YWxpZAo=",
	}

	for _, cacheContent := range cacheConents {
		t.Run("invalid "+cacheContent, func(t *testing.T) {
			cacheDir := ensureTmpDir(t)
			defer os.RemoveAll(cacheDir)

			exp := time.Now().Add(time.Hour * 1).Truncate(0)
			exp2 := time.Now().Add(time.Hour * 10).Truncate(0)
			callCount := 0
			cp := NewFunctionProvider(func(ctx context.Context) (*Credentials, error) {
				callCount++
				cred := &Credentials{
					AccessKeyId:     "test",
					AccessKeySecret: "test",
					Expiration:      exp,
					SecurityToken:   "test",
				}
				if callCount > 1 {
					cred.Expiration = exp2
				}
				return cred, nil
			})

			fp := NewFileCacheProvider(cacheDir, cp, FileCacheProviderOptions{})
			cred1, err := fp.Credentials(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			err = os.WriteFile(fp.cacheFilePath(), []byte(cacheContent), 0644)
			if err != nil {
				t.Fatal(err)
			}

			// use cache but cache was broken
			cred2, err := fp.Credentials(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			if callCount != 2 {
				t.Fatal("callCount should be 2")
			}
			if !cred1.Expiration.Equal(exp) {
				t.Fatal("cred1 should be equal to exp")
			}
			if !cred2.Expiration.Equal(exp2) {
				t.Fatal("cred2 should be equal to exp2")
			}
		})
	}
}

func TestFileCacheProvider_Credentials_error(t *testing.T) {
	cacheDir := ensureTmpDir(t)
	defer os.RemoveAll(cacheDir)

	callCount := 0
	cp := NewFunctionProvider(func(ctx context.Context) (*Credentials, error) {
		callCount++
		return nil, fmt.Errorf("error for %d", callCount)
	})

	fp := NewFileCacheProvider(cacheDir, cp, FileCacheProviderOptions{})
	_, err1 := fp.Credentials(context.Background())
	if err1 == nil {
		t.Fatal("should have error")
	}
	if err1.Error() != "error for 1" {
		t.Fatal("error should be error for 1")
	}

	// use cache but cache was broken
	_, err2 := fp.Credentials(context.Background())
	if err2 == nil {
		t.Fatal("should have error")
	}
	if err2.Error() != "error for 2" {
		t.Fatal("error should be error for 2")
	}

	if callCount != 2 {
		t.Fatal("callCount should be 2")
	}
}

func ensureTmpDir(t *testing.T) string {
	dir := path.Join(os.TempDir(), "foobar")
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		t.Fatal(err)
	}
	tmpDir, err := os.MkdirTemp(dir, "test")
	return tmpDir
}
