package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	LLM        MyLLMConfig
	Qdrant     QdrantConfig
	DataSource string
}

type MyLLMConfig struct {
	Url   string
	Model string
}

type QdrantConfig struct {
	Url            string
	CollectionName string
	Model          string
}
