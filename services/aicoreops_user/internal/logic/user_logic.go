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
	"errors"

	"aicoreops_user/internal/domain"
	"aicoreops_user/internal/svc"
	"aicoreops_user/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserLogic struct {
	ctx    context.Context
	domain *domain.UserDomain
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		domain: domain.NewUserDomain(svcCtx.DB, svcCtx.JWT),
	}
}

// CreateUser 创建用户
func (l *UserLogic) CreateUser(ctx context.Context, req *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	// 参数校验
	if err := l.domain.VerifyUsernameAndPassword(req.Username, req.Password); err != nil {
		l.Logger.Errorf("验证用户名密码格式失败: %v", err)
		return nil, err
	}

	// 检查用户名是否已存在
	exists, err := l.domain.CheckUsernameExists(ctx, req.Username)
	if err != nil {
		l.Logger.Errorf("检查用户名是否存在失败: %v", err)
		return nil, errors.New("系统错误")
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 加密密码
	encryptedPwd, err := l.domain.EncryptPassword(req.Password)
	if err != nil {
		l.Logger.Errorf("密码加密失败: %v", err)
		return nil, errors.New("系统错误")
	}

	// 构建用户模型
	user := l.domain.BuildUserModel(req, encryptedPwd)

	// 生成雪花ID
	user.ID = int(l.svcCtx.Snowflake.Generate().Int64())

	// 注册用户
	if err := l.domain.Register(ctx, user); err != nil {
		l.Logger.Errorf("用户注册失败: %v", err)
		return nil, errors.New("注册失败")
	}

	return &types.CreateUserResponse{
		Code:    0,
		Message: "创建用户成功",
	}, nil
}

// Login 登录
func (l *UserLogic) Login(ctx context.Context, req *types.LoginRequest) (*types.LoginResponse, error) {
	// 验证参数
	if err := l.domain.VerifyUsernameAndPassword(req.Username, req.Password); err != nil {
		l.Logger.Errorf("验证用户名密码格式失败: %v", err)
		return nil, err
	}

	// 加密密码
	encryptedPwd, err := l.domain.EncryptPassword(req.Password)
	if err != nil {
		l.Logger.Errorf("密码加密失败: %v", err)
		return nil, errors.New("系统错误")
	}

	// 查询用户
	user, err := l.domain.GetUserByUsernameAndPassword(ctx, req.Username, encryptedPwd)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户名或密码错误")
		}
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errors.New("系统错误")
	}

	// 生成token
	jwtToken, refreshToken, err := l.domain.GenerateToken(ctx, user.ID)
	if err != nil {
		l.Logger.Errorf("生成token失败: %v", err)
		return nil, errors.New("系统错误")
	}

	// 更新用户最后登录时间
	if err := l.domain.UpdateLastLoginTime(ctx, user.ID); err != nil {
		l.Logger.Errorf("更新用户最后登录时间失败: %v", err)
	}

	return &types.LoginResponse{
		Code:    0,
		Message: "登录成功",
		Data: &types.LoginResponseData{
			JwtToken:     jwtToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

// Logout 登出
func (l *UserLogic) Logout(ctx context.Context, req *types.LogoutRequest) (*types.LogoutResponse, error) {
	// 清除token
	if err := l.svcCtx.JWT.ClearToken(ctx, req.JwtToken, req.RefreshToken); err != nil {
		l.Logger.Errorf("清除token失败: %v", err)
		return nil, errors.New("登出失败")
	}

	return &types.LogoutResponse{
		Code:    0,
		Message: "登出成功",
	}, nil
}

// GetUser 获取用户详情
func (l *UserLogic) GetUser(ctx context.Context, req *types.GetUserRequest) (*types.GetUserResponse, error) {
	// 查询用户
	user, err := l.domain.GetUserById(ctx, int(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("用户不存在: %v", err)
			return nil, errors.New("用户不存在")
		}
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errors.New("系统错误")
	}

	// 构建响应
	return &types.GetUserResponse{
		Code:    0,
		Message: "获取用户成功",
		Data: &types.User{
			Id:            int64(user.ID),
			Username:      user.Username,
			Email:         user.Email,
			Phone:         user.Phone,
			Nickname:      user.Nickname,
			Avatar:        user.Avatar,
			LastLoginTime: user.LastLoginTime,
			CreateTime:    user.CreateTime,
			UpdateTime:    user.UpdateTime,
			Status:        types.UserStatus(user.Status),
		},
	}, nil
}

// UpdateUser 更新用户
func (l *UserLogic) UpdateUser(ctx context.Context, req *types.UpdateUserRequest) (*types.UpdateUserResponse, error) {
	// 检查用户是否存在
	user, err := l.domain.GetUserById(ctx, int(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("用户不存在: %v", err)
			return nil, errors.New("用户不存在")
		}
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errors.New("系统错误")
	}

	// 更新用户信息
	if req.Username != "" {
		l.Logger.Errorf("不允许修改用户名")
		return nil, errors.New("不允许修改用户名")
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Status != types.UserStatus_STATUS_UNSPECIFIED {
		user.Status = int(req.Status)
	}
	if req.IsDeleted != 0 {
		user.IsDeleted = int(req.IsDeleted)
	}

	// 调用领域层更新用户
	if err := l.domain.UpdateUser(ctx, user); err != nil {
		l.Logger.Errorf("更新用户失败: %v", err)
		return nil, errors.New("更新用户失败")
	}

	return &types.UpdateUserResponse{
		Code:    0,
		Message: "更新用户成功",
	}, nil
}

// DeleteUser 删除用户
func (l *UserLogic) DeleteUser(ctx context.Context, req *types.DeleteUserRequest) (*types.DeleteUserResponse, error) {
	// 检查用户是否存在
	user, err := l.domain.GetUserById(ctx, int(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("用户不存在: %v", err)
			return nil, errors.New("用户不存在")
		}
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errors.New("系统错误")
	}

	// 软删除用户
	user.IsDeleted = 1
	if err := l.domain.UpdateUser(ctx, user); err != nil {
		l.Logger.Errorf("删除用户失败: %v", err)
		return nil, errors.New("删除用户失败")
	}

	return &types.DeleteUserResponse{
		Code:    0,
		Message: "删除用户成功",
	}, nil
}

// ListUsers 列出用户
func (l *UserLogic) ListUsers(ctx context.Context, req *types.ListUsersRequest) (*types.ListUsersResponse, error) {
	// 参数校验
	if req.PageNumber < 1 {
		req.PageNumber = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}

	// 获取用户列表
	users, total, err := l.domain.ListUsers(ctx, int(req.PageNumber), int(req.PageSize))
	if err != nil {
		l.Logger.Errorf("查询用户列表失败: %v", err)
		return nil, errors.New("查询用户列表失败")
	}

	// 构建响应数据
	userList := make([]*types.User, 0, len(users))
	for _, user := range users {
		userList = append(userList, &types.User{
			Id:            int64(user.ID),
			Username:      user.Username,
			Email:         user.Email,
			Phone:         user.Phone,
			Nickname:      user.Nickname,
			Avatar:        user.Avatar,
			LastLoginTime: user.LastLoginTime,
			CreateTime:    user.CreateTime,
			UpdateTime:    user.UpdateTime,
			Status:        types.UserStatus(user.Status),
			IsDeleted:     int32(user.IsDeleted),
		})
	}

	return &types.ListUsersResponse{
		Code:    0,
		Message: "获取用户列表成功",
		Data: &types.ListUsersData{
			Users:      userList,
			Total:      int32(total),
			PageNumber: req.PageNumber,
			PageSize:   req.PageSize,
		},
	}, nil
}
