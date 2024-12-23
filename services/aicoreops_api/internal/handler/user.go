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
 * File: user.go
 */

package handler

import (
	"net/http"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/logic"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/types"

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
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// 创建逻辑处理器并调用登录方法
	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.Login(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	// 设置响应头部
	w.Header().Set("x-jwt-token", resp.Data.JwtToken)
	w.Header().Set("x-refresh-token", resp.Data.RefreshToken)

	// 返回响应结果
	httpx.OkJsonCtx(r.Context(), w, resp)
}

// CreateUser 处理用户注册请求
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req types.CreateUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.CreateUser(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// GetUser 获取用户信息
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var req types.GetUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.GetUser(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var req types.DeleteUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.DeleteUser(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// UpdateUser 更新用户信息
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req types.UpdateUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.UpdateUser(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// ListUsers 获取用户列表
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	var req types.GetUserListRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.GetUserList(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// Logout 处理用户登出请求
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req types.LogoutRequest

	req.JWTToken = r.Header.Get("x-jwt-token")
	req.RefreshToken = r.Header.Get("x-refresh-token")

	l := logic.NewUserLogic(r.Context(), h.svcCtx)
	resp, err := l.Logout(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}
