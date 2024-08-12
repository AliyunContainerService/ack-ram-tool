package provider

import (
	"context"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestSemaphoreProvider_Credentials(t *testing.T) {
	cp := NewFunctionProvider(func(ctx context.Context) (*Credentials, error) {
		time.Sleep(time.Millisecond * 200)
		return &Credentials{}, nil
	})

	p := NewSemaphoreProvider(cp, SemaphoreProviderOptions{})

	go p.Credentials(context.TODO())
	runtime.Gosched()

	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*100)
	defer cancel()

	// no free lock
	_, err := p.Credentials(ctx)
	if err == nil {
		t.Error("err should not be nil")
	} else {
		if !strings.Contains(err.Error(), "acquire semaphore: context deadline exceeded") {
			t.Log(err)
			t.Error("err should include 'context deadline exceeded'")
		}
	}

	time.Sleep(time.Millisecond * 300)
	// has free lock
	ctx2, cancel2 := context.WithTimeout(context.TODO(), time.Millisecond*100)
	defer cancel2()
	cred, err := p.Credentials(ctx2)
	if err != nil {
		t.Log(err)
		t.Error("err should be nil")
	}
	if cred == nil {
		t.Error("cred should not be nil")
	}
}
