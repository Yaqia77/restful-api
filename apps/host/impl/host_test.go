package impl_test

import (
	"context"
	"fmt"
	"restful-api/apps/host"
	"restful-api/apps/host/impl"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
)

var (
	service host.Service
)

func TestCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

func TestInit(t *testing.T) {

	fmt.Println(zap.DevelopmentSetup())

	service = impl.NewHostServiceImpl()

	TestCreate(t)
}
