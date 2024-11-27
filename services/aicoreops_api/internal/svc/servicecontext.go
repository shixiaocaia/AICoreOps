package svc

import (
	"aicoreops_api/internal/config"
	"aicoreops_common/types/user"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserServiceClient
	RDB     redis.Cmdable
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpc).Conn()),
		RDB: redis.NewClient(&redis.Options{
			Addr: c.MyRedis.Addr,
		}),
	}
}
