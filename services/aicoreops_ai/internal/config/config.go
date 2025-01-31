package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	LLM    LLMConfig
	Qdrant QdrantConfig
	MySQL  string
}

type LLMConfig struct {
	Url   string
	Model string
}

type QdrantConfig struct {
	Url            string
	CollectionName string
	Model          string
}
