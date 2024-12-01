package handler

import (
	"aicoreops_api/internal/logic"
	"aicoreops_api/internal/svc"
	"aicoreops_api/internal/types"
	"aicoreops_common"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type MenuHandler struct {
	svcCtx *svc.ServiceContext
}

func NewMenuHandler(svcCtx *svc.ServiceContext) *MenuHandler {
	return &MenuHandler{
		svcCtx: svcCtx,
	}
}

// CreateMenu 创建菜单
func (h *MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var req types.CreateMenuRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewMenuLogic(r.Context(), h.svcCtx)
	resp, err := l.CreateMenu(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// GetMenu 获取菜单详情
func (h *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	var req types.GetMenuRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewMenuLogic(r.Context(), h.svcCtx)
	resp, err := l.GetMenu(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// UpdateMenu 更新菜单
func (h *MenuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	var req types.UpdateMenuRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewMenuLogic(r.Context(), h.svcCtx)
	resp, err := l.UpdateMenu(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// DeleteMenu 删除菜单
func (h *MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	var req types.DeleteMenuRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewMenuLogic(r.Context(), h.svcCtx)
	resp, err := l.DeleteMenu(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// ListMenus 获取菜单列表
func (h *MenuHandler) ListMenus(w http.ResponseWriter, r *http.Request) {
	var req types.ListMenusRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewMenuLogic(r.Context(), h.svcCtx)
	resp, err := l.ListMenus(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}
