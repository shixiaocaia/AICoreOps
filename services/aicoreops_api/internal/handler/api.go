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
 * File: api.go
 */

package handler

import (
	"aicoreops_api/internal/logic"
	"aicoreops_api/internal/svc"
	"aicoreops_api/internal/types"
	"aicoreops_common"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type ApiHandler struct {
	svcCtx *svc.ServiceContext
}

func NewApiHandler(svcCtx *svc.ServiceContext) *ApiHandler {
	return &ApiHandler{
		svcCtx: svcCtx,
	}
}

// CreateApi 创建API
func (h *ApiHandler) CreateApi(w http.ResponseWriter, r *http.Request) {
	var req types.CreateApiRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewApiLogic(r.Context(), h.svcCtx)
	resp, err := l.CreateApi(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// GetApi 获取API详情
func (h *ApiHandler) GetApi(w http.ResponseWriter, r *http.Request) {
	var req types.GetApiRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewApiLogic(r.Context(), h.svcCtx)
	resp, err := l.GetApi(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// UpdateApi 更新API
func (h *ApiHandler) UpdateApi(w http.ResponseWriter, r *http.Request) {
	var req types.UpdateApiRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewApiLogic(r.Context(), h.svcCtx)
	resp, err := l.UpdateApi(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// DeleteApi 删除API
func (h *ApiHandler) DeleteApi(w http.ResponseWriter, r *http.Request) {
	var req types.DeleteApiRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewApiLogic(r.Context(), h.svcCtx)
	resp, err := l.DeleteApi(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}

// ListApis 获取API列表
func (h *ApiHandler) ListApis(w http.ResponseWriter, r *http.Request) {
	var req types.ListApisRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.Error(w, err)
		return
	}

	l := logic.NewApiLogic(r.Context(), h.svcCtx)
	resp, err := l.ListApis(&req)
	result := aicoreops_common.NewResultResponse().HandleResponse(&resp, err)

	httpx.OkJsonCtx(r.Context(), w, result)
}
