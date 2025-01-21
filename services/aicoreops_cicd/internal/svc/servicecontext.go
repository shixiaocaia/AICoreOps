package svc

import "github.com/GoSimplicity/AICoreOps/services/aicoreops_cicd/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
