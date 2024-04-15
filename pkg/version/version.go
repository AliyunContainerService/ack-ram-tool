package version

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

var (
	Version            = "unknown"
	GitCommit          = ""
	defaultProgramName = "ack-ram-tool"
	binName            = ""
)

func init() {
	if len(os.Args) > 0 {
		binName = path.Base(os.Args[0])
	}
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
