package impl

import (
	"restful-api/apps/host"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)



// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{

		l: zap.L().Named("Host"),
	}
}

type HostServiceImpl struct {
	l logger.Logger
}
