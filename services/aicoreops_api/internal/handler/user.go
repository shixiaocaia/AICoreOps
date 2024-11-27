package handler

import (
	"aicoreops_api/internal/logic"
	"aicoreops_api/internal/svc"
	"aicoreops_api/internal/types"
	"aicoreops_common"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type UserHandler struct {
	svcCtx *svc.ServiceContext
}

func NewUserHandler(svcCtx *svc.ServiceContext) *UserHandler {
	return &UserHandler{
		svcCtx: svcCtx,
	}
}

// Login 处理用户登录请求
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	var req types.LoginRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	// 创建逻辑处理器并调用登录方法
	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.Login(&req)
	// 处理响应结果
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	// 返回响应结果
	httpx.OkJsonCtx(r.Context(), w, result)
}

// CreateUser 处理用户注册请求
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req types.CreateUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	err := l.CreateUser(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(nil, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// GetUser 获取用户信息
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var req types.GetUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.GetUser(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var req types.DeleteUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	err := l.DeleteUser(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(nil, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// UpdateUser 更新用户信息
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req types.UpdateUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	err := l.UpdateUser(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(nil, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// ListUsers 获取用户列表
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	var req types.GetUserListRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.GetUserList(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// Logout 处理用户登出请求
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req types.LogoutRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	err := l.Logout(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(nil, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}
