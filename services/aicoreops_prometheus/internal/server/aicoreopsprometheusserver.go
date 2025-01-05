package server

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/logic"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
)

type AicoreopsPrometheusServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedPrometheusRpcServer
}

func NewAicoreopsPrometheusServer(svcCtx *svc.ServiceContext) *AicoreopsPrometheusServer {
	return &AicoreopsPrometheusServer{
		svcCtx: svcCtx,
	}
}

func (s *AicoreopsPrometheusServer) CreateMonitorAlertManagerPool(ctx context.Context, req *types.CreateMonitorAlertManagerPoolRequest) (*types.CreateMonitorAlertManagerPoolResponse, error) {
	l := logic.NewAlertmanagerLogic(ctx, s.svcCtx)
	return l.CreateMonitorAlertManagerPool(ctx, req)
}

func (s *AicoreopsPrometheusServer) GetMonitorAlertmanagerPoolList(ctx context.Context, req *types.GetAlertmanagerPoolListRequest) (*types.GetAlertmanagerPoolListResponse, error) {
	l := logic.NewAlertmanagerLogic(ctx, s.svcCtx)
	return l.GetMonitorAlertmanagerPoolList(ctx)
}

func (s *AicoreopsPrometheusServer) UpdateMonitorAlertManagerPool(ctx context.Context, req *types.UpdateMonitorAlertManagerPoolRequest) (*types.UpdateMonitorAlertManagerPoolResponse, error) {
	l := logic.NewAlertmanagerLogic(ctx, s.svcCtx)
	return l.UpdateMonitorAlertManagerPool(ctx, req)
}

func (s *AicoreopsPrometheusServer) DeleteMonitorAlertManagerPool(ctx context.Context, req *types.DeleteMonitorAlertManagerPoolRequest) (*types.DeleteMonitorAlertManagerPoolResponse, error) {
	l := logic.NewAlertmanagerLogic(ctx, s.svcCtx)
	return l.DeleteMonitorAlertManagerPool(ctx, req)
}
