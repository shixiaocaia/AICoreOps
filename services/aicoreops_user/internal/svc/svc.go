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
	"aicoreops_user/internal/config"
	"aicoreops_user/internal/pkg"

	sf "github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	Redis     redis.Cmdable
	JWT       pkg.JWTHandler
	Snowflake *sf.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := pkg.InitDB(c.Mysql.Addr)
	r := pkg.InitRedis(c.MyRedis)
	jwt := pkg.NewJWTHandler(r, c.JWT)
	snowflakeNode := pkg.InitializeSnowflakeNode()
	return &ServiceContext{
		Config:    c,
		DB:        db,
		Redis:     r,
		JWT:       jwt,
		Snowflake: snowflakeNode,
	}
}
