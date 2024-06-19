package version

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Version    string // 版本
	Message    string // 版本信息
	BuildTime  string // 编译时间
	CommitHash string // commit hash
)

// String 版本信息.
func String() string {
	kv := [][2]string{
		{"Go", runtime.Version()},
		{"Version", Version},
		{"Message", Message},
		{"BuildTime", BuildTime},
		{"CommitHash", CommitHash},
	}
	var str strings.Builder
	for _, v := range kv {
		str.WriteString(fmt.Sprintf("%10s: %s\n", v[0], v[1]))
	}
	return str.String()
}

// Command 版本命令.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "show the version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(String())
		},
	}
	return cmd
}
