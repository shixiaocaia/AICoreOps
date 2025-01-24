package model

import "github.com/prometheus/alertmanager/template"

// MonitorAlertEvent 告警事件与相关实体的关系
type MonitorAlertEvent struct {
	ID            int        `json:"id" gorm:"primaryKey;autoIncrement;comment:告警事件ID"`
	AlertName     string     `json:"alertName" binding:"required,min=1,max=200" gorm:"size:200;comment:告警名称"`
	Fingerprint   string     `json:"fingerprint" binding:"required,min=1,max=50" gorm:"uniqueIndex;size:100;comment:告警唯一ID"`
	Status        string     `json:"status,omitempty" gorm:"size:50;comment:告警状态（如告警中、已屏蔽、已认领、已恢复）"`
	RuleID        int        `json:"ruleId" gorm:"comment:关联的告警规则ID"`
	SendGroupID   int        `json:"sendGroupId" gorm:"comment:关联的发送组ID"`
	EventTimes    int        `json:"eventTimes" gorm:"comment:触发次数"`
	SilenceID     string     `json:"silenceId,omitempty" gorm:"size:100;comment:AlertManager返回的静默ID"`
	RenLingUserID int        `json:"renLingUserId" gorm:"comment:认领告警的用户ID"`
	Labels        StringList `json:"labels,omitempty" gorm:"type:text;comment:标签组，格式为 key=v"`
	CreateTime    int64      `gorm:"column:create_time;type:int;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime    int64      `gorm:"column:update_time;type:int;autoUpdateTime" json:"update_time"` // 更新时间
	IsDeleted     int        `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`    // 软删除标志（0:否, 1:是）

	// 前端使用字段
	Key           string            `json:"key" gorm:"-"`
	AlertRuleName string            `json:"alertRuleName,omitempty" gorm:"-"`
	SendGroupName string            `json:"sendGroupName,omitempty" gorm:"-"`
	Alert         template.Alert    `json:"alert,omitempty" gorm:"-"`
	SendGroup     *MonitorSendGroup `json:"sendGroup,omitempty" gorm:"-"`
	// RenLingUser   *User             `json:"renLingUser,omitempty" gorm:"-"`
	Rule          *AlertRule        `json:"rule,omitempty" gorm:"-"`
	LabelsMatcher map[string]string `json:"labelsMatcher,omitempty" gorm:"-"`
	// AnnotationsMatcher map[string]string `json:"annotationsMatcher,omitempty" gorm:"-"`
}

func (MonitorAlertEvent) TableName() string {
	return "monitor_alert_event"
}
