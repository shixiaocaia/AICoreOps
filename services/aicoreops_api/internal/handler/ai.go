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
 * File: ai.go
 */

package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/logic"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/types"
	"github.com/gorilla/websocket"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AiHandler struct {
	svcCtx *svc.ServiceContext
}

func NewAiHandler(svcCtx *svc.ServiceContext) *AiHandler {
	return &AiHandler{
		svcCtx: svcCtx,
	}
}

func (h *AiHandler) GetChatList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	pageInt, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "page 必须是正整数",
		})
		return
	}

	pageSize := r.URL.Query().Get("page_size")
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 32)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "page_size 必须是正整数",
		})
		return
	}

	l := logic.NewAiLogic(r.Context(), h.svcCtx)
	resp, err := l.GetChatList(r.Context(), int32(pageInt), int32(pageSizeInt))
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

func (h *AiHandler) GetChatHistory(w http.ResponseWriter, r *http.Request) {
	var req types.GetChatHistoryRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	l := logic.NewAiLogic(r.Context(), h.svcCtx)
	resp, err := l.GetChatHistory(&req)
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

func (h *AiHandler) UploadDocument(w http.ResponseWriter, r *http.Request) {
	var req types.UploadDocumentRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	l := logic.NewAiLogic(r.Context(), h.svcCtx)
	resp, err := l.UploadDocument(&req)
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

func (h *AiHandler) AskQuestion(w http.ResponseWriter, r *http.Request) {
	sessionId := r.URL.Query().Get("session_id")
	title := r.URL.Query().Get("title")
	scoreThreshold := r.URL.Query().Get("score_threshold")
	scoreThresholdFloat, err := strconv.ParseFloat(scoreThreshold, 32)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "score_threshold 必须是数字",
		})
		return
	}
	topK := r.URL.Query().Get("top_k")
	topKInt, err := strconv.Atoi(topK)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "top_k 必须是数字",
		})
		return
	}

	// 检查是否为 ws 连接
	if !websocket.IsWebSocketUpgrade(r) {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "请求不是WebSocket升级请求",
		})
		return
	}

	l := logic.NewAiLogic(r.Context(), h.svcCtx)
	resp, err := l.AskQuestion(w, r, sessionId, title, float32(scoreThresholdFloat), int32(topKInt))
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		fmt.Println(err)
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

func (h *AiHandler) NewChat(w http.ResponseWriter, r *http.Request) {
	l := logic.NewAiLogic(r.Context(), h.svcCtx)
	resp, err := l.NewChat()
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
