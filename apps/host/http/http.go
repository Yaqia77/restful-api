package http

import (
	"restful-api/apps"
	"restful-api/apps/host"

	"github.com/gin-gonic/gin"
)

// 面向接口，真正Service的实现，在服务实例化的时候传递进行
var handler = &Handler{}

// Handler 通过写一个实例类，把内部的接口通过http协议暴露出去
type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {

	//从IOC里面获取HostService的实例对象
	h.svc = apps.GetImpl(host.AppName).(host.Service)
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
	r.GET("/hosts", h.queryHost)
	r.GET("/hosts/:id", h.describeHost)
}

func (h *Handler) Name() string {
	return host.AppName
}

func init() {
	apps.RegistryGin(handler)
}
