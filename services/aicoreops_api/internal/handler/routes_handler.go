package handler

import (
	"aicoreops_api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, "pong")
	})
}
