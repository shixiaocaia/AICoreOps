package logic

import (
	"context"

	"role_rpc/internal/svc"
	"role_rpc/role_rpc"

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

func (l *PingLogic) Ping(in *role_rpc.Request) (*role_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &role_rpc.Response{}, nil
}
