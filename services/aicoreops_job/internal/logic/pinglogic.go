package logic

import (
	"context"

	"aicoreops_job/aicoreops_job"
	"aicoreops_job/internal/svc"

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

func (l *PingLogic) Ping(in *aicoreops_job.Request) (*aicoreops_job.Response, error) {
	// todo: add your logic here and delete this line

	return &aicoreops_job.Response{}, nil
}
