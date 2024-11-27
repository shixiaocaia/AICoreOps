package logic

import (
	"context"
	"time"

	"aicoreops_api/internal/svc"
	"aicoreops_api/internal/types"
	"aicoreops_common/types/user"

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
func (l *UserLogic) CreateUser(req *types.CreateUserRequest) (err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	createReq := &user.CreateUserRequest{}
	if err := copier.Copy(createReq, req); err != nil {
		return err
	}

	_, err = l.svcCtx.UserRpc.CreateUser(ctx, createReq)
	if err != nil {
		return err
	}

	return nil
}

// Logout 用户登出
func (l *UserLogic) Logout(req *types.LogoutRequest) (err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	logoutReq := &user.LogoutRequest{}
	if err := copier.Copy(logoutReq, req); err != nil {
		return err
	}

	_, err = l.svcCtx.UserRpc.Logout(ctx, logoutReq)
	if err != nil {
		return err
	}

	return nil
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
func (l *UserLogic) UpdateUser(req *types.UpdateUserRequest) (err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	updateReq := &user.UpdateUserRequest{}
	if err := copier.Copy(updateReq, req); err != nil {
		return err
	}

	_, err = l.svcCtx.UserRpc.UpdateUser(ctx, updateReq)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser 删除用户
func (l *UserLogic) DeleteUser(req *types.DeleteUserRequest) (err error) {
	ctx, cancel := context.WithTimeout(l.ctx, time.Second*5)
	defer cancel()

	deleteReq := &user.DeleteUserRequest{}
	if err := copier.Copy(deleteReq, req); err != nil {
		return err
	}

	_, err = l.svcCtx.UserRpc.DeleteUser(ctx, deleteReq)
	if err != nil {
		return err
	}

	return nil
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
