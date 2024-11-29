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
 * File: init.go
 * Description:
 */

package domain

import (
	"aicoreops_role/internal/dao"
	"aicoreops_role/internal/model"
	"aicoreops_role/internal/repo"
	"context"
	"errors"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

var (
	ErrRoleNotFound     = errors.New("角色不存在")
	ErrRoleNameRequired = errors.New("角色名称不能为空")
	ErrInvalidPageSize  = errors.New("无效的分页大小")
)

type RoleDomain struct {
	repo repo.RoleRepo
}

func NewRoleDomain(db *gorm.DB, enforcer *casbin.Enforcer) *RoleDomain {
	return &RoleDomain{
		repo: dao.NewRoleDao(db, enforcer),
	}
}

// CreateRole 创建角色
func (d *RoleDomain) CreateRole(ctx context.Context, role *model.Role) error {
	// 参数校验
	if err := d.validateRole(role); err != nil {
		return err
	}

	return d.repo.CreateRole(ctx, role)
}

// GetRole 获取角色
func (d *RoleDomain) GetRole(ctx context.Context, id int) (*model.Role, error) {
	// 参数校验
	if id <= 0 {
		return nil, errors.New("无效的角色ID")
	}

	role, err := d.repo.GetRoleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, ErrRoleNotFound
	}

	return role, nil
}

// UpdateRole 更新角色
func (d *RoleDomain) UpdateRole(ctx context.Context, role *model.Role) error {
	// 参数校验
	if err := d.validateRole(role); err != nil {
		return err
	}

	// 检查角色是否存在
	existingRole, err := d.GetRole(ctx, int(role.ID))
	if err != nil {
		return err
	}
	if existingRole == nil {
		return ErrRoleNotFound
	}

	return d.repo.UpdateRole(ctx, role)
}

// DeleteRole 删除角色
func (d *RoleDomain) DeleteRole(ctx context.Context, id int) error {
	// 参数校验
	if id <= 0 {
		return errors.New("无效的角色ID")
	}

	// 检查角色是否存在
	role, err := d.GetRole(ctx, id)
	if err != nil {
		return err
	}

	if role == nil {
		return ErrRoleNotFound
	}

	return d.repo.DeleteRole(ctx, id)
}

// ListRoles 获取角色列表
func (d *RoleDomain) ListRoles(ctx context.Context, page, pageSize int) ([]*model.Role, int, error) {
	// 参数校验
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 100 {
		return nil, 0, ErrInvalidPageSize
	}

	return d.repo.ListRoles(ctx, page, pageSize)
}

// AssignPermissions 分配权限
func (d *RoleDomain) AssignPermissions(ctx context.Context, roleId int, menuIds []int, apiIds []int) error {
	// 参数校验
	if roleId <= 0 {
		return errors.New("无效的角色ID")
	}

	// 检查角色是否存在
	role, err := d.GetRole(ctx, roleId)
	if err != nil {
		return err
	}

	if role == nil {
		return ErrRoleNotFound
	}

	// 校验菜单ID和API ID
	if len(menuIds) == 0 && len(apiIds) == 0 {
		return errors.New("菜单ID和API ID不能同时为空")
	}

	return d.repo.AssignPermissions(ctx, roleId, menuIds, apiIds)
}

// validateRole 校验角色信息
func (d *RoleDomain) validateRole(role *model.Role) error {
	if role == nil {
		return errors.New("角色信息不能为空")
	}

	// 使用模型自带的验证
	if err := role.Validate(); err != nil {
		return err
	}

	// 判断是否为管理员角色
	if role.RoleType == model.RoleTypeSystem {
		return errors.New("不允许创建系统角色")
	}

	return nil
}

