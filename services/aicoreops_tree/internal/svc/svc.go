package svc

import (
	"aicoreops_tree/internal/config"
	"aicoreops_tree/internal/pkg"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := pkg.InitDB(c.Mysql.Addr)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
