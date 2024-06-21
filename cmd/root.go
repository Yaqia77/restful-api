package cmd

import (
	"fmt"
	"restful-api/version"

	"github.com/spf13/cobra"
)

var vers bool
var RootCmd = &cobra.Command{
	Use:   "host-api",
	Short: "host-api 后端API服务",
	Long:  "host-api 后端API服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return nil
	},
}

// 初始化
func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "Print version information")
}
