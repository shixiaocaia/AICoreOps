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
 * File: auth.go
 */

package middleware

import (
	"context"
	"net/http"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_common"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_common/tools"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// type contextKey string

// const userIDKey contextKey = "uid"

type UidKey struct{}

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
			// 如果是 ws 协议，从sec-websocket-protocol获取 token
			token = r.Header.Get("Sec-Websocket-Protocol")
		}
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
		} else if r.URL.Path != "/api/ai/ask" {
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

		reqCtx := r.Context()

		// 设置用户ID到上下文
		ctx := context.WithValue(reqCtx, UidKey{}, uid)
		r = r.WithContext(ctx)

		next(w, r)
	}
}

// func UserIDKey() interface{} {
// 	return userIDKey
// }
