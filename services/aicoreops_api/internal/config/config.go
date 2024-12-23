/*
 * Copyright 2024 Bamboo
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * File: config.go
 */

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	ApiRpc  zrpc.RpcClientConf
	MenuRpc zrpc.RpcClientConf
	RoleRpc zrpc.RpcClientConf
	AiRpc   zrpc.RpcClientConf
	JWT     JWTConfig
	MyRedis MyRedis
	Casbin  CasbinConfig
	Mysql   MysqlConfig
}

type JWTConfig struct {
	Secret string
}

type MyRedis struct {
	Addr string
}

type CasbinConfig struct {
	Path string
}

type MysqlConfig struct {
	Addr string
}
