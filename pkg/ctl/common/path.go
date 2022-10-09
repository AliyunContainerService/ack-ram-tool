package common

import (
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	"os"
	"strings"
)

func EnsureDir(dir string) (string, error) {
	realDir, err := utils.ExpandPath(dir)
	if err != nil {
		return dir, err
	}
	f, err := os.Open(realDir)
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
