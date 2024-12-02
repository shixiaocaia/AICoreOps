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
 * File: menu_logic.go
 */

package logic

import (
	"aicoreops_api/internal/svc"
	"aicoreops_api/internal/types"
	"aicoreops_common/types/role"
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type MenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuLogic {
	return &MenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateMenu 创建菜单
func (l *MenuLogic) CreateMenu(req *types.CreateMenuRequest) (*role.CreateMenuResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	createReq := &role.CreateMenuRequest{}
	if err := copier.Copy(createReq, req); err != nil {
		return nil, err
	}

	createResp, err := l.svcCtx.MenuRpc.CreateMenu(ctx, createReq)
	if err != nil {
		return nil, err
	}

	return createResp, nil
}

// GetMenu 获取菜单详情
func (l *MenuLogic) GetMenu(req *types.GetMenuRequest) (*role.GetMenuResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	getReq := &role.GetMenuRequest{}
	if err := copier.Copy(getReq, req); err != nil {
		return nil, err
	}

	getResp, err := l.svcCtx.MenuRpc.GetMenu(ctx, getReq)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

// UpdateMenu 更新菜单
func (l *MenuLogic) UpdateMenu(req *types.UpdateMenuRequest) (*role.UpdateMenuResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	updateReq := &role.UpdateMenuRequest{}
	if err := copier.Copy(updateReq, req); err != nil {
		return nil, err
	}

	updateResp, err := l.svcCtx.MenuRpc.UpdateMenu(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return updateResp, nil
}

// DeleteMenu 删除菜单
func (l *MenuLogic) DeleteMenu(req *types.DeleteMenuRequest) (*role.DeleteMenuResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	deleteReq := &role.DeleteMenuRequest{}
	if err := copier.Copy(deleteReq, req); err != nil {
		return nil, err
	}

	deleteResp, err := l.svcCtx.MenuRpc.DeleteMenu(ctx, deleteReq)
	if err != nil {
		return nil, err
	}

	return deleteResp, nil
}

// ListMenus 获取菜单列表
func (l *MenuLogic) ListMenus(req *types.ListMenusRequest) (*role.ListMenusResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	listReq := &role.ListMenusRequest{}
	if err := copier.Copy(listReq, req); err != nil {
		return nil, err
	}

	listResp, err := l.svcCtx.MenuRpc.ListMenus(ctx, listReq)
	if err != nil {
		return nil, err
	}

	return listResp, nil
}
