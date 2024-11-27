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
 * File: role.go
 * Description: 角色数据访问层实现
 */

package dao

import (
	"aicoreops_role/internal/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db: db}
}

// CreateRole 创建角色
func (r RoleDao) CreateRole(ctx context.Context, role *model.Role) error {
	if role == nil {
		return gorm.ErrRecordNotFound
	}

	now := time.Now().Unix()
	role.CreateTime = now
	role.UpdateTime = now

	return r.db.WithContext(ctx).Create(role).Error
}

// GetRoleById 根据ID获取角色
func (r RoleDao) GetRoleById(ctx context.Context, id int) (*model.Role, error) {
	var role model.Role

	if err := r.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", id).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

// UpdateRole 更新角色
func (r RoleDao) UpdateRole(ctx context.Context, role *model.Role) error {
	if role == nil {
		return gorm.ErrRecordNotFound
	}

	role.UpdateTime = time.Now().Unix()

	return r.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", role.ID).Updates(role).Error
}

// DeleteRole 删除角色(软删除)
func (r RoleDao) DeleteRole(ctx context.Context, id int) error {
	updates := map[string]interface{}{
		"is_deleted":  1,
		"update_time": time.Now().Unix(),
	}

	return r.db.WithContext(ctx).Model(&model.Role{}).Where("id = ? AND is_deleted = 0", id).Updates(updates).Error
}

// ListRoles 获取角色列表
func (r RoleDao) ListRoles(ctx context.Context, page, pageSize int) ([]*model.Role, int, error) {
	var roles []*model.Role
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Role{}).Where("is_deleted = 0")

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Order("id DESC").Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, int(total), nil
}

// AssignPermissions 分配权限
func (r RoleDao) AssignPermissions(ctx context.Context, roleId int, menuIds []int, apiIds []int) error {
	// 检查角色是否存在
	if _, err := r.GetRoleById(ctx, roleId); err != nil {
		return err
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除原有的菜单权限
		if err := tx.Where("role_id = ?", roleId).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}

		// 批量添加新的菜单权限
		if len(menuIds) > 0 {
			roleMenus := make([]*model.RoleMenu, 0, len(menuIds))
			now := time.Now().Unix()
			for _, menuId := range menuIds {
				roleMenus = append(roleMenus, &model.RoleMenu{
					RoleID:    int64(roleId),
					MenuID:    int64(menuId),
					CreatedAt: now,
				})
			}
			if err := tx.Create(&roleMenus).Error; err != nil {
				return err
			}
		}

		// 删除原有的API权限
		if err := tx.Where("role_id = ?", roleId).Delete(&model.RoleApi{}).Error; err != nil {
			return err
		}

		// 批量添加新的API权限
		if len(apiIds) > 0 {
			roleApis := make([]*model.RoleApi, 0, len(apiIds))
			now := time.Now().Unix()
			for _, apiId := range apiIds {
				roleApis = append(roleApis, &model.RoleApi{
					RoleID:    int64(roleId),
					ApiID:     int64(apiId),
					CreatedAt: now,
				})
			}
			if err := tx.Create(&roleApis).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
