package server

import (
	"aicoreops_tree/internal/svc"
	"aicoreops_tree/types"
)

type EcsServiceServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedEcsServiceServer
}

func NewEcsServiceServer(svcCtx *svc.ServiceContext) *EcsServiceServer {
	return &EcsServiceServer{
		svcCtx: svcCtx,
	}
}
