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

type MenuServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedMenuServiceServer
}

func NewMenuServer(svcCtx *svc.ServiceContext) *MenuServer {
	return &MenuServer{
		svcCtx: svcCtx,
	}
}

func (s *MenuServer) CreateMenu(ctx context.Context, request *types.CreateMenuRequest) (*types.CreateMenuResponse, error) {
	l := logic.NewMenuLogic(ctx, s.svcCtx)
	return l.CreateMenu(ctx, request)
}

func (s *MenuServer) GetMenu(ctx context.Context, request *types.GetMenuRequest) (*types.GetMenuResponse, error) {
	l := logic.NewMenuLogic(ctx, s.svcCtx)
	return l.GetMenu(ctx, request)
}

func (s *MenuServer) UpdateMenu(ctx context.Context, request *types.UpdateMenuRequest) (*types.UpdateMenuResponse, error) {
	l := logic.NewMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(ctx, request)
}

func (s *MenuServer) DeleteMenu(ctx context.Context, request *types.DeleteMenuRequest) (*types.DeleteMenuResponse, error) {
	l := logic.NewMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(ctx, request)
}

func (s *MenuServer) ListMenus(ctx context.Context, request *types.ListMenusRequest) (*types.ListMenusResponse, error) {
	l := logic.NewMenuLogic(ctx, s.svcCtx)
	return l.ListMenus(ctx, request)
}
