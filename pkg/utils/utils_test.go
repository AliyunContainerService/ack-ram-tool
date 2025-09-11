package utils

import (
	"reflect"
	"testing"
)

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

func TestOr(t *testing.T) {
	type args[T comparable] struct {
		a         T
		fallbacks []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[string]{
		{
			name: "String: a has value",
			args: args[string]{a: "hello", fallbacks: []string{"world"}},
			want: "hello",
		},
		{
			name: "String: a is empty, use first fallback",
			args: args[string]{a: "", fallbacks: []string{"default1", "default2"}},
			want: "default1",
		},
		{
			name: "String: a and all fallbacks are empty",
			args: args[string]{a: "", fallbacks: []string{"", ""}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Or(tt.args.a, tt.args.fallbacks...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
