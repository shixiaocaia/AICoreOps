package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
)

type ScrapePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	domain *domain.MonitorScrapePoolDomain
}

func NewScrapePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScrapePoolLogic {
	return &ScrapePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		domain: domain.NewMonitorScrapePoolDomain(svcCtx),
	}
}

func (s *ScrapePoolLogic) GetMonitorScrapePoolList(req *types.GetMonitorScrapePoolListRequest) (*types.GetMonitorScrapePoolListResponse, error) {
	pools, err := s.domain.GetMonitorScrapePoolList(s.ctx)
	if err != nil {
		s.Logger.Errorf("获取采集池列表失败: %v", err)
		return nil, err
	}

	return &types.GetMonitorScrapePoolListResponse{
		Code:    200,
		Message: "获取采集池列表成功",
		Data:    s.domain.BuildScrapePoolRespModel(pools),
	}, nil
}

func (s *ScrapePoolLogic) CreateMonitorScrapePool(req *types.CreateMonitorScrapePoolRequest) (*types.CreateMonitorScrapePoolResponse, error) {
	// 创建采集池
	pool := domain.BuildMonitorScrapePoolModel(req.Pool)
	err := s.domain.CreateMonitorScrapePool(s.ctx, pool)
	if err != nil {
		s.Logger.Errorf("创建采集池失败: %v", err)
		return nil, err
	}

	// 更新缓存

	s.Logger.Infof("创建采集池成功: %v", pool)

	return &types.CreateMonitorScrapePoolResponse{
		Code:    200,
		Message: "创建采集池成功",
	}, nil
}

func (s *ScrapePoolLogic) UpdateMonitorScrapePool(req *types.UpdateMonitorScrapePoolRequest) (*types.UpdateMonitorScrapePoolResponse, error) {
	pool := domain.BuildMonitorScrapePoolModel(req.Pool)
	err := s.domain.UpdateMonitorScrapePool(s.ctx, pool)
	if err != nil {
		s.Logger.Errorf("更新采集池失败: %v", err)
		return nil, err
	}

	// 更新缓存

	s.Logger.Infof("更新采集池成功: %v", pool)

	return &types.UpdateMonitorScrapePoolResponse{
		Code:    200,
		Message: "更新采集池成功",
	}, nil
}

func (s *ScrapePoolLogic) DeleteMonitorScrapePool(req *types.DeleteMonitorScrapePoolRequest) (*types.DeleteMonitorScrapePoolResponse, error) {
	err := s.domain.DeleteMonitorScrapePool(s.ctx, req.Id)
	if err != nil {
		s.Logger.Errorf("删除采集池失败: %v", err)
		return nil, err
	}

	// 更新缓存

	s.Logger.Infof("删除采集池成功: %v", req.Id)

	return &types.DeleteMonitorScrapePoolResponse{
		Code:    200,
		Message: "删除采集池成功",
	}, nil
}
