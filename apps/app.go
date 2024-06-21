package apps

import (
	"restful-api/apps/host"

	"github.com/gin-gonic/gin"
)

// IOC 容器层：管理所有的服务的实例

// 1. HostService 服务实例必须注册过来，HostService 才会有具体的实例
// 2. HTTP 暴露模块，依赖IOC 里面的HostService
var (
	HostService host.Service

	ImplApp = map[string]ImplService{}
	ginApps = map[string]GinService{}
)

func RegistryImpl(svc ImplService) {
	//服务实例注册到svcs map当中
	if _, ok := ImplApp[svc.Name()]; ok {
		panic("service already registered: " + svc.Name())
	}

	ImplApp[svc.Name()] = svc

	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 已经加载完成的Gin App由哪些
func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}

	return
}

func GetImpl(name string) interface{} {
	for k, v := range ImplApp {
		if k == name {
			return v
		}
	}
	return nil
}

func RegistryGin(svc GinService) {
	//服务实例注册到svcs map当中
	if _, ok := ginApps[svc.Name()]; ok {
		panic("service already registered: " + svc.Name())
	}

	ginApps[svc.Name()] = svc

}

// 用户初始化 注册到IOC容器里面的所有服务
func InitImpl() {
	for _, v := range ImplApp {
		v.Config()
	}
}

func InitGin(r gin.IRouter) {
	//先初始化好所有对象
	for _, v := range ginApps {
		v.Config()
	}

	//完成HTTP Handler注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}
