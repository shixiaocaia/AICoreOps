package model

type HistorySession struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:历史ID"`
	SessionID string `json:"session_id" gorm:"comment:会话ID"`
	UserID    int64  `json:"user_id" gorm:"comment:用户ID"`
	Title     string `json:"title" gorm:"comment:标题"`

	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt int64 `json:"deleted_at" gorm:"comment:删除时间"`
}

func (m *HistorySession) TableName() string {
	return "history_session"
}
