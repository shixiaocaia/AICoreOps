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
 * File: ai_logic.go
 */

package logic

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"time"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/middleware"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/types"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_common/types/ai"
	"google.golang.org/grpc/metadata"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type AiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AiLogic {
	return &AiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	timeoutDuration = 600 * time.Second
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// GetHistoryList 获取历史记录列表
func (l *AiLogic) GetHistoryList(ctx context.Context) (*ai.GetHistoryListResponse, error) {
	uidValue := ctx.Value(middleware.UidKey{})
	uid, ok := uidValue.(int64)
	if !ok {
		return nil, fmt.Errorf("无效的用户ID类型或未找到用户ID")
	}

	resp, err := l.svcCtx.AiRpc.GetHistoryList(l.ctx, &ai.GetHistoryListRequest{
		UserId: strconv.FormatInt(uid, 10),
	})
	if err != nil {
		return nil, fmt.Errorf("获取历史会话失败: %v", err)
	}

	return resp, nil
}

// GetChatHistory 获取聊天历史
func (l *AiLogic) GetChatHistory(req *types.GetChatHistoryRequest) (*ai.GetChatHistoryResponse, error) {
	resp, err := l.svcCtx.AiRpc.GetChatHistory(l.ctx, &ai.GetChatHistoryRequest{
		SessionId: req.SessionId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取聊天历史失败: %v", err)
	}
	return resp, nil
}

// UploadDocument 上传文档
func (l *AiLogic) UploadDocument(req *types.UploadDocumentRequest) (*ai.UploadDocumentResponse, error) {
	resp, err := l.svcCtx.AiRpc.UploadDocument(l.ctx, &ai.UploadDocumentRequest{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("上传文档失败: %v", err)
	}
	return resp, nil
}

// AskQuestion 提问
func (l *AiLogic) AskQuestion(w http.ResponseWriter, r *http.Request, sessionId string) (*ai.AskQuestionResponse, error) {
	// 1. 验证会话有效性
	uidValue := l.ctx.Value(middleware.UidKey{})
	uid, ok := uidValue.(int64)
	if !ok {
		return nil, fmt.Errorf("无效的用户ID类型或未找到用户ID")
	}

	// 2. 建立 ws 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		l.Logger.Errorf("建立 ws 连接失败: %v", err)
		return nil, err
	}
	defer conn.Close()

	conn.SetReadLimit(32 * 1024) // 32KB
	conn.SetReadDeadline(time.Now().Add(timeoutDuration))
	conn.SetWriteDeadline(time.Now().Add(timeoutDuration))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		return nil
	})

	// 3. 建立 stream 双向流 RPC
	md := metadata.Pairs("sessionId", sessionId, "userId", strconv.FormatInt(uid, 10))
	ctx, cancel := context.WithCancel(metadata.NewOutgoingContext(l.ctx, md))
	defer cancel()

	stream, err := l.svcCtx.AiRpc.AskQuestion(l.ctx)
	if err != nil {
		l.Logger.Errorf("建立 gRPC 双向流失败: %v", err)
		l.sendWebSocketError(conn, "gRPC 连接失败")
		conn.Close()
		return nil, err
	}

	// 4. ws -- streaming rpc -- ws
	opCtx, opCancel := context.WithCancel(ctx)
	defer opCancel()
	timeout := time.NewTimer(timeoutDuration)
	defer timeout.Stop()

	var wg sync.WaitGroup
	wg.Add(2)

	// 4.1 接收 gRPC 响应并发送到 WebSocket
	go func() {
		defer wg.Done()
		l.receiveResponses(conn, stream, opCancel, sessionId)
	}()

	// 4.2从 WebSocket 接收消息并发送到 gRPC
	go func() {
		defer wg.Done()
		l.sendMessages(conn, stream, opCtx, timeout, sessionId)
	}()

	// 4.3监控超时
	go func() {
		select {
		case <-timeout.C:
			l.Logger.Infof("sessionId: %s, 连接超时，关闭连接", sessionId)
			l.sendWebSocketError(conn, "连接超时")
			opCancel()
			conn.Close()
			stream.CloseSend()
		case <-opCtx.Done():
			l.Logger.Infof("sessionId: %s, 正常结束", sessionId)
		}
	}()

	wg.Wait()
	l.Logger.Infof("sessionId: %s, 连接已关闭", sessionId)
	return nil, nil
}

func (l *AiLogic) receiveResponses(conn *websocket.Conn, stream ai.AIHelper_AskQuestionClient, cancel context.CancelFunc, sessionId string) {
	defer func() {
		l.Logger.Infof("sessionId: %s, 接收响应 goroutine 退出", sessionId)
		cancel() // 通知其他 goroutine 退出
	}()

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				l.Logger.Infof("sessionId: %s, gRPC 流已关闭", sessionId)
			} else {
				l.Logger.Errorf("sessionId: %s, 接收 gRPC 响应失败: %v", sessionId, err)
				l.sendWebSocketError(conn, fmt.Sprintf("接收响应失败: %v", err))
			}
			return
		}

		message := fmt.Sprintf("收到回答: %s", resp.GetData().GetAnswer())
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			l.Logger.Errorf("sessionId: %s, 发送 WebSocket 消息失败: %v", sessionId, err)
			return
		}
	}
}

func (l *AiLogic) sendMessages(conn *websocket.Conn, stream ai.AIHelper_AskQuestionClient, ctx context.Context, timeout *time.Timer, sessionId string) {
	for {
		select {
		case <-ctx.Done():
			l.Logger.Infof("sessionId: %s, 上下文已取消", sessionId)
			return

		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err,
					websocket.CloseNormalClosure,
					websocket.CloseGoingAway,
					websocket.CloseNoStatusReceived) {
					l.Logger.Infof("sessionId: %s, WebSocket 正常关闭: %v", sessionId, err)
				} else {
					l.Logger.Errorf("sessionId: %s, WebSocket 异常关闭: %v", sessionId, err)
				}
				// 通知其他 goroutine 关闭
				stream.CloseSend()
				return
			}

			l.Logger.Debugf("sessionId: %s, 接收到 WebSocket 消息: %s", sessionId, string(message))

			// 重置超时计时器
			if !timeout.Stop() {
				select {
				case <-timeout.C:
				default:
				}
			}
			timeout.Reset(timeoutDuration)
			l.Logger.Debugf("sessionId: %s, 重置超时计时器", sessionId)

			// 发送消息到 gRPC 流
			req := &ai.AskQuestionRequest{
				Question:  string(message),
				SessionId: sessionId,
			}

			if err := stream.Send(req); err != nil {
				l.Logger.Errorf("sessionId: %s, 发送 gRPC 请求失败: %v", sessionId, err)
				l.sendWebSocketError(conn, fmt.Sprintf("发送请求失败: %v", err))
				return
			}
		}
	}
}

func (l *AiLogic) sendWebSocketError(conn *websocket.Conn, errorMsg string) {
	if err := conn.WriteMessage(websocket.TextMessage, []byte(errorMsg)); err != nil {
		l.Logger.Errorf("发送 WebSocket 错误消息失败: %v", err)
	}
}
