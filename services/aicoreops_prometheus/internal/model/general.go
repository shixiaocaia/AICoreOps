package model

import (
	"database/sql/driver"
	"strings"
	"time"

	"gorm.io/plugin/soft_delete"
)

// StringList 封装了 []string 类型，用于与数据库中的逗号分隔字符串进行转换
type StringList []string

func (m *StringList) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), "|")
	*m = ss
	return nil
}

func (m StringList) Value() (driver.Value, error) {
	str := strings.Join(m, "|")
	return str, nil
}

type Model struct {
	ID        int                   `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"` // 主键ID，自增
	CreatedAt time.Time             `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`   // 创建时间，自动记录
	UpdatedAt time.Time             `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`   // 更新时间，自动记录
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"index;comment:删除时间"`            // 软删除时间，使用普通索引
}
