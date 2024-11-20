package repo

import (
	"aicoreops_user/internal/model"
	"context"
)

type UserRepo interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, user *model.User) error
	// GetUser 根据ID获取用户
	GetUser(ctx context.Context, id int) (*model.User, error)
	// GetUserByUsername 根据用户名获取用户
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	// UpdateUser 更新用户信息
	UpdateUser(ctx context.Context, user *model.User) error
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, id int) error
	// ListUsers 获取用户列表
	ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int, error)
	// UpdatePassword 更新用户密码
	UpdatePassword(ctx context.Context, id int, newPassword string) error
	// UpdateStatus 更新用户状态
	UpdateStatus(ctx context.Context, id int, status int) error
}
