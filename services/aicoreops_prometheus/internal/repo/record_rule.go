package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

type RecordRuleRepo interface {
	GetMonitorRecordRuleByPoolId(ctx context.Context, poolId int64) ([]*model.MonitorRecordRule, error)
	SearchMonitorRecordRuleByName(ctx context.Context, name string) ([]*model.MonitorRecordRule, error)
	GetMonitorRecordRuleList(ctx context.Context) ([]*model.MonitorRecordRule, error)
	CreateMonitorRecordRule(ctx context.Context, recordRule *model.MonitorRecordRule) error
	GetMonitorRecordRuleById(ctx context.Context, id int64) (*model.MonitorRecordRule, error)
	UpdateMonitorRecordRule(ctx context.Context, recordRule *model.MonitorRecordRule) error
	DeleteMonitorRecordRule(ctx context.Context, ruleID int64) error
	BatchDeleteMonitorRecordRule(ctx context.Context, ruleIDs []int64) error
	EnableSwitchMonitorRecordRule(ctx context.Context, ruleID int64) error
	BatchEnableSwitchMonitorRecordRule(ctx context.Context, ruleIDs []int64) error
	CheckMonitorRecordRuleExists(ctx context.Context, recordRule *model.MonitorRecordRule) (bool, error)
	CheckMonitorRecordRuleNameExists(ctx context.Context, recordRule *model.MonitorRecordRule) (bool, error)
}
