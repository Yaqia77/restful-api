package impl

import (
	"context"
	"restful-api/apps/host"

	"github.com/infraboard/mcube/logger"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	i.l.Debug("CreateHost")
	i.l.Debugf("CreateHost request: %v", ins)
	i.l.With(logger.NewAny("request", ins)).Debug("CreateHost")
	return nil, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.HostSet, error) {
	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.Host, error) {
	return nil, nil
}
func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (
	*host.Host, error) {
	return nil, nil
}
func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (
	*host.Host, error) {
	return nil, nil
}
