package impl_test

import (
	"context"
	"fmt"
	"restful-api/apps/host"
	"restful-api/apps/host/impl"
	"restful-api/conf"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
)

var (
	service host.Service
)

//查询详情信息
func TestDescribe(t *testing.T) {
	should := assert.New(t)

	// 创建一个Host实例
	// ins := host.NewQueryHostRequest()

	// id := &host.DescribeHostRequest{
	// 	Id: "test-02",
	// }
	id := host.NewDescribeHostRequestWithId("test-02")
	result, err := service.DescribeHost(context.Background(), id)
	if should.NoError(err) {
		fmt.Println(result.Id)
	}
}

func TestInit(t *testing.T) {

	err := conf.LoadConfigFromToml("../../../etc/demo.toml")
	if err != nil {
		panic(err)
	}
	//需要初始化全局Logger
	//为什么不涉及韦默认打印，因为性能
	fmt.Println(zap.DevelopmentSetup())
	// host service的具体实现
	service = impl.NewHostServiceImpl()

	TestDescribe(t)
}

//查询主机列表测试
// func TestQuery(t *testing.T) {
// 	should := assert.New(t)

// 	// 创建一个Host实例
// 	// ins := host.NewQueryHostRequest()

// 	// 设置Host的属性
// 	ins := &host.QueryHostRequest{
// 		PageSize:   10,
// 		PageNumber: 1,
// 		Keywords:   "11.10",
// 	}
// 	result, err := service.QueryHost(context.Background(), ins)
// 	if should.NoError(err) {
// 		for i := range result.Items {
// 			fmt.Println(result.Items[i].Id)
// 		}
// 	}
// }

// func TestInit(t *testing.T) {

// 	err := conf.LoadConfigFromToml("../../../etc/demo.toml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	//需要初始化全局Logger
// 	//为什么不涉及韦默认打印，因为性能
// 	fmt.Println(zap.DevelopmentSetup())
// 	// host service的具体实现
// 	service = impl.NewHostServiceImpl()

// 	TestQuery(t)
// }

//创建主机测试
// func TestCreate(t *testing.T) {
// 	should := assert.New(t)

// 	// 创建一个Host实例
// 	ins := host.NewHost()

// 	// 设置Host的属性
// 	ins.Id = "test-02"
// 	ins.Region = "广州"
// 	ins.Type = "small"
// 	ins.Name = "接口测试主机"
// 	ins.ResourceID = "test-02"
// 	ins.CPU = 1
// 	ins.Memory = 2048

// 	ins, err := service.CreateHost(context.Background(), ins)
// 	if should.NoError(err) {
// 		fmt.Println(ins)
// 	}
// }

// func TestInit(t *testing.T) {

// 	err := conf.LoadConfigFromToml("../../../etc/demo.toml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	//需要初始化全局Logger
// 	//为什么不涉及韦默认打印，因为性能
// 	fmt.Println(zap.DevelopmentSetup())
// 	// host service的具体实现
// 	service = impl.NewHostServiceImpl()

// 	TestCreate(t)
// }
