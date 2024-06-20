package cmd

import (
	"restful-api/apps"
	_ "restful-api/apps/all"
	"restful-api/conf"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	confType string
	confFile string
	confETCD string
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Long:  `Start the server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}
		//加载我们Host Service的实体类
		// host service 的具体实现
		// service := impl.NewHostServiceImpl()

		//注册Host Service 的实例到IOC容器中
		//采用 _ "restful-api/apps/host/impl" 方式，完成注册
		// apps.HostService = impl.NewHostServiceImpl()

		apps.InitImpl()

		r := gin.Default()
		//注册IOC容器中所有http handler
		apps.InitGin(r)

		return r.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "Host api配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
