package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"aicoreops_ai/types"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	grpcAddress      = "localhost:8083"
	webSocketAddress = ":8080"
	sessionID        = "1"
	timeoutDuration  = 600 * time.Second
)

// 使用 TLS 连接 gRPC 服务器的示例。如果不需要 TLS，可以继续使用 insecure 连接。
// 请确保正确配置 TLS 证书路径和相关信息。
// var grpcOptions = []grpc.DialOption{
// 	grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
// }

var grpcOptions = []grpc.DialOption{
	grpc.WithInsecure(), // 注意：WithInsecure 已被弃用，建议使用安全连接
	grpc.WithBlock(),
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 根据需要调整跨域策略
		return true
	},
}

func main() {
	// 捕获系统中断信号以实现优雅关闭
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	http.HandleFunc("/ws", handleWebSocket)

	server := &http.Server{
		Addr: webSocketAddress,
	}

	go func() {
		log.Printf("HTTP 服务启动，监听端口 %s", webSocketAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP 服务启动失败: %v", err)
		}
	}()

	<-stopChan
	log.Println("收到中断信号，正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务器关闭失败: %v", err)
	}

	log.Println("服务器已优雅关闭")
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("接收到 WebSocket 请求")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("无法升级为 WebSocket: %v", err)
		return
	}
	defer conn.Close()

	grpcConn, err := grpc.Dial(grpcAddress, grpcOptions...)
	if err != nil {
		log.Printf("无法连接到 gRPC 服务器: %v", err)
		sendWebSocketError(conn, "内部服务器错误")
		return
	}
	defer grpcConn.Close()

	client := types.NewAIHelperClient(grpcConn)

	md := metadata.Pairs("sessionID", sessionID)
	ctx, cancel := context.WithCancel(metadata.NewOutgoingContext(context.Background(), md))
	defer cancel()

	stream, err := client.AskQuestion(ctx)
	if err != nil {
		log.Printf("无法创建 gRPC 流: %v", err)
		sendWebSocketError(conn, "内部服务器错误")
		return
	}

	// 使用 WaitGroup 确保所有 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 创建一个用于取消所有操作的上下文
	opCtx, opCancel := context.WithCancel(context.Background())
	defer opCancel()

	// 设置超时计时器
	timeout := time.NewTimer(timeoutDuration)
	defer timeout.Stop()

	// 接收 gRPC 响应并发送到 WebSocket
	go func() {
		defer wg.Done()
		receiveResponses(conn, stream, opCancel)
	}()

	// 从 WebSocket 接收消息并发送到 gRPC
	go func() {
		defer wg.Done()
		sendMessages(conn, stream, opCtx, timeout)
	}()

	// 监控超时
	go func() {
		select {
		case <-timeout.C:
			log.Println("超时未收到消息，断开连接")
			sendWebSocketError(conn, "连接超时")
			opCancel()
			conn.Close()
			stream.CloseSend()
		case <-opCtx.Done():
			// 正常结束
		}
	}()

	// 等待所有 goroutine 完成
	wg.Wait()

	log.Println("WebSocket 连接已关闭")
}

func receiveResponses(conn *websocket.Conn, stream types.AIHelper_AskQuestionClient, cancel context.CancelFunc) {
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("gRPC 流已关闭")
			} else {
				log.Printf("接收 gRPC 响应失败: %v", err)
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

func sendMessages(conn *websocket.Conn, stream types.AIHelper_AskQuestionClient, ctx context.Context, timeout *time.Timer) {
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

			req := &types.AskQuestionRequest{
				Question:  string(message),
				SessionId: sessionID,
			}

			if err := stream.Send(req); err != nil {
				log.Printf("发送 gRPC 请求失败: %v", err)
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
