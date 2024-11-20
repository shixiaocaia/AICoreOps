package logic

import (
	"context"

	"aicoreops_role/aicoreops_role"
	"aicoreops_role/internal/svc"

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

func (l *PingLogic) Ping(in *aicoreops_role.Request) (*aicoreops_role.Response, error) {
	// todo: add your logic here and delete this line

	return &aicoreops_role.Response{}, nil
}
