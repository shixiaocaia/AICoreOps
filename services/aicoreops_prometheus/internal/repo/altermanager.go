package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

type MonitorAltermanagerRepo interface {
	CreateMonitorAlertmanagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error
	GetMonitorAlertmanagerPool(ctx context.Context, name string) (bool, error)
	// GetMonitorAlertmanagerPoolList(ctx context.Context) ([]model.MonitorAlertManagerPool, error)
	// UpdateMonitorAlertmanagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error
	// DeleteMonitorAlertmanagerPool(ctx context.Context, poolId int64) error
}
