package impl

import (
	"restful-api/apps"
	"restful-api/apps/host"
	"restful-api/conf"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

// 接口实现的静态检查
// var _ host.Service = (*HostServiceImpl)(nil)

var impl = &HostServiceImpl{}

// NewHostServiceImpl 保证调用该函数之前，全局conf对象已经初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{

		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

func (i *HostServiceImpl) Config() {

	//Host service 服务的子Loggger
	//封装的Zap让其满足 Logger接口
	i.l = zap.L().Named("Host")
	i.db = conf.C().MySQL.GetDB()

}

type HostServiceImpl struct {
	l  logger.Logger
	db *gorm.DB
}

// 返回服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

func init() {
	//对象注册到IOC层
	apps.RegistryImpl(impl)

	// apps.HostService = impl
}
