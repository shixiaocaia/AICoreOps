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

// alertRule
func (s *AicoreopsPrometheusServer) GetAlertRuleList(ctx context.Context, req *types.GetAlertRuleListRequest) (*types.GetAlertRuleListResponse, error) {
	l := logic.NewAlertRuleLogic(ctx, s.svcCtx)
	return l.GetAlertRuleList(ctx)
}

func (s *AicoreopsPrometheusServer) CreateAlertRule(ctx context.Context, req *types.CreateAlertRuleRequest) (*types.CreateAlertRuleResponse, error) {
	l := logic.NewAlertRuleLogic(ctx, s.svcCtx)
	return l.CreateAlertRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) UpdateAlertRule(ctx context.Context, req *types.UpdateAlertRuleRequest) (*types.UpdateAlertRuleResponse, error) {
	l := logic.NewAlertRuleLogic(ctx, s.svcCtx)
	return l.UpdateAlertRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) DeleteAlertRule(ctx context.Context, req *types.DeleteAlertRuleRequest) (*types.DeleteAlertRuleResponse, error) {
	l := logic.NewAlertRuleLogic(ctx, s.svcCtx)
	return l.DeleteAlertRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) EnableSwitchAlertRule(ctx context.Context, req *types.EnableSwitchAlertRuleRequest) (*types.EnableSwitchAlertRuleResponse, error) {
	l := logic.NewAlertRuleLogic(ctx, s.svcCtx)
	return l.EnableSwitchAlertRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) BatchEnableSwitchAlertRule(ctx context.Context, req *types.BatchEnableSwitchAlertRuleRequest) (*types.BatchEnableSwitchAlertRuleResponse, error) {
	l := logic.NewAlertRuleLogic(ctx, s.svcCtx)
	return l.BatchEnableSwitchAlertRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) BatchDeleteAlertRule(ctx context.Context, req *types.BatchDeleteAlertRuleRequest) (*types.BatchDeleteAlertRuleResponse, error) {
	l := logic.NewAlertRuleLogic(ctx, s.svcCtx)
	return l.BatchDeleteAlertRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) CheckPromqlExpr(ctx context.Context, req *types.CheckPromqlExprRequest) (*types.CheckPromqlExprResponse, error) {
	l := logic.NewAlertRuleLogic(ctx, s.svcCtx)
	return l.CheckPromqlExpr(ctx, req)
}

// record rule
func (s *AicoreopsPrometheusServer) GetRecordRuleList(ctx context.Context, req *types.GetRecordRuleListRequest) (*types.GetRecordRuleListResponse, error) {
	l := logic.NewRecordRuleLogic(ctx, s.svcCtx)
	return l.GetRecordRuleList(ctx)
}

func (s *AicoreopsPrometheusServer) CreateRecordRule(ctx context.Context, req *types.CreateRecordRuleRequest) (*types.CreateRecordRuleResponse, error) {
	l := logic.NewRecordRuleLogic(ctx, s.svcCtx)
	return l.CreateRecordRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) UpdateRecordRule(ctx context.Context, req *types.UpdateRecordRuleRequest) (*types.UpdateRecordRuleResponse, error) {
	l := logic.NewRecordRuleLogic(ctx, s.svcCtx)
	return l.UpdateRecordRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) DeleteRecordRule(ctx context.Context, req *types.DeleteRecordRuleRequest) (*types.DeleteRecordRuleResponse, error) {
	l := logic.NewRecordRuleLogic(ctx, s.svcCtx)
	return l.DeleteRecordRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) BatchDeleteRecordRule(ctx context.Context, req *types.BatchDeleteRecordRuleRequest) (*types.BatchDeleteRecordRuleResponse, error) {
	l := logic.NewRecordRuleLogic(ctx, s.svcCtx)
	return l.BatchDeleteRecordRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) EnableSwitchRecordRule(ctx context.Context, req *types.EnableSwitchRecordRuleRequest) (*types.EnableSwitchRecordRuleResponse, error) {
	l := logic.NewRecordRuleLogic(ctx, s.svcCtx)
	return l.EnableSwitchRecordRule(ctx, req)
}

func (s *AicoreopsPrometheusServer) BatchEnableSwitchRecordRule(ctx context.Context, req *types.BatchEnableSwitchRecordRuleRequest) (*types.BatchEnableSwitchRecordRuleResponse, error) {
	l := logic.NewRecordRuleLogic(ctx, s.svcCtx)
	return l.BatchEnableSwitchRecordRule(ctx, req)
}
