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

// Menu 菜单模型
type Menu struct {
	ID         int64   `json:"id" gorm:"primaryKey;column:id;comment:主键ID"`
	Name       string  `json:"name" gorm:"column:name;type:varchar(50);not null;comment:菜单名称"`
	ParentID   int64   `json:"parent_id" gorm:"column:parent_id;default:0;comment:父菜单ID"`
	Path       string  `json:"path" gorm:"column:path;type:varchar(255);not null;comment:路由路径"`
	Component  string  `json:"component" gorm:"column:component;type:varchar(255);not null;comment:前端组件路径"`
	Icon       string  `json:"icon" gorm:"column:icon;type:varchar(50);default:'';comment:菜单图标"`
	SortOrder  int32   `json:"sort_order" gorm:"column:sort_order;default:0;comment:排序值(数值越小排序越靠前)"`
	RouteName  string  `json:"route_name" gorm:"column:route_name;type:varchar(50);not null;comment:路由名称"`
	Hidden     int32   `json:"hidden" gorm:"column:hidden;type:tinyint(1);default:0;comment:是否隐藏(0:显示,1:隐藏)"`
	CreateTime int64   `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdateTime int64   `json:"update_time" gorm:"column:update_time;autoUpdateTime;comment:更新时间"`
	IsDeleted  int32   `json:"is_deleted" gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除(0:否,1:是)"`
	Children   []*Menu `json:"children" gorm:"-"`
}

func (m *Menu) TableName() string {
	return "menus"
}
