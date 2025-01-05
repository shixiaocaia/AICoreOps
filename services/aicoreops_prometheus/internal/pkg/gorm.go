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
 * File: gorm.go
 */

package pkg

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化数据库
func InitDB(addr string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	// 初始化表
	if err = InitTables(db); err != nil {
		panic(err)
	}

	return db
}

// HandleList 处理搜索或获取所有记录
func HandleList[T any](ctx context.Context, search *string, searchFunc func(ctx context.Context, name string) ([]*T, error), listFunc func(ctx context.Context) ([]*T, error)) ([]*T, error) {
	if search != nil && *search != "" {
		return searchFunc(ctx, *search)
	}
	return listFunc(ctx)
}
