package logic

import "aicoreops_tree/internal/svc"

type ElbLogic struct {
	svcCtx *svc.ServiceContext
}

func NewElbLogic(svcCtx *svc.ServiceContext) *ElbLogic {
	return &ElbLogic{
		svcCtx: svcCtx,
	}
}
