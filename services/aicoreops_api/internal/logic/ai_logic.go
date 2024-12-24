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
	"strconv"
	"sync"
	"time"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_common/types/ai"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc/metadata"

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
	timeoutDuration = 60 * time.Second
)

// AskQuestion 提问
func (l *AiLogic) AskQuestion(conn *websocket.Conn, sessionId string) (*ai.AskQuestionResponse, error) {
	// 1. 验证会话有效性
	var userId int64
	if userId, ok := l.validateSession(sessionId); !ok {
		l.Logger.Errorf("无效的 sessionId: %s, userId: %d", sessionId, userId)
		l.sendWebSocketError(conn, "无效的会话")
		conn.Close()
		return nil, fmt.Errorf("无效的会话")
	}

	// 2. 建立 ws 连接，stream 双向流 RPC
	md := metadata.Pairs("sessionId", sessionId, "userId", strconv.FormatInt(userId, 10))
	ctx, cancel := context.WithCancel(metadata.NewOutgoingContext(l.ctx, md))
	defer cancel()

	stream, err := l.svcCtx.AiRpc.AskQuestion(ctx)
	if err != nil {
		l.Logger.Errorf("建立 gRPC 双向流失败: %v", err)
		l.sendWebSocketError(conn, "gRPC 连接失败")
		conn.Close()
		return nil, err
	}

	opCtx, opCancel := context.WithCancel(ctx)
	defer opCancel()
	timeout := time.NewTimer(timeoutDuration)
	defer timeout.Stop()
	var wg sync.WaitGroup
	wg.Add(2)

	// 3.1 接收 gRPC 响应并发送到 WebSocket
	go func() {
		defer wg.Done()
		l.receiveResponses(conn, stream, opCancel, sessionId)
	}()

	// 3.2从 WebSocket 接收消息并发送到 gRPC
	go func() {
		defer wg.Done()
		l.sendMessages(conn, stream, opCtx, timeout, sessionId)
	}()

	// 3.3监控超时
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

// validateSession 验证会话有效性
func (l *AiLogic) validateSession(sessionId string) (int64, bool) {
	// TODO: 实现会话验证逻辑
	l.Logger.Debugf("sessionId: %s, 会话验证成功", sessionId)
	return 0, true
}

func (l *AiLogic) receiveResponses(conn *websocket.Conn, stream ai.AIHelper_AskQuestionClient, cancel context.CancelFunc, sessionId string) {
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				l.Logger.Infof("sessionId: %s, gRPC 流已关闭", sessionId)
			} else {
				l.Logger.Errorf("%s 接收 gRPC 响应失败: %v", sessionId, err)
			}
			cancel()
			return
		}

		message := fmt.Sprintf("收到回答: %s", resp.GetData().GetAnswer())
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			l.Logger.Errorf("sessionId: %s, 发送 WebSocket 消息失败: %v", sessionId, err)
			cancel()
			return
		}
	}
}

func (l *AiLogic) sendMessages(conn *websocket.Conn, stream ai.AIHelper_AskQuestionClient, ctx context.Context, timeout *time.Timer, sessionId string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					l.Logger.Errorf("WebSocket 异常关闭: %v", err)
				} else {
					l.Logger.Errorf("读取 WebSocket 消息失败: %v", err)
				}
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

			req := &ai.AskQuestionRequest{
				Question:  string(message),
				SessionId: sessionId,
			}

			if err := stream.Send(req); err != nil {
				l.Logger.Errorf("sessionId: %s, 发送 gRPC 请求失败: %v", sessionId, err)
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
