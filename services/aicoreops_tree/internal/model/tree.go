package model

import (
	"database/sql/driver"
	"strings"
)

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

// TreeNode 服务树节点模型
type TreeNode struct {
	ID          int64       `json:"id" gorm:"primaryKey;column:id;comment:主键ID"`
	CreateTime  int64       `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdateTime  int64       `json:"update_time" gorm:"column:update_time;autoUpdateTime;comment:更新时间"`
	IsDeleted   int         `json:"is_deleted" gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除(0:否,1:是)"`
	Title       string      `json:"title" gorm:"column:title;type:varchar(255);not null;comment:节点标题"`
	Pid         int         `json:"pid" gorm:"column:pid;type:int(11);not null;default:0;comment:父节点ID(0表示根节点)"`
	Level       int         `json:"level" gorm:"column:level;type:int(11);not null;default:1;comment:节点层级(从1开始)"`
	IsLeaf      int         `json:"is_leaf" gorm:"column:is_leaf;type:tinyint(1);not null;default:0;comment:是否为叶子节点(0:否,1:是)"`
	Description string      `json:"description" gorm:"column:description;type:varchar(500);comment:节点描述"`
	Children    []*TreeNode `json:"children" gorm:"-"`
	CMDBID      string      `json:"cmdb_id" gorm:"column:cmdb_id;type:varchar(64);uniqueIndex;comment:CMDB资源唯一标识"`
	CMDBType    string      `json:"cmdb_type" gorm:"column:cmdb_type;type:varchar(32);index;comment:CMDB资源类型(如:host,container,service等)"`
	CMDBAttrs   StringList  `json:"cmdb_attrs" gorm:"column:cmdb_attrs;type:text;comment:CMDB资源属性(如:IP,CPU,内存等)"`
	Creator     string      `json:"creator" gorm:"column:creator;type:varchar(64);not null;comment:创建者"`
	Updater     string      `json:"updater" gorm:"column:updater;type:varchar(64);comment:更新者"`
	Status      string      `json:"status" gorm:"column:status;type:varchar(16);default:normal;comment:节点状态(normal-正常,disabled-禁用,deleted-已删除)"`
}

// TableName 返回表名
func (t *TreeNode) TableName() string {
	return "tree_nodes"
}
