package model

// MonitorRecordRule 记录规则的配置
type MonitorRecordRule struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:记录规则ID"`
	Name        string `json:"name" binding:"required,min=1,max=50" gorm:"uniqueIndex;size:100;comment:记录规则名称，支持使用通配符*进行模糊搜索"`
	RecordName  string `json:"recordName" binding:"required,min=1,max=500" gorm:"uniqueIndex;size:500;comment:记录名称，支持使用通配符*进行模糊搜索"`
	UserID      int64  `json:"userId" gorm:"comment:创建该记录规则的用户ID"`
	PoolID      int64  `json:"poolId" gorm:"comment:关联的Prometheus实例池ID"`
	TreeNodeID  int64  `json:"treeNodeId" gorm:"comment:绑定的树节点ID"`
	Enable      int32  `json:"enable" gorm:"type:int;comment:是否启用记录规则：1启用，2禁用"`
	ForDuration string `json:"forDuration,omitempty" gorm:"size:50;comment:持续时间，达到此时间才触发记录规则"`
	Expr        string `json:"expr" gorm:"type:text;comment:记录规则表达式"`
	CreateTime  int64  `gorm:"column:create_time;type:int;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime  int64  `gorm:"column:update_time;type:int;autoUpdateTime" json:"update_time"` // 更新时间
	IsDeleted   int32  `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`    // 软删除标志（0:否, 1:是）

	// 前端使用字段
	NodePath         string            `json:"nodePath,omitempty" gorm:"-"`
	TreeNodeIDs      []int             `json:"treeNodeIds,omitempty" gorm:"-"`
	Key              string            `json:"key" gorm:"-"`
	PoolName         string            `json:"poolName,omitempty" gorm:"-"`
	SendGroupName    string            `json:"sendGroupName,omitempty" gorm:"-"`
	CreateUserName   string            `json:"createUserName,omitempty" gorm:"-"`
	LabelsFront      string            `json:"labelsFront,omitempty" gorm:"-"`
	AnnotationsFront string            `json:"annotationsFront,omitempty" gorm:"-"`
	LabelsM          map[string]string `json:"labelsM,omitempty" gorm:"-"`
	AnnotationsM     map[string]string `json:"annotationsM,omitempty" gorm:"-"`
}

func (MonitorRecordRule) TableName() string {
	return "monitor_record_rule"
}
