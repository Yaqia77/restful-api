package host

import (
	"fmt"
	"net/http"
	"strconv"

	"dario.cat/mergo"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	validate = validator.New()
)

func NewQueryHostFromHTTP(r *http.Request) *QueryHostRequest {
	req := NewQueryHostRequest()
	qs := r.URL.Query()

	pss := qs.Get("page_size")
	if pss != "" {
		req.PageSize, _ = strconv.Atoi(pss)
	}
	pns := qs.Get("page_number")
	if pns != "" {
		req.PageNumber, _ = strconv.Atoi(pns)
	}
	req.Keywords = qs.Get("kws")

	return req
}

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

type HostSet struct {
	Total int     `json:"total"`
	Items []*Host `json:"items"`
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
	Id          string   `json:"id" gorm:"primarykey;"`       // 全局唯一Id主键ID
	Vendor      Vendor   `json:"vendor"`                      // 厂商
	Region      string   `json:"region"  validate:"required"` // 地域
	CreateAt    int64    `json:"create_at"`                   // 创建时间
	ExpireAt    int64    `json:"expire_at"`                   // 过期时间
	Type        string   `json:"type"  validate:"required"`   // 规格
	Name        string   `json:"name"  validate:"required"`   // 名称
	Description string   `json:"description"`                 // 描述
	Status      string   `json:"status"`                      // 服务商中的状态
	Tags        string   `json:"tags"`                        // 标签
	UpdateAt    int64    `json:"update_at"`                   // 更新时间
	SyncAt      int64    `json:"sync_at"`                     // 同步时间
	Account     string   `json:"accout"`                      // 资源的所属账号
	PublicIP    string   `json:"public_ip"`                   // 公网IP
	PrivateIP   string   `json:"private_ip"`                  // 内网IP
	Describe    Describe `json:"describe" gorm:"foreignKey:ResourceID;references:Id"`
}

type Describe struct {
	ResourceID   string `json:"resource_id" gorm:"primarykey"` // 主机ID
	CPU          int    `json:"cpu" validate:"-"`              // 核数
	Memory       int    `json:"memory" validate:"-"`           // 内存
	GPUAmount    int    `json:"gpu_amount"`                    // GPU数量
	GPUSpec      string `json:"gpu_spec"`                      // GPU类型
	OSType       string `json:"os_type"`                       // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                       // 操作系统名称
	SerialNumber string `json:"serial_number"`                 // 序列号
}

func (req *QueryHostRequest) GetPageSize() uint {
	return uint(req.PageSize)
}

func (req *QueryHostRequest) OffSet() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

type QueryHostRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
}

// ValidateStruct validates the struct using the validator library
// func (u *UpdateHostRequest) ValidateStruct() error {
// 	validate := validator.New()
// 	return validate.Struct(u)
// }

// updateStruct 使用反射更新结构体的字段
// func UpdateStruct(dest, src interface{}) error {
// 	destVal := reflect.ValueOf(dest).Elem()
// 	srcVal := reflect.ValueOf(src).Elem()

//		for i := 0; i < srcVal.NumField(); i++ {
//			field := srcVal.Field(i)
//			if !field.IsNil() {
//				destVal.Field(i).Set(field.Elem())
//			}
//		}
//		return nil
//	}
//
// 对象全量更新
func (h *Host) Put(obj *Host) error {
	if obj.Id != h.Id {
		return fmt.Errorf("id not equal")
	}
	*h.Resource = *obj.Resource
	*h.Describe = *obj.Describe
	fmt.Println("*h.Describe的值为:", *h.Describe)
	return nil
}

// 对象的局部更新
func (h *Host) Patch(obj *Host) error {
	// 保留当前的 CPU 和 Memory 值
	currentCPU := h.Resource.Describe.CPU
	currentMemory := h.Resource.Describe.Memory
	fmt.Println("currentCPU的值为:", currentCPU)
	// 使用传入的值替换现有值（如果有）
	if obj.Describe.CPU != 0 {
		h.Describe.CPU = obj.Describe.CPU
	}
	if obj.Describe.Memory != 0 {
		h.Describe.Memory = obj.Describe.Memory
	}

	// 使用 mergo.MergeWithOverwrite 进行合并
	if err := mergo.MergeWithOverwrite(h, obj); err != nil {
		return err
	}

	// 如果 CPU 和 Memory 没有传入新值，则恢复原值
	if obj.Describe.CPU == 0 {
		h.Describe.CPU = currentCPU
	}
	if obj.Describe.Memory == 0 {
		h.Describe.Memory = currentMemory
	}
	fmt.Println("h.Describe.CPU的值为:", h.Describe.CPU)
	return nil
}

func (h *Host) Validate() error {
	return validate.Struct(h)
}

func NewPutUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_MODE_PUT,
		Host:       h,
	}
}

func NewPatchUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_MODE_PATCH,
		Host:       h,
	}
}

type UpdateHostRequest struct {
	UpdateMode UPDATE_MODE `json:"update_mode"`
	*Host
}
type UPDATE_MODE string

const (
	// 全量更新
	UPDATE_MODE_PUT UPDATE_MODE = "put"
	// 局部更新
	UPDATE_MODE_PATCH UPDATE_MODE = "patch"
)

func NewDescribeHostRequestWithId(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}

type DescribeHostRequest struct {
	Id string
}

func NewDeleteHostRequestWithId(id string) *DeleteHostRequest {
	return &DeleteHostRequest{
		Id: id,
	}
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
