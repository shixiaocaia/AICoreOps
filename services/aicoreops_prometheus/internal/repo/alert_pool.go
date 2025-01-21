package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

// MonitorAlterManagerPoolRepo 监控告警池Repo
type MonitorAlterManagerPoolRepo interface {
	CreateMonitorAlertManagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error
	SearchMonitorAlertManagerPoolByName(ctx context.Context, name string) ([]*model.MonitorAlertManagerPool, error)
	GetMonitorAlertManagerPoolList(ctx context.Context) ([]*model.MonitorAlertManagerPool, error)
	UpdateMonitorAlertManagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error
	DeleteMonitorAlertManagerPool(ctx context.Context, poolId int64) error
	CheckMonitorAlertManagerPoolExist(ctx context.Context, name string) (bool, error)
}
