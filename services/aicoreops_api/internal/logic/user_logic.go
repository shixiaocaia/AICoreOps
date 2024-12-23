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
 * File: user_logic.go
 */

package logic

import (
	"context"
	"time"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/types"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_common/types/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 用户登录
func (l *UserLogic) Login(req *types.LoginRequest) (resp *user.LoginResponse, err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	loginReq := &user.LoginRequest{}
	if err := copier.Copy(loginReq, req); err != nil {
		return nil, err
	}

	loginResp, err := l.svcCtx.UserRpc.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}

	return loginResp, nil
}

// CreateUser 创建用户
func (l *UserLogic) CreateUser(req *types.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	createReq := &user.CreateUserRequest{}
	if err := copier.Copy(createReq, req); err != nil {
		return nil, err
	}

	createResp, err := l.svcCtx.UserRpc.CreateUser(ctx, createReq)
	if err != nil {
		return nil, err
	}

	return createResp, nil
}

// Logout 用户登出
func (l *UserLogic) Logout(req *types.LogoutRequest) (resp *user.LogoutResponse, err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	logoutReq := &user.LogoutRequest{}
	if err := copier.Copy(logoutReq, req); err != nil {
		return nil, err
	}

	logoutResp, err := l.svcCtx.UserRpc.Logout(ctx, logoutReq)
	if err != nil {
		return nil, err
	}

	return logoutResp, nil
}

// GetUser 获取用户信息
func (l *UserLogic) GetUser(req *types.GetUserRequest) (resp *user.GetUserResponse, err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	getUserReq := &user.GetUserRequest{}
	if err := copier.Copy(getUserReq, req); err != nil {
		return nil, err
	}

	userResp, err := l.svcCtx.UserRpc.GetUser(ctx, getUserReq)
	if err != nil {
		return nil, err
	}

	return userResp, nil
}

// UpdateUser 更新用户信息
func (l *UserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	updateReq := &user.UpdateUserRequest{}
	if err := copier.Copy(updateReq, req); err != nil {
		return nil, err
	}

	updateResp, err := l.svcCtx.UserRpc.UpdateUser(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return updateResp, nil
}

// DeleteUser 删除用户
func (l *UserLogic) DeleteUser(req *types.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	deleteReq := &user.DeleteUserRequest{}
	if err := copier.Copy(deleteReq, req); err != nil {
		return nil, err
	}

	deleteResp, err := l.svcCtx.UserRpc.DeleteUser(ctx, deleteReq)
	if err != nil {
		return nil, err
	}

	return deleteResp, nil
}

// GetUserList 获取用户列表
func (l *UserLogic) GetUserList(req *types.GetUserListRequest) (resp *user.ListUsersResponse, err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	listReq := &user.ListUsersRequest{}
	if err := copier.Copy(listReq, req); err != nil {
		return nil, err
	}

	listResp, err := l.svcCtx.UserRpc.ListUsers(ctx, listReq)
	if err != nil {
		return nil, err
	}

	return listResp, nil
}
