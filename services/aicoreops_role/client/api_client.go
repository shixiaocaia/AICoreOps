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
 * File: api_client.go
 */

package client

import (
	"context"

	"aicoreops_role/types"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Api               = types.Api
	CreateApiRequest  = types.CreateApiRequest
	CreateApiResponse = types.CreateApiResponse
	DeleteApiRequest  = types.DeleteApiRequest
	DeleteApiResponse = types.DeleteApiResponse
	GetApiRequest     = types.GetApiRequest
	GetApiResponse    = types.GetApiResponse
	ListApisData      = types.ListApisData
	ListApisRequest   = types.ListApisRequest
	ListApisResponse  = types.ListApisResponse
	UpdateApiRequest  = types.UpdateApiRequest
	UpdateApiResponse = types.UpdateApiResponse

	ApiService interface {
		// 创建 API
		CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*CreateApiResponse, error)
		// 获取 API 详情
		GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*GetApiResponse, error)
		// 更新 API
		UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*UpdateApiResponse, error)
		// 删除 API
		DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*DeleteApiResponse, error)
		// 列出 APIs
		ListApis(ctx context.Context, in *ListApisRequest, opts ...grpc.CallOption) (*ListApisResponse, error)
	}

	defaultApiService struct {
		cli zrpc.Client
	}
)

func NewApiService(cli zrpc.Client) ApiService {
	return &defaultApiService{
		cli: cli,
	}
}

// 创建 API
func (m *defaultApiService) CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*CreateApiResponse, error) {
	client := types.NewApiServiceClient(m.cli.Conn())
	return client.CreateApi(ctx, in, opts...)
}

// 获取 API 详情
func (m *defaultApiService) GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*GetApiResponse, error) {
	client := types.NewApiServiceClient(m.cli.Conn())
	return client.GetApi(ctx, in, opts...)
}

// 更新 API
func (m *defaultApiService) UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*UpdateApiResponse, error) {
	client := types.NewApiServiceClient(m.cli.Conn())
	return client.UpdateApi(ctx, in, opts...)
}

// 删除 API
func (m *defaultApiService) DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*DeleteApiResponse, error) {
	client := types.NewApiServiceClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

// 列出 APIs
func (m *defaultApiService) ListApis(ctx context.Context, in *ListApisRequest, opts ...grpc.CallOption) (*ListApisResponse, error) {
	client := types.NewApiServiceClient(m.cli.Conn())
	return client.ListApis(ctx, in, opts...)
}
