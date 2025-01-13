package model

// MonitorSendGroup 发送组的配置
type MonitorSendGroup struct {
	ID            int    `json:"id" gorm:"primaryKey;autoIncrement;comment:发送组ID"`
	Name          string `json:"name" binding:"required,min=1,max=50" gorm:"uniqueIndex;size:100;comment:发送组英文名称，供AlertManager配置文件使用，支持通配符*进行模糊搜索"`
	NameZh        string `json:"nameZh" binding:"required,min=1,max=50" gorm:"uniqueIndex;size:100;comment:发送组中文名称，供告警规则选择发送组时使用，支持通配符*进行模糊搜索"`
	Enable        int    `json:"enable" gorm:"type:int;comment:是否启用发送组：1启用，2禁用"`
	UserID        int    `json:"userId" gorm:"comment:创建该发送组的用户ID"`
	PoolID        int    `json:"poolId" gorm:"comment:关联的AlertManager实例ID"`
	OnDutyGroupID int    `json:"onDutyGroupId" gorm:"comment:值班组ID"`
	// StaticReceiveUsers  []*User    `json:"staticReceiveUsers" gorm:"many2many:static_receive_users;comment:静态配置的接收人列表，多对多关系"`
	FeiShuQunRobotToken string     `json:"feiShuQunRobotToken,omitempty" gorm:"size:255;comment:飞书机器人Token，对应IM群"`
	RepeatInterval      string     `json:"repeatInterval,omitempty" gorm:"size:50;comment:默认重复发送时间"`
	SendResolved        int        `json:"sendResolved" gorm:"type:int;comment:是否发送恢复通知：1发送，2不发送"`
	NotifyMethods       StringList `json:"notifyMethods,omitempty" gorm:"type:text;comment:通知方法，如：email, im, phone, sms"`
	NeedUpgrade         int        `json:"needUpgrade" gorm:"type:int;comment:是否需要告警升级：1需要，2不需要"`
	// FirstUpgradeUsers   []*User    `json:"firstUpgradeUsers" gorm:"many2many:first_upgrade_users;comment:第一升级人列表，多对多关系"`
	UpgradeMinutes int `json:"upgradeMinutes,omitempty" gorm:"type:int;comment:告警多久未恢复则升级（分钟）"`
	// SecondUpgradeUsers  []*User    `json:"secondUpgradeUsers" gorm:"many2many:second_upgrade_users;comment:第二升级人列表，多对多关系"`
	CreateTime int64 `gorm:"column:create_time;type:int;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime int64 `gorm:"column:update_time;type:int;autoUpdateTime" json:"update_time"` // 更新时间
	IsDeleted  int   `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`    // 软删除标志（0:否, 1:是）

	// 前端使用字段
	TreeNodeIDs     []int    `json:"treeNodeIds,omitempty" gorm:"-"`
	FirstUserNames  []string `json:"firstUserNames,omitempty" gorm:"-"`
	Key             string   `json:"key" gorm:"-"`
	PoolName        string   `json:"poolName,omitempty" gorm:"-"`
	OnDutyGroupName string   `json:"onDutyGroupName,omitempty" gorm:"-"`
	CreateUserName  string   `json:"createUserName,omitempty" gorm:"-"`
}

func (MonitorSendGroup) TableName() string {
	return "monitor_send_group"
}
