package svc

import (
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/config"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/pkg"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  redis.Cmdable
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := pkg.InitDB(c.Mysql)

	redis := pkg.InitRedis(c.XRedis)
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis,
	}
}
