package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"aicoreops_ai/types"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("HTTP 服务启动，监听端口 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP 服务启动失败: %v", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("接收到ws请求")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("无法升级为 WebSocket: %v", err)
		return
	}
	defer conn.Close()

	grpcConn, err := grpc.Dial("localhost:8083", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("无法连接到 gRPC 服务器: %v", err)
		return
	}
	defer grpcConn.Close()

	client := types.NewAIHelperClient(grpcConn)
	stream, err := client.AskQuestion(context.Background())
	if err != nil {
		log.Printf("无法创建 gRPC 流: %v", err)
		return
	}

	timeoutDuration := 60 * time.Second
	timeout := time.NewTimer(timeoutDuration)
	defer timeout.Stop()

	go receiveResponses(conn, stream)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取 WebSocket 消息失败: %v", err)
			break
		}

		log.Println("接收到ws消息", string(message))

		resetTimeout(timeout, timeoutDuration)

		req := &types.AskQuestionRequest{
			Question:  string(message),
			SessionId: "1",
		}

		if err := stream.Send(req); err != nil {
			log.Printf("发送 gRPC 请求失败: %v", err)
			break
		}
	}

	// 这里会阻塞等待吗
	closeStream(stream, conn, timeout)
}

func receiveResponses(conn *websocket.Conn, stream types.AIHelper_AskQuestionClient) {
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("gRPC 流已关闭")
				return
			}
			log.Printf("接收 gRPC 响应失败: %v", err)
			return
		}
		message := fmt.Sprintf("收到回答: %s", resp.GetData().GetAnswer())
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("发送 WebSocket 消息失败: %v", err)
			return
		}
	}
}

func resetTimeout(timeout *time.Timer, duration time.Duration) {
	if !timeout.Stop() {
		<-timeout.C
	}
	timeout.Reset(duration)
}

func closeStream(stream types.AIHelper_AskQuestionClient, conn *websocket.Conn, timeout *time.Timer) {
	if err := stream.CloseSend(); err != nil {
		log.Printf("关闭 gRPC 发送流失败: %v", err)
	}

	select {
	case <-timeout.C:
		log.Println("超时未收到消息，断开连接")
		conn.Close()
	}
}
