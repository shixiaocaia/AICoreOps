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

package repo

import (
	"aicoreops_role/internal/model"
	"context"
)

type ApiRepo interface {
	// CreateApi 创建API
	CreateApi(ctx context.Context, api *model.Api) error
	// GetApiById 根据ID获取API
	GetApiById(ctx context.Context, id int) (*model.Api, error)
	// UpdateApi 更新API信息
	UpdateApi(ctx context.Context, api *model.Api) error
	// DeleteApi 删除API
	DeleteApi(ctx context.Context, id int) error
	// ListApis 获取API列表
	ListApis(ctx context.Context, page, pageSize int) ([]*model.Api, int, error)
}
