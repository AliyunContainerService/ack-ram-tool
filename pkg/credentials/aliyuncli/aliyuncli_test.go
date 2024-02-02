package aliyuncli

import (
	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, got, tt.want)
		})
	}
}
