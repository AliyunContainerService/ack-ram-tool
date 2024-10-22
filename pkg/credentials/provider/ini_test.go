package provider

import (
	"reflect"
	"testing"
)

func Test_parseIniConfig(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want iniConfig
	}{
		{
			name: "test 1",
			args: args{
				input: []byte(`
; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.

[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62     
port = 143
file = "payroll.dat"
`),
			},
			want: iniConfig{
				items: []iniConfigItem{
					{
						section: "owner",
						config: map[string]iniConfigItemValue{
							"name":         "John Doe",
							"organization": "Acme Widgets Inc.",
						},
					},
					{
						section: "database",
						config: map[string]iniConfigItemValue{
							"server": "192.0.2.62",
							"port":   "143",
							"file":   "payroll.dat",
						},
					},
				},
			},
		},
		{
			name: "test 2",
			args: args{
				input: []byte(`
[default]                          # Default credential
type = access_key                  # Certification type: access_key
access_key_id = foo                # access key id
access_key_secret = bar            # access key secret

[ram]
type = ram_role_arn
access_key_id = foo
access_key_secret = bar
role_arn = role_arn
role_session_name = session_name
`),
			},
			want: iniConfig{
				items: []iniConfigItem{
					{
						section: "default",
						config: map[string]iniConfigItemValue{
							"type":              "access_key",
							"access_key_id":     "foo",
							"access_key_secret": "bar",
						},
					},
					{
						section: "ram",
						config: map[string]iniConfigItemValue{
							"type":              "ram_role_arn",
							"access_key_id":     "foo",
							"access_key_secret": "bar",
							"role_arn":          "role_arn",
							"role_session_name": "session_name",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseIniConfig(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIniConfig() = %#v, want %v", got, tt.want)
			}
		})
	}
}
