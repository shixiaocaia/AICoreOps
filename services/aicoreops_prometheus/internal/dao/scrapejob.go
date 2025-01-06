package dao

import (
	"context"
	"errors"
	"fmt"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type ScrapeJobDao struct {
	db *gorm.DB
}

func NewScrapeJobDao(db *gorm.DB) *ScrapeJobDao {
	return &ScrapeJobDao{db: db}
}

func (d *ScrapeJobDao) GetMonitorScrapeJobList(ctx context.Context) ([]*model.MonitorScrapeJob, error) {
	var jobs []*model.MonitorScrapeJob
	if err := d.db.WithContext(ctx).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (d *ScrapeJobDao) SearchMonitorScrapeJobByName(ctx context.Context, name string) ([]*model.MonitorScrapeJob, error) {
	var jobs []*model.MonitorScrapeJob
	if err := d.db.WithContext(ctx).Where("name = ?", name).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (d *ScrapeJobDao) CreateMonitorScrapeJob(ctx context.Context, job *model.MonitorScrapeJob) error {
	if job == nil {
		return errors.New("job 不能为空")
	}

	return d.db.WithContext(ctx).Create(job).Error
}

func (d *ScrapeJobDao) UpdateMonitorScrapeJob(ctx context.Context, job *model.MonitorScrapeJob) error {
	if job == nil {
		return errors.New("job 不能为空")
	}

	if job.ID <= 0 {
		return fmt.Errorf("无效的 jobId：%d", job.ID)
	}

	return d.db.WithContext(ctx).Model(&model.MonitorScrapeJob{}).Where("id = ?", job.ID).Updates(job).Error
}

func (d *ScrapeJobDao) DeleteMonitorScrapeJob(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("无效的 jobId：%d", id)
	}

	return d.db.WithContext(ctx).Delete(&model.MonitorScrapeJob{}, id).Error
}
