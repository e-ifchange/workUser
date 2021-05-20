package svc

import "github.com/suhanyujie/workUser/service/workUser/cmd/rpc/user/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
