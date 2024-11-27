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

type MenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuLogic {
	return &MenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateMenu 创建菜单
func (l *MenuLogic) CreateMenu(ctx context.Context, request *types.CreateMenuRequest) (*types.CreateMenuResponse, error) {
	return nil, nil
}

// GetMenu 获取菜单详情
func (l *MenuLogic) GetMenu(ctx context.Context, request *types.GetMenuRequest) (*types.GetMenuResponse, error) {
	return nil, nil
}

// UpdateMenu 更新菜单
func (l *MenuLogic) UpdateMenu(ctx context.Context, request *types.UpdateMenuRequest) (*types.UpdateMenuResponse, error) {
	return nil, nil
}

// DeleteMenu 删除菜单
func (l *MenuLogic) DeleteMenu(ctx context.Context, request *types.DeleteMenuRequest) (*types.DeleteMenuResponse, error) {
	return nil, nil
}

// ListMenus 获取菜单列表
func (l *MenuLogic) ListMenus(ctx context.Context, request *types.ListMenusRequest) (*types.ListMenusResponse, error) {
	return nil, nil
}
