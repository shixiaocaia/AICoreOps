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

type ApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateApi 创建API
func (l *ApiLogic) CreateApi(ctx context.Context, request *types.CreateApiRequest) (*types.CreateApiResponse, error) {
	return nil, nil
}

// GetApi 获取API详情
func (l *ApiLogic) GetApi(ctx context.Context, request *types.GetApiRequest) (*types.GetApiResponse, error) {
	return nil, nil
}

// UpdateApi 更新API
func (l *ApiLogic) UpdateApi(ctx context.Context, request *types.UpdateApiRequest) (*types.UpdateApiResponse, error) {
	return nil, nil
}

// DeleteApi 删除API
func (l *ApiLogic) DeleteApi(ctx context.Context, request *types.DeleteApiRequest) (*types.DeleteApiResponse, error) {
	return nil, nil
}

// ListApis 获取API列表
func (l *ApiLogic) ListApis(ctx context.Context, request *types.ListApisRequest) (*types.ListApisResponse, error) {
	return nil, nil
}
