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
	r.POST("/hosts", h.createHost)       //创建
	r.GET("/hosts", h.queryHost)         //查询列表
	r.GET("/hosts/:id", h.describeHost)  //查询单个
	r.PUT("/hosts/:id", h.putHost)       //全量更新
	r.PATCH("/hosts/:id", h.patchHost)   //局部更新
	r.DELETE("/hosts/:id", h.deleteHost) //删除
}

func (h *Handler) Name() string {
	return host.AppName
}

func init() {
	apps.RegistryGin(handler)
}
