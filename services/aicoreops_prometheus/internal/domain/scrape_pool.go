package domain

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/repo"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
)

type MonitorScrapePoolDomain struct {
	repo          repo.MonitorScrapePoolRepo
	scrapeJobRepo repo.MonitorScrapeJobRepo
}

func NewMonitorScrapePoolDomain(svcCtx *svc.ServiceContext) *MonitorScrapePoolDomain {
	return &MonitorScrapePoolDomain{
		repo:          dao.NewMonitorScrapePoolDAO(svcCtx.DB),
		scrapeJobRepo: dao.NewMonitorScrapeJobDAO(svcCtx.DB),
	}
}

func (d *MonitorScrapePoolDomain) GetMonitorScrapePoolList(ctx context.Context) ([]*model.MonitorScrapePool, error) {
	return d.repo.GetMonitorScrapePoolList(ctx)
}

func (d *MonitorScrapePoolDomain) CreateMonitorScrapePool(ctx context.Context, pool *model.MonitorScrapePool) error {
	// 检查采集池是否存在
	exist, err := d.repo.CheckMonitorScrapePoolExist(ctx, pool.Name)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("采集池 %s 已存在", pool.Name)
	}

	// 检查实例是否存在
	if err = d.checkInstanceExist(ctx, pool.ID, pool.PrometheusInstances, pool.AlertManagerInstances); err != nil {
		return err
	}

	return d.repo.CreateMonitorScrapePool(ctx, pool)
}

func (d *MonitorScrapePoolDomain) UpdateMonitorScrapePool(ctx context.Context, pool *model.MonitorScrapePool) error {
	// 检查采集池是否存在
	exist, err := d.repo.CheckMonitorScrapePoolExist(ctx, pool.Name)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("采集池 %s 不存在", pool.Name)
	}

	// 检查实例是否存在
	if err := d.checkInstanceExist(ctx, pool.ID, pool.PrometheusInstances, pool.AlertManagerInstances); err != nil {
		return err
	}

	return d.repo.UpdateMonitorScrapePool(ctx, pool)
}

func (d *MonitorScrapePoolDomain) DeleteMonitorScrapePool(ctx context.Context, id int64) error {
	// 检查采集任务是否存在
	exist, err := d.scrapeJobRepo.SearchMonitorScrapeJobByID(ctx, id)
	if err != nil {
		return err
	}
	if len(exist) > 0 {
		return errors.New("采集池关联采集任务，无法删除")
	}

	return d.repo.DeleteMonitorScrapePool(ctx, id)
}

func (d *MonitorScrapePoolDomain) checkInstanceExist(ctx context.Context, pid int64, instancesP, instancesA []string) error {
	pools, err := d.repo.GetMonitorScrapePoolList(ctx)
	if err != nil {
		return err
	}

	ipPrometheus := make([]string, 0)
	ipAlertManager := make([]string, 0)
	for _, pool := range pools {
		if pool.ID == pid {
			continue
		}

		ipPrometheus = append(ipPrometheus, pool.PrometheusInstances...)
		ipAlertManager = append(ipAlertManager, pool.AlertManagerInstances...)
	}

	for _, ip := range instancesP {
		if slices.Contains(ipPrometheus, ip) {
			return fmt.Errorf("prometheus 实例 %s 已存在", ip)
		}
	}

	for _, ip := range instancesA {
		if slices.Contains(ipAlertManager, ip) {
			return fmt.Errorf("alertmanager 实例 %s 已存在", ip)
		}
	}
	return nil
}

func (d *MonitorScrapePoolDomain) BuildScrapePoolRespModel(pools []*model.MonitorScrapePool) []*types.ScrapePool {
	var data []*types.ScrapePool
	for _, pool := range pools {
		data = append(data, &types.ScrapePool{
			Id:                    pool.ID,
			Name:                  pool.Name,
			PrometheusInstances:   pool.PrometheusInstances,
			AlertmanagerInstances: pool.AlertManagerInstances,
			UserId:                pool.UserID,
			ScrapeInterval:        pool.ScrapeInterval,
			ScrapeTimeout:         pool.ScrapeTimeout,
			ExternalLabels:        pool.ExternalLabels,
			SupportAlert:          pool.SupportAlert,
			SupportRecord:         pool.SupportRecord,
			RemoteReadUrl:         pool.RemoteReadUrl,
			AlertmanagerUrl:       pool.AlertManagerUrl,
			RuleFilePath:          pool.RuleFilePath,
			RecordFilePath:        pool.RecordFilePath,
			RemoteWriteUrl:        pool.RemoteWriteUrl,
			RemoteTimeoutSeconds:  pool.RemoteTimeoutSeconds,
		})
	}
	return data
}

func BuildMonitorScrapePoolModel(pool *types.ScrapePool) *model.MonitorScrapePool {
	return &model.MonitorScrapePool{
		ID:                    pool.Id,
		Name:                  pool.Name,
		PrometheusInstances:   pool.PrometheusInstances,
		AlertManagerInstances: pool.AlertmanagerInstances,
		UserID:                pool.UserId,
		ScrapeInterval:        pool.ScrapeInterval,
		ScrapeTimeout:         pool.ScrapeTimeout,
		ExternalLabels:        pool.ExternalLabels,
		SupportAlert:          pool.SupportAlert,
		SupportRecord:         pool.SupportRecord,
		RemoteReadUrl:         pool.RemoteReadUrl,
		AlertManagerUrl:       pool.AlertmanagerUrl,
		RuleFilePath:          pool.RuleFilePath,
		RecordFilePath:        pool.RecordFilePath,
		RemoteWriteUrl:        pool.RemoteWriteUrl,
		RemoteTimeoutSeconds:  pool.RemoteTimeoutSeconds,
	}
}
