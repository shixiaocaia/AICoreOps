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
 * File: api_logic.go
 */

package logic

import (
	"context"
	"time"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/types"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_common/types/api"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateApi 创建接口
func (l *ApiLogic) CreateApi(req *types.CreateApiRequest) (*api.CreateApiResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	createReq := &api.CreateApiRequest{}
	if err := copier.Copy(createReq, req); err != nil {
		return nil, err
	}

	createResp, err := l.svcCtx.ApiRpc.CreateApi(ctx, createReq)
	if err != nil {
		return nil, err
	}

	return createResp, nil
}

// GetApi 获取接口详情
func (l *ApiLogic) GetApi(req *types.GetApiRequest) (*api.GetApiResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	getReq := &api.GetApiRequest{}
	if err := copier.Copy(getReq, req); err != nil {
		return nil, err
	}

	getResp, err := l.svcCtx.ApiRpc.GetApi(ctx, getReq)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

// UpdateApi 更新接口
func (l *ApiLogic) UpdateApi(req *types.UpdateApiRequest) (*api.UpdateApiResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	updateReq := &api.UpdateApiRequest{}
	if err := copier.Copy(updateReq, req); err != nil {
		return nil, err
	}

	updateResp, err := l.svcCtx.ApiRpc.UpdateApi(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return updateResp, nil
}

// DeleteApi 删除接口
func (l *ApiLogic) DeleteApi(req *types.DeleteApiRequest) (*api.DeleteApiResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	deleteReq := &api.DeleteApiRequest{}
	if err := copier.Copy(deleteReq, req); err != nil {
		return nil, err
	}

	deleteResp, err := l.svcCtx.ApiRpc.DeleteApi(ctx, deleteReq)
	if err != nil {
		return nil, err
	}

	return deleteResp, nil
}

// ListApis 获取接口列表
func (l *ApiLogic) ListApis(req *types.ListApisRequest) (*api.ListApisResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	listReq := &api.ListApisRequest{}
	if err := copier.Copy(listReq, req); err != nil {
		return nil, err
	}

	listResp, err := l.svcCtx.ApiRpc.ListApis(ctx, listReq)
	if err != nil {
		return nil, err
	}

	return listResp, nil
}
