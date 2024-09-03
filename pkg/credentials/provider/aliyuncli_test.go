package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
)

func Test_tryToParseProfileFromOutput(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name string
		args args
		want *Profile
	}{
		{
			name: "include profile 1",
			args: args{
				output: `

{
  "access_key_id": "accessKeyId",
  "mode": "AK",
  "access_key_secret": "accessKeySecret"
}

`,
			},
			want: &Profile{
				Mode:            "AK",
				AccessKeyId:     "accessKeyId",
				AccessKeySecret: "accessKeySecret",
			},
		},
		{
			name: "include profile 2",
			args: args{
				output: `afbb

{
  "mode": "StsToken",
  "access_key_id": "accessKeyId",
  "access_key_secret": "accessKeySecret",
  "sts_token": "stsToken"
}
xxxx
`,
			},
			want: &Profile{
				Mode:            "StsToken",
				AccessKeyId:     "accessKeyId",
				AccessKeySecret: "accessKeySecret",
				StsToken:        "stsToken",
			},
		},
		{
			name: "include profile 3",
			args: args{
				output: `
XXX
{
  "mode": "RamRoleArn",
  "access_key_id": "accessKeyId",
  "access_key_secret": "accessKeySecret",
  "ram_role_arn": "ramRoleArn",
  "ram_session_name": "ramSessionName"
}XXX
xxxx
`,
			},
			want: &Profile{
				Mode:            "RamRoleArn",
				AccessKeyId:     "accessKeyId",
				AccessKeySecret: "accessKeySecret",
				RamRoleArn:      "ramRoleArn",
				RoleSessionName: "ramSessionName",
			},
		},
		{
			name: "not include profile 1",
			args: args{
				output: "",
			},
			want: nil,
		},
		{
			name: "not include profile 2",
			args: args{
				output: "foo {}",
			},
			want: nil,
		},
		{
			name: "not include profile 3",
			args: args{
				output: `foo { "mode": XXXX }`,
			},
			want: nil,
		},
		{
			name: "not include profile 4",
			args: args{
				output: `
{
  "mode": "",
  "access_key_id": "accessKeyId",
  "access_key_secret": "accessKeySecret"
}
`,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tryToParseProfileFromOutput(tt.args.output)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("should return %#v but got %#v", tt.want, got)
			}
		})
	}
}

func Test_AK(t *testing.T) {
	configPath := "testdata/ak.json"
	p, err := NewCLIConfigProvider(CLIConfigProviderOptions{
		ConfigPath:  configPath,
		ProfileName: "default",
		STSEndpoint: "127.0.0.1",
		Logger:      DefaultLogger,
	})
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	if cred.AccessKeyId != "accessKeyId_AK" {
		t.Errorf("unexpected AccessKeyId: %q", cred.AccessKeyId)
	}
}

func Test_StsToken(t *testing.T) {
	configPath := "testdata/sts.json"
	p, err := NewCLIConfigProvider(CLIConfigProviderOptions{
		ConfigPath:  configPath,
		ProfileName: "default",
		STSEndpoint: "127.0.0.1",
		Logger:      DefaultLogger,
	})
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	if cred.AccessKeyId != "accessKeyId_STS" {
		t.Errorf("unexpected AccessKeyId: %q", cred.AccessKeyId)
	}
}

