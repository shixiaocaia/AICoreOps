package tools

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

// 定义常量和错误
const (
	bearerPrefix = "Bearer "
	blacklistKey = "aicoreops:user:blacklist:ssid:%s"
)

var (
	ErrEmptyToken       = errors.New("token不能为空")
	ErrEmptySecret      = errors.New("secret不能为空")
	ErrInvalidFormat    = errors.New("token格式无效")
	ErrInvalidToken     = errors.New("无效的token")
	ErrTokenExpired     = errors.New("token已过期")
	ErrInvalidUserID    = errors.New("无效的用户ID")
	ErrMissingClaims    = errors.New("token缺少必要的声明信息")
	ErrTokenBlacklisted = errors.New("token已被加入黑名单")
)

// Claims 定义token的声明结构
type Claims struct {
	jwt.RegisteredClaims
	Uid int `json:"uid"`
}

func ParseToken(tokenString string, secret string) (int64, error) {
	// 基础验证
	if tokenString == "" {
		return 0, ErrEmptyToken
	}

	if secret == "" {
		return 0, ErrEmptySecret
	}

	tokenString = strings.TrimPrefix(tokenString, bearerPrefix)

	// 解析token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名算法: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, ErrTokenExpired
		}
		return 0, fmt.Errorf("解析token失败: %w", err)
	}

	if !token.Valid {
		return 0, ErrInvalidToken
	}

	// 验证必要字段
	if claims.Uid <= 0 {
		return 0, ErrInvalidUserID
	}

	return int64(claims.Uid), nil
}

func ValidateTokenBlacklist(ctx context.Context, rdb redis.Cmdable, tokenString string) error {
	if rdb == nil {
		return errors.New("redis客户端不能为空")
	}

	key := fmt.Sprintf(blacklistKey, tokenString)
	exists, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("检查黑名单失败: %w", err)
	}

	if exists > 0 {
		return ErrTokenBlacklisted
	}

	return nil
}
