package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql   DBConfig
	JWT     JWTConfig
	MyRedis MyRedisConfig
}

type DBConfig struct {
	Addr string
}

type JWTConfig struct {
	Secret string
	Expire int64
}

type MyRedisConfig struct {
	Addr string
}
