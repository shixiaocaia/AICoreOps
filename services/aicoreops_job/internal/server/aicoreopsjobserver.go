// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_job.proto

package server

import (
	"context"

	"aicoreops_job/aicoreops_job"
	"aicoreops_job/internal/logic"
	"aicoreops_job/internal/svc"
)

type AicoreopsJobServer struct {
	svcCtx *svc.ServiceContext
	aicoreops_job.UnimplementedAicoreopsJobServer
}

func NewAicoreopsJobServer(svcCtx *svc.ServiceContext) *AicoreopsJobServer {
	return &AicoreopsJobServer{
		svcCtx: svcCtx,
	}
}

func (s *AicoreopsJobServer) Ping(ctx context.Context, in *aicoreops_job.Request) (*aicoreops_job.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
