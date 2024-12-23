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
 * File: role_logic.go
 */

package logic

import (
	"context"
	"time"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/types"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_common/types/role"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleLogic {
	return &RoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateRole 创建角色
func (l *RoleLogic) CreateRole(req *types.CreateRoleRequest) (*role.CreateRoleResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	createReq := &role.CreateRoleRequest{}
	if err := copier.Copy(createReq, req); err != nil {
		return nil, err
	}

	createResp, err := l.svcCtx.RoleRpc.CreateRole(ctx, createReq)
	if err != nil {
		return nil, err
	}

	return createResp, nil
}

// GetRole 获取角色详情
func (l *RoleLogic) GetRole(req *types.GetRoleRequest) (*role.GetRoleResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	getReq := &role.GetRoleRequest{}
	if err := copier.Copy(getReq, req); err != nil {
		return nil, err
	}

	getResp, err := l.svcCtx.RoleRpc.GetRole(ctx, getReq)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

// UpdateRole 更新角色
func (l *RoleLogic) UpdateRole(req *types.UpdateRoleRequest) (*role.UpdateRoleResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	updateReq := &role.UpdateRoleRequest{}
	if err := copier.Copy(updateReq, req); err != nil {
		return nil, err
	}

	updateResp, err := l.svcCtx.RoleRpc.UpdateRole(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return updateResp, nil
}

// DeleteRole 删除角色
func (l *RoleLogic) DeleteRole(req *types.DeleteRoleRequest) (*role.DeleteRoleResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	deleteReq := &role.DeleteRoleRequest{}
	if err := copier.Copy(deleteReq, req); err != nil {
		return nil, err
	}

	deleteResp, err := l.svcCtx.RoleRpc.DeleteRole(ctx, deleteReq)
	if err != nil {
		return nil, err
	}

	// 加载最新的策略
	if err := l.svcCtx.Enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return deleteResp, nil
}

// ListRoles 获取角色列表
func (l *RoleLogic) ListRoles(req *types.ListRolesRequest) (*role.ListRolesResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	listReq := &role.ListRolesRequest{}
	if err := copier.Copy(listReq, req); err != nil {
		return nil, err
	}

	listResp, err := l.svcCtx.RoleRpc.ListRoles(ctx, listReq)
	if err != nil {
		return nil, err
	}

	return listResp, nil
}

// AssignPermissions 分配权限
func (l *RoleLogic) AssignPermissions(req *types.AssignPermissionsRequest) (*role.AssignPermissionsResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	assignReq := &role.AssignPermissionsRequest{}
	if err := copier.Copy(assignReq, req); err != nil {
		return nil, err
	}

	assignResp, err := l.svcCtx.RoleRpc.AssignPermissions(ctx, assignReq)
	if err != nil {
		return nil, err
	}

	// 加载最新的策略
	if err := l.svcCtx.Enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return assignResp, nil
}

// AssignRoleToUser 分配角色给用户
func (l *RoleLogic) AssignRoleToUser(req *types.AssignRoleToUserRequest) (*role.AssignRoleToUserResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	assignReq := &role.AssignRoleToUserRequest{}
	if err := copier.Copy(assignReq, req); err != nil {
		return nil, err
	}

	assignResp, err := l.svcCtx.RoleRpc.AssignRoleToUser(ctx, assignReq)
	if err != nil {
		return nil, err
	}

	// 加载最新的策略
	if err := l.svcCtx.Enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return assignResp, nil
}

// RemoveUserPermissions 移除用户权限
func (l *RoleLogic) RemoveUserPermissions(req *types.RemoveUserPermissionsRequest) (*role.RemoveUserPermissionsResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	removeReq := &role.RemoveUserPermissionsRequest{}
	if err := copier.Copy(removeReq, req); err != nil {
		return nil, err
	}

	removeResp, err := l.svcCtx.RoleRpc.RemoveUserPermissions(ctx, removeReq)
	if err != nil {
		return nil, err
	}

	// 加载最新的策略
	if err := l.svcCtx.Enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return removeResp, nil
}

// RemoveRoleFromUser 移除用户角色
func (l *RoleLogic) RemoveRoleFromUser(req *types.RemoveRoleFromUserRequest) (*role.RemoveRoleFromUserResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	removeReq := &role.RemoveRoleFromUserRequest{}
	if err := copier.Copy(removeReq, req); err != nil {
		return nil, err
	}

	removeResp, err := l.svcCtx.RoleRpc.RemoveRoleFromUser(ctx, removeReq)
	if err != nil {
		return nil, err
	}

	// 加载最新的策略
	if err := l.svcCtx.Enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return removeResp, nil
}
