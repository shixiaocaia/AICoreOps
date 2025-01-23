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

// ScrapePool
func (s *AicoreopsPrometheusServer) GetMonitorScrapePoolList(ctx context.Context, req *types.GetMonitorScrapePoolListRequest) (*types.GetMonitorScrapePoolListResponse, error) {
	l := logic.NewScrapePoolLogic(ctx, s.svcCtx)
	return l.GetMonitorScrapePoolList(req)
}

func (s *AicoreopsPrometheusServer) CreateMonitorScrapePool(ctx context.Context, req *types.CreateMonitorScrapePoolRequest) (*types.CreateMonitorScrapePoolResponse, error) {
	l := logic.NewScrapePoolLogic(ctx, s.svcCtx)
	return l.CreateMonitorScrapePool(req)
}

func (s *AicoreopsPrometheusServer) UpdateMonitorScrapePool(ctx context.Context, req *types.UpdateMonitorScrapePoolRequest) (*types.UpdateMonitorScrapePoolResponse, error) {
	l := logic.NewScrapePoolLogic(ctx, s.svcCtx)
	return l.UpdateMonitorScrapePool(req)
}

func (s *AicoreopsPrometheusServer) DeleteMonitorScrapePool(ctx context.Context, req *types.DeleteMonitorScrapePoolRequest) (*types.DeleteMonitorScrapePoolResponse, error) {
	l := logic.NewScrapePoolLogic(ctx, s.svcCtx)
	return l.DeleteMonitorScrapePool(req)
}

// Alertmanager
func (s *AicoreopsPrometheusServer) CreateMonitorAlertManagerPool(ctx context.Context, req *types.CreateMonitorAlertManagerPoolRequest) (*types.CreateMonitorAlertManagerPoolResponse, error) {
	l := logic.NewAlertManagerPoolLogic(ctx, s.svcCtx)
	return l.CreateMonitorAlertManagerPool(ctx, req)
}

func (s *AicoreopsPrometheusServer) GetMonitorAlertManagerPoolList(ctx context.Context, req *types.GetAlertManagerPoolListRequest) (*types.GetAlertManagerPoolListResponse, error) {
	l := logic.NewAlertManagerPoolLogic(ctx, s.svcCtx)
	return l.GetMonitorAlertManagerPoolList(ctx)
}

func (s *AicoreopsPrometheusServer) UpdateMonitorAlertManagerPool(ctx context.Context, req *types.UpdateMonitorAlertManagerPoolRequest) (*types.UpdateMonitorAlertManagerPoolResponse, error) {
	l := logic.NewAlertManagerPoolLogic(ctx, s.svcCtx)
	return l.UpdateMonitorAlertManagerPool(ctx, req)
}

func (s *AicoreopsPrometheusServer) DeleteMonitorAlertManagerPool(ctx context.Context, req *types.DeleteMonitorAlertManagerPoolRequest) (*types.DeleteMonitorAlertManagerPoolResponse, error) {
	l := logic.NewAlertManagerPoolLogic(ctx, s.svcCtx)
	return l.DeleteMonitorAlertManagerPool(ctx, req)
}

// ScrapeJob
func (s *AicoreopsPrometheusServer) GetMonitorScrapeJobList(ctx context.Context, req *types.GetMonitorScrapeJobListRequest) (*types.GetMonitorScrapeJobListResponse, error) {
	l := logic.NewScrapeJobLogic(ctx, s.svcCtx)
	return l.GetMonitorScrapeJobList(ctx)
}

func (s *AicoreopsPrometheusServer) CreateMonitorScrapeJob(ctx context.Context, req *types.CreateMonitorScrapeJobRequest) (*types.CreateMonitorScrapeJobResponse, error) {
	l := logic.NewScrapeJobLogic(ctx, s.svcCtx)
	return l.CreateMonitorScrapeJob(ctx, req)
}

func (s *AicoreopsPrometheusServer) UpdateMonitorScrapeJob(ctx context.Context, req *types.UpdateMonitorScrapeJobRequest) (*types.UpdateMonitorScrapeJobResponse, error) {
	l := logic.NewScrapeJobLogic(ctx, s.svcCtx)
	return l.UpdateMonitorScrapeJob(ctx, req)
}

func (s *AicoreopsPrometheusServer) DeleteMonitorScrapeJob(ctx context.Context, req *types.DeleteMonitorScrapeJobRequest) (*types.DeleteMonitorScrapeJobResponse, error) {
	l := logic.NewScrapeJobLogic(ctx, s.svcCtx)
	return l.DeleteMonitorScrapeJob(ctx, req)
}
