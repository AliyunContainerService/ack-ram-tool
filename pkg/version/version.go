package version

import (
	"fmt"
	"os"
	"path"
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
	if binName == "" || binName == defaultProgramName {
		return fmt.Sprintf("ack-ram-tool/%s", Version)
	}
	return fmt.Sprintf("%s ack-ram-tool/%s", binName, Version)
}
