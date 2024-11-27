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
 * File: init.go
 * Description:
 */

package model

// Api API模型
type Api struct {
	ID          int64  `json:"id" gorm:"primaryKey;column:id;comment:主键ID"`
	Name        string `json:"name" gorm:"column:name;not null;comment:API名称"`
	Path        string `json:"path" gorm:"column:path;not null;comment:API路径"`
	Method      int32  `json:"method" gorm:"column:method;not null;comment:HTTP请求方法(1:GET,2:POST,3:PUT,4:DELETE)"`
	Description string `json:"description" gorm:"column:description;comment:API描述"`
	Version     string `json:"version" gorm:"column:version;default:v1;comment:API版本"`
	Category    int32  `json:"category" gorm:"column:category;not null;comment:API分类(1:系统,2:业务)"`
	IsPublic    int32  `json:"is_public" gorm:"column:is_public;default:0;comment:是否公开(0:否,1:是)"`
	CreateTime  int64  `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdateTime  int64  `json:"update_time" gorm:"column:update_time;autoUpdateTime;comment:更新时间"`
	IsDeleted   int32  `json:"is_deleted" gorm:"column:is_deleted;default:0;comment:是否删除(0:否,1:是)"`
}

func (a *Api) TableName() string {
	return "apis"
}
