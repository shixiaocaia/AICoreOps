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
	"aicoreops_role/internal/constant"
	"aicoreops_role/internal/model"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

// RoleDao 角色数据访问层结构体
type RoleDao struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
	apiDao   *ApiDao
}

// NewRoleDao 创建角色数据访问层实例
func NewRoleDao(db *gorm.DB, enforcer *casbin.Enforcer) *RoleDao {
	return &RoleDao{
		db:       db,
		enforcer: enforcer,
		apiDao:   NewApiDao(db),
	}
}

// CreateRole 创建角色
func (r *RoleDao) CreateRole(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查角色名是否已存在
		var count int64
		if err := tx.Model(&model.Role{}).Where("name = ? AND is_deleted = ?", role.Name, constant.DeletedNo).Count(&count).Error; err != nil {
			return fmt.Errorf("检查角色名称失败: %v", err)
		}
		if count > 0 {
			return errors.New("角色名称已存在")
		}

		// 创建角色
		if err := tx.Create(role).Error; err != nil {
			return fmt.Errorf("创建角色失败: %v", err)
		}

		return nil
	})
}

// GetRoleById 根据ID获取角色
func (r *RoleDao) GetRoleById(ctx context.Context, id int) (*model.Role, error) {
	if id <= 0 {
		return nil, errors.New("无效的角色ID")
	}

	var role model.Role
	if err := r.db.WithContext(ctx).Where("id = ? AND is_deleted = ?", id, constant.DeletedNo).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("角色不存在: %v", err)
		}
		return nil, fmt.Errorf("查询角色失败: %v", err)
	}

	return &role, nil
}

// UpdateRole 更新角色
func (r *RoleDao) UpdateRole(ctx context.Context, role *model.Role) error {
	if role == nil {
		return errors.New("角色对象不能为空")
	}
	if role.ID <= 0 {
		return errors.New("无效的角色ID")
	}

	updates := map[string]interface{}{
		"name":        role.Name,
		"description": role.Description,
		"role_type":   role.RoleType,
		"is_default":  role.IsDefault,
		"update_time": time.Now().Unix(),
	}

	result := r.db.WithContext(ctx).
		Model(&model.Role{}).
		Where("id = ? AND is_deleted = ?", role.ID, constant.DeletedNo).
		Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("更新角色失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("角色不存在或已被删除")
	}

	return nil
}

// DeleteRole 删除角色(软删除)
func (r *RoleDao) DeleteRole(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("无效的角色ID")
	}

	updates := map[string]interface{}{
		"is_deleted":  constant.DeletedYes,
		"update_time": time.Now().Unix(),
	}

	result := r.db.WithContext(ctx).Model(&model.Role{}).Where("id = ? AND is_deleted = ?", id, constant.DeletedNo).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("删除角色失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("角色不存在或已被删除")
	}

	return nil
}

// ListRoles 获取角色列表
func (r *RoleDao) ListRoles(ctx context.Context, page, pageSize int) ([]*model.Role, int, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, 0, errors.New("无效的分页参数")
	}

	var roles []*model.Role
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Role{}).Where("is_deleted = 0")

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("获取角色总数失败: %v", err)
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Order("id ASC").Find(&roles).Error; err != nil {
		return nil, 0, fmt.Errorf("获取角色列表失败: %v", err)
	}

	return roles, int(total), nil
}

// AssignPermissions 分配权限
func (r *RoleDao) AssignPermissions(ctx context.Context, roleId int, menuIds []int, apiIds []int) error {
	// 使用常量替代魔法数字
	const batchSize = 1000

	// 参数校验
	if roleId <= 0 {
		return errors.New("无效的角色ID")
	}

	// 检查角色是否存在
	role, err := r.GetRoleById(ctx, roleId)
	if err != nil {
		return fmt.Errorf("获取角色失败: %v", err)
	}
	if role == nil {
		return errors.New("角色不存在")
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除原有的casbin规则
		if _, err := r.enforcer.DeleteRolesForUser(role.Name); err != nil {
			return fmt.Errorf("删除原有权限失败: %v", err)
		}

		// 添加菜单权限
		if err := r.assignMenuPermissions(role.Name, menuIds, batchSize); err != nil {
			return err
		}

		// 添加API权限
		if err := r.assignAPIPermissions(ctx, role.Name, apiIds, batchSize); err != nil {
			return err
		}

		// 加载最新的策略
		if err := r.enforcer.LoadPolicy(); err != nil {
			return fmt.Errorf("加载策略失败: %v", err)
		}

		return nil
	})
}

// assignMenuPermissions 分配菜单权限
func (r *RoleDao) assignMenuPermissions(roleName string, menuIds []int, batchSize int) error {
	// 如果菜单ID列表为空,直接返回
	if len(menuIds) == 0 {
		return nil
	}

	// 构建casbin策略规则
	policies := make([][]string, 0, len(menuIds))
	for _, menuId := range menuIds {
		if menuId <= 0 {
			return fmt.Errorf("无效的菜单ID: %d", menuId)
		}
		policies = append(policies, []string{roleName, fmt.Sprintf("menu:%d", menuId), "read"})
	}

	// 批量添加策略
	return r.batchAddPolicies(policies, batchSize)
}

// assignAPIPermissions 分配API权限
func (r *RoleDao) assignAPIPermissions(ctx context.Context, roleName string, apiIds []int, batchSize int) error {
	// 如果API ID列表为空,直接返回
	if len(apiIds) == 0 {
		return nil
	}

	// HTTP方法映射表
	methodMap := map[int]string{
		1: "GET",
		2: "POST",
		3: "PUT",
		4: "DELETE",
		5: "PATCH",
		6: "OPTIONS",
		7: "HEAD",
	}

	// 构建casbin策略规则
	policies := make([][]string, 0, len(apiIds))
	for _, apiId := range apiIds {
		if apiId <= 0 {
			return fmt.Errorf("无效的API ID: %d", apiId)
		}

		// 获取API信息
		api, err := r.apiDao.GetApiById(ctx, apiId)
		if err != nil {
			return fmt.Errorf("获取API信息失败: %v", err)
		}
		if api == nil {
			return fmt.Errorf("API不存在: %d", apiId)
		}

		// 获取HTTP方法
		method, ok := methodMap[api.Method]
		if !ok {
			return fmt.Errorf("无效的HTTP方法: %d", api.Method)
		}

		policies = append(policies, []string{roleName, api.Path, strings.ToLower(method)})
	}

	// 批量添加策略
	return r.batchAddPolicies(policies, batchSize)
}

// batchAddPolicies 批量添加策略
func (r *RoleDao) batchAddPolicies(policies [][]string, batchSize int) error {
	// 按批次处理策略规则
	for i := 0; i < len(policies); i += batchSize {
		end := i + batchSize
		if end > len(policies) {
			end = len(policies)
		}

		// 添加一批策略规则
		if _, err := r.enforcer.AddPolicies(policies[i:end]); err != nil {
			return fmt.Errorf("添加权限策略失败: %v", err)
		}
	}

	return nil
}
