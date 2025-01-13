package model

import (
	"database/sql/driver"
	"strings"
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
