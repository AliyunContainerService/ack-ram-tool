package utils

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func JSONEqual(a interface{}, b interface{}) bool {
	av, err := json.Marshal(a)
	if err != nil {
		return false
	}
	bv, err := json.Marshal(b)
	if err != nil {
		return false
	}
	return bytes.Equal(av, bv)
}

func StringInterfaceMapEqual(a map[string]interface{}, b map[string]interface{}) bool {
	aKeys := []string{}
	bKeys := []string{}
	for k := range a {
		aKeys = append(aKeys, k)
	}
	for k := range b {
		bKeys = append(bKeys, k)
	}
	if !reflect.DeepEqual(aKeys, bKeys) {
		return false
	}
	for k, av := range a {
		bv := b[k]
		if !JSONEqual(av, bv) {
			return false
		}
	}
	return true
}

func JSONValue(o interface{}) []byte {
	d, _ := json.MarshalIndent(o, " ", "")
	return d
}

func ReplaceNewLine(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	return strings.ReplaceAll(s, "\n", " ")
}

func ExpandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[1:])
	}
	return path, nil
}
