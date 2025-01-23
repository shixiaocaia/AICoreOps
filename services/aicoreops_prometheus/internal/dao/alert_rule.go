package dao

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type MonitorAlertRuleDAO struct {
	db *gorm.DB
}

func NewMonitorAlertRuleDAO(db *gorm.DB) *MonitorAlertRuleDAO {
	return &MonitorAlertRuleDAO{
		db: db,
	}
}

// GetMonitorAlertRuleByPoolId 根据池ID获取告警规则列表
func (d *MonitorAlertRuleDAO) GetMonitorAlertRuleByPoolId(ctx context.Context, poolId int) ([]*model.MonitorAlertRule, error) {
	var rules []*model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("pool_id = ? AND is_deleted = 0", poolId).Find(&rules).Error
	return rules, err
}

// SearchMonitorAlertRuleByName 根据名称搜索告警规则
func (d *MonitorAlertRuleDAO) SearchMonitorAlertRuleByName(ctx context.Context, name string) ([]*model.MonitorAlertRule, error) {
	var rules []*model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("name LIKE ? AND is_deleted = 0", "%"+name+"%").Find(&rules).Error
	return rules, err
}

// GetMonitorAlertRuleList 获取所有告警规则列表
func (d *MonitorAlertRuleDAO) GetMonitorAlertRuleList(ctx context.Context) ([]*model.MonitorAlertRule, error) {
	var rules []*model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("is_deleted = 0").Find(&rules).Error
	return rules, err
}

// CreateMonitorAlertRule 创建告警规则
func (d *MonitorAlertRuleDAO) CreateMonitorAlertRule(ctx context.Context, monitorAlertRule *model.MonitorAlertRule) error {
	return d.db.WithContext(ctx).Create(monitorAlertRule).Error
}

// GetMonitorAlertRuleById 根据ID获取告警规则
func (d *MonitorAlertRuleDAO) GetMonitorAlertRuleById(ctx context.Context, id int) (*model.MonitorAlertRule, error) {
	var rule model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", id).First(&rule).Error
	return &rule, err
}

// UpdateMonitorAlertRule 更新告警规则
func (d *MonitorAlertRuleDAO) UpdateMonitorAlertRule(ctx context.Context, monitorAlertRule *model.MonitorAlertRule) error {
	return d.db.WithContext(ctx).Save(monitorAlertRule).Error
}

// EnableSwitchMonitorAlertRule 启用/禁用告警规则
func (d *MonitorAlertRuleDAO) EnableSwitchMonitorAlertRule(ctx context.Context, ruleID int) error {
	return d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).Where("id = ?", ruleID).
		Update("enable", gorm.Expr("CASE WHEN enable = 1 THEN 2 ELSE 1 END")).Error
}

// BatchEnableSwitchMonitorAlertRule 批量启用告警规则
func (d *MonitorAlertRuleDAO) BatchEnableSwitchMonitorAlertRule(ctx context.Context, ruleIDs []int) error {
	return d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).Where("id IN ?", ruleIDs).
		Update("enable", 1).Error
}

// DeleteMonitorAlertRule 删除告警规则（软删除）
func (d *MonitorAlertRuleDAO) DeleteMonitorAlertRule(ctx context.Context, ruleID int) error {
	return d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).Where("id = ?", ruleID).
		Update("is_deleted", 1).Error
}

// GetAssociatedResourcesBySendGroupId 获取发送组关联的告警规则
func (d *MonitorAlertRuleDAO) GetAssociatedResourcesBySendGroupId(ctx context.Context, sendGroupId int) ([]*model.MonitorAlertRule, error) {
	var rules []*model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("send_group_id = ? AND is_deleted = 0", sendGroupId).Find(&rules).Error
	return rules, err
}

// CheckMonitorAlertRuleExists 检查告警规则是否存在
func (d *MonitorAlertRuleDAO) CheckMonitorAlertRuleExists(ctx context.Context, alertRule *model.MonitorAlertRule) (bool, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).
		Where("id = ? AND is_deleted = 0", alertRule.ID).Count(&count).Error
	return count > 0, err
}

// CheckMonitorAlertRuleNameExists 检查告警规则名称是否存在
func (d *MonitorAlertRuleDAO) CheckMonitorAlertRuleNameExists(ctx context.Context, alertRule *model.MonitorAlertRule) (bool, error) {
	var count int64
	query := d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).
		Where("name = ? AND is_deleted = 0", alertRule.Name)
	if alertRule.ID > 0 {
		query = query.Where("id != ?", alertRule.ID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}
