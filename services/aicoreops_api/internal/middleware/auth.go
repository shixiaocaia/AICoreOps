package middleware

import (
	"aicoreops_common"
	"aicoreops_common/tools"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type contextKey string

const userIDKey contextKey = "uid"

func AuthMiddleware(secret string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			response := aicoreops_common.NewResultResponse()

			// 获取 token
			token := r.Header.Get("Authorization")
			if token == "" {
				logx.Errorf("token为空")
				response.SetFailResponse(aicoreops_common.BizCodeUnauthorized, tools.ErrEmptyToken.Error())
				httpx.WriteJson(w, http.StatusUnauthorized, response)
				return
			}

			// 解析token获取用户ID
			uid, err := tools.ParseToken(token, secret)
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
}
