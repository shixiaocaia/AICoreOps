package server

import (
	"aicoreops_tree/internal/svc"
	"aicoreops_tree/types"
)

type RdsServiceServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedRdsServiceServer
}

func NewRdsServiceServer(svcCtx *svc.ServiceContext) *RdsServiceServer {
	return &RdsServiceServer{
		svcCtx: svcCtx,
	}
}
