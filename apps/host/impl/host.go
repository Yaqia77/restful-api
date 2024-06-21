package impl

import (
	"context"
	"restful-api/apps/host"
	"restful-api/conf"
	"restful-api/pkg/utils"
	"time"

	"github.com/infraboard/mcube/logger"
	"gorm.io/gorm"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	i.l.Debug("CreateHost")
	i.l.Debugf("CreateHost request: %v", ins)
	i.l.With(logger.NewAny("request", ins)).Debug("CreateHost")
	ins.Id = utils.GenerateId(8)
	ins.ResourceID = ins.Id

	//添加创建时间
	ins.CreateAt = time.Now().Unix()
	ins.UpdateAt = ins.CreateAt

	//插入资源和描述
	insertResource := func(db *gorm.DB) error {
		return db.Create(ins.Resource).Error
	}

	insertDescribe := func(db *gorm.DB) error {
		return db.Create(ins.Describe).Error
	}

	i.l.Debug("CreateHost", "transfer funds")

	err := conf.TransferFunds(i.db, insertResource, insertDescribe)
	if err != nil {
		i.l.Error("create host failed", err)
		return nil, err
	}
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
