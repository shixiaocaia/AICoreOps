package dao

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type AlertRuleDAO struct {
	db *gorm.DB
}

func NewAlertRuleDAO(db *gorm.DB) *AlertRuleDAO {
	return &AlertRuleDAO{
		db: db,
	}
}

// GetAlertRuleByPoolId 根据池ID获取告警规则列表
func (d *AlertRuleDAO) GetAlertRuleByPoolId(ctx context.Context, poolId int64) ([]*model.AlertRule, error) {
	var rules []*model.AlertRule
	err := d.db.WithContext(ctx).Where("pool_id = ? AND is_deleted = 0", poolId).Find(&rules).Error
	return rules, err
}

// SearchAlertRuleByName 根据名称搜索告警规则
func (d *AlertRuleDAO) SearchAlertRuleByName(ctx context.Context, name string) ([]*model.AlertRule, error) {
	var rules []*model.AlertRule
	err := d.db.WithContext(ctx).Where("name LIKE ? AND is_deleted = 0", "%"+name+"%").Find(&rules).Error
	return rules, err
}

// GetAlertRuleList 获取所有告警规则列表
func (d *AlertRuleDAO) GetAlertRuleList(ctx context.Context) ([]*model.AlertRule, error) {
	var rules []*model.AlertRule
	err := d.db.WithContext(ctx).Where("is_deleted = 0").Find(&rules).Error
	return rules, err
}

// CreateAlertRule 创建告警规则
func (d *AlertRuleDAO) CreateAlertRule(ctx context.Context, monitorAlertRule *model.AlertRule) error {
	return d.db.WithContext(ctx).Create(monitorAlertRule).Error
}

// GetAlertRuleById 根据ID获取告警规则
func (d *AlertRuleDAO) GetAlertRuleById(ctx context.Context, id int64) (*model.AlertRule, error) {
	var rule model.AlertRule
	err := d.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", id).First(&rule).Error
	return &rule, err
}

// UpdateAlertRule 更新告警规则
func (d *AlertRuleDAO) UpdateAlertRule(ctx context.Context, monitorAlertRule *model.AlertRule) error {
	return d.db.WithContext(ctx).Save(monitorAlertRule).Error
}

// EnableSwitchAlertRule 启用/禁用告警规则
func (d *AlertRuleDAO) EnableSwitchAlertRule(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Model(&model.AlertRule{}).Where("id = ?", id).
		Update("enable", gorm.Expr("CASE WHEN enable = 1 THEN 2 ELSE 1 END")).Error
}

// BatchEnableSwitchAlertRule 批量启用告警规则
func (d *AlertRuleDAO) BatchEnableSwitchAlertRule(ctx context.Context, ids []int64) error {
	return d.db.WithContext(ctx).Model(&model.AlertRule{}).Where("id IN ?", ids).
		Update("enable", 1).Error
}

// DeleteAlertRule 删除告警规则（软删除）
func (d *AlertRuleDAO) DeleteAlertRule(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Model(&model.AlertRule{}).Where("id = ?", id).
		Update("is_deleted", 1).Error
}

// BatchDeleteAlertRule 批量删除告警规则
func (d *AlertRuleDAO) BatchDeleteAlertRule(ctx context.Context, ids []int64) error {
	return d.db.WithContext(ctx).Model(&model.AlertRule{}).Where("id IN ?", ids).
		Update("is_deleted", 1).Error
}

// GetAssociatedResourcesBySendGroupId 获取发送组关联的告警规则
func (d *AlertRuleDAO) GetAssociatedResourcesBySendGroupId(ctx context.Context, sendGroupId int64) ([]*model.AlertRule, error) {
	var rules []*model.AlertRule
	err := d.db.WithContext(ctx).Where("send_group_id = ? AND is_deleted = 0", sendGroupId).Find(&rules).Error
	return rules, err
}

// CheckAlertRuleExists 检查告警规则是否存在
func (d *AlertRuleDAO) CheckAlertRuleExists(ctx context.Context, id int64) (bool, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.AlertRule{}).
		Where("id = ? AND is_deleted = 0", id).Count(&count).Error
	return count > 0, err
}

// CheckAlertRuleNameExists 检查告警规则名称是否存在
func (d *AlertRuleDAO) CheckAlertRuleNameExists(ctx context.Context, name string) (bool, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.AlertRule{}).
		Where("name = ? AND is_deleted = 0", name).Count(&count).Error
	return count > 0, err
}
