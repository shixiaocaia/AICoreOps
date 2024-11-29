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
	"aicoreops_role/internal/constant"
	"aicoreops_role/internal/model"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var (
	ErrMenuNotFound = errors.New("菜单不存在")
	ErrInvalidMenu  = errors.New("无效的菜单参数")
)

type MenuDao struct {
	db *gorm.DB
}

func NewMenuDao(db *gorm.DB) *MenuDao {
	return &MenuDao{db: db}
}

// CreateMenu 创建菜单
func (m *MenuDao) CreateMenu(ctx context.Context, menu *model.Menu) error {
	if menu == nil {
		return ErrInvalidMenu
	}

	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查父菜单是否存在
		if menu.ParentID != 0 {
			var count int64
			if err := tx.Model(&model.Menu{}).Where("id = ? AND is_deleted = ?", menu.ParentID, constant.DeletedNo).Count(&count).Error; err != nil {
				return fmt.Errorf("检查父菜单失败: %v", err)
			}
			if count == 0 {
				return errors.New("父菜单不存在")
			}
		}

		now := time.Now().Unix()
		menu.CreateTime = now
		menu.UpdateTime = now

		return tx.Create(menu).Error
	})
}

// GetMenuById 根据ID获取菜单
func (m *MenuDao) GetMenuById(ctx context.Context, id int) (*model.Menu, error) {
	if id <= 0 {
		return nil, errors.New("无效的菜单ID")
	}

	var menu model.Menu
	if err := m.db.WithContext(ctx).Where("id = ? AND is_deleted = ?", id, constant.DeletedNo).First(&menu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMenuNotFound
		}
		return nil, fmt.Errorf("查询菜单失败: %v", err)
	}

	return &menu, nil
}

// UpdateMenu 更新菜单
func (m *MenuDao) UpdateMenu(ctx context.Context, menu *model.Menu) error {
	if menu == nil {
		return errors.New("菜单对象不能为空")
	}
	if menu.ID <= 0 {
		return errors.New("无效的菜单ID")
	}

	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查父菜单是否存在且不能将菜单设置为自己的子菜单
		if menu.ParentID != 0 {
			if menu.ParentID == menu.ID {
				return errors.New("不能将菜单设置为自己的子菜单")
			}
			var count int64
			if err := tx.Model(&model.Menu{}).Where("id = ? AND is_deleted = ?", menu.ParentID, constant.DeletedNo).Count(&count).Error; err != nil {
				return fmt.Errorf("检查父菜单失败: %v", err)
			}
			if count == 0 {
				return errors.New("父菜单不存在")
			}
		}

		updates := map[string]interface{}{
			"name":        menu.Name,
			"parent_id":   menu.ParentID,
			"path":        menu.Path,
			"component":   menu.Component,
			"icon":        menu.Icon,
			"sort_order":  menu.SortOrder,
			"route_name":  menu.RouteName,
			"hidden":      menu.Hidden,
			"update_time": time.Now().Unix(),
		}

		result := tx.Model(&model.Menu{}).
			Where("id = ? AND is_deleted = ?", menu.ID, constant.DeletedNo).
			Updates(updates)
		if result.Error != nil {
			return fmt.Errorf("更新菜单失败: %v", result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.New("菜单不存在或已被删除")
		}

		return nil
	})
}

// DeleteMenu 删除菜单(软删除)
func (m *MenuDao) DeleteMenu(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("无效的菜单ID")
	}

	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查是否有子菜单
		var count int64
		if err := tx.Model(&model.Menu{}).Where("parent_id = ? AND is_deleted = ?", id, constant.DeletedNo).Count(&count).Error; err != nil {
			return fmt.Errorf("检查子菜单失败: %v", err)
		}
		if count > 0 {
			return errors.New("存在子菜单,不能删除")
		}

		updates := map[string]interface{}{
			"is_deleted":  constant.DeletedYes,
			"update_time": time.Now().Unix(),
		}
		result := tx.Model(&model.Menu{}).Where("id = ? AND is_deleted = ?", id, constant.DeletedNo).Updates(updates)
		if result.Error != nil {
			return fmt.Errorf("删除菜单失败: %v", result.Error)
		}
		if result.RowsAffected == 0 {
			return ErrMenuNotFound
		}
		return nil
	})
}

// ListMenus 获取菜单列表
func (m *MenuDao) ListMenus(ctx context.Context, page, pageSize int) ([]*model.Menu, int, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, 0, errors.New("无效的分页参数")
	}

	var menus []*model.Menu
	var total int64

	db := m.db.WithContext(ctx).Model(&model.Menu{}).Where("is_deleted = ?", constant.DeletedNo)

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("获取菜单总数失败: %v", err)
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Order("sort_order ASC, id DESC").Find(&menus).Error; err != nil {
		return nil, 0, fmt.Errorf("查询菜单列表失败: %v", err)
	}

	return menus, int(total), nil
}

// GetMenuTree 获取菜单树
func (m *MenuDao) GetMenuTree(ctx context.Context) ([]*model.Menu, error) {
	// 预分配合适的初始容量
	menus := make([]*model.Menu, 0, 50)

	// 使用索引字段优化查询
	if err := m.db.WithContext(ctx).
		Select("id, name, parent_id, path, component, icon, sort_order, route_name, hidden, create_time, update_time").
		Where("is_deleted = ?", constant.DeletedNo).
		Order("sort_order ASC, id ASC").
		Find(&menus).Error; err != nil {
		return nil, fmt.Errorf("查询菜单列表失败: %v", err)
	}

	// 预分配map容量
	menuMap := make(map[int64]*model.Menu, len(menus))
	rootMenus := make([]*model.Menu, 0, len(menus)/3)

	// 单次遍历构建树形结构
	for _, menu := range menus {
		// 初始化Children切片
		menu.Children = make([]*model.Menu, 0, 4)
		menuMap[menu.ID] = menu

		if menu.ParentID == 0 {
			rootMenus = append(rootMenus, menu)
		} else if parent, exists := menuMap[menu.ParentID]; exists {
			parent.Children = append(parent.Children, menu)
		}
	}

	return rootMenus, nil
}
