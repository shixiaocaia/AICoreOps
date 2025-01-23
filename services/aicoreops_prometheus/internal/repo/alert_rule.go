package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

type RuleRepo interface {
	GetMonitorAlertRuleByPoolId(ctx context.Context, poolId int) ([]*model.MonitorAlertRule, error)
	SearchMonitorAlertRuleByName(ctx context.Context, name string) ([]*model.MonitorAlertRule, error)
	GetMonitorAlertRuleList(ctx context.Context) ([]*model.MonitorAlertRule, error)
	CreateMonitorAlertRule(ctx context.Context, monitorAlertRule *model.MonitorAlertRule) error
	GetMonitorAlertRuleById(ctx context.Context, id int) (*model.MonitorAlertRule, error)
	UpdateMonitorAlertRule(ctx context.Context, monitorAlertRule *model.MonitorAlertRule) error
	EnableSwitchMonitorAlertRule(ctx context.Context, ruleID int) error
	BatchEnableSwitchMonitorAlertRule(ctx context.Context, ruleIDs []int) error
	DeleteMonitorAlertRule(ctx context.Context, ruleID int) error
	GetAssociatedResourcesBySendGroupId(ctx context.Context, sendGroupId int) ([]*model.MonitorAlertRule, error)
	CheckMonitorAlertRuleExists(ctx context.Context, alertRule *model.MonitorAlertRule) (bool, error)
	CheckMonitorAlertRuleNameExists(ctx context.Context, alertRule *model.MonitorAlertRule) (bool, error)
}
