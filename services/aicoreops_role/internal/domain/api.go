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

package domain

import (
	"aicoreops_role/internal/dao"
	"aicoreops_role/internal/model"
	"aicoreops_role/internal/repo"
	"aicoreops_role/types"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ApiDomain struct {
	repo repo.ApiRepo
}

func NewApiDomain(db *gorm.DB) *ApiDomain {
	return &ApiDomain{
		repo: dao.NewApiDao(db),
	}
}

// CreateApi 创建API
func (d *ApiDomain) CreateApi(ctx context.Context, api *model.Api) error {
	return d.repo.CreateApi(ctx, api)
}

// GetApi 获取API详情
func (d *ApiDomain) GetApi(ctx context.Context, id int) (*model.Api, error) {
	return d.repo.GetApiById(ctx, id)
}

// UpdateApi 更新API
func (d *ApiDomain) UpdateApi(ctx context.Context, api *model.Api) error {
	return d.repo.UpdateApi(ctx, api)
}

// DeleteApi 删除API
func (d *ApiDomain) DeleteApi(ctx context.Context, id int) error {
	return d.repo.DeleteApi(ctx, id)
}

// ListApis 获取API列表
func (d *ApiDomain) ListApis(ctx context.Context, page, pageSize int) ([]*model.Api, int, error) {
	return d.repo.ListApis(ctx, page, pageSize)
}

// ValidateApi 验证API参数
func (d *ApiDomain) ValidateApi(api *types.UpdateApiRequest) error {
	if api == nil {
		return errors.New("API对象不能为空")
	}

	if api.Name == "" {
		return errors.New("API名称不能为空")
	}

	if api.Path == "" {
		return errors.New("API路径不能为空")
	}

	if api.Method <= 0 {
		return errors.New("请指定有效的HTTP方法")
	}

	return nil
}

// BuildApi 构建API对象
func (d *ApiDomain) BuildApi(api *types.UpdateApiRequest) *model.Api {
	return &model.Api{
		ID:          api.Id,
		Name:        api.Name,
		Path:        api.Path,
		Method:      int(api.Method),
		Description: api.Description,
		Version:     api.Version,
		Category:    int(api.Category),
		IsPublic:    int(api.IsPublic),
	}
}
