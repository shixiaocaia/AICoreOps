package logic

import (
	"context"

	"aicoreops_prometheus/aicoreops_prometheus"
	"aicoreops_prometheus/internal/svc"

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

func (l *PingLogic) Ping(in *aicoreops_prometheus.Request) (*aicoreops_prometheus.Response, error) {
	// todo: add your logic here and delete this line

	return &aicoreops_prometheus.Response{}, nil
}
