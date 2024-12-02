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
 * File: api_server.go
 */

package server

import (
	"aicoreops_role/internal/logic"
	"aicoreops_role/internal/svc"
	"aicoreops_role/types"
	"context"
)

type ApiServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedApiServiceServer
}

func NewApiServer(svcCtx *svc.ServiceContext) *ApiServer {
	return &ApiServer{
		svcCtx: svcCtx,
	}
}

func (s *ApiServer) CreateApi(ctx context.Context, request *types.CreateApiRequest) (*types.CreateApiResponse, error) {
	l := logic.NewApiLogic(ctx, s.svcCtx)
	return l.CreateApi(ctx, request)
}

func (s *ApiServer) GetApi(ctx context.Context, request *types.GetApiRequest) (*types.GetApiResponse, error) {
	l := logic.NewApiLogic(ctx, s.svcCtx)
	return l.GetApi(ctx, request)
}

func (s *ApiServer) UpdateApi(ctx context.Context, request *types.UpdateApiRequest) (*types.UpdateApiResponse, error) {
	l := logic.NewApiLogic(ctx, s.svcCtx)
	return l.UpdateApi(ctx, request)
}

func (s *ApiServer) DeleteApi(ctx context.Context, request *types.DeleteApiRequest) (*types.DeleteApiResponse, error) {
	l := logic.NewApiLogic(ctx, s.svcCtx)
	return l.DeleteApi(ctx, request)
}

func (s *ApiServer) ListApis(ctx context.Context, request *types.ListApisRequest) (*types.ListApisResponse, error) {
	l := logic.NewApiLogic(ctx, s.svcCtx)
	return l.ListApis(ctx, request)
}
