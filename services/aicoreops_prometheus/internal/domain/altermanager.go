package domain

import (
	"context"
	"errors"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/repo"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
)

type AltermanagerDomain struct {
	repo repo.MonitorAltermanagerRepo
}

func NewAltermanagerDomain(ctx *svc.ServiceContext) *AltermanagerDomain {
	return &AltermanagerDomain{repo: dao.NewAltermanagerDao(ctx.DB)}
}

func (a *AltermanagerDomain) CreateMonitorAlertmanagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	// 检查实例是否存在
	exist, err := a.repo.GetMonitorAlertmanagerPool(ctx, pool.Name)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("AlertmanagerPool 实例已存在")
	}
	return a.repo.CreateMonitorAlertmanagerPool(ctx, pool)
}

// func (a *AltermanagerDomain) GetMonitorAlertmanagerPool(ctx context.Context, poolId int64) (*model.MonitorAlertManagerPool, error) {
// 	return a.repo.GetMonitorAlertmanagerPool(ctx, poolId)
// }

func (a *AltermanagerDomain) GetMonitorAlertmanagerPoolList(ctx context.Context) ([]model.MonitorAlertManagerPool, error) {
	panic("not implemented")
}

func (a *AltermanagerDomain) UpdateMonitorAlertmanagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	panic("not implemented")
}

func (a *AltermanagerDomain) DeleteMonitorAlertmanagerPool(ctx context.Context, poolId int64) error {
	panic("not implemented")
}

func (a *AltermanagerDomain) BuildMonitorAlertmanagerPoolModel(ctx context.Context, pool *types.AlertmanagerPool) *model.MonitorAlertManagerPool {
	return &model.MonitorAlertManagerPool{
		Name:                  pool.Name,
		AlertManagerInstances: pool.AlertmanagerInstances,
		UserID:                int(pool.UserId),
		ResolveTimeout:        pool.ResolveTimeout,
		GroupWait:             pool.GroupWait,
		GroupInterval:         pool.GroupInterval,
		RepeatInterval:        pool.RepeatInterval,
		GroupBy:               pool.GroupBy,
		Receiver:              pool.Receiver,
	}
}
