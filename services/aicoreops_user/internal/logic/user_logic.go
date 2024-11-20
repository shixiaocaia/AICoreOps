package logic

import (
	"context"
	"errors"

	"aicoreops_user/internal/domain"
	"aicoreops_user/internal/model"
	"aicoreops_user/internal/svc"
	"aicoreops_user/types"

	"github.com/zeromicro/go-zero/core/logx"
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
		domain: domain.NewUserDomain(svcCtx.DB),
	}
}

// CreateUser 创建用户
func (l *UserLogic) CreateUser(ctx context.Context, req *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	// 验证username(长度至少6位且不能包含特殊字符)，密码(至少8位，包含大小写字母、数字、特殊字符)
	isValid, err := l.domain.VerifyUsernameAndPassword(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("用户名或密码格式不正确")
	}

	// 检查用户名是否已存在
	exists, err := l.domain.CheckUsernameExists(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 加密密码
	encryptedPwd, err := l.domain.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username: req.Username,
		Password: encryptedPwd,
		Email:    req.Email,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Status:   int(req.Status),
	}

	// 注册用户
	err = l.domain.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &types.CreateUserResponse{
		Code:    0,
		Message: "创建用户成功",
	}, nil
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
