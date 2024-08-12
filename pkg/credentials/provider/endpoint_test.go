package provider

import (
	"os"
	"testing"
)

func TestGetSTSEndpoint(t *testing.T) {
	type args struct {
		region     string
		vpcNetwork bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				region:     "cn-hangzhou",
				vpcNetwork: false,
			},
			want: "sts.cn-hangzhou.aliyuncs.com",
		},
		{
			name: "default vpc",
			args: args{
				region:     "cn-hangzhou",
				vpcNetwork: true,
			},
			want: "sts-vpc.cn-hangzhou.aliyuncs.com",
		},
		{
			name: "cn-hangzhou-finance",
			args: args{
				region:     "cn-hangzhou-finance",
				vpcNetwork: false,
			},
			want: "sts.cn-hangzhou.aliyuncs.com",
		},
		{
			name: "cn-hangzhou-finance vpc",
			args: args{
				region:     "cn-hangzhou-finance",
				vpcNetwork: true,
			},
			want: "sts-vpc.cn-hangzhou.aliyuncs.com",
		},
		{
			name: "unknown vpc",
			args: args{
				region:     "cn-unknown",
				vpcNetwork: true,
			},
			want: "sts-vpc.cn-unknown.aliyuncs.com",
		},
		{
			name: "unknown",
			args: args{
				region:     "cn-unknown",
				vpcNetwork: false,
			},
			want: "sts.cn-unknown.aliyuncs.com",
		},
		{
			name: "us-west-1",
			args: args{
				region:     "us-west-1",
				vpcNetwork: false,
			},
			want: "sts.us-west-1.aliyuncs.com",
		},
		{
			name: "us-west-1 vpc",
			args: args{
				region:     "us-west-1",
				vpcNetwork: true,
			},
			want: "sts-vpc.us-west-1.aliyuncs.com",
		},
		{
			name: "empty vpc",
			args: args{
				region:     "",
				vpcNetwork: true,
			},
			want: "sts.aliyuncs.com",
		},
		{
			name: "empty",
			args: args{
				region:     "",
				vpcNetwork: true,
			},
			want: "sts.aliyuncs.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSTSEndpoint(tt.args.region, tt.args.vpcNetwork); got != tt.want {
				t.Errorf("GetSTSEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSTSEndpointFromEnv(t *testing.T) {
	os.Setenv("ALIBABA_CLOUD_STS_ENDPOINT", "sts.cn-hangzhou.aliyuncs.com")
	defer os.Unsetenv("ALIBABA_CLOUD_STS_ENDPOINT")

	want := "sts.cn-hangzhou.aliyuncs.com"
	if got := GetSTSEndpoint("foo", true); got != want {
		t.Errorf("GetSTSEndpoint() = %v, want %v", got, want)
	}
}
