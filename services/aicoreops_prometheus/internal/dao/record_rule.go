package dao

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type MonitorRecordRuleDAO struct {
	db *gorm.DB
}

func NewMonitorRecordRuleDAO(db *gorm.DB) *MonitorRecordRuleDAO {
	return &MonitorRecordRuleDAO{
		db: db,
	}
}

// GetMonitorRecordRuleByPoolId 根据池ID获取记录规则列表
func (d *MonitorRecordRuleDAO) GetMonitorRecordRuleByPoolId(ctx context.Context, poolId int64) ([]*model.MonitorRecordRule, error) {
	var rules []*model.MonitorRecordRule
	err := d.db.WithContext(ctx).Where("pool_id = ? AND is_deleted = 0", poolId).Find(&rules).Error
	return rules, err
}

// SearchMonitorRecordRuleByName 根据名称搜索记录规则
func (d *MonitorRecordRuleDAO) SearchMonitorRecordRuleByName(ctx context.Context, name string) ([]*model.MonitorRecordRule, error) {
	var rules []*model.MonitorRecordRule
	err := d.db.WithContext(ctx).Where("name LIKE ? AND is_deleted = 0", "%"+name+"%").Find(&rules).Error
	return rules, err
}

// GetMonitorRecordRuleList 获取所有记录规则列表
func (d *MonitorRecordRuleDAO) GetMonitorRecordRuleList(ctx context.Context) ([]*model.MonitorRecordRule, error) {
	var rules []*model.MonitorRecordRule
	err := d.db.WithContext(ctx).Where("is_deleted = 0").Find(&rules).Error
	return rules, err
}

// CreateMonitorRecordRule 创建记录规则
func (d *MonitorRecordRuleDAO) CreateMonitorRecordRule(ctx context.Context, recordRule *model.MonitorRecordRule) error {
	return d.db.WithContext(ctx).Create(recordRule).Error
}

// GetMonitorRecordRuleById 根据ID获取记录规则
func (d *MonitorRecordRuleDAO) GetMonitorRecordRuleById(ctx context.Context, id int64) (*model.MonitorRecordRule, error) {
	var rule model.MonitorRecordRule
	err := d.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", id).First(&rule).Error
	return &rule, err
}

// UpdateMonitorRecordRule 更新记录规则
func (d *MonitorRecordRuleDAO) UpdateMonitorRecordRule(ctx context.Context, recordRule *model.MonitorRecordRule) error {
	return d.db.WithContext(ctx).Save(recordRule).Error
}

// DeleteMonitorRecordRule 删除记录规则（软删除）
func (d *MonitorRecordRuleDAO) DeleteMonitorRecordRule(ctx context.Context, ruleID int64) error {
	return d.db.WithContext(ctx).Model(&model.MonitorRecordRule{}).Where("id = ?", ruleID).
		Update("is_deleted", 1).Error
}

// BatchDeleteMonitorRecordRule 批量删除记录规则
func (d *MonitorRecordRuleDAO) BatchDeleteMonitorRecordRule(ctx context.Context, ruleIDs []int64) error {
	return d.db.WithContext(ctx).Model(&model.MonitorRecordRule{}).Where("id IN (?)", ruleIDs).
		Update("is_deleted", 1).Error
}

// EnableSwitchMonitorRecordRule 启用/禁用记录规则
func (d *MonitorRecordRuleDAO) EnableSwitchMonitorRecordRule(ctx context.Context, ruleID int64) error {
	return d.db.WithContext(ctx).Model(&model.MonitorRecordRule{}).Where("id = ?", ruleID).
		Update("enable", gorm.Expr("CASE WHEN enable = 1 THEN 2 ELSE 1 END")).Error
}

// BatchEnableSwitchMonitorRecordRule 批量启用/禁用记录规则
func (d *MonitorRecordRuleDAO) BatchEnableSwitchMonitorRecordRule(ctx context.Context, ruleIDs []int64) error {
	return d.db.WithContext(ctx).Model(&model.MonitorRecordRule{}).Where("id IN (?)", ruleIDs).
		Update("enable", gorm.Expr("CASE WHEN enable = 1 THEN 2 ELSE 1 END")).Error
}

// CheckMonitorRecordRuleExists 检查记录规则是否存在
func (d *MonitorRecordRuleDAO) CheckMonitorRecordRuleExists(ctx context.Context, recordRule *model.MonitorRecordRule) (bool, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.MonitorRecordRule{}).
		Where("id = ? AND is_deleted = 0", recordRule.ID).Count(&count).Error
	return count > 0, err
}

// CheckMonitorRecordRuleNameExists 检查记录规则名称是否存在
func (d *MonitorRecordRuleDAO) CheckMonitorRecordRuleNameExists(ctx context.Context, recordRule *model.MonitorRecordRule) (bool, error) {
	var count int64
	query := d.db.WithContext(ctx).Model(&model.MonitorRecordRule{}).
		Where("name = ? AND is_deleted = 0", recordRule.Name)
	if recordRule.ID > 0 {
		query = query.Where("id != ?", recordRule.ID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}

// ExampleRecordRules 记录规则示例
func (d *MonitorRecordRuleDAO) ExampleRecordRules() []*model.MonitorRecordRule {
	return []*model.MonitorRecordRule{
		{
			Name:        "node_memory_usage_bytes",
			RecordName:  "node:memory:usage:bytes",
			Expr:        "node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "node_cpu_usage_percent",
			RecordName:  "node:cpu:usage:percent",
			Expr:        "100 - (avg by(instance) (rate(node_cpu_seconds_total{mode='idle'}[5m])) * 100)",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "node_disk_usage_percent",
			RecordName:  "node:disk:usage:percent",
			Expr:        "100 - ((node_filesystem_avail_bytes{mountpoint='/'} * 100) / node_filesystem_size_bytes{mountpoint='/'})",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "http_request_total_5m",
			RecordName:  "http:request:total:5m",
			Expr:        "sum(rate(http_requests_total[5m])) by (instance)",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "http_request_errors_5m",
			RecordName:  "http:request:errors:5m",
			Expr:        "sum(rate(http_requests_total{status=~'5..'}[5m])) by (instance)",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "container_memory_usage_bytes",
			RecordName:  "container:memory:usage:bytes",
			Expr:        "sum(container_memory_usage_bytes) by (container)",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "container_cpu_usage_seconds",
			RecordName:  "container:cpu:usage:seconds",
			Expr:        "sum(rate(container_cpu_usage_seconds_total[5m])) by (container)",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "mysql_queries_rate",
			RecordName:  "mysql:queries:rate",
			Expr:        "rate(mysql_global_status_queries[5m])",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "redis_connected_clients",
			RecordName:  "redis:connected:clients",
			Expr:        "redis_connected_clients",
			ForDuration: "5m",
			Enable:      1,
		},
		{
			Name:        "redis_commands_per_second",
			RecordName:  "redis:commands:per:second",
			Expr:        "rate(redis_commands_processed_total[5m])",
			ForDuration: "5m",
			Enable:      1,
		},
	}
}
