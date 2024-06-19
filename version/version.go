package version

import "fmt"

//服务器名称
const (
	ServiceName = "host-api"
)

var (
	GitTag    string
	GitCommit string
	GitBranch string
	BuildTime string
	GoVersion string
)

func FullVersion() string {
	version := fmt.Sprintf("Version: %s\n Commit: %s\n Branch: %s\n , BuildTime: %s\n , GoVersion: %s\n ")
	return version
}

//Short 版本缩写
func Short() string {
	return fmt.Sprintf("%s[%s %s]", GitTag, GitBranch, GitCommit)
}
