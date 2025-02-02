package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// SDResponse 符合 Prometheus HTTP SD 格式的响应
type SDResponse []struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

func getTreeNodeBindIps(w http.ResponseWriter, r *http.Request) {
	// 获取查询参数
	port := r.URL.Query().Get("port")
	leafNodeIds := r.URL.Query()["leafNodeIds"]

	log.Printf("Received request with port: %s, leafNodeIds: %v", port, leafNodeIds)

	// 模拟的服务发现响应
	response := SDResponse{
		{
			Targets: []string{
				"target-service-1:9100",
				"target-service-2:9100",
			},
			Labels: map[string]string{
				"env": "prod",
				"job": "TestJob1",
			},
		},
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 输出响应内容到日志
	responseBytes, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("Responding with: %s", string(responseBytes))

	// 发送响应
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// 添加一个简单的健康检查端点
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/api/not_auth/getTreeNodeBindIps", getTreeNodeBindIps)

	log.Println("Starting server on :8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
