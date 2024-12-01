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
 * Description: 角色领域层实现
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
	ErrInvalidUserId    = errors.New("无效的用户ID")
	ErrInvalidRoleIds   = errors.New("无效的角色ID列表")
	ErrInvalidRoleId    = errors.New("无效的角色ID")
	ErrEmptyRole        = errors.New("角色信息不能为空")
	ErrSystemRole       = errors.New("不允许创建系统角色")
	ErrEmptyPermissions = errors.New("菜单ID和API ID不能同时为空")
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
	if err := d.validateRole(role); err != nil {
		return err
	}
	return d.repo.CreateRole(ctx, role)
}

// GetRole 获取角色
func (d *RoleDomain) GetRole(ctx context.Context, id int) (*model.Role, error) {
	if id <= 0 {
		return nil, ErrInvalidRoleId
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
	if err := d.validateRole(role); err != nil {
		return err
	}

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
	if id <= 0 {
		return ErrInvalidRoleId
	}

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
	if roleId <= 0 {
		return ErrInvalidRoleId
	}

	role, err := d.GetRole(ctx, roleId)
	if err != nil {
		return err
	}

	if role == nil {
		return ErrRoleNotFound
	}

	if len(menuIds) == 0 && len(apiIds) == 0 {
		return ErrEmptyPermissions
	}

	return d.repo.AssignPermissions(ctx, roleId, menuIds, apiIds)
}

// AssignRoleToUser 分配角色给用户
func (d *RoleDomain) AssignRoleToUser(ctx context.Context, userId int, roleIds []int) error {
	if userId <= 0 {
		return ErrInvalidUserId
	}

	if len(roleIds) == 0 {
		return ErrInvalidRoleIds
	}

	for _, roleId := range roleIds {
		role, err := d.GetRole(ctx, roleId)
		if err != nil {
			return err
		}
		if role == nil {
			return ErrRoleNotFound
		}
	}

	return d.repo.AssignRoleToUser(ctx, userId, roleIds)
}

// RemoveUserPermissions 移除用户权限
func (d *RoleDomain) RemoveUserPermissions(ctx context.Context, userId int) error {
	if userId <= 0 {
		return ErrInvalidUserId
	}

	return d.repo.RemoveUserPermissions(ctx, userId)
}

// RemoveRoleFromUser 移除用户角色
func (d *RoleDomain) RemoveRoleFromUser(ctx context.Context, userId int, roleIds []int) error {
	if userId <= 0 {
		return ErrInvalidUserId
	}

	if len(roleIds) == 0 {
		return ErrInvalidRoleIds
	}

	for _, roleId := range roleIds {
		role, err := d.GetRole(ctx, roleId)
		if err != nil {
			return err
		}
		if role == nil {
			return ErrRoleNotFound
		}
	}

	return d.repo.RemoveRoleFromUser(ctx, userId, roleIds)
}

// validateRole 校验角色信息
func (d *RoleDomain) validateRole(role *model.Role) error {
	if role == nil {
		return ErrEmptyRole
	}

	if err := role.Validate(); err != nil {
		return err
	}

	if role.RoleType == model.RoleTypeSystem {
		return ErrSystemRole
	}

	return nil
}
