package model

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
 * File: user.go
 * Description:
 */

type User struct {
	ID            int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                     // 用户ID
	Username      string `gorm:"column:username;type:varchar(50);not null;unique" json:"username"` // 用户名
	Password      string `gorm:"column:password;type:varchar(100);not null" json:"password"`       // 加密后的密码
	Email         string `gorm:"column:email;type:varchar(100)" json:"email"`                      // 用户邮箱
	Phone         string `gorm:"column:phone;type:varchar(20)" json:"phone"`                       // 用户电话
	Nickname      string `gorm:"column:nickname;type:varchar(50)" json:"nickname"`                 // 用户昵称
	Avatar        string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`                    // 用户头像URL
	Status        int    `gorm:"column:status;type:tinyint;default:1" json:"status"`               // 用户状态（1:活跃, 2:不活跃, 3:被封禁）
	LastLoginTime int64  `gorm:"column:last_login_time;type:int" json:"last_login_time"`           // 最后登录时间
	CreateTime    int64  `gorm:"column:create_time;type:int;autoCreateTime" json:"create_time"`    // 创建时间
	UpdateTime    int64  `gorm:"column:update_time;type:int;autoUpdateTime" json:"update_time"`    // 更新时间
	IsDeleted     int    `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`       // 软删除标志（0:否, 1:是）
}

func (*User) TableName() string {
	return "users"
}
