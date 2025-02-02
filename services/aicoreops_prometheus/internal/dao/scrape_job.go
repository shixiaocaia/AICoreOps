package dao

import (
	"context"
	"errors"
	"fmt"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type MonitorScrapeJobDAO struct {
	db *gorm.DB
}

func NewMonitorScrapeJobDAO(db *gorm.DB) *MonitorScrapeJobDAO {
	return &MonitorScrapeJobDAO{db: db}
}

// GetMonitorScrapeJobList 获取采集任务列表
func (d *MonitorScrapeJobDAO) GetMonitorScrapeJobList(ctx context.Context) ([]*model.MonitorScrapeJob, error) {
	var jobs []*model.MonitorScrapeJob
	if err := d.db.WithContext(ctx).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

// CreateMonitorScrapeJob 创建采集任务
func (d *MonitorScrapeJobDAO) CreateMonitorScrapeJob(ctx context.Context, job *model.MonitorScrapeJob) error {
	if job == nil {
		return errors.New("job 不能为空")
	}

	return d.db.WithContext(ctx).Create(job).Error
}

// UpdateMonitorScrapeJob 更新采集任务
func (d *MonitorScrapeJobDAO) UpdateMonitorScrapeJob(ctx context.Context, job *model.MonitorScrapeJob) error {
	if job == nil {
		return errors.New("job 不能为空")
	}

	if job.ID <= 0 {
		return fmt.Errorf("无效的 jobId：%d", job.ID)
	}

	return d.db.WithContext(ctx).Model(&model.MonitorScrapeJob{}).Where("id = ?", job.ID).Updates(job).Error
}

// DeleteMonitorScrapeJob 删除采集任务
func (d *MonitorScrapeJobDAO) DeleteMonitorScrapeJob(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("无效的 jobId：%d", id)
	}

	return d.db.WithContext(ctx).Delete(&model.MonitorScrapeJob{}, id).Error
}

// SearchMonitorScrapeJobByName 根据名称搜索采集任务
func (d *MonitorScrapeJobDAO) SearchMonitorScrapeJobByName(ctx context.Context, name string) ([]*model.MonitorScrapeJob, error) {
	var jobs []*model.MonitorScrapeJob
	if err := d.db.WithContext(ctx).Where("name = ?", name).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

// SearchMonitorScrapeJobByID 根据poolID获取采集任务
func (d *MonitorScrapeJobDAO) SearchMonitorScrapeJobByID(ctx context.Context, id int64) ([]*model.MonitorScrapeJob, error) {
	var jobs []*model.MonitorScrapeJob
	if err := d.db.WithContext(ctx).Where("pool_id = ?", id).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}
