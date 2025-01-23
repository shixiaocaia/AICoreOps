package cache

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/config"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/pkg"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/repo"
	pm "github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// RuleConfigCache 告警规则缓存接口
type RuleConfigCache interface {
	// GetPrometheusAlertRuleConfigYamlByIp 根据IP地址获取Prometheus的告警规则配置YAML
	GetPrometheusAlertRuleConfigYamlByIp(ip string) string
	// GenerateAlertRuleConfigYaml 生成并更新所有Prometheus的告警规则配置YAML
	GenerateAlertRuleConfigYaml(ctx context.Context) error
	// GeneratePrometheusAlertRuleConfigYamlOnePool 根据单个采集池生成Prometheus的告警规则配置YAML
	GeneratePrometheusAlertRuleConfigYamlOnePool(ctx context.Context, pool *model.MonitorScrapePool) map[string]string
}

type ruleConfigCache struct {
	logx.Logger
	mu             sync.RWMutex      // 读写锁，保护缓存数据
	AlertRuleMap   map[string]string // 存储告警规则
	localYamlDir   string            // 本地YAML目录
	scrapePoolRepo repo.MonitorScrapePoolRepo
	alertRuleRepo  repo.RuleRepo
}

// RuleGroups 告警规则组
type RuleGroups struct {
	Groups []RuleGroup `yaml:"groups"`
}

// RuleGroup 单个告警规则组
type RuleGroup struct {
	Name  string         `yaml:"name"`
	Rules []rulefmt.Rule `yaml:"rules"`
}

// NewRuleConfigCache 创建新的告警规则缓存实例
func NewRuleConfigCache(ctx context.Context, db *gorm.DB, config *config.Config) RuleConfigCache {
	return &ruleConfigCache{
		Logger:         logx.WithContext(ctx),
		mu:             sync.RWMutex{},
		localYamlDir:   config.AlertManagerConfig.LocalYamlDir,
		AlertRuleMap:   make(map[string]string),
		scrapePoolRepo: dao.NewMonitorScrapePoolDAO(db),
		alertRuleRepo:  dao.NewMonitorAlertRuleDAO(db),
	}
}

// GetPrometheusAlertRuleConfigYamlByIp 根据IP地址获取Prometheus的告警规则配置YAML
func (r *ruleConfigCache) GetPrometheusAlertRuleConfigYamlByIp(ip string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.AlertRuleMap[ip]
}

func (r *ruleConfigCache) GenerateAlertRuleConfigYaml(ctx context.Context) error {
	// 获取支持告警配置的所有采集池
	pools, err := r.scrapePoolRepo.GetMonitorScrapePoolList(ctx)
	if err != nil {
		r.Logger.Error("[监控模块] 获取支持告警的采集池失败: %v", err)
		return err
	}

	if len(pools) == 0 {
		r.Logger.Info("没有找到支持告警的采集池")
		return nil
	}

	// 过滤出支持告警的采集池
	ruleConfigMap := r.AlertRuleMap
	for _, pool := range pools {
		if pool.SupportAlert == 1 { // 1表示支持告警
			oneMap := r.GeneratePrometheusAlertRuleConfigYamlOnePool(ctx, pool)
			for ip, out := range oneMap {
				ruleConfigMap[ip] = out
			}
		}
	}

	r.mu.Lock()
	r.AlertRuleMap = ruleConfigMap
	r.mu.Unlock()

	return nil
}

func (r *ruleConfigCache) GeneratePrometheusAlertRuleConfigYamlOnePool(ctx context.Context, pool *model.MonitorScrapePool) map[string]string {
	rules, err := r.alertRuleRepo.GetMonitorAlertRuleByPoolId(ctx, int(pool.ID))
	if err != nil {
		r.Logger.Errorf("[监控模块] 根据采集池ID [%d] 获取告警规则失败: %v", pool.ID, err)
		return nil
	}
	if len(rules) == 0 {
		r.Logger.Infof("[监控模块] 采集池 [%s] 没有告警规则", pool.Name)
		return nil
	}

	numInstances := len(pool.PrometheusInstances)
	if numInstances == 0 {
		r.Logger.Infof("[监控模块] 采集池 [%s] 中没有Prometheus实例", pool.Name)
		return nil
	}

	var ruleGroups RuleGroups

	// 构建规则组
	for _, rule := range rules {
		ft, err := pm.ParseDuration(rule.ForTime)
		if err != nil {
			r.Logger.Errorf("[监控模块] 解析告警规则持续时间失败，使用默认值: %v", err)
			ft, _ = pm.ParseDuration("5s")
		}
		labels := pkg.FromSliceTuMap(rule.Labels)
		annotations := pkg.FromSliceTuMap(rule.Annotations)

		oneRule := rulefmt.Rule{
			Alert:       rule.Name,   // 告警名称
			Expr:        rule.Expr,   // 告警表达式
			For:         ft,          // 持续时间
			Labels:      labels,      // 标签组
			Annotations: annotations, // 注解组
		}

		ruleGroup := RuleGroup{
			Name:  rule.Name,
			Rules: []rulefmt.Rule{oneRule}, // 一个规则组可以包含多个规则
		}
		ruleGroups.Groups = append(ruleGroups.Groups, ruleGroup)
	}

	ruleMap := make(map[string]string)

	// 分片逻辑，将规则分配给不同的Prometheus实例，以减少服务器的负载
	for i, ip := range pool.PrometheusInstances {
		var myRuleGroups RuleGroups

		for j, group := range ruleGroups.Groups {
			// 按顺序平均分片
			if j%numInstances == i {
				myRuleGroups.Groups = append(myRuleGroups.Groups, group)
			}
		}

		// 序列化规则组为YAML
		yamlData, err := yaml.Marshal(&myRuleGroups)
		if err != nil {
			r.Logger.Errorf("[监控模块] 序列化告警规则YAML失败: %v", err)
			continue
		}

		fileName := fmt.Sprintf("%s/prometheus_rule_%s_%s.yml",
			r.localYamlDir,
			pool.Name,
			ip,
		)
		if err := os.WriteFile(fileName, yamlData, 0644); err != nil {
			r.Logger.Errorf("[监控模块] 写入告警规则文件失败: %v", err)
			continue
		}

		ruleMap[ip] = string(yamlData)
	}

	r.Logger.Infof("[监控模块] 生成告警规则文件成功，共生成 %d 个文件", len(ruleMap))

	return ruleMap
}
