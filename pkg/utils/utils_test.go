package utils

import "testing"

func TestReplaceNewLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "\n",
			args: args{
				s: "a\nb",
			},
			want: "a b",
		},
		{
			name: "\r\n",
			args: args{
				s: "a\r\nb",
			},
			want: "a b",
		},
		{
			name: "\r\n v2",
			args: args{
				s: "a\rc\n\r\nb",
			},
			want: "ac  b",
		},
		{
			name: "dont replace",
			args: args{
				s: "abc",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceNewLine(tt.args.s); got != tt.want {
				t.Errorf("ReplaceNewLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
