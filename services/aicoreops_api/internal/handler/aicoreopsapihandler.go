package handler

import (
	"net/http"

	"aicoreops_api/internal/logic"
	"aicoreops_api/internal/svc"
	"aicoreops_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Aicoreops_apiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAicoreops_apiLogic(r.Context(), svcCtx)
		resp, err := l.Aicoreops_api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
