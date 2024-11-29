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
 * File: api.go
 * Description: API逻辑层实现
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

type ApiLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	apiDomain *domain.ApiDomain
	logx.Logger
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		apiDomain: domain.NewApiDomain(svcCtx.DB),
		Logger:    logx.WithContext(ctx),
	}
}

// CreateApi 创建API
func (l *ApiLogic) CreateApi(ctx context.Context, req *types.CreateApiRequest) (*types.CreateApiResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	api := &model.Api{
		Name:        req.Name,
		Path:        req.Path,
		Method:      int(req.Method),
		Description: req.Description,
		Version:     req.Version,
		Category:    int(req.Category),
		IsPublic:    int(req.IsPublic),
	}

	// 验证API数据有效性
	if err := api.Validate(); err != nil {
		return nil, err
	}

	apiDomain := domain.NewApiDomain(l.svcCtx.DB)
	if err := apiDomain.CreateApi(ctx, api); err != nil {
		l.Logger.Errorf("创建API失败: %v", err)
		return nil, err
	}

	return &types.CreateApiResponse{
		Code:    0,
		Message: "创建API成功",
	}, nil
}

// GetApi 获取API详情
func (l *ApiLogic) GetApi(ctx context.Context, req *types.GetApiRequest) (*types.GetApiResponse, error) {
	if req == nil || req.Id <= 0 {
		return nil, errors.New("无效的API ID")
	}

	apiDomain := domain.NewApiDomain(l.svcCtx.DB)
	api, err := apiDomain.GetApi(ctx, int(req.Id))
	if err != nil {
		l.Logger.Errorf("获取API失败: %v", err)
		return nil, err
	}

	return &types.GetApiResponse{
		Code:    0,
		Message: "获取API成功",
		Data: &types.Api{
			Id:          api.ID,
			Name:        api.Name,
			Path:        api.Path,
			Method:      types.HttpMethod(api.Method),
			Description: api.Description,
			Version:     api.Version,
			Category:    types.ApiCategory(api.Category),
			IsPublic:    int32(api.IsPublic),
			CreateTime:  api.CreateTime,
			UpdateTime:  api.UpdateTime,
		},
	}, nil
}

// UpdateApi 更新API
func (l *ApiLogic) UpdateApi(ctx context.Context, req *types.UpdateApiRequest) (*types.UpdateApiResponse, error) {
	// 参数校验
	if err := l.apiDomain.ValidateApi(req); err != nil {
		return nil, err
	}

	// 构建API对象
	api := l.apiDomain.BuildApi(req)

	// 验证API数据有效性
	if err := api.Validate(); err != nil {
		return nil, err
	}

	// 更新API
	apiDomain := domain.NewApiDomain(l.svcCtx.DB)
	if err := apiDomain.UpdateApi(ctx, api); err != nil {
		l.Logger.Errorf("更新API失败: %v", err)
		return nil, fmt.Errorf("更新API失败: %v", err)
	}

	return &types.UpdateApiResponse{
		Code:    0,
		Message: "更新API成功",
	}, nil
}

// DeleteApi 删除API
func (l *ApiLogic) DeleteApi(ctx context.Context, req *types.DeleteApiRequest) (*types.DeleteApiResponse, error) {
	if req == nil || req.Id <= 0 {
		return nil, errors.New("无效的API ID")
	}

	apiDomain := domain.NewApiDomain(l.svcCtx.DB)
	if err := apiDomain.DeleteApi(ctx, int(req.Id)); err != nil {
		l.Logger.Errorf("删除API失败: %v", err)
		return nil, err
	}

	return &types.DeleteApiResponse{
		Code:    0,
		Message: "删除API成功",
	}, nil
}

// ListApis 获取API列表
func (l *ApiLogic) ListApis(ctx context.Context, req *types.ListApisRequest) (*types.ListApisResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	// 设置默认值和参数验证
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if req.PageSize <= 0 || req.PageSize > 100 { // 添加最大页面大小限制
		req.PageSize = 10
	}

	apiDomain := domain.NewApiDomain(l.svcCtx.DB)
	apis, total, err := apiDomain.ListApis(ctx, int(req.PageNumber), int(req.PageSize))
	if err != nil {
		l.Logger.Errorf("获取API列表失败: %v", err)
		return nil, fmt.Errorf("获取API列表失败: %v", err)
	}

	apiList := make([]*types.Api, 0, len(apis))
	for _, api := range apis {
		apiList = append(apiList, &types.Api{
			Id:          api.ID,
			Name:        api.Name,
			Path:        api.Path,
			Method:      types.HttpMethod(api.Method),
			Description: api.Description,
			Version:     api.Version,
			Category:    types.ApiCategory(api.Category),
			IsPublic:    int32(api.IsPublic),
			CreateTime:  api.CreateTime,
			UpdateTime:  api.UpdateTime,
		})
	}

	return &types.ListApisResponse{
		Code:    0,
		Message: "获取API列表成功",
		Data: &types.ListApisData{
			Total:      int32(total),
			Apis:       apiList,
			PageNumber: req.PageNumber,
			PageSize:   req.PageSize,
		},
	}, nil
}
