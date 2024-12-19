package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	LLM    MyLLMConfig
	Qdrant QdrantConfig
	MySQL  MySQLConfig
}

type MyLLMConfig struct {
	Url   string
	Model string
}

type QdrantConfig struct {
	Url            string
	CollectionName string
	Model          string
	DocumentPath   string
}

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
