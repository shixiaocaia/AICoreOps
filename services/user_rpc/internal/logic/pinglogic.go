package logic

import (
	"context"

	"user_rpc/internal/svc"
	"user_rpc/user_rpc"

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

func (l *PingLogic) Ping(in *user_rpc.Request) (*user_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &user_rpc.Response{}, nil
}
