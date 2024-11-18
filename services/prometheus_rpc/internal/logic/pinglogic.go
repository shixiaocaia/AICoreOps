package logic

import (
	"context"

	"prometheus_rpc/internal/svc"
	"prometheus_rpc/prometheus_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *prometheus_rpc.Request) (*prometheus_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &prometheus_rpc.Response{}, nil
}
