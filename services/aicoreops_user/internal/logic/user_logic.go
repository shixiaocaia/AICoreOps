package logic

import (
	"context"

	"aicoreops_user/internal/svc"
	"aicoreops_user/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUser 创建用户
func (l *UserLogic) CreateUser(ctx context.Context, req *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	return &types.CreateUserResponse{}, nil
}

// GetUser 获取用户详情
func (l *UserLogic) GetUser(ctx context.Context, req *types.GetUserRequest) (*types.GetUserResponse, error) {
	return &types.GetUserResponse{}, nil
}

// UpdateUser 更新用户
func (l *UserLogic) UpdateUser(ctx context.Context, req *types.UpdateUserRequest) (*types.UpdateUserResponse, error) {
	return &types.UpdateUserResponse{}, nil
}

// DeleteUser 删除用户
func (l *UserLogic) DeleteUser(ctx context.Context, req *types.DeleteUserRequest) (*types.DeleteUserResponse, error) {
	return &types.DeleteUserResponse{}, nil
}

// ListUsers 列出用户
func (l *UserLogic) ListUsers(ctx context.Context, req *types.ListUsersRequest) (*types.ListUsersResponse, error) {
	return &types.ListUsersResponse{}, nil
}
