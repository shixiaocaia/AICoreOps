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
 * Description: 菜单数据访问层实现
 */

package dao

import (
	"aicoreops_role/internal/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type MenuDao struct {
	db *gorm.DB
}

func NewMenuDao(db *gorm.DB) *MenuDao {
	return &MenuDao{db: db}
}

// CreateMenu 创建菜单
func (m MenuDao) CreateMenu(ctx context.Context, menu *model.Menu) error {
	if menu == nil {
		return gorm.ErrRecordNotFound
	}

	now := time.Now().Unix()
	menu.CreateTime = now
	menu.UpdateTime = now

	return m.db.WithContext(ctx).Create(menu).Error
}

// GetMenuById 根据ID获取菜单
func (m MenuDao) GetMenuById(ctx context.Context, id int) (*model.Menu, error) {
	var menu model.Menu

	if err := m.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", id).First(&menu).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}

// UpdateMenu 更新菜单
func (m MenuDao) UpdateMenu(ctx context.Context, menu *model.Menu) error {
	if menu == nil {
		return gorm.ErrRecordNotFound
	}

	menu.UpdateTime = time.Now().Unix()

	return m.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", menu.ID).Updates(menu).Error
}

// DeleteMenu 删除菜单(软删除)
func (m MenuDao) DeleteMenu(ctx context.Context, id int) error {
	updates := map[string]interface{}{
		"is_deleted":  1,
		"update_time": time.Now().Unix(),
	}

	return m.db.WithContext(ctx).Model(&model.Menu{}).Where("id = ? AND is_deleted = 0", id).Updates(updates).Error
}

// ListMenus 获取菜单列表
func (m MenuDao) ListMenus(ctx context.Context, page, pageSize int) ([]*model.Menu, int, error) {
	var menus []*model.Menu
	var total int64

	db := m.db.WithContext(ctx).Model(&model.Menu{}).Where("is_deleted = 0")

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Order("sort_order, id DESC").Find(&menus).Error; err != nil {
		return nil, 0, err
	}

	return menus, int(total), nil
}

// GetMenuTree 获取菜单树
func (m MenuDao) GetMenuTree(ctx context.Context) ([]*model.Menu, error) {
	var menus []*model.Menu

	// 先获取所有未删除的菜单
	if err := m.db.WithContext(ctx).Where("is_deleted = 0").Order("sort_order, id").Find(&menus).Error; err != nil {
		return nil, err
	}

	// 构建菜单树
	menuMap := make(map[int64]*model.Menu, len(menus))
	var rootMenus []*model.Menu

	// 将所有菜单放入map中
	for _, menu := range menus {
		menuMap[menu.ID] = menu
		menu.Children = make([]*model.Menu, 0) // 初始化Children切片
	}

	// 构建树形结构
	for _, menu := range menus {
		if menu.ParentID == 0 {
			rootMenus = append(rootMenus, menu)
		} else {
			if parent, ok := menuMap[menu.ParentID]; ok {
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	return rootMenus, nil
}
