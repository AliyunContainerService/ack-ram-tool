package provider

import (
	"context"
	"testing"
)

func TestNewIniConfigProvider(t *testing.T) {
	fp := "testdata/ak.ini"
	cp, err := NewIniConfigProvider(INIConfigProviderOptions{
		ConfigPath: fp, SectionName: "default"})
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	cred, err := cp.Credentials(context.TODO())
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	if cred.AccessKeyId != "foo_from_ini" ||
		cred.AccessKeySecret != "bar_from_ini" {
		t.Errorf("unexpected cred: %+v", *cred)
	}
}

func TestNewIniConfigProvider_load(t *testing.T) {
	files := []string{
		"testdata/ak.ini",
		"testdata/sts.ini",
		"testdata/uri.ini",
		"testdata/ext.ini",
		"testdata/role.ini",
		"testdata/ecs.ini",
	}

	for _, fp := range files {
		t.Run(fp, func(t *testing.T) {
			cp, err := NewIniConfigProvider(INIConfigProviderOptions{
				ConfigPath: fp, SectionName: "default"})
			if err != nil {
				t.Errorf("err is not nil: %+v", err)
			}
			if cp == nil {
				t.Errorf("cp is nil")
			}
		})
	}
}
