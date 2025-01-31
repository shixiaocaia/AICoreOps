package model

type HistorySession struct {
	ID        int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:历史ID"`
	SessionID int64
	UserID    int64
	Title     string

	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	DeleteAt  int64 `json:"delete_at" gorm:"comment:删除时间"`
}

func (m *HistorySession) TableName() string {
	return "history_session"
}
