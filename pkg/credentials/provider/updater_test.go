package provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync/atomic"
	"testing"
	"time"
)

type TLogger struct {
	t *testing.T
}

func (d TLogger) Info(msg string) {
	defer catchPanic()
	d.t.Logf(fmt.Sprintf("%s, %s", time.Now().Format(time.RFC3339), msg))
}

func (d TLogger) Debug(msg string) {
	defer catchPanic()
	d.t.Logf(fmt.Sprintf("%s, %s", time.Now().Format(time.RFC3339), msg))
}

func (d TLogger) Error(err error, msg string) {
	defer catchPanic()
	d.t.Logf(fmt.Sprintf("%s, %s", time.Now().Format(time.RFC3339), msg))
}

func TestUpdater_refreshCredForLoop_refresh(t *testing.T) {
	var callCount int32
	fakeCred := Credentials{
		Expiration: time.Now().Add(time.Minute),
	}
	u := NewUpdater(func(ctx context.Context) (*Credentials, error) {
		atomic.AddInt32(&callCount, 1)
		return &fakeCred, nil
	}, UpdaterOptions{
		ExpiryWindow:  0,
		RefreshPeriod: 0,
		Logger:        TLogger{t: t},
	})

	u.refreshCredForLoop(context.TODO())

	cv := atomic.LoadInt32(&callCount)
	if cv != 1 {
		t.Errorf("callCount should be 1 but got %d", cv)
	}
	ret := u.Expired()
	if ret {
		t.Errorf("should not expired")
	}

	u.refreshCredForLoop(context.TODO())
	cv = atomic.LoadInt32(&callCount)
	if cv != 1 {
		t.Errorf("callCount should be 1 but got %d", cv)
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
	cv = atomic.LoadInt32(&callCount)
	if cv != 2 {
		t.Errorf("callCount should be 2 but got %d", cv)
	}
	ret = u.Expired()
	if ret {
		t.Errorf("should not expired")
	}
}

func TestUpdater_refreshCredForLoop_erorr(t *testing.T) {
	var callCount int32

	u := NewUpdater(func(ctx context.Context) (*Credentials, error) {
		atomic.AddInt32(&callCount, 1)
		return nil, errors.New("error message")
	}, UpdaterOptions{
		ExpiryWindow:  0,
		RefreshPeriod: 0,
		Logger:        TLogger{t: t},
	})

	u.refreshCredForLoop(context.TODO())
	cv := atomic.LoadInt32(&callCount)
	if cv != 5 {
		t.Errorf("callCount should be 5 but got %d", cv)
	}
	ret := u.Expired()
	if !ret {
		t.Errorf("should expired")
	}
}

func TestUpdater_Credentials_refresh(t *testing.T) {
	var callCount int32
	fakeCred := Credentials{
		Expiration: time.Now().Add(time.Minute),
	}
	u := NewUpdater(func(ctx context.Context) (*Credentials, error) {
		atomic.AddInt32(&callCount, 1)
		return &fakeCred, nil
	}, UpdaterOptions{
		ExpiryWindow:  0,
		RefreshPeriod: 0,
		Logger:        TLogger{t: t},
	})

	t.Run("Credentials use cache", func(t *testing.T) {
		u.Credentials(context.TODO())
		cv := atomic.LoadInt32(&callCount)
		if cv != 1 {
			t.Errorf("callCount should be 1 but got %d", cv)
		}
		ret := u.Expired()
		if ret {
			t.Errorf("should not expired")
		}

		u.Credentials(context.TODO())
		cv = atomic.LoadInt32(&callCount)
		if cv != 1 {
			t.Errorf("callCount should be 1 but got %d", cv)
		}
	})

	t.Run("Credentials expired", func(t *testing.T) {
		u.nowFunc = func() time.Time {
			return time.Now().Add(time.Minute * 2)
		}
		ret := u.Expired()
		if !ret {
			t.Errorf("should expired")
		}
	})

	t.Run("not expire, should not refresh", func(t *testing.T) {
		fakeCred.Expiration = time.Now().Add(time.Minute * 5)
		u.Credentials(context.TODO())
		cv := atomic.LoadInt32(&callCount)
		if cv != 2 {
			t.Errorf("callCount should be 2 but got %d", cv)
		}
		ret := u.Expired()
		if ret {
			t.Errorf("should not expired")
		}
	})
}

func TestUpdater_expired(t *testing.T) {
	u := &Updater{}
	u.setCred(&Credentials{Expiration: time.Now().Add(time.Minute)})

	t.Run("expiryDelta=0", func(t *testing.T) {
		ret := u.expired(0)
		if ret {
			t.Errorf("should be false")
		}
	})

	t.Run("expiryDelta > 0", func(t *testing.T) {
		ret := u.expired(time.Minute * 5)
		if !ret {
			t.Errorf("should be true")
		}
	})
}

func TestUpdater_stop(t *testing.T) {
	var callCount int32
	fakeCred := Credentials{
		Expiration: time.Now().Add(-time.Minute),
	}
	u := NewUpdater(func(ctx context.Context) (*Credentials, error) {
		atomic.AddInt32(&callCount, 1)
		return &fakeCred, nil
	}, UpdaterOptions{
		ExpiryWindow:  0,
		RefreshPeriod: time.Millisecond * 100,
		Logger:        TLogger{t: t},
	})

	u.Start(context.TODO())

	t.Run("test-refresh", func(t *testing.T) {
		time.Sleep(time.Second)
		cv := atomic.LoadInt32(&callCount)
		if cv < 1 {
			t.Errorf("callCount should be >1 but got %d", cv)
		}
	})

	t.Run("test-stop", func(t *testing.T) {
		u.Stop(context.TODO())
		time.Sleep(time.Second)

		curr := atomic.LoadInt32(&callCount)
		time.Sleep(time.Second)

		cv := atomic.LoadInt32(&callCount)
		if cv != curr {
			t.Errorf("callCount should be %d but got %d", curr, cv)
		}
	})

	t.Run("test-stop-multiple-times", func(t *testing.T) {
		u.Stop(context.TODO())
		u.Stop(context.TODO())
		u.Stop(context.TODO())
	})
}

func TestUpdater_stop_no_start(t *testing.T) {
	var callCount int32
	fakeCred := Credentials{
		Expiration: time.Now().Add(-time.Minute),
	}
	u := NewUpdater(func(ctx context.Context) (*Credentials, error) {
		atomic.AddInt32(&callCount, 1)
		return &fakeCred, nil
	}, UpdaterOptions{
		ExpiryWindow:  0,
		RefreshPeriod: 0,
		Logger:        TLogger{t: t},
	})

	u.Stop(context.TODO())
	u.Stop(context.TODO())
	u.Stop(context.TODO())
}

func catchPanic() {
	if err := recover(); err != nil {
		log.Printf("catch panic:\n%+v", err)
	}
}
