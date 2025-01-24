package logic

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type AlertRuleLogic struct {
	ctx    context.Context
	domain *domain.AlertRuleDomain
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlertRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlertRuleLogic {
	return &AlertRuleLogic{
		ctx:    ctx,
		domain: domain.NewAlertRuleDomain(svcCtx),
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (a *AlertRuleLogic) GetAlertRuleList(ctx context.Context) (*types.GetMonitorAlertRuleListResponse, error) {
	rules, err := a.domain.GetAlertRuleList(ctx, nil)
	if err != nil {
		a.Logger.Errorf("获取告警规则列表失败: %v", err)
		return nil, err
	}

	return &types.GetMonitorAlertRuleListResponse{
		Code:    0,
		Message: "获取告警规则列表成功",
		Data:    a.domain.BuildAlertRuleRespModel(rules),
	}, nil
}

func (a *AlertRuleLogic) CreateAlertRule(ctx context.Context, req *types.CreateMonitorAlertRuleRequest) (*types.CreateMonitorAlertRuleResponse, error) {
	// 创建告警规则
	rule := a.domain.BuildAlertRuleModel(req.Rule)
	if err := a.domain.CreateAlertRule(ctx, rule); err != nil {
		a.Logger.Errorf("创建告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.CreateMonitorAlertRuleResponse{
		Code:    0,
		Message: "创建告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) UpdateAlertRule(ctx context.Context, req *types.UpdateMonitorAlertRuleRequest) (*types.UpdateMonitorAlertRuleResponse, error) {
	// 更新告警规则
	rule := a.domain.BuildAlertRuleModel(req.Rule)
	if err := a.domain.UpdateAlertRule(ctx, rule); err != nil {
		a.Logger.Errorf("更新告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.UpdateMonitorAlertRuleResponse{
		Code:    0,
		Message: "更新告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) DeleteAlertRule(ctx context.Context, req *types.DeleteMonitorAlertRuleRequest) (*types.DeleteMonitorAlertRuleResponse, error) {
	// 删除告警规则
	if err := a.domain.DeleteAlertRule(ctx, req.Id); err != nil {
		a.Logger.Errorf("删除告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.DeleteMonitorAlertRuleResponse{
		Code:    0,
		Message: "删除告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) BatchDeleteAlertRule(ctx context.Context, req *types.BatchDeleteMonitorAlertRuleRequest) (*types.BatchDeleteMonitorAlertRuleResponse, error) {
	// 批量删除告警规则
	if err := a.domain.BatchDeleteAlertRule(ctx, req.Ids); err != nil {
		a.Logger.Errorf("批量删除告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.BatchDeleteMonitorAlertRuleResponse{
		Code:    0,
		Message: "批量删除告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) EnableSwitchAlertRule(ctx context.Context, req *types.EnableSwitchMonitorAlertRuleRequest) (*types.EnableSwitchMonitorAlertRuleResponse, error) {
	// 启用或禁用告警规则
	if err := a.domain.EnableSwitchAlertRule(ctx, req.Id); err != nil {
		a.Logger.Errorf("启用或禁用告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.EnableSwitchMonitorAlertRuleResponse{
		Code:    0,
		Message: "启用或禁用告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) BatchEnableSwitchAlertRule(ctx context.Context, req *types.BatchEnableSwitchMonitorAlertRuleRequest) (*types.BatchEnableSwitchMonitorAlertRuleResponse, error) {
	// 批量启用或禁用告警规则
	if err := a.domain.BatchEnableSwitchAlertRule(ctx, req.Ids); err != nil {
		a.Logger.Errorf("批量启用或禁用告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.BatchEnableSwitchMonitorAlertRuleResponse{
		Code:    0,
		Message: "批量启用或禁用告警规则成功",
	}, nil
}
