package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql              string
	XRedis             string
	PrometheusConfig   PrometheusConfig
	AlertManagerConfig AlertManagerConfig
}

type PrometheusConfig struct {
	LocalYamlDir string
	HttpSdAPI    string
}

type AlertManagerConfig struct {
	LocalYamlDir     string
	AlertWebhookAddr string
}
