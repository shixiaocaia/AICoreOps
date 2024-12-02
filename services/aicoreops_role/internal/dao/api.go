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
 */

package dao

import (
	"aicoreops_role/internal/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type ApiDao struct {
	db *gorm.DB
}

func NewApiDao(db *gorm.DB) *ApiDao {
	return &ApiDao{
		db: db,
	}
}

// CreateApi 创建API
func (a *ApiDao) CreateApi(ctx context.Context, api *model.Api) error {
	if api == nil {
		return gorm.ErrRecordNotFound
	}

	api.CreateTime = time.Now().Unix()
	api.UpdateTime = time.Now().Unix()

	return a.db.WithContext(ctx).Create(api).Error
}

// GetApiById 根据ID获取API
func (a *ApiDao) GetApiById(ctx context.Context, id int) (*model.Api, error) {
	var api model.Api

	if err := a.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", id).First(&api).Error; err != nil {
		return nil, err
	}

	return &api, nil
}

// UpdateApi 更新API
func (a *ApiDao) UpdateApi(ctx context.Context, api *model.Api) error {
	if api == nil {
		return gorm.ErrRecordNotFound
	}

	updates := map[string]interface{}{
		"name":        api.Name,
		"path":        api.Path,
		"method":      api.Method,
		"description": api.Description,
		"version":     api.Version,
		"category":    api.Category,
		"is_public":   api.IsPublic,
		"update_time": time.Now().Unix(),
	}

	return a.db.WithContext(ctx).
		Model(&model.Api{}).
		Where("id = ? AND is_deleted = 0", api.ID).
		Updates(updates).Error
}

// DeleteApi 删除API(软删除)
func (a *ApiDao) DeleteApi(ctx context.Context, id int) error {
	updates := map[string]interface{}{
		"is_deleted":  1,
		"update_time": time.Now().Unix(),
	}

	return a.db.WithContext(ctx).Model(&model.Api{}).Where("id = ? AND is_deleted = 0", id).Updates(updates).Error
}

// ListApis 获取API列表
func (a *ApiDao) ListApis(ctx context.Context, page, pageSize int) ([]*model.Api, int, error) {
	var apis []*model.Api
	var total int64

	db := a.db.WithContext(ctx).Model(&model.Api{}).Where("is_deleted = 0")

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Order("id ASC").Find(&apis).Error; err != nil {
		return nil, 0, err
	}

	return apis, int(total), nil
}
