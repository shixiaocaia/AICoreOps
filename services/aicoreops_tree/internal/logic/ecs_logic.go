package logic

import "aicoreops_tree/internal/svc"

type EcsLogic struct {
	svcCtx *svc.ServiceContext
}

func NewEcsLogic(svcCtx *svc.ServiceContext) *EcsLogic {
	return &EcsLogic{
		svcCtx: svcCtx,
	}
}
