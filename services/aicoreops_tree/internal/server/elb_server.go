package server

import (
	"aicoreops_tree/internal/svc"
	"aicoreops_tree/types"
)

type ElbServiceServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedElbServiceServer
}

func NewElbServiceServer(svcCtx *svc.ServiceContext) *ElbServiceServer {
	return &ElbServiceServer{
		svcCtx: svcCtx,
	}
}
