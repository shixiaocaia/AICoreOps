package logic

import (
	"context"

	"aicoreops_api/internal/svc"
	"aicoreops_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Aicoreops_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAicoreops_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Aicoreops_apiLogic {
	return &Aicoreops_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Aicoreops_apiLogic) Aicoreops_api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
