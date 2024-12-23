## 环境准备
```shell
# 安装向量数据库
docker run -d -p 6333:6333 qdrant/qdrant:latest

curl -X PUT <http://localhost:6333/collections/aicoreops> \\
-H "Content-Type: application/json" \\
-d '{
  "vectors": {
    "size": 3584,
    "distance": "Cosine"
  }
}'

# 安装ollama并拉取模型
<https://ollama.com/download/mac>

ollama pull qwen:7b
ollama pull nomic-embed-text:latest
ollama run qwen2.5:latest


# 安装mysql, 执行 model/sql 建库建表
docker run --name aicoreops -e MYSQL_ROOT_PASSWORD=root -d -p 3306:3306 -v /my/own/datadir:/var/lib/mysql mysql:latest

# etcd

```

## 测试对话助手

启动服务
```
go run main.go
```

启动api客户端
```
go run client/apiclient.go
```

调用api, 建立websocket连接，发送消息
```
http://localhost:8080/ws
```

TODO
- 管理 memoryBuf，及时释放，避免 OOM
    - 现有memoryBuf 取决于 model，没有实现 token 有效管理
- 管理 streaming RPC 连接