package main

import (
	"flag"
	"fmt"

	"aicoreops_api/internal/config"
	"aicoreops_api/internal/handler"
	"aicoreops_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	logx.MustSetup(logx.LogConf{Stat: false, Encoding: "plain"})

	// 加载配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 创建服务器实例
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	// 创建服务上下文
	ctx := svc.NewServiceContext(c)

	// 初始化路由
	routers := handler.NewRouters(server)
	handler.RegisterHandlers(routers, ctx)

	// 启动服务器
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
