package logic

import (
	"context"

	"k8s_rpc/internal/svc"
	"k8s_rpc/k8s_rpc"

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

func (l *PingLogic) Ping(in *k8s_rpc.Request) (*k8s_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &k8s_rpc.Response{}, nil
}
