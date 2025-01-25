package domain

import (
	"context"
	"errors"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/pkg"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/repo"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
)

type RecordRuleDomain struct {
	repo repo.RecordRuleRepo
}

func NewRecordRuleDomain(svcCtx *svc.ServiceContext) *RecordRuleDomain {
	return &RecordRuleDomain{
		repo: dao.NewMonitorRecordRuleDAO(svcCtx.DB),
	}
}

func (r *RecordRuleDomain) GetRecordRuleList(ctx context.Context) ([]*model.MonitorRecordRule, error) {
	rules, err := r.repo.GetMonitorRecordRuleList(ctx)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (r *RecordRuleDomain) GetRecordRuleByPoolId(ctx context.Context, poolId int64) ([]*model.MonitorRecordRule, error) {
	return r.repo.GetMonitorRecordRuleByPoolId(ctx, poolId)
}

func (r *RecordRuleDomain) CreateRecordRule(ctx context.Context, rule *model.MonitorRecordRule) error {
	// 检查记录规则名称是否存在
	exists, err := r.repo.CheckMonitorRecordRuleNameExists(ctx, rule)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("记录规则名称已存在")
	}

	// 检查PromQL表达式
	correct, err := pkg.PromqlExprCheck(rule.Expr)
	if err != nil {
		return err
	}
	if !correct {
		return errors.New("PromQL表达式不正确")
	}

	return r.repo.CreateMonitorRecordRule(ctx, rule)
}

func (r *RecordRuleDomain) UpdateRecordRule(ctx context.Context, rule *model.MonitorRecordRule) error {
	// 检查记录规则是否存在
	exists, err := r.repo.CheckMonitorRecordRuleExists(ctx, rule)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("记录规则不存在")
	}

	// 检查PromQL表达式
	correct, err := pkg.PromqlExprCheck(rule.Expr)
	if err != nil {
		return err
	}
	if !correct {
		return errors.New("PromQL表达式不正确")
	}

	return r.repo.UpdateMonitorRecordRule(ctx, rule)
}

func (r *RecordRuleDomain) DeleteRecordRule(ctx context.Context, id int64) error {
	return r.repo.DeleteMonitorRecordRule(ctx, id)
}

func (r *RecordRuleDomain) BatchDeleteRecordRule(ctx context.Context, ids []int64) error {
	return r.repo.BatchDeleteMonitorRecordRule(ctx, ids)
}

func (r *RecordRuleDomain) EnableSwitchRecordRule(ctx context.Context, id int64) error {
	return r.repo.EnableSwitchMonitorRecordRule(ctx, id)
}

func (r *RecordRuleDomain) BatchEnableSwitchRecordRule(ctx context.Context, ids []int64) error {
	return r.repo.BatchEnableSwitchMonitorRecordRule(ctx, ids)
}

func (r *RecordRuleDomain) BuildRecordRuleRespModel(rules []*model.MonitorRecordRule) []*types.RecordRule {
	vec := make([]*types.RecordRule, 0)
	for _, rule := range rules {
		vec = append(vec, &types.RecordRule{
			Id:          rule.ID,
			Name:        rule.Name,
			RecordName:  rule.RecordName,
			UserId:      rule.UserID,
			PoolId:      rule.PoolID,
			TreeNodeId:  rule.TreeNodeID,
			Enable:      rule.Enable,
			ForDuration: rule.ForDuration,
			Expr:        rule.Expr,
		})
	}
	return vec
}

func (r *RecordRuleDomain) BuildRecordRuleModel(rule *types.RecordRule) *model.MonitorRecordRule {
	return &model.MonitorRecordRule{
		ID:          rule.Id,
		Name:        rule.Name,
		RecordName:  rule.RecordName,
		UserID:      rule.UserId,
		PoolID:      rule.PoolId,
		TreeNodeID:  rule.TreeNodeId,
		Enable:      rule.Enable,
		ForDuration: rule.ForDuration,
		Expr:        rule.Expr,
	}
}
