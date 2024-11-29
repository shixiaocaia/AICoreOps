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
 * Description: 菜单领域层实现
 */

package domain

import (
	"aicoreops_role/internal/dao"
	"aicoreops_role/internal/model"
	"aicoreops_role/internal/repo"
	"context"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrMenuNotFound     = errors.New("菜单不存在")
	ErrMenuNameRequired = errors.New("菜单名称不能为空")
	ErrInvalidMenuID    = errors.New("无效的菜单ID")
)

type MenuDomain struct {
	repo repo.MenuRepo
}

func NewMenuDomain(db *gorm.DB) *MenuDomain {
	return &MenuDomain{
		repo: dao.NewMenuDao(db),
	}
}

// CreateMenu 创建菜单
func (d *MenuDomain) CreateMenu(ctx context.Context, menu *model.Menu) error {
	if err := d.validateMenu(menu); err != nil {
		return err
	}
	return d.repo.CreateMenu(ctx, menu)
}

// GetMenu 获取菜单详情
func (d *MenuDomain) GetMenu(ctx context.Context, id int) (*model.Menu, error) {
	if id <= 0 {
		return nil, ErrInvalidMenuID
	}
	menu, err := d.repo.GetMenuById(ctx, id)
	if err != nil {
		return nil, err
	}
	if menu == nil {
		return nil, ErrMenuNotFound
	}
	return menu, nil
}

// UpdateMenu 更新菜单
func (d *MenuDomain) UpdateMenu(ctx context.Context, menu *model.Menu) error {
	if err := d.validateMenu(menu); err != nil {
		return err
	}
	// 检查菜单是否存在
	existingMenu, err := d.GetMenu(ctx, int(menu.ID))
	if err != nil {
		return err
	}
	if existingMenu == nil {
		return ErrMenuNotFound
	}
	return d.repo.UpdateMenu(ctx, menu)
}

// DeleteMenu 删除菜单
func (d *MenuDomain) DeleteMenu(ctx context.Context, id int) error {
	if id <= 0 {
		return ErrInvalidMenuID
	}
	// 检查菜单是否存在
	menu, err := d.GetMenu(ctx, id)
	if err != nil {
		return err
	}
	if menu == nil {
		return ErrMenuNotFound
	}
	return d.repo.DeleteMenu(ctx, id)
}

// ListMenus 获取菜单列表
func (d *MenuDomain) ListMenus(ctx context.Context, page, pageSize int) ([]*model.Menu, int, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, 0, errors.New("无效的分页参数")
	}
	return d.repo.ListMenus(ctx, page, pageSize)
}

// GetMenuTree 获取菜单树
func (d *MenuDomain) GetMenuTree(ctx context.Context) ([]*model.Menu, error) {
	return d.repo.GetMenuTree(ctx)
}

// validateMenu 验证菜单参数
func (d *MenuDomain) validateMenu(menu *model.Menu) error {
	if menu == nil {
		return errors.New("菜单对象不能为空")
	}
	
	if menu.Name == "" {
		return ErrMenuNameRequired
	}

	if menu.Path == "" {
		return errors.New("路由路径不能为空")
	}

	if menu.Component == "" {
		return errors.New("组件路径不能为空")
	}

	if menu.RouteName == "" {
		return errors.New("路由名称不能为空")
	}

	if menu.Hidden != model.MenuHiddenNo && menu.Hidden != model.MenuHiddenYes {
		return errors.New("菜单隐藏状态值无效")
	}

	return nil
}
