package domain

import (
	"context"
	"errors"
	"slices"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/pkg"
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
	// 检查 AlertManagerPool是否存在
	pools, err := a.repo.SearchMonitorAlertManagerPoolByName(ctx, pool.Name)
	if err != nil {
		return err
	}
	if len(pools) > 0 {
		return errors.New("AlertmanagerPool 实例已存在")
	}

	// 检查 instances 是否存在
	exist, err := a.checkAlertmanagerIpExist(ctx, pool.ID, pool.AlertManagerInstances)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("instances 已存在其他 AlertmanagerPool 中")
	}
	return a.repo.CreateMonitorAlertmanagerPool(ctx, pool)
}

func (a *AltermanagerDomain) GetMonitorAlertmanagerPoolList(ctx context.Context, searchName *string) ([]*model.MonitorAlertManagerPool, error) {
	return pkg.HandleList(ctx, searchName,
		a.repo.SearchMonitorAlertManagerPoolByName, // 搜索函数
		a.repo.GetMonitorAlertmanagerPoolList)      // 获取所有函数
}

func (a *AltermanagerDomain) UpdateMonitorAlertmanagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	// 检查 instances 是否存在
	exist, err := a.checkAlertmanagerIpExist(ctx, pool.ID, pool.AlertManagerInstances)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("instances 已存在其他 AlertmanagerPool 中")
	}

	return a.repo.UpdateMonitorAlertmanagerPool(ctx, pool)
}

func (a *AltermanagerDomain) DeleteMonitorAlertmanagerPool(ctx context.Context, poolId int64) error {
	return a.repo.DeleteMonitorAlertmanagerPool(ctx, poolId)
}

func (a *AltermanagerDomain) checkAlertmanagerIpExist(ctx context.Context, poolId int64, ip []string) (bool, error) {
	pools, err := a.repo.GetMonitorAlertmanagerPoolList(ctx)
	if err != nil {
		return false, err
	}

	existIps := make([]string, 0)
	for _, pool := range pools {
		if pool.ID == poolId {
			continue
		}

		existIps = append(existIps, pool.AlertManagerInstances...)
	}

	for _, i := range ip {
		if slices.Contains(existIps, i) {
			return true, nil
		}
	}
	return false, nil
}

func (a *AltermanagerDomain) BuildMonitorAlertmanagerPoolModel(pool *types.AlertmanagerPool) *model.MonitorAlertManagerPool {
	return &model.MonitorAlertManagerPool{
		ID:                    pool.Id,
		Name:                  pool.Name,
		AlertManagerInstances: pool.AlertmanagerInstances,
		UserID:                pool.UserId,
		ResolveTimeout:        pool.ResolveTimeout,
		GroupWait:             pool.GroupWait,
		GroupInterval:         pool.GroupInterval,
		RepeatInterval:        pool.RepeatInterval,
		GroupBy:               pool.GroupBy,
		Receiver:              pool.Receiver,
	}
}

func (a *AltermanagerDomain) BuildAlertmanagerPoolRespModel(pools []*model.MonitorAlertManagerPool) []*types.AlertmanagerPool {
	list := make([]*types.AlertmanagerPool, 0)
	for _, pool := range pools {
		list = append(list, &types.AlertmanagerPool{
			Id:                    int64(pool.ID),
			Name:                  pool.Name,
			AlertmanagerInstances: pool.AlertManagerInstances,
			UserId:                int64(pool.UserID),
			ResolveTimeout:        pool.ResolveTimeout,
			GroupWait:             pool.GroupWait,
			GroupInterval:         pool.GroupInterval,
			RepeatInterval:        pool.RepeatInterval,
			GroupBy:               pool.GroupBy,
			Receiver:              pool.Receiver,
		})
	}
	return list
}
