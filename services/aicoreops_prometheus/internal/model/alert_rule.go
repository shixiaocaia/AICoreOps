package model

// MonitorAlertRule 告警规则的配置
type MonitorAlertRule struct {
	ID          int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:告警规则ID"`
	Name        string     `json:"name" binding:"required,min=1,max=50" gorm:"uniqueIndex;size:100;comment:告警规则名称，支持通配符*进行模糊搜索"`
	UserID      int64      `json:"userId" gorm:"comment:创建该告警规则的用户ID"`
	PoolID      int64      `json:"poolId" gorm:"comment:关联的Prometheus实例池ID"`
	SendGroupID int64      `json:"sendGroupId" gorm:"comment:关联的发送组ID"`
	TreeNodeID  int64      `json:"treeNodeId" gorm:"comment:绑定的树节点ID"`
	Enable      int32      `json:"enable" gorm:"type:int;comment:是否启用告警规则：1启用，2禁用"`
	Expr        string     `json:"expr" gorm:"type:text;comment:告警规则表达式"`
	Severity    string     `json:"severity,omitempty" gorm:"size:50;comment:告警级别，如critical、warning"`
	GrafanaLink string     `json:"grafanaLink,omitempty" gorm:"type:text;comment:Grafana大盘链接"`
	ForTime     string     `json:"forTime,omitempty" gorm:"size:50;comment:持续时间，达到此时间才触发告警"`
	Labels      StringList `json:"labels,omitempty" gorm:"type:text;comment:标签组，格式为 key=v"`
	Annotations StringList `json:"annotations,omitempty" gorm:"type:text;comment:注解，格式为 key=v"`
	CreateTime  int64      `gorm:"column:create_time;type:int;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime  int64      `gorm:"column:update_time;type:int;autoUpdateTime" json:"update_time"` // 更新时间
	IsDeleted   int32      `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`    // 软删除标志（0:否, 1:是）

	// 前端使用字段
	NodePath       string  `json:"nodePath,omitempty" gorm:"-"`
	TreeNodeIDs    []int64 `json:"treeNodeIds,omitempty" gorm:"-"`
	Key            string  `json:"key" gorm:"-"`
	PoolName       string  `json:"poolName,omitempty" gorm:"-"`
	SendGroupName  string  `json:"sendGroupName,omitempty" gorm:"-"`
	CreateUserName string  `json:"createUserName,omitempty" gorm:"-"`
	LabelsFront    string  `json:"labelsFront,omitempty" gorm:"-"`
}

func (MonitorAlertRule) TableName() string {
	return "monitor_alert_rule"
}
