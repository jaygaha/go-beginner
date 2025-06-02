package svc

import (
	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_10_go-zero/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
