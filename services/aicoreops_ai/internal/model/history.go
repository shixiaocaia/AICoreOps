package model

type History struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:历史ID"`
	SessionID string `json:"session_id" gorm:"comment:会话ID"`
	Question  string `json:"question" gorm:"comment:问题"`
	Answer    string `json:"answer" gorm:"comment:答案"`

	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	DeleteAt  int64 `json:"delete_at" gorm:"comment:删除时间"`
}

func (m *History) TableName() string {
	return "history"
}
