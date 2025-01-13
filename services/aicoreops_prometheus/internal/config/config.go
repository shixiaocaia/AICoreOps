package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql            string
	XRedis           string
	PrometheusConfig PrometheusConfig
}

type PrometheusConfig struct {
	LocalYamlDir string
	HttpSdAPI    string
}
