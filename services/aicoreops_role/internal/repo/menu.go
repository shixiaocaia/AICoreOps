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
 */

package repo

import (
	"aicoreops_role/internal/model"
	"context"
)

type MenuRepo interface {
	// CreateMenu 创建菜单
	CreateMenu(ctx context.Context, menu *model.Menu) error
	// GetMenuById 根据ID获取菜单
	GetMenuById(ctx context.Context, id int) (*model.Menu, error)
	// UpdateMenu 更新菜单信息
	UpdateMenu(ctx context.Context, menu *model.Menu) error
	// DeleteMenu 删除菜单
	DeleteMenu(ctx context.Context, id int) error
	// ListMenus 获取菜单列表
	ListMenus(ctx context.Context, page, pageSize int) ([]*model.Menu, int, error)
	// GetMenuTree 获取菜单树
	GetMenuTree(ctx context.Context) ([]*model.Menu, error)
}
