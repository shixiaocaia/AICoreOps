package logic

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_cicd/aicoreops_cicd"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_cicd/internal/svc"

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

func (l *PingLogic) Ping(in *aicoreops_cicd.Request) (*aicoreops_cicd.Response, error) {
	// todo: add your logic here and delete this line

	return &aicoreops_cicd.Response{}, nil
}
