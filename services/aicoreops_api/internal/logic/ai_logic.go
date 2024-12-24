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
	"log"
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
	// 1. 建立 ws 连接，stream 双向流 RPC
	md := metadata.Pairs("sessionId", sessionId)
	ctx, cancel := context.WithCancel(metadata.NewOutgoingContext(l.ctx, md))
	defer cancel()

	l.Logger.Debugf("请求头信息: %v", md)

	stream, err := l.svcCtx.AiRpc.AskQuestion(ctx)
	if err != nil {
		l.Logger.Errorf("建立 gRPC 双向流失败: %v", err)
		return nil, err
	}
	// 使用 WaitGroup 确保所有 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 创建一个用于取消所有操作的上下文
	opCtx, opCancel := context.WithCancel(ctx)
	defer opCancel()

	// 设置超时计时器
	timeout := time.NewTimer(timeoutDuration)
	defer timeout.Stop()

	// 接收 gRPC 响应并发送到 WebSocket
	go func() {
		defer wg.Done()
		l.receiveResponses(conn, stream, opCancel, sessionId)
	}()

	// 从 WebSocket 接收消息并发送到 gRPC
	go func() {
		defer wg.Done()
		l.sendMessages(conn, stream, opCtx, timeout, sessionId)
	}()

	// 监控超时
	go func() {
		select {
		case <-timeout.C:
			l.Logger.Infof("%s 连接超时", sessionId)
			sendWebSocketError(conn, "连接超时")
			opCancel()
			conn.Close()
			stream.CloseSend()
		case <-opCtx.Done():
			l.Logger.Infof("%s 正常结束", sessionId)
		}
	}()

	// 等待所有 goroutine 完成
	wg.Wait()

	l.Logger.Infof("%s 连接已关闭", sessionId)
	return nil, nil
}

func (l *AiLogic) receiveResponses(conn *websocket.Conn, stream ai.AIHelper_AskQuestionClient, cancel context.CancelFunc, sessionId string) {
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				l.Logger.Infof("%s gRPC 流已关闭", sessionId)
			} else {
				l.Logger.Errorf("%s 接收 gRPC 响应失败: %v", sessionId, err)
			}
			cancel()
			return
		}

		message := fmt.Sprintf("收到回答: %s", resp.GetData().GetAnswer())
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("发送 WebSocket 消息失败: %v", err)
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
					log.Printf("WebSocket 异常关闭: %v", err)
				} else {
					log.Printf("读取 WebSocket 消息失败: %v", err)
				}
				return
			}

			log.Println("接收到 WebSocket 消息:", string(message))

			// 重置超时计时器
			if !timeout.Stop() {
				select {
				case <-timeout.C:
				default:
				}
			}
			timeout.Reset(timeoutDuration)
			log.Println("重置超时计时器")

			req := &ai.AskQuestionRequest{
				Question:  string(message),
				SessionId: sessionId,
			}

			if err := stream.Send(req); err != nil {
				l.Logger.Errorf("%s 发送 gRPC 请求失败: %v", sessionId, err)
				return
			}
		}
	}
}

func sendWebSocketError(conn *websocket.Conn, errorMessage string) {
	errMsg := fmt.Sprintf("错误: %s", errorMessage)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(errMsg)); err != nil {
		log.Printf("发送 WebSocket 错误消息失败: %v", err)
	}
}
