package credentialplugin

import (
	"os"
	"reflect"
	"testing"
)

func Test_getExtraTokenQuery(t *testing.T) {
	os.Setenv(envTokenExtraQueryKeyPrefix, "FOO_")
	defer os.Unsetenv(envTokenExtraQueryKeyPrefix)
	os.Setenv("FOO_BAR", "test1")
	os.Setenv("FOO_FUZZ", "test2")
	tests := []struct {
		name string
		want map[string]string
	}{
		{
			name: "test",
			want: map[string]string{
				"bar":  "test1",
				"fuzz": "test2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getExtraTokenQuery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExtraTokenQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
