package version

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

var (
	Version            = "unknown"
	GitCommit          = ""
	defaultProgramName = "ack-ram-tool"
	binName            = defaultProgramName
)

func init() {
	if len(os.Args) > 0 {
		binName = getBaseBinName(os.Args[0])
	}
}

func getBaseBinName(rawName string) string {
	rawName = strings.TrimSpace(rawName)
	binName := path.Base(rawName)
	binName = strings.Trim(binName, "./\\")
	if strings.Contains(binName, "/") {
		parts := strings.Split(binName, "/")
		binName = parts[len(parts)-1]
	}
	if strings.Contains(binName, "\\") {
		parts := strings.Split(binName, "\\")
		binName = parts[len(parts)-1]
	}
	binName = strings.TrimSpace(binName)
	if binName == "" {
		return defaultProgramName
	}
	return binName
}

func BinName() string {
	if binName == "" {
		return defaultProgramName
	}
	return binName
}

func UserAgent() string {
	ua := fmt.Sprintf("%s ack-ram-tool/%s", binName, Version)
	if binName == "" || binName == defaultProgramName {
		ua = fmt.Sprintf("ack-ram-tool/%s", Version)
	}

	goInfo := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	ua = fmt.Sprintf("%s (%s)", ua, goInfo)
	return ua
}
