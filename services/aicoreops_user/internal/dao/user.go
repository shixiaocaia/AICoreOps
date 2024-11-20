package dao

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

import (
	"aicoreops_user/internal/model"
	"context"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

// CreateUser 创建用户
func (d *UserDao) CreateUser(ctx context.Context, user *model.User) error {
	return d.db.WithContext(ctx).Create(user).Error
}

// GetUser 根据ID获取用户
func (d *UserDao) GetUser(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户
func (d *UserDao) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).Where("username = ? AND is_deleted = 0", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (d *UserDao) UpdateUser(ctx context.Context, user *model.User) error {
	return d.db.WithContext(ctx).Save(user).Error
}

// DeleteUser 删除用户
func (d *UserDao) DeleteUser(ctx context.Context, id int) error {
	return d.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

// ListUsers 获取用户列表
func (d *UserDao) ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int, error) {
	var users []*model.User
	var total int64

	offset := (page - 1) * pageSize

	err := d.db.WithContext(ctx).Model(&model.User{}).Where("is_deleted = 0").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = d.db.WithContext(ctx).Where("is_deleted = 0").Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

// UpdatePassword 更新用户密码
func (d *UserDao) UpdatePassword(ctx context.Context, id int, newPassword string) error {
	return d.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("password", newPassword).Error
}

// UpdateStatus 更新用户状态
func (d *UserDao) UpdateStatus(ctx context.Context, id int, status int) error {
	return d.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("status", status).Error
}
