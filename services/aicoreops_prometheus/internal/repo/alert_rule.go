package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

type AlertRuleRepo interface {
	GetAlertRuleByPoolId(ctx context.Context, poolId int64) ([]*model.AlertRule, error)
	SearchAlertRuleByName(ctx context.Context, name string) ([]*model.AlertRule, error)
	GetAlertRuleList(ctx context.Context) ([]*model.AlertRule, error)
	CreateAlertRule(ctx context.Context, monitorAlertRule *model.AlertRule) error
	GetAlertRuleById(ctx context.Context, id int64) (*model.AlertRule, error)
	UpdateAlertRule(ctx context.Context, monitorAlertRule *model.AlertRule) error
	EnableSwitchAlertRule(ctx context.Context, ruleID int64) error
	BatchEnableSwitchAlertRule(ctx context.Context, ruleIDs []int64) error
	DeleteAlertRule(ctx context.Context, id int64) error
	BatchDeleteAlertRule(ctx context.Context, ids []int64) error
	GetAssociatedResourcesBySendGroupId(ctx context.Context, sendGroupId int64) ([]*model.AlertRule, error)
	CheckAlertRuleExists(ctx context.Context, id int64) (bool, error)
	CheckAlertRuleNameExists(ctx context.Context, name string) (bool, error)
}
