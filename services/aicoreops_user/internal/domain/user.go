package domain

import (
	"aicoreops_user/internal/dao"
	"aicoreops_user/internal/model"
	"aicoreops_user/internal/repo"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"

	regexp "github.com/dlclark/regexp2"

	"gorm.io/gorm"
)

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

const (
	UsernameRegexPattern = `^[a-zA-Z0-9_]{6,}$`                                                     // 用户名正则:至少6位字母数字下划线
	PasswordRegexPattern = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$` // 密码正则:至少8位,包含大小写字母、数字、特殊字符
)

type UserDomain struct {
	Username *regexp.Regexp
	PassWord *regexp.Regexp
	repo     repo.UserRepo
}

func NewUserDomain(db *gorm.DB) *UserDomain {
	return &UserDomain{
		Username: regexp.MustCompile(UsernameRegexPattern, regexp.None),
		PassWord: regexp.MustCompile(PasswordRegexPattern, regexp.None),
		repo:     dao.NewUserDao(db),
	}
}

func (u *UserDomain) VerifyUsernameAndPassword(username string, password string) (bool, error) {
	// 验证username(长度至少6位且不能包含特殊字符)
	isMatch, err := u.Username.MatchString(username)
	if err != nil {
		return false, errors.New(fmt.Sprintf("验证用户名失败: %v", err))
	}
	if !isMatch {
		return false, errors.New("用户名格式不正确,需至少6位字母数字下划线")
	}

	// 验证密码(至少8位，包含大小写字母、数字、特殊字符)
	isMatch, err = u.PassWord.MatchString(password)
	if err != nil {
		return false, errors.New(fmt.Sprintf("验证密码失败: %v", err))
	}
	if !isMatch {
		return false, errors.New("密码格式不正确,需至少8位且包含大小写字母、数字、特殊字符")
	}

	return true, nil
}

func (u *UserDomain) EncryptPassword(password string) (string, error) {
	hash := md5.New()
	magicStr := "AICoreOps@2024"
	hash.Write([]byte(magicStr + password + magicStr))
	// 二次加密
	firstHash := hex.EncodeToString(hash.Sum(nil))
	hash.Reset()
	hash.Write([]byte(magicStr + firstHash + magicStr))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 检查用户名是否已存在
func (u *UserDomain) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return user != nil, nil
}

func (u *UserDomain) Register(ctx context.Context, user model.User) error {
	if err := u.repo.CreateUser(ctx, &user); err != nil {
		return errors.New(fmt.Sprintf("创建用户失败: %v", err))
	}

	return nil
}
