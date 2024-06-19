package cmd

import (
	"restful-api/apps/host/http"
	"restful-api/apps/host/impl"
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
		service := impl.NewHostServiceImpl()
		api := http.NewHostHttpHandler(service)

		r := gin.Default()
		api.Registry(r)
		return r.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile,"config","f","etc/demo.toml","Host api配置文件路径")
	RootCmd.AddCommand(StartCmd)
}