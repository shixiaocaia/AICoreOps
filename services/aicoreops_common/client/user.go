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
 * File: user.go
 */

package client

import (
	"context"

	"aicoreops_user/types"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateUserRequest  = types.CreateUserRequest
	CreateUserResponse = types.CreateUserResponse
	GetUserRequest     = types.GetUserRequest
	GetUserResponse    = types.GetUserResponse
	UpdateUserRequest  = types.UpdateUserRequest
	UpdateUserResponse = types.UpdateUserResponse
	DeleteUserRequest  = types.DeleteUserRequest
	DeleteUserResponse = types.DeleteUserResponse
	ListUsersRequest   = types.ListUsersRequest
	ListUsersResponse  = types.ListUsersResponse
	LoginRequest       = types.LoginRequest
	LoginResponse      = types.LoginResponse
	LogoutRequest      = types.LogoutRequest
	LogoutResponse     = types.LogoutResponse

	AicoreopsUser interface {
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
		UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
		DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
		ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	}

	defaultAicoreopsUser struct {
		cli zrpc.Client
	}
)

func NewAicoreopsUser(cli zrpc.Client) AicoreopsUser {
	return &defaultAicoreopsUser{
		cli: cli,
	}
}

func (m *defaultAicoreopsUser) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := types.NewUserServiceClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultAicoreopsUser) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	client := types.NewUserServiceClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultAicoreopsUser) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	client := types.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultAicoreopsUser) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	client := types.NewUserServiceClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultAicoreopsUser) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	client := types.NewUserServiceClient(m.cli.Conn())
	return client.ListUsers(ctx, in, opts...)
}

func (m *defaultAicoreopsUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := types.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultAicoreopsUser) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	client := types.NewUserServiceClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}
