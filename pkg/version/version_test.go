package version

import "testing"

func Test_getBaseBinName(t *testing.T) {
	type args struct {
		rawName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "bar => bar",
			args: args{
				rawName: "bar",
			},
			want: "bar",
		},
		{
			name: "/bar => bar",
			args: args{
				rawName: "/bar",
			},
			want: "bar",
		},
		{
			name: "/foo/bar => bar",
			args: args{
				rawName: "/foo/bar",
			},
			want: "bar",
		},
		{
			name: "\\foo\\bar => bar",
			args: args{
				rawName: "\\foo\\bar",
			},
			want: "bar",
		},
		{
			name: "\\\\foo\\\\bar => bar",
			args: args{
				rawName: "\\\\foo\\\\bar",
			},
			want: "bar",
		},
		{
			name: "/ => default",
			args: args{
				rawName: "/",
			},
			want: "ack-ram-tool",
		},
		{
			name: "./ => default",
			args: args{
				rawName: "./",
			},
			want: "ack-ram-tool",
		},
		{
			name: "empty => default",
			args: args{
				rawName: "",
			},
			want: "ack-ram-tool",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBaseBinName(tt.args.rawName); got != tt.want {
				t.Errorf("getBaseBinName() = %v, want %v", got, tt.want)
			}
		})
	}
}
