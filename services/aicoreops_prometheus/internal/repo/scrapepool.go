package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

// MonitorScrapePoolRepo 采集池仓储接口
type MonitorScrapePoolRepo interface {
	GetMonitorScrapePoolList(ctx context.Context) ([]*model.MonitorScrapePool, error)
	CreateMonitorScrapePool(ctx context.Context, pool *model.MonitorScrapePool) error
	UpdateMonitorScrapePool(ctx context.Context, pool *model.MonitorScrapePool) error
	DeleteMonitorScrapePool(ctx context.Context, id int64) error
	SearchMonitorScrapePoolByName(ctx context.Context, name string) ([]*model.MonitorScrapePool, error)
}
