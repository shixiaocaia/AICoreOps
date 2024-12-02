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
 * File: jwt.go
 */

package pkg

import (
	"aicoreops_user/internal/config"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// 定义常量
const (
	bearerPrefix  = "Bearer "
	blacklistKey  = "aicoreops:user:blacklist:ssid:%s"
	defaultExpire = time.Hour * 24 * 7
)

// 定义错误
var (
	ErrInvalidUserID    = errors.New("无效的用户ID")
	ErrEmptySessionID   = errors.New("会话ID不能为空")
	ErrInvalidToken     = errors.New("无效的授权令牌")
	ErrInvalidFormat    = errors.New("授权令牌格式无效")
	ErrEmptyContext     = errors.New("上下文不能为空")
	ErrTokenExpired     = errors.New("令牌已过期")
	ErrEmptyRedisClient = errors.New("Redis客户端不能为空")
	ErrEmptyJWTSecret   = errors.New("JWT密钥不能为空")
)

// JWTHandler 定义 JWT 处理接口
type JWTHandler interface {
	SetLoginToken(uid int) (accessToken string, refreshToken string, err error)
	SetJWTToken(uid int, ssid string) (string, error)
	ClearToken(ctx context.Context, token string, refreshToken string) error
	setRefreshToken(uid int, ssid string) (string, error)
}

// UserClaims 用户声明结构体
type UserClaims struct {
	jwt.RegisteredClaims
	Uid         int    `json:"uid"`
	Ssid        string `json:"ssid"`
	UserAgent   string `json:"user_agent"`
	ContentType string `json:"content_type"`
}

// RefreshClaims 刷新令牌声明结构体
type RefreshClaims struct {
	jwt.RegisteredClaims
	Uid  int    `json:"uid"`
	Ssid string `json:"ssid"`
}

// jwtHandler JWT处理器实现
type jwtHandler struct {
	client        redis.Cmdable
	signingMethod jwt.SigningMethod
	rcExpiration  time.Duration
	secret        []byte
	refreshSecret []byte
	expire        int64
}

// NewJWTHandler 创建新的JWT处理器
func NewJWTHandler(c redis.Cmdable, jwtConfig config.JWTConfig) JWTHandler {
	if c == nil {
		panic(ErrEmptyRedisClient)
	}
	if len(jwtConfig.Secret) == 0 {
		panic(ErrEmptyJWTSecret)
	}

	refreshSecret := generateRefreshSecret([]byte(jwtConfig.Secret))

	return &jwtHandler{
		client:        c,
		signingMethod: jwt.SigningMethodHS512,
		rcExpiration:  defaultExpire,
		secret:        []byte(jwtConfig.Secret),
		refreshSecret: refreshSecret,
		expire:        jwtConfig.Expire,
	}
}

// generateRefreshSecret 生成刷新令牌密钥
func generateRefreshSecret(secret []byte) []byte {
	refreshSecret := make([]byte, len(secret))
	copy(refreshSecret, secret)
	for i := range refreshSecret {
		refreshSecret[i] = refreshSecret[i] ^ 0x55
	}
	return refreshSecret
}

// SetLoginToken 设置登录令牌
func (j *jwtHandler) SetLoginToken(uid int) (string, string, error) {
	if uid <= 0 {
		return "", "", ErrInvalidUserID
	}

	ssid := uuid.New().String()
	refreshToken, err := j.setRefreshToken(uid, ssid)
	if err != nil {
		return "", "", fmt.Errorf("设置刷新令牌失败: %w", err)
	}

	jwtToken, err := j.SetJWTToken(uid, ssid)
	if err != nil {
		return "", "", fmt.Errorf("设置JWT令牌失败: %w", err)
	}

	return jwtToken, refreshToken, nil
}

// SetJWTToken 设置JWT令牌
func (j *jwtHandler) SetJWTToken(uid int, ssid string) (string, error) {
	if uid <= 0 {
		return "", ErrInvalidUserID
	}
	if ssid == "" {
		return "", ErrEmptySessionID
	}

	claims := UserClaims{
		Uid:  uid,
		Ssid: ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(j.expire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	return j.generateToken(claims, j.secret)
}

// setRefreshToken 设置刷新令牌
func (j *jwtHandler) setRefreshToken(uid int, ssid string) (string, error) {
	if uid <= 0 {
		return "", ErrInvalidUserID
	}
	if ssid == "" {
		return "", ErrEmptySessionID
	}

	claims := RefreshClaims{
		Uid:  uid,
		Ssid: ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.rcExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	return j.generateToken(claims, j.refreshSecret)
}

// generateToken 生成令牌
func (j *jwtHandler) generateToken(claims jwt.Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(j.signingMethod, claims)
	signedString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("令牌签名失败: %w", err)
	}
	return signedString, nil
}

// ClearToken 清除令牌
func (j *jwtHandler) ClearToken(ctx context.Context, token string, refreshToken string) error {
	claims := &UserClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != j.signingMethod {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return j.secret, nil
	})

	if err != nil || !parsedToken.Valid {
		return ErrInvalidToken
	}

	if err := j.addToBlacklist(ctx, token, claims.ExpiresAt.Time); err != nil {
		return fmt.Errorf("添加访问令牌到黑名单失败: %w", err)
	}

	refreshClaims := &RefreshClaims{}
	parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, refreshClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != j.signingMethod {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return j.refreshSecret, nil
	})

	if err != nil || !parsedRefreshToken.Valid {
		return ErrInvalidToken
	}

	if err := j.addToBlacklist(ctx, refreshToken, refreshClaims.ExpiresAt.Time); err != nil {
		return fmt.Errorf("添加刷新令牌到黑名单失败: %w", err)
	}

	return nil
}

// addToBlacklist 添加令牌到黑名单
func (j *jwtHandler) addToBlacklist(ctx context.Context, ssid string, expiresAt time.Time) error {
	if ctx == nil {
		return ErrEmptyContext
	}
	if ssid == "" {
		return ErrInvalidToken
	}

	remainingTime := time.Until(expiresAt)
	if remainingTime <= 0 {
		return ErrTokenExpired
	}

	key := fmt.Sprintf(blacklistKey, ssid)
	if err := j.client.Set(ctx, key, "invalid", remainingTime).Err(); err != nil {
		return fmt.Errorf("设置黑名单失败: %w", err)
	}

	return nil
}
