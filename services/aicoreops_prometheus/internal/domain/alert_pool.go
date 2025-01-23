package domain

import (
	"context"
	"fmt"
	"slices"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/pkg"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/repo"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
)

type AlterManagerPoolDomain struct {
	repo repo.MonitorAlterManagerPoolRepo
}

func NewAlterManagerPoolDomain(ctx *svc.ServiceContext) *AlterManagerPoolDomain {
	return &AlterManagerPoolDomain{repo: dao.NewAlertManagerPoolDAO(ctx.DB)}
}

func (a *AlterManagerPoolDomain) CreateMonitorAlertManagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	// 检查 AlertManagerPool是否存在
	exist, err := a.repo.CheckMonitorAlertManagerPoolExist(ctx, pool.Name)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("AlertmanagerPool %s 已存在", pool.Name)
	}

	// 检查 instances 是否存在
	if err := a.checkAlertManagerIpExist(ctx, pool.ID, pool.AlertManagerInstances); err != nil {
		return err
	}

	return a.repo.CreateMonitorAlertManagerPool(ctx, pool)
}

func (a *AlterManagerPoolDomain) GetMonitorAlertManagerPoolList(ctx context.Context, searchName *string) ([]*model.MonitorAlertManagerPool, error) {
	return pkg.HandleList(ctx, searchName,
		a.repo.SearchMonitorAlertManagerPoolByName, // 搜索函数
		a.repo.GetMonitorAlertManagerPoolList)      // 获取所有函数
}

func (a *AlterManagerPoolDomain) UpdateMonitorAlertManagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	// 检查 AlertmanagerPool 是否存在
	exist, err := a.repo.CheckMonitorAlertManagerPoolExist(ctx, pool.Name)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("AlertmanagerPool %s 不存在", pool.Name)
	}

	// 检查 instances 是否存在
	if err := a.checkAlertManagerIpExist(ctx, pool.ID, pool.AlertManagerInstances); err != nil {
		return err
	}

	return a.repo.UpdateMonitorAlertManagerPool(ctx, pool)
}

func (a *AlterManagerPoolDomain) DeleteMonitorAlertManagerPool(ctx context.Context, poolId int64) error {
	return a.repo.DeleteMonitorAlertManagerPool(ctx, poolId)
}

func (a *AlterManagerPoolDomain) checkAlertManagerIpExist(ctx context.Context, poolId int64, ip []string) error {
	pools, err := a.repo.GetMonitorAlertManagerPoolList(ctx)
	if err != nil {
		return err
	}

	ips := make([]string, 0)
	for _, pool := range pools {
		if pool.ID == poolId {
			continue
		}

		ips = append(ips, pool.AlertManagerInstances...)
	}

	for _, i := range ip {
		if slices.Contains(ips, i) {
			return fmt.Errorf("alertmanager 实例 %s 已存在", i)
		}
	}
	return nil
}

func (a *AlterManagerPoolDomain) BuildMonitorAlertManagerPoolModel(pool *types.AlertManagerPool) *model.MonitorAlertManagerPool {
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

func (a *AlterManagerPoolDomain) BuildAlertManagerPoolRespModel(pools []*model.MonitorAlertManagerPool) []*types.AlertManagerPool {
	list := make([]*types.AlertManagerPool, 0)
	for _, pool := range pools {
		list = append(list, &types.AlertManagerPool{
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
