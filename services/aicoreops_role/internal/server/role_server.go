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

package server

import (
	"aicoreops_role/internal/logic"
	"aicoreops_role/internal/svc"
	"aicoreops_role/types"
	"context"
)

type RoleServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedRoleServiceServer
}

func NewRoleServer(svcCtx *svc.ServiceContext) *RoleServer {
	return &RoleServer{
		svcCtx: svcCtx,
	}
}

func (s *RoleServer) CreateRole(ctx context.Context, request *types.CreateRoleRequest) (*types.CreateRoleResponse, error) {
	l := logic.NewRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(ctx, request)
}

func (s *RoleServer) GetRole(ctx context.Context, request *types.GetRoleRequest) (*types.GetRoleResponse, error) {
	l := logic.NewRoleLogic(ctx, s.svcCtx)
	return l.GetRole(ctx, request)
}

func (s *RoleServer) UpdateRole(ctx context.Context, request *types.UpdateRoleRequest) (*types.UpdateRoleResponse, error) {
	l := logic.NewRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(ctx, request)
}

func (s *RoleServer) DeleteRole(ctx context.Context, request *types.DeleteRoleRequest) (*types.DeleteRoleResponse, error) {
	l := logic.NewRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(ctx, request)
}

func (s *RoleServer) ListRoles(ctx context.Context, request *types.ListRolesRequest) (*types.ListRolesResponse, error) {
	l := logic.NewRoleLogic(ctx, s.svcCtx)
	return l.ListRoles(ctx, request)
}

func (s *RoleServer) AssignPermissions(ctx context.Context, request *types.AssignPermissionsRequest) (*types.AssignPermissionsResponse, error) {
	l := logic.NewRoleLogic(ctx, s.svcCtx)
	return l.AssignPermissions(ctx, request)
}
