package dao

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

// MonitorScrapePoolDAO 采集池数据访问对象
type MonitorScrapePoolDAO struct {
	db *gorm.DB
}

// NewMonitorScrapePoolDAO 创建采集池DAO
func NewMonitorScrapePoolDAO(db *gorm.DB) *MonitorScrapePoolDAO {
	return &MonitorScrapePoolDAO{
		db: db,
	}
}

// GetMonitorScrapePoolList 获取采集池列表
func (d *MonitorScrapePoolDAO) GetMonitorScrapePoolList(ctx context.Context) ([]*model.MonitorScrapePool, error) {
	var pools []*model.MonitorScrapePool
	err := d.db.WithContext(ctx).Find(&pools).Error
	if err != nil {
		return nil, err
	}
	return pools, nil
}

// CreateMonitorScrapePool 创建采集池
func (d *MonitorScrapePoolDAO) CreateMonitorScrapePool(ctx context.Context, pool *model.MonitorScrapePool) error {
	return d.db.WithContext(ctx).Create(pool).Error
}

// UpdateMonitorScrapePool 更新采集池
func (d *MonitorScrapePoolDAO) UpdateMonitorScrapePool(ctx context.Context, pool *model.MonitorScrapePool) error {
	if pool == nil {
		return errors.New("pool 不能为空")
	}

	if pool.ID <= 0 {
		return fmt.Errorf("无效的 poolId：%d", pool.ID)
	}

	return d.db.WithContext(ctx).Model(&model.MonitorScrapePool{}).Where("id = ?", pool.ID).Updates(pool).Error
}

// DeleteMonitorScrapePool 删除采集池
func (d *MonitorScrapePoolDAO) DeleteMonitorScrapePool(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Delete(&model.MonitorScrapePool{}, id).Error
}

// SearchMonitorScrapePoolByName 根据名称搜索采集池
func (d *MonitorScrapePoolDAO) SearchMonitorScrapePoolByName(ctx context.Context, name string) ([]*model.MonitorScrapePool, error) {
	var pools []*model.MonitorScrapePool
	err := d.db.WithContext(ctx).Where("name = ?", name).Find(&pools).Error
	if err != nil {
		return nil, err
	}
	return pools, nil
}
