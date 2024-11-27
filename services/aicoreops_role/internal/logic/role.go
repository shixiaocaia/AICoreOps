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

package logic

import (
	"aicoreops_role/internal/svc"
	"aicoreops_role/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleLogic {
	return &RoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateRole 创建角色
func (l *RoleLogic) CreateRole(ctx context.Context, request *types.CreateRoleRequest) (*types.CreateRoleResponse, error) {
	return nil, nil
}

// GetRole 获取角色详情
func (l *RoleLogic) GetRole(ctx context.Context, request *types.GetRoleRequest) (*types.GetRoleResponse, error) {
	return nil, nil
}

// UpdateRole 更新角色
func (l *RoleLogic) UpdateRole(ctx context.Context, request *types.UpdateRoleRequest) (*types.UpdateRoleResponse, error) {
	return nil, nil
}

// DeleteRole 删除角色
func (l *RoleLogic) DeleteRole(ctx context.Context, request *types.DeleteRoleRequest) (*types.DeleteRoleResponse, error) {
	return nil, nil
}

// ListRoles 获取角色列表
func (l *RoleLogic) ListRoles(ctx context.Context, request *types.ListRolesRequest) (*types.ListRolesResponse, error) {
	return nil, nil
}
