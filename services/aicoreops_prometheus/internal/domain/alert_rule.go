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

type AlertRuleDomain struct {
	repo repo.AlertRuleRepo
}

func NewAlertRuleDomain(svcCtx *svc.ServiceContext) *AlertRuleDomain {
	return &AlertRuleDomain{
		repo: dao.NewAlertRuleDAO(svcCtx.DB),
	}
}

func (a *AlertRuleDomain) GetAlertRuleList(ctx context.Context, req *types.GetAlertRuleListRequest) ([]*model.AlertRule, error) {
	rules, err := a.repo.GetAlertRuleList(ctx)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (a *AlertRuleDomain) CreateAlertRule(ctx context.Context, rule *model.AlertRule) error {
	// 检查告警规则是否存在
	exists, err := a.repo.CheckAlertRuleNameExists(ctx, rule.Name)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("告警规则已存在")
	}

	// 检查PromQL表达式
	if err := a.CheckPromqlExpr(ctx, rule.Expr); err != nil {
		return err
	}

	return a.repo.CreateAlertRule(ctx, rule)
}

func (a *AlertRuleDomain) UpdateAlertRule(ctx context.Context, rule *model.AlertRule) error {
	// 检查告警规则是否存在
	exists, err := a.repo.CheckAlertRuleExists(ctx, rule.ID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("告警规则不存在")
	}

	// 检查PromQL表达式
	if err := a.CheckPromqlExpr(ctx, rule.Expr); err != nil {
		return err
	}

	return a.repo.UpdateAlertRule(ctx, rule)
}

func (a *AlertRuleDomain) DeleteAlertRule(ctx context.Context, id int64) error {
	return a.repo.DeleteAlertRule(ctx, id)
}

func (a *AlertRuleDomain) BatchDeleteAlertRule(ctx context.Context, ids []int64) error {
	return a.repo.BatchDeleteAlertRule(ctx, ids)
}

func (a *AlertRuleDomain) EnableSwitchAlertRule(ctx context.Context, id int64) error {
	return a.repo.EnableSwitchAlertRule(ctx, id)
}

func (a *AlertRuleDomain) BatchEnableSwitchAlertRule(ctx context.Context, ids []int64) error {
	return a.repo.BatchEnableSwitchAlertRule(ctx, ids)
}

func (a *AlertRuleDomain) CheckPromqlExpr(ctx context.Context, expr string) error {
	correct, err := pkg.PromqlExprCheck(expr)
	if err != nil {
		return err
	}
	if !correct {
		return errors.New("PromQL表达式不正确")
	}
	return nil
}

func (a *AlertRuleDomain) BuildAlertRuleRespModel(rules []*model.AlertRule) []*types.AlertRule {
	vec := make([]*types.AlertRule, 0)
	for _, rule := range rules {
		vec = append(vec, &types.AlertRule{
			Id:          rule.ID,
			Name:        rule.Name,
			UserId:      rule.UserID,
			PoolId:      rule.PoolID,
			SendGroupId: rule.SendGroupID,
			TreeNodeId:  rule.TreeNodeID,
			Enable:      rule.Enable,
			Expr:        rule.Expr,
			Severity:    rule.Severity,
			GrafanaLink: rule.GrafanaLink,
			ForDuration: rule.ForDuration,
			Labels:      rule.Labels,
			Annotations: rule.Annotations,
		})
	}
	return vec
}

func (a *AlertRuleDomain) BuildAlertRuleModel(rule *types.AlertRule) *model.AlertRule {
	return &model.AlertRule{
		ID:          rule.Id,
		Name:        rule.Name,
		UserID:      rule.UserId,
		PoolID:      rule.PoolId,
		SendGroupID: rule.SendGroupId,
		TreeNodeID:  rule.TreeNodeId,
		Enable:      rule.Enable,
		Expr:        rule.Expr,
		Severity:    rule.Severity,
		GrafanaLink: rule.GrafanaLink,
		ForDuration: rule.ForDuration,
		Labels:      rule.Labels,
		Annotations: rule.Annotations,
	}
}
