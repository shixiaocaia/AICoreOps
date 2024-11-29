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
 * File: api.go
 * Description: API模型定义及相关常量、方法
 */

package model

import (
	"errors"
	"fmt"
)

const (
	// 长度限制常量
	MaxApiNameLength        = 50
	MaxApiPathLength        = 255
	MaxApiDescriptionLength = 500
	MaxApiVersionLength     = 20

	// HTTP方法常量
	MethodGet    = 1
	MethodPost   = 2
	MethodPut    = 3
	MethodDelete = 4

	// API分类常量
	CategorySystem = 1 // 系统API
	CategoryBiz    = 2 // 业务API

	// 是否公开常量
	PublicNo  = 0 // 不公开
	PublicYes = 1 // 公开
)

// Api API模型
type Api struct {
	ID          int64  `json:"id" gorm:"primaryKey;column:id;comment:主键ID"`
	Name        string `json:"name" gorm:"column:name;type:varchar(50);not null;comment:API名称"`
	Path        string `json:"path" gorm:"column:path;type:varchar(255);not null;comment:API路径"`
	Method      int    `json:"method" gorm:"column:method;type:tinyint(1);not null;comment:HTTP请求方法(1:GET,2:POST,3:PUT,4:DELETE)"`
	Description string `json:"description" gorm:"column:description;type:varchar(500);comment:API描述"`
	Version     string `json:"version" gorm:"column:version;type:varchar(20);default:v1;comment:API版本"`
	Category    int    `json:"category" gorm:"column:category;type:tinyint(1);not null;comment:API分类(1:系统,2:业务)"`
	IsPublic    int    `json:"is_public" gorm:"column:is_public;type:tinyint(1);default:0;comment:是否公开(0:否,1:是)"`
	CreateTime  int64  `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdateTime  int64  `json:"update_time" gorm:"column:update_time;autoUpdateTime;comment:更新时间"`
	IsDeleted   int    `json:"is_deleted" gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除(0:否,1:是)"`
}

// TableName 返回表名
func (a *Api) TableName() string {
	return "apis"
}

// Validate 验证API数据有效性
func (a *Api) Validate() error {
	// 验证名称
	if a.Name == "" {
		return errors.New("API名称不能为空")
	}
	if len(a.Name) > MaxApiNameLength {
		return fmt.Errorf("API名称长度不能超过%d个字符", MaxApiNameLength)
	}

	// 验证路径
	if a.Path == "" {
		return errors.New("API路径不能为空")
	}
	if len(a.Path) > MaxApiPathLength {
		return fmt.Errorf("API路径长度不能超过%d个字符", MaxApiPathLength)
	}

	// 验证描述
	if len(a.Description) > MaxApiDescriptionLength {
		return fmt.Errorf("API描述长度不能超过%d个字符", MaxApiDescriptionLength)
	}

	// 验证版本
	if len(a.Version) > MaxApiVersionLength {
		return fmt.Errorf("API版本长度不能超过%d个字符", MaxApiVersionLength)
	}

	// 验证HTTP方法
	if a.Method < MethodGet || a.Method > MethodDelete {
		return errors.New("无效的HTTP请求方法")
	}

	// 验证分类
	if a.Category != CategorySystem && a.Category != CategoryBiz {
		return errors.New("无效的API分类")
	}

	// 验证是否公开
	if a.IsPublic != PublicNo && a.IsPublic != PublicYes {
		return errors.New("无效的公开状态")
	}

	return nil
}
