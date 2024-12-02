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

package domain

import (
	"aicoreops_user/internal/dao"
	"aicoreops_user/internal/model"
	"aicoreops_user/internal/pkg"
	"aicoreops_user/internal/repo"
	"aicoreops_user/types"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"

	regexp "github.com/dlclark/regexp2"

	"gorm.io/gorm"
)

const (
	UsernameRegexPattern = `^[a-zA-Z0-9_]{6,}$`                                                     // 用户名正则:至少6位字母数字下划线
	PasswordRegexPattern = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$` // 密码正则:至少8位,包含大小写字母、数字、特殊字符
	MagicStr             = "AICoreOps@2024"                                                         // 密码加密魔法字符串
)

type UserDomain struct {
	Username *regexp.Regexp
	PassWord *regexp.Regexp
	repo     repo.UserRepo
	ijwt     pkg.JWTHandler
}

func NewUserDomain(db *gorm.DB, ijwt pkg.JWTHandler) *UserDomain {
	return &UserDomain{
		Username: regexp.MustCompile(UsernameRegexPattern, regexp.None),
		PassWord: regexp.MustCompile(PasswordRegexPattern, regexp.None),
		repo:     dao.NewUserDao(db),
		ijwt:     ijwt,
	}
}

func (u *UserDomain) VerifyUsernameAndPassword(username string, password string) error {
	if username == "" || password == "" {
		return errors.New("用户名或密码不能为空")
	}

	// 验证username(长度至少6位且不能包含特殊字符)
	isMatch, err := u.Username.MatchString(username)
	if err != nil {
		return fmt.Errorf("验证用户名失败: %v", err)
	}
	if !isMatch {
		return errors.New("用户名格式不正确,需至少6位字母数字下划线")
	}

	// 验证密码(至少8位，包含大小写字母、数字、特殊字符)
	isMatch, err = u.PassWord.MatchString(password)
	if err != nil {
		return fmt.Errorf("验证密码失败: %v", err)
	}
	if !isMatch {
		return errors.New("密码格式不正确,需至少8位且包含大小写字母、数字、特殊字符")
	}

	return nil
}

func (u *UserDomain) EncryptPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("密码不能为空")
	}

	hash := md5.New()
	hash.Write([]byte(MagicStr + password + MagicStr))
	// 二次加密
	firstHash := hex.EncodeToString(hash.Sum(nil))
	hash.Reset()
	hash.Write([]byte(MagicStr + firstHash + MagicStr))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// CheckUsernameExists 检查用户名是否已存在
func (u *UserDomain) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	if username == "" {
		return false, errors.New("用户名不能为空")
	}

	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, fmt.Errorf("查询用户失败: %v", err)
	}

	return user != nil, nil
}

func (u *UserDomain) Register(ctx context.Context, user model.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("用户名或密码不能为空")
	}

	if err := u.repo.CreateUser(ctx, &user); err != nil {
		return fmt.Errorf("创建用户失败: %v", err)
	}

	return nil
}

func (u *UserDomain) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("用户名不能为空")
	}
	return u.repo.GetUserByUsername(ctx, username)
}

func (u *UserDomain) GetUserById(ctx context.Context, id int) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("无效的用户ID")
	}
	return u.repo.GetUserById(ctx, id)
}

func (u *UserDomain) ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int, error) {
	return u.repo.ListUsers(ctx, page, pageSize)
}

func (u *UserDomain) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}
	return u.repo.GetUserByUsernameAndPassword(ctx, username, password)
}

func (u *UserDomain) GenerateToken(ctx context.Context, id int) (string, string, error) {
	if id <= 0 {
		return "", "", errors.New("无效的用户ID")
	}
	return u.ijwt.SetLoginToken(id)
}

func (u *UserDomain) UpdateUser(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.New("用户信息不能为空")
	}
	return u.repo.UpdateUser(ctx, user)
}

func (u *UserDomain) UpdateLastLoginTime(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("无效的用户ID")
	}
	return u.repo.UpdateLastLoginTime(ctx, id)
}

// BuildUserModel 构建用户模型
func (u *UserDomain) BuildUserModel(req *types.CreateUserRequest, encryptedPwd string) model.User {
	if req == nil {
		return model.User{}
	}

	return model.User{
		Username: req.Username,
		Password: encryptedPwd,
		Email:    req.Email,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Status:   int(req.Status),
	}
}
