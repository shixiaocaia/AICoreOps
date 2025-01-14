# Run
docker-compose up -d --build # --build 强制make image
# 简要流程
1. Prometheus 调用 SD 服务 (http://sd-service:8888/api/not_auth/getTreeNodeBindIps)
2. SD 服务返回目标列表
3. Prometheus 对每个目标应用 relabel 规则
4. Prometheus 使用最终的地址去抓取 metrics
