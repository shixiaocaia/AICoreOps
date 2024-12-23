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
 * File: user.go
 */

package main

import (
	"flag"
	"fmt"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/config"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/handler"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"

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
