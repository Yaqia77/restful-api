package impl

import (
	"context"
	"fmt"
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
	// var hosts []*host.Host
	hosts := host.NewHostSet()
	var totalCount int64

	// Base query
	// query := i.db.Model(&host.Resource{}).Joins("left join host on host.resource_id = id").Preload("Describe")
	query := i.db.Model(&host.Resource{}).Joins("left join host on host.resource_id = resource.id").Preload("Describe")

	// Apply keyword filter if provided
	if req.Keywords != "" {
		i.l.Debug("QueryHost", "keywords", req.Keywords)
		query = query.Where("resource.name LIKE ? OR resource.description LIKE ? OR resource.public_ip LIKE ? OR resource.private_ip LIKE ?", "%"+req.Keywords+"%", "%"+req.Keywords+"%", "%"+req.Keywords+"%", "%"+req.Keywords+"%")
	}

	// Count total records
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (req.PageNumber - 1) * req.PageSize
	if err := query.Limit(req.PageSize).Offset(offset).Find(&hosts.Items).Error; err != nil {
		return nil, err
	}
	fmt.Println("111111111111", hosts.Items)
	hostSet := &host.HostSet{
		Items: hosts.Items,
		Total: int(totalCount),
	}

	return hostSet, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	hosts := host.NewHost()
	query := i.db.Model(&host.Resource{}).Joins("left join host on host.resource_id = id").Preload("Describe")

	query = query.Where("resource.id = ? ", req.Id)

	if err := query.Find(&hosts).Error; err != nil {
		return nil, err
	}

	return hosts, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	// 获取已有对象
	ins, err := i.DescribeHost(ctx, host.NewDescribeHostRequestWithId(req.Id))
	if err != nil {
		return nil, err
	}
	// 设置 Resource 和 Describe 中的 ResourceID 为相同的值
	// ins.ResourceID = ins.Id

	fmt.Println("33333333", ins.Describe)
	// 根据更新的模式, 更新对象
	switch req.UpdateMode {
	case host.UPDATE_MODE_PUT:
		if err := ins.Put(req.Host); err != nil {
			return nil, err
		}
		// 整个对象的局部更新
	case host.UPDATE_MODE_PATCH:
		if err := ins.Patch(req.Host); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("update_mode only requred put/patch")
	}
	ins.Describe.ResourceID = ins.Resource.Id
	// 检查更新后的数据是否合法
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	fmt.Println("111", ins.Resource)

	fmt.Println("222", ins.Describe)
	// 获取要更新的资源ID
	// resourceID := ins.Resource.Id
	// 更新数据库里面的数据
	// 更新数据库里面的数据
	if err := i.db.Save(ins.Resource).Error; err != nil {
		return nil, err
	}
	if err := i.db.Save(ins.Describe).Error; err != nil {
		return nil, err
	}

	// 返回更新后的对象
	return ins, nil
}
func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	hosts := host.NewHost()
	query := i.db.Model(&host.Resource{}).Joins("left join host on host.resource_id = id").Preload("Describe")

	query = query.Where("resource.id = ? ", req.Id)

	if err := query.First(&hosts).Error; err != nil {
		return nil, err
	}
	fmt.Println("id1111111111", req.Id)
	//删除资源和描述
	deleteDescribe := func(tx *gorm.DB) error {
		if err := tx.Where("resource_id = ?", req.Id).Delete(&host.Describe{}).Error; err != nil {
			return err
		}
		return nil
	}

	deleteResource := func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", req.Id).Delete(&host.Resource{}).Error; err != nil {
			return err
		}
		return nil
	}

	i.l.Debug("DeleteHost", "transfer funds")

	// 使用事务删除资源和描述
	err := i.db.Transaction(func(tx *gorm.DB) error {
		if err := deleteDescribe(tx); err != nil {
			return err
		}
		if err := deleteResource(tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		i.l.Error("delete host failed", err)
		return nil, err
	}
	return hosts, nil
}
