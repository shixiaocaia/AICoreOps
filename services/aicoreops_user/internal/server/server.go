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
 * File: server.go
 */

package server

import (
	"aicoreops_user/internal/logic"
	"context"

	"aicoreops_user/internal/svc"
	"aicoreops_user/types"
)

type AicoreopsUserServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedUserServiceServer
}

func NewAicoreopsUserServer(svcCtx *svc.ServiceContext) *AicoreopsUserServer {
	return &AicoreopsUserServer{
		svcCtx: svcCtx,
	}
}

func (s *AicoreopsUserServer) CreateUser(ctx context.Context, request *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.CreateUser(ctx, request)
}

func (s *AicoreopsUserServer) GetUser(ctx context.Context, request *types.GetUserRequest) (*types.GetUserResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.GetUser(ctx, request)
}

func (s *AicoreopsUserServer) UpdateUser(ctx context.Context, request *types.UpdateUserRequest) (*types.UpdateUserResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(ctx, request)
}

func (s *AicoreopsUserServer) DeleteUser(ctx context.Context, request *types.DeleteUserRequest) (*types.DeleteUserResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(ctx, request)
}

func (s *AicoreopsUserServer) ListUsers(ctx context.Context, request *types.ListUsersRequest) (*types.ListUsersResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.ListUsers(ctx, request)
}

func (s *AicoreopsUserServer) Login(ctx context.Context, request *types.LoginRequest) (*types.LoginResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.Login(ctx, request)
}

func (s *AicoreopsUserServer) Logout(ctx context.Context, request *types.LogoutRequest) (*types.LogoutResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.Logout(ctx, request)
}
