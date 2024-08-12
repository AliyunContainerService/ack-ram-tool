package openapi

import (
	"reflect"
	"testing"
)

func TestNewEndpoints(t *testing.T) {
	type args struct {
		region string
		vpc    bool
	}
	tests := []struct {
		name string
		args args
		want Endpoints
	}{
		{
			name: "cn-hangzhou",
			args: args{
				region: "cn-hangzhou",
				vpc:    false,
			},
			want: Endpoints{
				RAM: "ram.aliyuncs.com",
				STS: "sts.cn-hangzhou.aliyuncs.com",
				CS:  "cs.cn-hangzhou.aliyuncs.com",
			},
		},
		{
			name: "cn-hangzhou with vpc",
			args: args{
				region: "cn-hangzhou",
				vpc:    true,
			},
			want: Endpoints{
				RAM: "ram.vpc-proxy.aliyuncs.com",
				STS: "sts-vpc.cn-hangzhou.aliyuncs.com",
				CS:  "cs-anony-vpc.cn-hangzhou.aliyuncs.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEndpoints(tt.args.region, tt.args.vpc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEndpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
