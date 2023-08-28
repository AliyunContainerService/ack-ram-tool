package provider

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type TLogger struct {
	t *testing.T
}

func (d TLogger) Info(msg string) {
	d.t.Logf(fmt.Sprintf("%s, %s", time.Now().Format(time.RFC3339), msg))
}

func (d TLogger) Debug(msg string) {
	d.t.Logf(fmt.Sprintf("%s, %s", time.Now().Format(time.RFC3339), msg))
}

func (d TLogger) Error(err error, msg string) {
	d.t.Logf(fmt.Sprintf("%s, %s", time.Now().Format(time.RFC3339), msg))
}

func TestUpdater_refreshCredForLoop_refresh(t *testing.T) {
	var callCount int
	fakeCred := Credentials{
		Expiration: time.Now().Add(time.Minute),
	}
	u := NewUpdater(func(ctx context.Context) (*Credentials, error) {
		callCount++
		return &fakeCred, nil
	}, UpdaterOptions{
		ExpiryWindow:  0,
		RefreshPeriod: 0,
		Logger:        TLogger{t: t},
	})

	u.refreshCredForLoop(context.TODO())
	if callCount != 1 {
		t.Errorf("callCount should be 1 but got %d", callCount)
	}
	ret := u.Expired()
	if ret {
		t.Errorf("should not expired")
	}

	u.refreshCredForLoop(context.TODO())
	if callCount != 1 {
		t.Errorf("callCount should be 1 but got %d", callCount)
	}

	u.nowFunc = func() time.Time {
		return time.Now().Add(time.Minute)
	}
	ret = u.Expired()
	if !ret {
		t.Errorf("should expired")
	}

	fakeCred.Expiration = time.Now().Add(time.Minute * 5)
	u.refreshCredForLoop(context.TODO())
	if callCount != 2 {
		t.Errorf("callCount should be 2 but got %d", callCount)
	}
	ret = u.Expired()
	if ret {
		t.Errorf("should not expired")
	}
}

func TestUpdater_refreshCredForLoop_erorr(t *testing.T) {
	var callCount int

	u := NewUpdater(func(ctx context.Context) (*Credentials, error) {
		callCount++
		return nil, errors.New("error message")
	}, UpdaterOptions{
		ExpiryWindow:  0,
		RefreshPeriod: 0,
		Logger:        TLogger{t: t},
	})

	u.refreshCredForLoop(context.TODO())
	if callCount != 5 {
		t.Errorf("callCount should be 5 but got %d", callCount)
	}
	ret := u.Expired()
	if !ret {
		t.Errorf("should expired")
	}
}

func TestUpdater_Credentials_refresh(t *testing.T) {
	var callCount int
	fakeCred := Credentials{
		Expiration: time.Now().Add(time.Minute),
	}
	u := NewUpdater(func(ctx context.Context) (*Credentials, error) {
		callCount++
		return &fakeCred, nil
	}, UpdaterOptions{
		ExpiryWindow:  0,
		RefreshPeriod: 0,
		Logger:        TLogger{t: t},
	})

	u.Credentials(context.TODO())
	if callCount != 1 {
		t.Errorf("callCount should be 1 but got %d", callCount)
	}
	ret := u.Expired()
	if ret {
		t.Errorf("should not expired")
	}

	u.Credentials(context.TODO())
	if callCount != 1 {
		t.Errorf("callCount should be 1 but got %d", callCount)
	}

	u.nowFunc = func() time.Time {
		return time.Now().Add(time.Minute)
	}
	ret = u.Expired()
	if !ret {
		t.Errorf("should expired")
	}

	fakeCred.Expiration = time.Now().Add(time.Minute * 5)
	u.Credentials(context.TODO())
	if callCount != 2 {
		t.Errorf("callCount should be 2 but got %d", callCount)
	}
	ret = u.Expired()
	if ret {
		t.Errorf("should not expired")
	}
}
