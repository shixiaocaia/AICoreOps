package dao

import (
	"context"

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

func (d *AltermanagerDao) GetMonitorAlertmanagerPool(ctx context.Context, name string) (bool, error) {
	var pool model.MonitorAlertManagerPool
	if err := d.db.WithContext(ctx).Where("name = ?", name).First(&pool).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *AltermanagerDao) GetMonitorAlertmanagerPoolList(ctx context.Context) ([]model.MonitorAlertManagerPool, error) {
	var pools []model.MonitorAlertManagerPool
	if err := d.db.WithContext(ctx).Find(&pools).Error; err != nil {
		return nil, err
	}
	return pools, nil
}
