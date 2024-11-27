package middleware

import (
	"aicoreops_common"
	"aicoreops_common/tools"
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type contextKey string

const userIDKey contextKey = "uid"

type AuthMiddleware struct {
	secret string
	rdb    redis.Cmdable
}

func NewAuthMiddleware(secret string, rdb redis.Cmdable) *AuthMiddleware {
	return &AuthMiddleware{
		secret: secret,
		rdb:    rdb,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := aicoreops_common.NewResultResponse()

		// 获取 token
		token := r.Header.Get("Authorization")
		if token == "" {
			logx.Error("token为空")
			response.SetFailResponse(aicoreops_common.BizCodeUnauthorized, tools.ErrEmptyToken.Error())
			httpx.WriteJson(w, http.StatusUnauthorized, response)
			return
		}

		// 提取token
		const bearerPrefix = "Bearer "
		if len(token) > len(bearerPrefix) && token[:len(bearerPrefix)] == bearerPrefix {
			token = token[len(bearerPrefix):]
		} else {
			logx.Error("token格式错误")
			response.SetFailResponse(aicoreops_common.BizCodeUnauthorized, "invalid token format")
			httpx.WriteJson(w, http.StatusUnauthorized, response)
			return
		}

		// 检查token是否在黑名单中
		if err := tools.ValidateTokenBlacklist(r.Context(), m.rdb, token); err != nil {
			logx.Errorf("token黑名单验证失败: %v", err)
			response.SetFailResponse(aicoreops_common.BizCodeUnauthorized, err.Error())
			httpx.WriteJson(w, http.StatusUnauthorized, response)
			return
		}

		// 解析token获取用户ID
		uid, err := tools.ParseToken(token, m.secret)
		if err != nil {
			logx.Errorf("解析token失败: %v", err)
			response.SetFailResponse(aicoreops_common.BizCodeUnauthorized, err.Error())
			httpx.WriteJson(w, http.StatusUnauthorized, response)
			return
		}

		// 设置用户ID到上下文
		ctx := context.WithValue(r.Context(), userIDKey, uid)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
