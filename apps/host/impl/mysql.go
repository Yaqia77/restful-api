package impl

import (
	"restful-api/apps/host"
	"restful-api/conf"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

// NewHostServiceImpl 保证调用该函数之前，全局conf对象已经初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{

		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *gorm.DB
}
