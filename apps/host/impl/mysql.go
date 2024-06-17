package impl

import (
	"restful-api/apps/host"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// "restful-api/apps/host"

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostService(l logger.Logger) *HostServiceImpl {
	return &HostServiceImpl{

		l: zap.L().Named("Host"),
	}
}

type HostServiceImpl struct {
	l logger.Logger
}
