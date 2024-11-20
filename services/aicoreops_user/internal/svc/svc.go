package svc

import (
	"aicoreops_user/internal/config"
	"aicoreops_user/internal/pkg"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB := pkg.InitDB(c.Mysql.Addr)
	return &ServiceContext{
		Config: c,
		DB:     DB,
	}
}
