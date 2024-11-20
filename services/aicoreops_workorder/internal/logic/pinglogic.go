package logic

import (
	"context"

	"aicoreops_workorder/aicoreops_workorder"
	"aicoreops_workorder/internal/svc"

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

func (l *PingLogic) Ping(in *aicoreops_workorder.Request) (*aicoreops_workorder.Response, error) {
	// todo: add your logic here and delete this line

	return &aicoreops_workorder.Response{}, nil
}
