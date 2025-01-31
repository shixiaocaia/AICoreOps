package dao

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/model"
	"gorm.io/gorm"
)

type HistoryDAO struct {
	db *gorm.DB
}

func NewHistoryDAO(db *gorm.DB) *HistoryDAO {
	return &HistoryDAO{db: db}
}

// CreateHistory 创建历史记录
func (d *HistoryDAO) CreateHistory(ctx context.Context, history *model.History) error {
	return d.db.WithContext(ctx).Create(history).Error
}

// GetHistoryByID 获取历史记录
func (d *HistoryDAO) GetHistoryByID(ctx context.Context, id int64) (*model.History, error) {
	var history model.History
	if err := d.db.Where("id = ?", id).First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

// GetHistoryList 获取历史记录列表
func (d *HistoryDAO) GetHistoryList(ctx context.Context, userId int64, offset, limit int) ([]*model.History, error) {
	var histories []*model.History
	if err := d.db.WithContext(ctx).Where("user_id = ?", userId).Offset(offset).Limit(limit).Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

// UpdateHistory 更新历史记录
func (d *HistoryDAO) UpdateHistory(ctx context.Context, history *model.History) error {
	return d.db.WithContext(ctx).Save(history).Error
}

// DeleteHistory 删除历史记录
func (d *HistoryDAO) DeleteHistory(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.History{}).Error
}
