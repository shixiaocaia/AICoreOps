package logic

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type RecordRuleLogic struct {
	ctx    context.Context
	domain *domain.RecordRuleDomain
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecordRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordRuleLogic {
	return &RecordRuleLogic{
		ctx:    ctx,
		domain: domain.NewRecordRuleDomain(svcCtx),
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (r *RecordRuleLogic) GetRecordRuleList(ctx context.Context) (*types.GetRecordRuleListResponse, error) {
	rules, err := r.domain.GetRecordRuleList(ctx)
	if err != nil {
		r.Logger.Errorf("获取预聚合规则列表失败: %v", err)
		return nil, err
	}

	return &types.GetRecordRuleListResponse{
		Code:    0,
		Message: "获取预聚合规则列表成功",
		Data:    r.domain.BuildRecordRuleRespModel(rules),
	}, nil
}

func (r *RecordRuleLogic) CreateRecordRule(ctx context.Context, req *types.CreateRecordRuleRequest) (*types.CreateRecordRuleResponse, error) {
	// 创建记录规则
	rule := r.domain.BuildRecordRuleModel(req.Rule)
	if err := r.domain.CreateRecordRule(ctx, rule); err != nil {
		r.Logger.Errorf("创建预聚合规则失败: %v", err)
		return nil, err
	}

	return &types.CreateRecordRuleResponse{
		Code:    0,
		Message: "创建预聚合规则成功",
	}, nil
}

func (r *RecordRuleLogic) UpdateRecordRule(ctx context.Context, req *types.UpdateRecordRuleRequest) (*types.UpdateRecordRuleResponse, error) {
	// 更新记录规则
	rule := r.domain.BuildRecordRuleModel(req.Rule)
	if err := r.domain.UpdateRecordRule(ctx, rule); err != nil {
		r.Logger.Errorf("更新预聚合规则失败: %v", err)
		return nil, err
	}

	return &types.UpdateRecordRuleResponse{
		Code:    0,
		Message: "更新预聚合规则成功",
	}, nil
}

func (r *RecordRuleLogic) DeleteRecordRule(ctx context.Context, req *types.DeleteRecordRuleRequest) (*types.DeleteRecordRuleResponse, error) {
	// 删除记录规则
	if err := r.domain.DeleteRecordRule(ctx, req.Id); err != nil {
		r.Logger.Errorf("删除预聚合规则失败: %v", err)
		return nil, err
	}

	return &types.DeleteRecordRuleResponse{
		Code:    0,
		Message: "删除预聚合规则成功",
	}, nil
}

func (r *RecordRuleLogic) BatchDeleteRecordRule(ctx context.Context, req *types.BatchDeleteRecordRuleRequest) (*types.BatchDeleteRecordRuleResponse, error) {
	// 批量删除记录规则
	if err := r.domain.BatchDeleteRecordRule(ctx, req.Ids); err != nil {
		r.Logger.Errorf("批量删除预聚合规则失败: %v", err)
		return nil, err
	}

	return &types.BatchDeleteRecordRuleResponse{
		Code:    0,
		Message: "批量删除预聚合规则成功",
	}, nil
}

func (r *RecordRuleLogic) EnableSwitchRecordRule(ctx context.Context, req *types.EnableSwitchRecordRuleRequest) (*types.EnableSwitchRecordRuleResponse, error) {
	// 启用或禁用记录规则
	if err := r.domain.EnableSwitchRecordRule(ctx, req.Id); err != nil {
		r.Logger.Errorf("启用或禁用预聚合规则失败: %v", err)
		return nil, err
	}

	return &types.EnableSwitchRecordRuleResponse{
		Code:    0,
		Message: "启用或禁用预聚合规则成功",
	}, nil
}

func (r *RecordRuleLogic) BatchEnableSwitchRecordRule(ctx context.Context, req *types.BatchEnableSwitchRecordRuleRequest) (*types.BatchEnableSwitchRecordRuleResponse, error) {
	// 批量启用或禁用记录规则
	if err := r.domain.BatchEnableSwitchRecordRule(ctx, req.Ids); err != nil {
		r.Logger.Errorf("批量启用或禁用预聚合规则失败: %v", err)
		return nil, err
	}

	return &types.BatchEnableSwitchRecordRuleResponse{
		Code:    0,
		Message: "批量启用或禁用预聚合规则成功",
	}, nil
}
