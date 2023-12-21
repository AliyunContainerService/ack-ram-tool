package binding

import (
	"testing"
)

func Test_getAliUidFromSubjectName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "valid 1",
			args: args{
				name: "12345-1676970192",
			},
			want:    12345,
			wantErr: false,
		},
		{
			name: "valid 2",
			args: args{
				name: "12345",
			},
			want:    12345,
			wantErr: false,
		},
		{
			name: "invalid 1",
			args: args{
				name: "12345-",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "invalid 2",
			args: args{
				name: "foobar",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAliUidFromSubjectName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAliUidFromSubjectName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getAliUidFromSubjectName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