func Test_ChainableRamRoleArn(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "ak_ChainableRamRoleArn",
     "AccessKeySecret": "sk",
     "SecurityToken": "tt",
     "Expiration": "2206-01-02T15:04:05Z"
  }
}
`)
	})
	defer s.Close()

	configPath := "testdata/chain.json"
	p, err := NewCLIConfigProvider(CLIConfigProviderOptions{
		ConfigPath:  configPath,
		ProfileName: "default",
		STSEndpoint: s.URL,
		Logger:      DefaultLogger,
	})
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	if cred.AccessKeyId != "ak_ChainableRamRoleArn" {
		t.Errorf("unexpected AccessKeyId: %q", cred.AccessKeyId)
	}
}

func Test_RamRoleArn(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
{
  "Credentials": {
     "AccessKeyId": "accessKeyId_RamRoleArn",
     "AccessKeySecret": "sk",
     "SecurityToken": "tt",
     "Expiration": "2206-01-02T15:04:05Z"
  }
}
`)
	})
	defer s.Close()

	configPath := "testdata/assumerole.json"
	p, err := NewCLIConfigProvider(CLIConfigProviderOptions{
		ConfigPath:  configPath,
		ProfileName: "default",
		STSEndpoint: s.URL,
		Logger:      DefaultLogger,
	})
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	if cred.AccessKeyId != "accessKeyId_RamRoleArn" {
		t.Errorf("unexpected AccessKeyId: %q", cred.AccessKeyId)
	}
}

func Test_CredentialsURI(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
{
  "Code": "Success",
  "AccessKeyId": "ak_CredentialsURI",
  "AccessKeySecret": "<ak secret>",
  "SecurityToken": "<security token>",
  "Expiration": "2006-01-02T15:04:05Z"
}
`)
	})
	defer s.Close()

	tpl := "testdata/uri.json"
	dir, _ := os.MkdirTemp("", "testuri")
	fpath := path.Join(dir, "uri.json")
	b, _ := os.ReadFile(tpl)
	data := strings.ReplaceAll(string(b), "<uri>", s.URL)
	os.WriteFile(fpath, []byte(data), 0644)

	p, err := NewCLIConfigProvider(CLIConfigProviderOptions{
		ConfigPath:  fpath,
		ProfileName: "default",
		STSEndpoint: s.URL,
		Logger:      DefaultLogger,
	})
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	if cred.AccessKeyId != "ak_CredentialsURI" {
		t.Errorf("unexpected AccessKeyId: %q", cred.AccessKeyId)
	}
}

func Test_EcsRamRole(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "token") {
			w.Write([]byte("fake token"))
			return
		}
		fmt.Fprint(w, `
{
  "Code": "Success",
  "AccessKeyId": "ak_EcsRamRole",
  "AccessKeySecret": "<ak secret>",
  "SecurityToken": "<security token>",
  "Expiration": "2006-01-02T15:04:05Z"
}
`)
	})
	defer s.Close()

	fpath := "testdata/ecs.json"
	os.Setenv("ALIBABA_CLOUD_IMDS_ENDPOINT", s.URL)
	defer os.Unsetenv("ALIBABA_CLOUD_IMDS_ENDPOINT")

	p, err := NewCLIConfigProvider(CLIConfigProviderOptions{
		ConfigPath:  fpath,
		ProfileName: "default",
		STSEndpoint: "127.0.0.1",
		Logger:      DefaultLogger,
	})
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	if cred.AccessKeyId != "ak_EcsRamRole" {
		t.Errorf("unexpected AccessKeyId: %q", cred.AccessKeyId)
	}
}

func Test_External(t *testing.T) {
	s := setupHttpTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
{
      "name": "default",
      "mode": "AK",
      "access_key_id": "ak_External",
      "access_key_secret": "accessKeySecret"
    }
`)
	})
	defer s.Close()

	tpl := "testdata/ext.json"
	dir, _ := os.MkdirTemp("", "test_ext")
	fpath := path.Join(dir, "ext.json")
	b, _ := os.ReadFile(tpl)
	data := strings.ReplaceAll(string(b), "<uri>", s.URL)
	os.WriteFile(fpath, []byte(data), 0644)

	p, err := NewCLIConfigProvider(CLIConfigProviderOptions{
		ConfigPath:  fpath,
		ProfileName: "default",
		STSEndpoint: s.URL,
		Logger:      DefaultLogger,
	})
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	cred, err := p.Credentials(context.TODO())
	if err != nil {
		t.Errorf("err is not nil: %+v", err)
	}
	if cred.AccessKeyId != "ak_External" {
		t.Errorf("unexpected AccessKeyId: %q", cred.AccessKeyId)
	}
}
