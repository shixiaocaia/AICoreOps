package main

import (
	"flag"
	"fmt"

	"aicoreops_tree/internal/config"
	"aicoreops_tree/internal/server"
	"aicoreops_tree/internal/svc"
	"aicoreops_tree/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		types.RegisterResourceTreeServiceServer(grpcServer, server.NewResourceTreeServiceServer(ctx))
		types.RegisterEcsServiceServer(grpcServer, server.NewEcsServiceServer(ctx))
		types.RegisterRdsServiceServer(grpcServer, server.NewRdsServiceServer(ctx))
		types.RegisterElbServiceServer(grpcServer, server.NewElbServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
