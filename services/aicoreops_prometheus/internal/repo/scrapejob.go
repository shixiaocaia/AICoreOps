package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

type MonitorScrapeJobRepo interface {
	GetMonitorScrapeJobList(ctx context.Context) ([]*model.MonitorScrapeJob, error)
	CreateMonitorScrapeJob(ctx context.Context, job *model.MonitorScrapeJob) error
	UpdateMonitorScrapeJob(ctx context.Context, job *model.MonitorScrapeJob) error
	DeleteMonitorScrapeJob(ctx context.Context, id int64) error
	SearchMonitorScrapeJobByName(ctx context.Context, name string) ([]*model.MonitorScrapeJob, error)
	SearchMonitorScrapeJobByID(ctx context.Context, id int64) ([]*model.MonitorScrapeJob, error)
}
