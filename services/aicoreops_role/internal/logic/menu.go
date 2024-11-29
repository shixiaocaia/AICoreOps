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
 * File: menu.go
 * Description: 菜单相关逻辑处理
 */

package logic

import (
	"aicoreops_role/internal/svc"
	"aicoreops_role/types"
	"context"
	"errors"
	"fmt"

	"aicoreops_role/internal/domain"
	"aicoreops_role/internal/model"

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
func (l *MenuLogic) CreateMenu(ctx context.Context, req *types.CreateMenuRequest) (*types.CreateMenuResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	menu := &model.Menu{
		Name:      req.Name,
		ParentID:  req.ParentId,
		Path:      req.Path,
		Component: req.Component,
		Icon:      req.Icon,
		SortOrder: int(req.SortOrder),
		RouteName: req.RouteName,
		Hidden:    int(req.Hidden),
	}

	menuDomain := domain.NewMenuDomain(l.svcCtx.DB)
	err := menuDomain.CreateMenu(ctx, menu)
	if err != nil {
		l.Logger.Errorf("创建菜单失败: %v", err)
		return nil, err
	}

	return &types.CreateMenuResponse{
		Code:    0,
		Message: "创建菜单成功",
	}, nil
}

// GetMenu 获取菜单详情
func (l *MenuLogic) GetMenu(ctx context.Context, req *types.GetMenuRequest) (*types.GetMenuResponse, error) {
	if req == nil || req.Id <= 0 {
		return nil, errors.New("无效的菜单ID")
	}

	menuDomain := domain.NewMenuDomain(l.svcCtx.DB)
	menu, err := menuDomain.GetMenu(ctx, int(req.Id))
	if err != nil {
		l.Logger.Errorf("获取菜单详情失败: %v", err)
		return nil, err
	}

	if menu == nil {
		return nil, errors.New("菜单不存在")
	}

	return &types.GetMenuResponse{
		Code:    0,
		Message: "获取菜单详情成功",
		Data:    convertMenuToType(menu),
	}, nil
}

// UpdateMenu 更新菜单
func (l *MenuLogic) UpdateMenu(ctx context.Context, req *types.UpdateMenuRequest) (*types.UpdateMenuResponse, error) {
	if req == nil || req.Id <= 0 {
		return nil, errors.New("无效的菜单ID")
	}

	menu := &model.Menu{
		ID:        req.Id,
		Name:      req.Name,
		ParentID:  req.ParentId,
		Path:      req.Path,
		Component: req.Component,
		Icon:      req.Icon,
		SortOrder: int(req.SortOrder),
		RouteName: req.RouteName,
		Hidden:    int(req.Hidden),
	}

	menuDomain := domain.NewMenuDomain(l.svcCtx.DB)
	err := menuDomain.UpdateMenu(ctx, menu)
	if err != nil {
		l.Logger.Errorf("更新菜单失败: %v", err)
		return nil, err
	}

	return &types.UpdateMenuResponse{
		Code:    0,
		Message: "更新菜单成功",
	}, nil
}

// DeleteMenu 删除菜单
func (l *MenuLogic) DeleteMenu(ctx context.Context, req *types.DeleteMenuRequest) (*types.DeleteMenuResponse, error) {
	if req == nil || req.Id <= 0 {
		return nil, errors.New("无效的菜单ID")
	}

	menuDomain := domain.NewMenuDomain(l.svcCtx.DB)
	err := menuDomain.DeleteMenu(ctx, int(req.Id))
	if err != nil {
		l.Logger.Errorf("删除菜单失败: %v", err)
		return nil, err
	}

	return &types.DeleteMenuResponse{
		Code:    0,
		Message: "删除菜单成功",
	}, nil
}

// ListMenus 获取菜单列表
func (l *MenuLogic) ListMenus(ctx context.Context, req *types.ListMenusRequest) (*types.ListMenusResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	menuDomain := domain.NewMenuDomain(l.svcCtx.DB)

	// 设置默认值和参数验证
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if req.PageSize <= 0 || req.PageSize > 100 { // 添加最大页面大小限制
		req.PageSize = 10
	}

	var (
		menus []*model.Menu
		total int
		err   error
	)

	// 根据请求类型获取数据
	if req.IsTree {
		menus, err = menuDomain.GetMenuTree(ctx)
	} else {
		menus, total, err = menuDomain.ListMenus(ctx, int(req.PageNumber), int(req.PageSize))
	}

	if err != nil {
		l.Logger.Errorf("获取菜单列表失败: %v", err)
		return nil, fmt.Errorf("获取菜单列表失败: %v", err)
	}

	return &types.ListMenusResponse{
		Code:    0,
		Message: "获取菜单列表成功",
		Data: &types.ListMenusData{
			Total:      int32(total),
			Menus:      convertMenusToTypes(menus),
			PageNumber: req.PageNumber,
			PageSize:   req.PageSize,
		},
	}, nil
}

// convertMenuToType 将单个model.Menu转换为types.Menu
func convertMenuToType(menu *model.Menu) *types.Menu {
	if menu == nil {
		return nil
	}

	result := &types.Menu{
		Id:         menu.ID,
		Name:       menu.Name,
		ParentId:   menu.ParentID,
		Path:       menu.Path,
		Component:  menu.Component,
		Icon:       menu.Icon,
		SortOrder:  int32(menu.SortOrder),
		RouteName:  menu.RouteName,
		Hidden:     int32(menu.Hidden),
		CreateTime: menu.CreateTime,
		UpdateTime: menu.UpdateTime,
		Children:   make([]*types.Menu, 0),
	}

	if len(menu.Children) > 0 {
		result.Children = convertMenusToTypes(menu.Children)
	}

	return result
}

// convertMenusToTypes 将model.Menu切片转换为types.Menu切片
func convertMenusToTypes(menus []*model.Menu) []*types.Menu {
	if len(menus) == 0 {
		return nil
	}

	result := make([]*types.Menu, 0, len(menus))
	for _, menu := range menus {
		if m := convertMenuToType(menu); m != nil {
			result = append(result, m)
		}
	}
	return result
}
