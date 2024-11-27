package tools

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// 定义常量和错误
const (
	bearerPrefix = "Bearer "
)

var (
	ErrEmptyToken    = errors.New("token不能为空")
	ErrEmptySecret   = errors.New("secret不能为空")
	ErrInvalidFormat = errors.New("token格式无效")
	ErrInvalidToken  = errors.New("无效的token")
	ErrTokenExpired  = errors.New("token已过期")
	ErrInvalidUserID = errors.New("无效的用户ID")
	ErrMissingClaims = errors.New("token缺少必要的声明信息")
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

	// 提取Bearer token
	if !strings.HasPrefix(tokenString, bearerPrefix) {
		return 0, ErrInvalidFormat
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
