package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql DBConfig
}

type DBConfig struct {
	Addr string
}
