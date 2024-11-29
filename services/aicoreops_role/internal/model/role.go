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
 * File: role.go
 * Description: 角色相关模型定义
 */

package model

import (
	"errors"
)

const (
	RoleTypeSystem = 1 // 系统角色
	RoleTypeCustom = 2 // 自定义角色

	DefaultRoleNo  = 0 // 非默认角色
	DefaultRoleYes = 1 // 默认角色
)

// Role 角色模型
type Role struct {
	ID          int64  `json:"id" gorm:"primaryKey;column:id;comment:主键ID"`
	Name        string `json:"name" gorm:"column:name;type:varchar(50);not null;unique;comment:角色名称"`
	Description string `json:"description" gorm:"column:description;type:varchar(255);comment:角色描述"`
	RoleType    int    `json:"role_type" gorm:"column:role_type;type:tinyint(1);not null;comment:角色类型(1:系统角色,2:自定义角色)"`
	IsDefault   int    `json:"is_default" gorm:"column:is_default;type:tinyint(1);default:0;comment:是否为默认角色(0:否,1:是)"`
	CreateTime  int64  `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdateTime  int64  `json:"update_time" gorm:"column:update_time;autoUpdateTime;comment:更新时间"`
	IsDeleted   int    `json:"is_deleted" gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除(0:否,1:是)"`
}

func (r *Role) TableName() string {
	return "roles"
}

// Validate 验证角色结构体
func (r *Role) Validate() error {
	if r.Name == "" {
		return errors.New("角色名称不能为空")
	}

	if len(r.Name) > 50 {
		return errors.New("角色名称长度不能超过50个字符")
	}

	if r.RoleType != RoleTypeSystem && r.RoleType != RoleTypeCustom {
		return errors.New("无效的角色类型")
	}

	if r.IsDefault != DefaultRoleNo && r.IsDefault != DefaultRoleYes {
		return errors.New("无效的默认角色标识")
	}

	return nil
}
