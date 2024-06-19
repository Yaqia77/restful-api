package http

import (
	"restful-api/apps/host"

	"github.com/gin-gonic/gin"
)

// 面向接口，真正Service的实现，在服务实例化的时候传递进行
var handler = &Handler{}

// Handler 通过写一个实例类，把内部的接口通过http协议暴露出去
type Handler struct {
	svc host.Service
}

// NewHandler 实例化Handler
func NewHostHttpHandler(svc host.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
