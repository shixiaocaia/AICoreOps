// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: prometheus_rpc.proto

package server

import (
	"context"

	"prometheus_rpc/internal/logic"
	"prometheus_rpc/internal/svc"
	"prometheus_rpc/prometheus_rpc"
)

type PrometheusRpcServer struct {
	svcCtx *svc.ServiceContext
	prometheus_rpc.UnimplementedPrometheusRpcServer
}

func NewPrometheusRpcServer(svcCtx *svc.ServiceContext) *PrometheusRpcServer {
	return &PrometheusRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *PrometheusRpcServer) Ping(ctx context.Context, in *prometheus_rpc.Request) (*prometheus_rpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
