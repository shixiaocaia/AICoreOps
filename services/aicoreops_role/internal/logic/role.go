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

// RoleLogic 角色管理逻辑结构体
type RoleLogic struct {
	ctx    context.Context
	domain *domain.RoleDomain
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewRoleLogic 创建角色管理逻辑实例
func NewRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleLogic {
	return &RoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		domain: domain.NewRoleDomain(svcCtx.DB, svcCtx.Casbin),
	}
}

// convertToRoleResponse 将角色模型转换为响应类型
func convertToRoleResponse(role *model.Role) *types.Role {
	if role == nil {
		return nil
	}
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
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}
	
	if req.Name == "" {
		return nil, errors.New("角色名称不能为空")
	}

	role := &model.Role{
		Name:        req.Name,
		Description: req.Description,
		RoleType:    int(req.RoleType),
		IsDefault:   int(req.IsDefault),
	}

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
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.Id <= 0 {
		return nil, errors.New("无效的角色ID")
	}

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
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.Id <= 0 {
		return nil, errors.New("无效的角色ID")
	}

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
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.Id <= 0 {
		return nil, errors.New("无效的角色ID")
	}

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
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.PageSize <= 0 || req.PageSize > 100 {
		return nil, errors.New("无效的分页大小")
	}

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
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.RoleId <= 0 {
		return nil, errors.New("无效的角色ID")
	}

	if len(req.MenuIds) == 0 && len(req.ApiIds) == 0 {
		return nil, errors.New("菜单ID和API ID不能同时为空")
	}

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

// AssignRoleToUser 分配角色给用户
func (l *RoleLogic) AssignRoleToUser(ctx context.Context, req *types.AssignRoleToUserRequest) (*types.AssignRoleToUserResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.UserId <= 0 {
		return nil, errors.New("无效的用户ID")
	}

	if len(req.RoleIds) == 0 {
		return nil, errors.New("角色ID列表不能为空")
	}

	roleIds := make([]int, len(req.RoleIds))
	for i, id := range req.RoleIds {
		roleIds[i] = int(id)
	}

	if err := l.domain.AssignRoleToUser(ctx, int(req.UserId), roleIds); err != nil {
		l.Errorf("分配角色给用户失败: %v", err)
		return nil, err
	}

	return &types.AssignRoleToUserResponse{
		Code:    0,
		Message: "分配角色给用户成功",
	}, nil
}

// RemoveUserPermissions 移除用户权限
func (l *RoleLogic) RemoveUserPermissions(ctx context.Context, req *types.RemoveUserPermissionsRequest) (*types.RemoveUserPermissionsResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.UserId <= 0 {
		return nil, errors.New("无效的用户ID")
	}

	if err := l.domain.RemoveUserPermissions(ctx, int(req.UserId)); err != nil {
		l.Errorf("移除用户权限失败: %v", err)
		return nil, err
	}

	return &types.RemoveUserPermissionsResponse{
		Code:    0,
		Message: "移除用户权限成功",
	}, nil
}

// RemoveRoleFromUser 移除用户角色
func (l *RoleLogic) RemoveRoleFromUser(ctx context.Context, req *types.RemoveRoleFromUserRequest) (*types.RemoveRoleFromUserResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.UserId <= 0 {
		return nil, errors.New("无效的用户ID")
	}

	if len(req.RoleIds) == 0 {
		return nil, errors.New("角色ID列表不能为空")
	}

	roleIds := make([]int, len(req.RoleIds))
	for i, id := range req.RoleIds {
		roleIds[i] = int(id)
	}

	if err := l.domain.RemoveRoleFromUser(ctx, int(req.UserId), roleIds); err != nil {
		l.Errorf("移除用户角色失败: %v", err)
		return nil, err
	}

	return &types.RemoveRoleFromUserResponse{
		Code:    0,
		Message: "移除用户角色成功",
	}, nil
}
