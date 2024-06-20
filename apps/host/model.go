package host

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	validate = validator.New()
)

type HostSet struct {
	Items []*Host
	Total int
}

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

type Host struct {
	//资源公共属性部分
	*Resource
	//资源独有属性部分
	*Describe
}

type Vendor int

const (
	//枚举的默认值
	PrivateIDC Vendor = iota + 1
	//阿里云
	ALIYUN
	//腾讯云
	TXYUN
)

type Resource struct {
	Id          string `json:"id" gorm:"primarykey"`       // 全局唯一Id主键ID
	Vendor      Vendor `json:"vendor"`                     // 厂商
	Region      string `json:"region"  binding:"required"` // 地域
	CreateAt    int64  `json:"create_at"`                  // 创建时间
	ExpireAt    int64  `json:"expire_at"`                  // 过期时间
	Type        string `json:"type"  binding:"required"`   // 规格
	Name        string `json:"name"  binding:"required"`   // 名称
	Description string `json:"description"`                // 描述
	Status      string `json:"status"`                     // 服务商中的状态
	Tags        string `json:"tags"`                       // 标签
	UpdateAt    int64  `json:"update_at"`                  // 更新时间
	SyncAt      int64  `json:"sync_at"`                    // 同步时间
	Account     string `json:"accout"`                     // 资源的所属账号
	PublicIP    string `json:"public_ip"`                  // 公网IP
	PrivateIP   string `json:"private_ip"`                 // 内网IP
}

type Describe struct {
	ResourceID   string `json:"resource_id" gorm:"primarykey"` // 主机ID
	CPU          int    `json:"cpu" binding:"required"`        // 核数
	Memory       int    `json:"memory" binding:"required"`     // 内存
	GPUAmount    int    `json:"gpu_amount"`                    // GPU数量
	GPUSpec      string `json:"gpu_spec"`                      // GPU类型
	OSType       string `json:"os_type"`                       // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                       // 操作系统名称
	SerialNumber string `json:"serial_number"`                 // 序列号
}

type QueryHostRequest struct {
}

type UpdateHostRequest struct {
	*Describe
}

type DescribeHostRequest struct {
	Id string
}
type DeleteHostRequest struct {
	Id string
}

func (Resource) TableName() string {
	return "resource"
}

func (Describe) TableName() string {
	return "host"
}

func AutoMigrateResource(db *gorm.DB) error {
	if err := db.AutoMigrate(&Resource{}); err != nil {
		return err
	}
	return nil
}

func AutoMigrateDescribe(db *gorm.DB) error {
	if err := db.AutoMigrate(&Describe{}); err != nil {
		return err
	}
	return nil
}