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
 * Description: 角色管理逻辑层
 */

package logic

import (
	"aicoreops_role/internal/domain"
	"aicoreops_role/internal/model"
	"aicoreops_role/internal/svc"
	"aicoreops_role/types"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleLogic struct {
	ctx    context.Context
	domain *domain.RoleDomain
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleLogic {
	return &RoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		domain: domain.NewRoleDomain(svcCtx.DB, svcCtx.Casbin),
	}
}

// 将角色模型转换为响应类型
func convertToRoleResponse(role *model.Role) *types.Role {
	return &types.Role{
		Id:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		RoleType:    types.RoleType(role.RoleType),
		IsDefault:   int32(role.IsDefault),
		CreateTime:  role.CreateTime,
		UpdateTime:  role.UpdateTime,
	}
}

// CreateRole 创建角色
func (l *RoleLogic) CreateRole(ctx context.Context, req *types.CreateRoleRequest) (*types.CreateRoleResponse, error) {
	// 参数验证
	if req.Name == "" {
		return nil, errors.New("角色名称不能为空")
	}

	role := &model.Role{
		Name:        req.Name,
		Description: req.Description,
		RoleType:    int(req.RoleType),

		IsDefault: int(req.IsDefault),
	}

	// 验证角色模型
	if err := role.Validate(); err != nil {
		l.Errorf("角色参数验证失败: %v", err)
		return nil, err
	}

	if err := l.domain.CreateRole(ctx, role); err != nil {
		l.Errorf("创建角色失败: %v", err)
		return nil, err
	}

	l.Infof("角色创建成功, ID: %d, 名称: %s", role.ID, role.Name)
	return &types.CreateRoleResponse{
		Code:    0,
		Message: "创建角色成功",
	}, nil
}

// GetRole 获取角色详情
func (l *RoleLogic) GetRole(ctx context.Context, req *types.GetRoleRequest) (*types.GetRoleResponse, error) {
	role, err := l.domain.GetRole(ctx, int(req.Id))
	if err != nil {
		l.Errorf("获取角色详情失败: %v", err)
		return nil, err
	}

	return &types.GetRoleResponse{
		Code:    0,
		Message: "获取角色详情成功",
		Data:    convertToRoleResponse(role),
	}, nil
}

// UpdateRole 更新角色
func (l *RoleLogic) UpdateRole(ctx context.Context, req *types.UpdateRoleRequest) (*types.UpdateRoleResponse, error) {
	role := &model.Role{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		RoleType:    int(req.RoleType),
		IsDefault:   int(req.IsDefault),
	}

	if err := l.domain.UpdateRole(ctx, role); err != nil {
		l.Errorf("更新角色失败: %v", err)
		return nil, err
	}

	return &types.UpdateRoleResponse{
		Code:    0,
		Message: "更新角色成功",
	}, nil
}

// DeleteRole 删除角色
func (l *RoleLogic) DeleteRole(ctx context.Context, req *types.DeleteRoleRequest) (*types.DeleteRoleResponse, error) {
	if err := l.domain.DeleteRole(ctx, int(req.Id)); err != nil {
		l.Errorf("删除角色失败: %v", err)
		return nil, err
	}

	return &types.DeleteRoleResponse{
		Code:    0,
		Message: "删除角色成功",
	}, nil
}

// ListRoles 获取角色列表
func (l *RoleLogic) ListRoles(ctx context.Context, req *types.ListRolesRequest) (*types.ListRolesResponse, error) {
	roles, total, err := l.domain.ListRoles(ctx, int(req.PageNumber), int(req.PageSize))
	if err != nil {
		l.Errorf("获取角色列表失败: %v", err)
		return nil, err
	}

	roleList := make([]*types.Role, 0, len(roles))
	for _, role := range roles {
		roleList = append(roleList, convertToRoleResponse(role))
	}

	return &types.ListRolesResponse{
		Code:    0,
		Message: "获取角色列表成功",
		Data: &types.ListRolesData{
			Roles:      roleList,
			Total:      int32(total),
			PageNumber: req.PageNumber,
			PageSize:   req.PageSize,
		},
	}, nil
}

// AssignPermissions 分配权限
func (l *RoleLogic) AssignPermissions(ctx context.Context, req *types.AssignPermissionsRequest) (*types.AssignPermissionsResponse, error) {
	menuIds := make([]int, len(req.MenuIds))
	for i, id := range req.MenuIds {
		menuIds[i] = int(id)
	}

	apiIds := make([]int, len(req.ApiIds))
	for i, id := range req.ApiIds {
		apiIds[i] = int(id)
	}

	if err := l.domain.AssignPermissions(ctx, int(req.RoleId), menuIds, apiIds); err != nil {
		l.Errorf("分配权限失败: %v", err)
		return nil, err
	}

	return &types.AssignPermissionsResponse{
		Code:    0,
		Message: "分配权限成功",
	}, nil
}
