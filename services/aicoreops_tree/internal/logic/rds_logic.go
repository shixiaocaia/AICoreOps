package logic

import "aicoreops_tree/internal/svc"

type RdsLogic struct {
	svcCtx *svc.ServiceContext
}

func NewRdsLogic(svcCtx *svc.ServiceContext) *RdsLogic {
	return &RdsLogic{
		svcCtx: svcCtx,
	}
}
