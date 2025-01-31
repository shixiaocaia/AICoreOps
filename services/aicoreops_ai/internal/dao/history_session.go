package dao

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/model"
	"gorm.io/gorm"
)

type HistorySessionDAO struct {
	db *gorm.DB
}

func NewHistorySessionDAO(db *gorm.DB) *HistorySessionDAO {
	return &HistorySessionDAO{db: db}
}

// CreateHistorySession 创建历史会话
func (d *HistorySessionDAO) CreateHistorySession(ctx context.Context, session *model.HistorySession) error {
	return d.db.WithContext(ctx).Create(session).Error
}

// GetHistorySessionByID 获取历史会话
func (d *HistorySessionDAO) GetHistorySessionByID(ctx context.Context, id int64) (*model.HistorySession, error) {
	var session model.HistorySession
	if err := d.db.WithContext(ctx).Where("id = ?", id).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

// GetHistorySessionList 获取历史会话列表
func (d *HistorySessionDAO) GetHistorySessionList(ctx context.Context, userId int64, offset, limit int) ([]*model.HistorySession, error) {
	var sessions []*model.HistorySession
	if err := d.db.WithContext(ctx).Where("user_id = ?", userId).Offset(offset).Limit(limit).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// UpdateHistorySession 更新历史会话
func (d *HistorySessionDAO) UpdateHistorySession(ctx context.Context, session *model.HistorySession) error {
	return d.db.WithContext(ctx).Save(session).Error
}

// DeleteHistorySession 删除历史会话
func (d *HistorySessionDAO) DeleteHistorySession(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.HistorySession{}).Error
}
