package dao

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type AlertRuleDao struct {
	db *gorm.DB
}

func NewAlertRuleDao(db *gorm.DB) *AlertRuleDao {
	return &AlertRuleDao{
		db: db,
	}
}

// GetMonitorAlertRuleByPoolId 根据池ID获取告警规则列表
func (d *AlertRuleDao) GetMonitorAlertRuleByPoolId(ctx context.Context, poolId int) ([]*model.MonitorAlertRule, error) {
	var rules []*model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("pool_id = ? AND is_deleted = 0", poolId).Find(&rules).Error
	return rules, err
}

// SearchMonitorAlertRuleByName 根据名称搜索告警规则
func (d *AlertRuleDao) SearchMonitorAlertRuleByName(ctx context.Context, name string) ([]*model.MonitorAlertRule, error) {
	var rules []*model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("name LIKE ? AND is_deleted = 0", "%"+name+"%").Find(&rules).Error
	return rules, err
}

// GetMonitorAlertRuleList 获取所有告警规则列表
func (d *AlertRuleDao) GetMonitorAlertRuleList(ctx context.Context) ([]*model.MonitorAlertRule, error) {
	var rules []*model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("is_deleted = 0").Find(&rules).Error
	return rules, err
}

// CreateMonitorAlertRule 创建告警规则
func (d *AlertRuleDao) CreateMonitorAlertRule(ctx context.Context, monitorAlertRule *model.MonitorAlertRule) error {
	return d.db.WithContext(ctx).Create(monitorAlertRule).Error
}

// GetMonitorAlertRuleById 根据ID获取告警规则
func (d *AlertRuleDao) GetMonitorAlertRuleById(ctx context.Context, id int) (*model.MonitorAlertRule, error) {
	var rule model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("id = ? AND is_deleted = 0", id).First(&rule).Error
	return &rule, err
}

// UpdateMonitorAlertRule 更新告警规则
func (d *AlertRuleDao) UpdateMonitorAlertRule(ctx context.Context, monitorAlertRule *model.MonitorAlertRule) error {
	return d.db.WithContext(ctx).Save(monitorAlertRule).Error
}

// EnableSwitchMonitorAlertRule 启用/禁用告警规则
func (d *AlertRuleDao) EnableSwitchMonitorAlertRule(ctx context.Context, ruleID int) error {
	return d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).Where("id = ?", ruleID).
		Update("enable", gorm.Expr("CASE WHEN enable = 1 THEN 2 ELSE 1 END")).Error
}

// BatchEnableSwitchMonitorAlertRule 批量启用告警规则
func (d *AlertRuleDao) BatchEnableSwitchMonitorAlertRule(ctx context.Context, ruleIDs []int) error {
	return d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).Where("id IN ?", ruleIDs).
		Update("enable", 1).Error
}

// DeleteMonitorAlertRule 删除告警规则（软删除）
func (d *AlertRuleDao) DeleteMonitorAlertRule(ctx context.Context, ruleID int) error {
	return d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).Where("id = ?", ruleID).
		Update("is_deleted", 1).Error
}

// GetAssociatedResourcesBySendGroupId 获取发送组关联的告警规则
func (d *AlertRuleDao) GetAssociatedResourcesBySendGroupId(ctx context.Context, sendGroupId int) ([]*model.MonitorAlertRule, error) {
	var rules []*model.MonitorAlertRule
	err := d.db.WithContext(ctx).Where("send_group_id = ? AND is_deleted = 0", sendGroupId).Find(&rules).Error
	return rules, err
}

// CheckMonitorAlertRuleExists 检查告警规则是否存在
func (d *AlertRuleDao) CheckMonitorAlertRuleExists(ctx context.Context, alertRule *model.MonitorAlertRule) (bool, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).
		Where("id = ? AND is_deleted = 0", alertRule.ID).Count(&count).Error
	return count > 0, err
}

