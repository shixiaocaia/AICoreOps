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
 */

package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_user/internal/model"
)

type UserRepo interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, user *model.User) error
	// GetUserById 根据ID获取用户
	GetUserById(ctx context.Context, id int) (*model.User, error)
	// GetUserByUsernameAndPassword 根据用户名和密码获取用户
	GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error)
	// GetUserByUsername 根据用户名获取用户
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	// UpdateUser 更新用户信息
	UpdateUser(ctx context.Context, user *model.User) error
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, id int) error
	// ListUsers 获取用户列表
	ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int, error)
	// UpdatePassword 更新用户密码
	UpdatePassword(ctx context.Context, id int, newPassword string) error
	// UpdateStatus 更新用户状态
	UpdateStatus(ctx context.Context, id int, status int) error
	// UpdateLastLoginTime 更新用户最后登录时间
	UpdateLastLoginTime(ctx context.Context, id int) error
	// GetUserAccessCodes 获取用户权限码
	GetUserAccessCodes(ctx context.Context, userId int) ([]string, error)
}
