package common

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
)

func EnsureDir(dir string) (string, error) {
	realDir, err := utils.ExpandPath(dir)
	if err != nil {
		return dir, err
	}
	f, err := os.Open(filepath.Clean(realDir))
	if err != nil {
		if os.IsNotExist(err) || strings.Contains(err.Error(), "no such file or directory") {
			err := os.MkdirAll(realDir, 0700)
			return realDir, err
		}
		return realDir, err
	}
	f.Close()
	return realDir, nil
}
