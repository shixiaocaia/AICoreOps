package model

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID        uint   `gorm:"primarykey"`
	SessionID string `gorm:"index;type:varchar(64)"`
	Question  string `gorm:"type:text"`
	Answer    string `gorm:"type:text"`
	CreatedAt time.Time
}

// TableName 指定表名
func (History) TableName() string {
	return "history"
}

type HistoryModel struct {
	DB *gorm.DB
}

func NewHistoryModel(db *gorm.DB) *HistoryModel {
	return &HistoryModel{
		DB: db,
	}
}

func (m *HistoryModel) GetBySessionID(sessionID string) ([]History, error) {
	var histories []History
	err := m.DB.Where("session_id = ?", sessionID).
		Order("created_at desc").
		Find(&histories).Error
	return histories, err
}

func (m *HistoryModel) Create(history *History) error {
	return m.DB.Create(history).Error
}

// 其他方法...
