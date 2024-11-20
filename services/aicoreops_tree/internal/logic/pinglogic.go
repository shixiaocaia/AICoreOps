package logic

import (
	"context"

	"aicoreops_tree/aicoreops_tree"
	"aicoreops_tree/internal/svc"

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

func (l *PingLogic) Ping(in *aicoreops_tree.Request) (*aicoreops_tree.Response, error) {
	// todo: add your logic here and delete this line

	return &aicoreops_tree.Response{}, nil
}
