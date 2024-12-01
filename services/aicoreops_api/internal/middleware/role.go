package middleware

import (
	"aicoreops_common"
	"aicoreops_common/tools"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type CasbinMiddleware struct {
	enforcer *casbin.Enforcer
}

func NewCasbinMiddleware(enforcer *casbin.Enforcer) *CasbinMiddleware {
	return &CasbinMiddleware{
		enforcer: enforcer,
	}
}

func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := aicoreops_common.NewResultResponse()

		// 从上下文获取用户ID
		uidValue := r.Context().Value(userIDKey)
		uid, ok := uidValue.(int64)
		if !ok || uid <= 0 {
			logx.Error("无效的用户ID类型或未找到用户ID")
			response.SetFailResponse(aicoreops_common.BizCodeForbidden, "无效的用户ID")
			httpx.WriteJson(w, http.StatusForbidden, response)
			return
		}

		// 获取请求方法和路径
		method := r.Method
		path := r.URL.Path
		// 测试 列出所有策略
		policies, err := m.enforcer.GetPolicy()
		if err != nil {
			logx.Errorf("获取策略失败: %v", err)
			response.SetFailResponse(aicoreops_common.BizCodeForbidden, "获取策略失败")
			httpx.WriteJson(w, http.StatusForbidden, response)
			return
		}
		logx.Info("策略: ", policies)

		// 检查权限
		allowed, err := tools.CheckPermission(m.enforcer, uid, path, method)
		if err != nil {
			logx.Errorf("权限检查失败: %v", err)
			response.SetFailResponse(aicoreops_common.BizCodeForbidden, "权限检查失败")
			httpx.WriteJson(w, http.StatusForbidden, response)
			return
		}

		if !allowed {
			logx.Errorf("用户 %d 没有访问权限: %s %s", uid, method, path)
			response.SetFailResponse(aicoreops_common.BizCodeForbidden, "没有访问权限")
			httpx.WriteJson(w, http.StatusForbidden, response)
			return
		}

		next(w, r)
	}
}