// CheckMonitorAlertRuleNameExists 检查告警规则名称是否存在
func (d *AlertRuleDao) CheckMonitorAlertRuleNameExists(ctx context.Context, alertRule *model.MonitorAlertRule) (bool, error) {
	var count int64
	query := d.db.WithContext(ctx).Model(&model.MonitorAlertRule{}).
		Where("name = ? AND is_deleted = 0", alertRule.Name)
	if alertRule.ID > 0 {
		query = query.Where("id != ?", alertRule.ID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}

// ExampleAlertRules 告警规则示例
func (d *AlertRuleDao) ExampleAlertRules() []*model.MonitorAlertRule {
	return []*model.MonitorAlertRule{
		{
			Name:        "HighCPUUsage",
			Expr:        "100 - (avg by(instance) (rate(node_cpu_seconds_total{mode='idle'}[5m])) * 100) > 80",
			ForTime:     "5m",
			Severity:    "warning",
			Enable:      1,
			Labels:      []string{"severity=warning", "type=resource"},
			Annotations: []string{"summary=CPU使用率过高", "description=实例 {{ $labels.instance }} CPU使用率超过80%，当前值: {{ $value }}%"},
		},
		{
			Name:        "HighMemoryUsage",
			Expr:        "(node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes) / node_memory_MemTotal_bytes * 100 > 85",
			ForTime:     "5m",
			Severity:    "warning",
			Enable:      1,
			Labels:      []string{"severity=warning", "type=resource"},
			Annotations: []string{"summary=内存使用率过高", "description=实例 {{ $labels.instance }} 内存使用率超过85%，当前值: {{ $value }}%"},
		},
		{
			Name:        "DiskSpaceRunningOut",
			Expr:        "100 - ((node_filesystem_avail_bytes{mountpoint='/'} * 100) / node_filesystem_size_bytes{mountpoint='/'}) > 85",
			ForTime:     "30m",
			Severity:    "critical",
			Enable:      1,
			Labels:      []string{"severity=critical", "type=resource"},
			Annotations: []string{"summary=磁盘空间不足", "description=实例 {{ $labels.instance }} 根分区使用率超过85%，当前值: {{ $value }}%"},
		},
		{
			Name:        "InstanceDown",
			Expr:        "up == 0",
			ForTime:     "5m",
			Severity:    "critical",
			Enable:      1,
			Labels:      []string{"severity=critical", "type=availability"},
			Annotations: []string{"summary=实例不可用", "description=实例 {{ $labels.instance }} 已经下线超过5分钟"},
		},
		{
			Name:        "HighLatency",
			Expr:        "histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le, instance)) > 1",
			ForTime:     "5m",
			Severity:    "warning",
			Enable:      1,
			Labels:      []string{"severity=warning", "type=performance"},
			Annotations: []string{"summary=接口延迟过高", "description=实例 {{ $labels.instance }} 的95%请求延迟超过1秒，当前值: {{ $value }}秒"},
		},
		{
			Name:        "HighErrorRate",
			Expr:        "sum(rate(http_requests_total{status=~'5..'}[5m])) by (instance) / sum(rate(http_requests_total[5m])) by (instance) * 100 > 5",
			ForTime:     "5m",
			Severity:    "critical",
			Enable:      1,
			Labels:      []string{"severity=critical", "type=reliability"},
			Annotations: []string{"summary=错误率过高", "description=实例 {{ $labels.instance }} 的HTTP 5xx错误率超过5%，当前值: {{ $value }}%"},
		},
		{
			Name:        "ContainerRestartFrequently",
			Expr:        "changes(kube_pod_container_status_restarts_total[1h]) > 3",
			ForTime:     "10m",
			Severity:    "warning",
			Enable:      1,
			Labels:      []string{"severity=warning", "type=availability"},
			Annotations: []string{"summary=容器频繁重启", "description=容器 {{ $labels.container }} 在过去1小时内重启超过3次"},
		},
		{
			Name:        "PodCrashLoopBackOff",
			Expr:        "kube_pod_container_status_waiting_reason{reason='CrashLoopBackOff'} == 1",
			ForTime:     "10m",
			Severity:    "critical",
			Enable:      1,
			Labels:      []string{"severity=critical", "type=availability"},
			Annotations: []string{"summary=Pod崩溃循环", "description=Pod {{ $labels.pod }} 处于CrashLoopBackOff状态"},
		},
		{
			Name:        "SlowQueries",
			Expr:        "rate(mysql_slow_queries[5m]) > 10",
			ForTime:     "5m",
			Severity:    "warning",
			Enable:      1,
			Labels:      []string{"severity=warning", "type=performance"},
			Annotations: []string{"summary=慢查询过多", "description=数据库实例 {{ $labels.instance }} 的慢查询数超过每秒10次，当前值: {{ $value }}"},
		},
		{
			Name:        "RedisHighMemoryUsage",
			Expr:        "redis_memory_used_bytes / redis_memory_max_bytes * 100 > 80",
			ForTime:     "5m",
			Severity:    "warning",
			Enable:      1,
			Labels:      []string{"severity=warning", "type=resource"},
			Annotations: []string{"summary=Redis内存使用率过高", "description=Redis实例 {{ $labels.instance }} 内存使用率超过80%，当前值: {{ $value }}%"},
		},
	}
}
