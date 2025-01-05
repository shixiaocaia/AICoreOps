package dao

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type AltermanagerDao struct {
	db *gorm.DB
}

func NewAltermanagerDao(db *gorm.DB) *AltermanagerDao {
	return &AltermanagerDao{db: db}
}

func (d *AltermanagerDao) CreateMonitorAlertmanagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	return d.db.WithContext(ctx).Create(pool).Error
}

func (d *AltermanagerDao) GetMonitorAlertmanagerPoolList(ctx context.Context) ([]*model.MonitorAlertManagerPool, error) {
	var pools []*model.MonitorAlertManagerPool
	if err := d.db.WithContext(ctx).Find(&pools).Error; err != nil {
		return nil, err
	}
	return pools, nil
}

func (d *AltermanagerDao) SearchMonitorAlertManagerPoolByName(ctx context.Context, name string) ([]*model.MonitorAlertManagerPool, error) {
	var pools []*model.MonitorAlertManagerPool

	if err := d.db.WithContext(ctx).
		Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%").
		Find(&pools).Error; err != nil {
		return nil, err
	}

	return pools, nil
}

func (d *AltermanagerDao) UpdateMonitorAlertmanagerPool(ctx context.Context, pool *model.MonitorAlertManagerPool) error {
	if pool == nil {
		return errors.New("pool 不能为空")
	}

	if pool.ID <= 0 {
		return fmt.Errorf("无效的 poolId：%d", pool.ID)
	}

	return d.db.WithContext(ctx).Model(&model.MonitorAlertManagerPool{}).Where("id = ?", pool.ID).Updates(pool).Error
}

func (d *AltermanagerDao) DeleteMonitorAlertmanagerPool(ctx context.Context, poolId int64) error {
	if poolId <= 0 {
		return fmt.Errorf("无效的 poolId：%d", poolId)
	}

	return d.db.WithContext(ctx).Delete(&model.MonitorAlertManagerPool{}, poolId).Error
}
