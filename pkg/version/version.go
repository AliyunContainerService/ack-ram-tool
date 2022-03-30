package version

import "fmt"

var (
	Version   = "unknown"
	GitCommit = ""
)

func UserAgent() string {
	return fmt.Sprintf("ack-ram-tool/%s", Version)
}
