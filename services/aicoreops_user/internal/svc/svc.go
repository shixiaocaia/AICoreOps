package svc

import (
	"aicoreops_user/internal/config"
	"aicoreops_user/internal/pkg"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  redis.Cmdable
	JWT    pkg.JWTHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := pkg.InitDB(c.Mysql.Addr)
	r := pkg.InitRedis(c.MyRedis)
	jwt := pkg.NewJWTHandler(r, c.JWT)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  r,
		JWT:    jwt,
	}
}
