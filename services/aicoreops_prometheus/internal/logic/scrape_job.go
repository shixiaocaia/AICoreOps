package logic

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScrapeJobLogic struct {
	ctx    context.Context
	domain *domain.ScrapeJobDomain
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScrapeJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScrapeJobLogic {
	return &ScrapeJobLogic{
		ctx:    ctx,
		domain: domain.NewScrapeJobDomain(svcCtx),
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (s *ScrapeJobLogic) GetMonitorScrapeJobList(ctx context.Context) (*types.GetMonitorScrapeJobListResponse, error) {
	jobs, err := s.domain.GetMonitorScrapeJobList(ctx)
	if err != nil {
		s.Logger.Errorf("获取 ScrapeJob 列表失败: %v", err)
		return nil, err
	}

	return &types.GetMonitorScrapeJobListResponse{
		Code:    0,
		Message: "获取 ScrapeJob 列表成功",
		Data:    s.domain.BuildScrapeJobRespModel(jobs),
	}, nil
}

func (s *ScrapeJobLogic) CreateMonitorScrapeJob(ctx context.Context, req *types.CreateMonitorScrapeJobRequest) (*types.CreateMonitorScrapeJobResponse, error) {
	job := s.domain.BuildMonitorScrapeJobModel(req.Job)
	err := s.domain.CreateMonitorScrapeJob(ctx, job)
	if err != nil {
		s.Logger.Errorf("创建 ScrapeJob 失败: %v", err)
		return nil, err
	}

	// TODO: 更新 Prometheus 配置

	s.Logger.Infof("创建 ScrapeJob 成功: %+v", job)

	return &types.CreateMonitorScrapeJobResponse{
		Code:    0,
		Message: "创建 ScrapeJob 成功",
	}, nil
}

func (s *ScrapeJobLogic) UpdateMonitorScrapeJob(ctx context.Context, req *types.UpdateMonitorScrapeJobRequest) (*types.UpdateMonitorScrapeJobResponse, error) {
	job := s.domain.BuildMonitorScrapeJobModel(req.Job)
	err := s.domain.UpdateMonitorScrapeJob(ctx, job)
	if err != nil {
		s.Logger.Errorf("更新 ScrapeJob 失败: %v", err)
		return nil, err
	}

	// TODO: 更新 Prometheus 配置

	s.Logger.Infof("更新 ScrapeJob 成功: %+v", job)

	return &types.UpdateMonitorScrapeJobResponse{
		Code:    0,
		Message: "更新 ScrapeJob 成功",
	}, nil
}

func (s *ScrapeJobLogic) DeleteMonitorScrapeJob(ctx context.Context, req *types.DeleteMonitorScrapeJobRequest) (*types.DeleteMonitorScrapeJobResponse, error) {
	err := s.domain.DeleteMonitorScrapeJob(ctx, req.Id)
	if err != nil {
		s.Logger.Errorf("删除 ScrapeJob 失败: %v", err)
		return nil, err
	}

	// TODO: 更新 Prometheus 配置

	s.Logger.Infof("删除 ScrapeJob 成功: %d", req.Id)

	return &types.DeleteMonitorScrapeJobResponse{
		Code:    0,
		Message: "删除 ScrapeJob 成功",
	}, nil
}
