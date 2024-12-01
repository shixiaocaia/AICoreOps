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
 * Description: 角色数据层接口
 */

package repo

import (
	"aicoreops_role/internal/model"
	"context"
)

type RoleRepo interface {
	// CreateRole 创建角色
	CreateRole(ctx context.Context, role *model.Role) error
	// GetRoleById 根据ID获取角色
	GetRoleById(ctx context.Context, id int) (*model.Role, error)
	// UpdateRole 更新角色信息
	UpdateRole(ctx context.Context, role *model.Role) error
	// DeleteRole 删除角色
	DeleteRole(ctx context.Context, id int) error
	// ListRoles 获取角色列表
	ListRoles(ctx context.Context, page, pageSize int) ([]*model.Role, int, error)
	// AssignPermissions 分配权限
	AssignPermissions(ctx context.Context, roleId int, menuIds []int, apiIds []int) error
	// AssignRoleToUser 分配角色给用户
	AssignRoleToUser(ctx context.Context, userId int, roleIds []int) error
	// RemoveUserPermissions 移除用户权限
	RemoveUserPermissions(ctx context.Context, userId int) error
	// RemoveRoleFromUser 移除用户角色
	RemoveRoleFromUser(ctx context.Context, userId int, roleIds []int) error
}
