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

func (a *AlertRuleLogic) GetAlertRuleList(ctx context.Context) (*types.GetAlertRuleListResponse, error) {
	rules, err := a.domain.GetAlertRuleList(ctx, nil)
	if err != nil {
		a.Logger.Errorf("获取告警规则列表失败: %v", err)
		return nil, err
	}

	return &types.GetAlertRuleListResponse{
		Code:    0,
		Message: "获取告警规则列表成功",
		Data:    a.domain.BuildAlertRuleRespModel(rules),
	}, nil
}

func (a *AlertRuleLogic) CreateAlertRule(ctx context.Context, req *types.CreateAlertRuleRequest) (*types.CreateAlertRuleResponse, error) {
	// 创建告警规则
	rule := a.domain.BuildAlertRuleModel(req.Rule)
	if err := a.domain.CreateAlertRule(ctx, rule); err != nil {
		a.Logger.Errorf("创建告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.CreateAlertRuleResponse{
		Code:    0,
		Message: "创建告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) UpdateAlertRule(ctx context.Context, req *types.UpdateAlertRuleRequest) (*types.UpdateAlertRuleResponse, error) {
	// 更新告警规则
	rule := a.domain.BuildAlertRuleModel(req.Rule)
	if err := a.domain.UpdateAlertRule(ctx, rule); err != nil {
		a.Logger.Errorf("更新告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.UpdateAlertRuleResponse{
		Code:    0,
		Message: "更新告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) DeleteAlertRule(ctx context.Context, req *types.DeleteAlertRuleRequest) (*types.DeleteAlertRuleResponse, error) {
	// 删除告警规则
	if err := a.domain.DeleteAlertRule(ctx, req.Id); err != nil {
		a.Logger.Errorf("删除告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.DeleteAlertRuleResponse{
		Code:    0,
		Message: "删除告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) BatchDeleteAlertRule(ctx context.Context, req *types.BatchDeleteAlertRuleRequest) (*types.BatchDeleteAlertRuleResponse, error) {
	// 批量删除告警规则
	if err := a.domain.BatchDeleteAlertRule(ctx, req.Ids); err != nil {
		a.Logger.Errorf("批量删除告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.BatchDeleteAlertRuleResponse{
		Code:    0,
		Message: "批量删除告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) EnableSwitchAlertRule(ctx context.Context, req *types.EnableSwitchAlertRuleRequest) (*types.EnableSwitchAlertRuleResponse, error) {
	// 启用或禁用告警规则
	if err := a.domain.EnableSwitchAlertRule(ctx, req.Id); err != nil {
		a.Logger.Errorf("启用或禁用告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.EnableSwitchAlertRuleResponse{
		Code:    0,
		Message: "启用或禁用告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) BatchEnableSwitchAlertRule(ctx context.Context, req *types.BatchEnableSwitchAlertRuleRequest) (*types.BatchEnableSwitchAlertRuleResponse, error) {
	// 批量启用或禁用告警规则
	if err := a.domain.BatchEnableSwitchAlertRule(ctx, req.Ids); err != nil {
		a.Logger.Errorf("批量启用或禁用告警规则失败: %v", err)
		return nil, err
	}

	// TODO 更新缓存

	return &types.BatchEnableSwitchAlertRuleResponse{
		Code:    0,
		Message: "批量启用或禁用告警规则成功",
	}, nil
}

func (a *AlertRuleLogic) CheckPromqlExpr(ctx context.Context, req *types.CheckPromqlExprRequest) (*types.CheckPromqlExprResponse, error) {
	if err := a.domain.CheckPromqlExpr(ctx, req.Expr); err != nil {
		a.Logger.Errorf("检查PromQL表达式失败: %v", err)
		return nil, err
	}

	return &types.CheckPromqlExprResponse{
		Code:    0,
		Message: "PromQL表达式正确",
	}, nil
}
