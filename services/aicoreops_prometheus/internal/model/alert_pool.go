package model

// MonitorAlertManagerPool 采集池
type MonitorAlertManagerPool struct {
	ID                    int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	Name                  string     `json:"name" binding:"required,min=1,max=50" gorm:"uniqueIndex;size:100;comment:AlertManager实例名称，支持使用通配符*进行模糊搜索"`
	AlertManagerInstances StringList `json:"alertManagerInstances" gorm:"type:text;comment:选择多个AlertManager实例"`
	UserID                int64      `json:"userId" gorm:"comment:创建该实例池的用户ID"`
	ResolveTimeout        string     `json:"resolveTimeout,omitempty" gorm:"size:50;comment:默认恢复时间"`
	GroupWait             string     `json:"groupWait,omitempty" gorm:"size:50;comment:默认分组第一次等待时间"`
	GroupInterval         string     `json:"groupInterval,omitempty" gorm:"size:50;comment:默认分组等待间隔"`
	RepeatInterval        string     `json:"repeatInterval,omitempty" gorm:"size:50;comment:默认重复发送时间"`
	GroupBy               StringList `json:"groupBy,omitempty" gorm:"type:text;comment:分组的标签"`
	Receiver              string     `json:"receiver,omitempty" gorm:"size:100;comment:兜底接收者"`
	CreateTime            int64      `gorm:"column:create_time;type:int;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime            int64      `gorm:"column:update_time;type:int;autoUpdateTime" json:"update_time"` // 更新时间
	IsDeleted             int        `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`    // 软删除标志（0:否, 1:是）

	// 前端使用字段
	GroupByFront   string `json:"groupByFront,omitempty" gorm:"-"`
	Key            string `json:"key" gorm:"-"`
	CreateUserName string `json:"createUserName,omitempty" gorm:"-"`
}

func (*MonitorAlertManagerPool) TableName() string {
	return "monitor_alertmanager_pool"
}
