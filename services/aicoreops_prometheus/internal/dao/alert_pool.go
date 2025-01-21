package dao

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type AlertManagerPoolDao struct {
	db *gorm.DB
}

func NewAlertManagerPoolDao(db *gorm.DB) *AlertManagerPoolDao {
	return &AlertManagerPoolDao{db: db}
}

// CreateMonitorAlertManagerPool 创建AlertManager池
func (d *AlertManagerPoolDao) CreateMonitorAlertManagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	if pool == nil {
		return errors.New("pool 不能为空")
	}

	return d.db.WithContext(ctx).Create(pool).Error
}

// GetMonitorAlertManagerPoolList 获取AlertManager池列表
func (d *AlertManagerPoolDao) GetMonitorAlertManagerPoolList(ctx context.Context) ([]*model.MonitorAlertManagerPool, error) {
	var pools []*model.MonitorAlertManagerPool
	if err := d.db.WithContext(ctx).Find(&pools).Error; err != nil {
		return nil, err
	}
	return pools, nil
}

// SearchMonitorAlertManagerPoolByName 根据名称搜索AlertManager池
func (d *AlertManagerPoolDao) SearchMonitorAlertManagerPoolByName(ctx context.Context, name string) ([]*model.MonitorAlertManagerPool, error) {
	var pools []*model.MonitorAlertManagerPool
	if err := d.db.WithContext(ctx).
		Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%").
		Find(&pools).Error; err != nil {
		return nil, err
	}

	return pools, nil
}

// UpdateMonitorAlertManagerPool 更新AlertManager池
func (d *AlertManagerPoolDao) UpdateMonitorAlertManagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	if pool == nil {
		return errors.New("pool 不能为空")
	}

	if pool.ID <= 0 {
		return fmt.Errorf("无效的 poolId：%d", pool.ID)
	}

	return d.db.WithContext(ctx).Model(&model.MonitorAlertManagerPool{}).Where("id = ?", pool.ID).Updates(pool).Error
}

// DeleteMonitorAlertManagerPool 删除AlertManager池
func (d *AlertManagerPoolDao) DeleteMonitorAlertManagerPool(ctx context.Context, poolId int64) error {
	if poolId <= 0 {
		return fmt.Errorf("无效的 poolId：%d", poolId)
	}

	return d.db.WithContext(ctx).Delete(&model.MonitorAlertManagerPool{}, poolId).Error
}

// CheckMonitorAlertManagerPoolExist 检查 AlertManagerPool 是否存在
func (d *AlertManagerPoolDao) CheckMonitorAlertManagerPoolExist(ctx context.Context, name string) (bool, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.MonitorAlertManagerPool{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
