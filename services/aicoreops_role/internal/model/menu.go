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
 * File: menu.go
 * Description: 菜单模型定义
 */

package model

import (
	"errors"
	"fmt"
)

// 定义菜单相关常量
const (
	MenuHiddenNo  = 0 // 菜单显示
	MenuHiddenYes = 1 // 菜单隐藏

	MaxMenuNameLength      = 50
	MaxMenuPathLength      = 255
	MaxMenuComponentLength = 255
	MaxMenuIconLength      = 50
	MaxMenuRouteNameLength = 50
)

// Menu 菜单模型
type Menu struct {
	ID         int64   `json:"id" gorm:"primaryKey;column:id;comment:菜单ID"`
	Name       string  `json:"name" gorm:"column:name;type:varchar(50);not null;comment:菜单显示名称"`
	ParentID   int64   `json:"parent_id" gorm:"column:parent_id;default:0;comment:上级菜单ID,0表示顶级菜单"`
	Path       string  `json:"path" gorm:"column:path;type:varchar(255);not null;comment:前端路由访问路径"`
	Component  string  `json:"component" gorm:"column:component;type:varchar(255);not null;comment:前端组件文件路径"`
	Icon       string  `json:"icon" gorm:"column:icon;type:varchar(50);default:'';comment:菜单显示图标"`
	SortOrder  int     `json:"sort_order" gorm:"column:sort_order;default:0;comment:菜单显示顺序,数值越小越靠前"`
	RouteName  string  `json:"route_name" gorm:"column:route_name;type:varchar(50);not null;comment:前端路由名称,需唯一"`
	Hidden     int     `json:"hidden" gorm:"column:hidden;type:tinyint(1);default:0;comment:菜单是否隐藏(0:显示 1:隐藏)"`
	CreateTime int64   `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:记录创建时间戳"`
	UpdateTime int64   `json:"update_time" gorm:"column:update_time;autoUpdateTime;comment:记录最后更新时间戳"`
	IsDeleted  int     `json:"is_deleted" gorm:"column:is_deleted;type:tinyint(1);default:0;comment:逻辑删除标记(0:未删除 1:已删除)"`
	Children   []*Menu `json:"children" gorm:"-"`
}

// TableName 返回表名
func (m *Menu) TableName() string {
	return "menus"
}

// Validate 验证菜单数据有效性
func (m *Menu) Validate() error {
	if m.Name == "" {
		return errors.New("菜单名称不能为空")
	}
	if len(m.Name) > MaxMenuNameLength {
		return fmt.Errorf("菜单名称长度不能超过%d个字符", MaxMenuNameLength)
	}

	if m.Path == "" {
		return errors.New("路由路径不能为空")
	}
	if len(m.Path) > MaxMenuPathLength {
		return fmt.Errorf("路由路径长度不能超过%d个字符", MaxMenuPathLength)
	}

	if m.Component == "" {
		return errors.New("组件路径不能为空")
	}
	if len(m.Component) > MaxMenuComponentLength {
		return fmt.Errorf("组件路径长度不能超过%d个字符", MaxMenuComponentLength)
	}

	if m.RouteName == "" {
		return errors.New("路由名称不能为空")
	}
	if len(m.RouteName) > MaxMenuRouteNameLength {
		return fmt.Errorf("路由名称长度不能超过%d个字符", MaxMenuRouteNameLength)
	}

	if m.Hidden != MenuHiddenNo && m.Hidden != MenuHiddenYes {
		return errors.New("无效的菜单隐藏状态")
	}

	return nil
}
