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
 * File: svc.go
 */

package svc

import (
	"aicoreops_api/internal/config"
	"aicoreops_common/types/role"
	"aicoreops_common/types/user"
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  user.UserServiceClient
	ApiRpc   role.ApiServiceClient
	MenuRpc  role.MenuServiceClient
	RoleRpc  role.RoleServiceClient
	RDB      redis.Cmdable
	Enforcer *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.MyRedis.Addr,
		Password: "",
		DB:       0,
	})

	// 初始化用户RPC客户端
	userRpc := user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpc).Conn())
	apiRpc := role.NewApiServiceClient(zrpc.MustNewClient(c.ApiRpc).Conn())
	menuRpc := role.NewMenuServiceClient(zrpc.MustNewClient(c.MenuRpc).Conn())
	roleRpc := role.NewRoleServiceClient(zrpc.MustNewClient(c.RoleRpc).Conn())

	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(c.Mysql.Addr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}

	// 初始化Casbin适配器
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("初始化Casbin适配器失败: %v", err))
	}

	// 初始化Casbin enforcer
	enforcer, err := casbin.NewEnforcer(c.Casbin.Path, adapter)
	if err != nil {
		panic(fmt.Sprintf("初始化Casbin失败: %v", err))
	}

	// 加载Casbin策略
	if err := enforcer.LoadPolicy(); err != nil {
		panic(fmt.Sprintf("加载Casbin策略失败: %v", err))
	}

	return &ServiceContext{
		Config:   c,
		UserRpc:  userRpc,
		ApiRpc:   apiRpc,
		MenuRpc:  menuRpc,
		RoleRpc:  roleRpc,
		RDB:      rdb,
		Enforcer: enforcer,
	}
}
