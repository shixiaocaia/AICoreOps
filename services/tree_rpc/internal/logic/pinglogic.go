package logic

import (
	"context"

	"tree_rpc/internal/svc"
	"tree_rpc/tree_rpc"

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

func (l *PingLogic) Ping(in *tree_rpc.Request) (*tree_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &tree_rpc.Response{}, nil
}
