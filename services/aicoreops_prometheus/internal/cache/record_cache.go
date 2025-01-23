package cache

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/config"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/repo"
	pm "github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

type RecordConfigCache interface {
	// GetPrometheusRecordRuleConfigYamlByIp 根据IP地址获取Prometheus的预聚合规则配置YAML
	GetPrometheusRecordRuleConfigYamlByIp(ip string) string
	// GenerateRecordRuleConfigYaml 生成并更新所有Prometheus的预聚合规则配置YAML
	GenerateRecordRuleConfigYaml(ctx context.Context) error
	// GeneratePrometheusRecordRuleConfigYamlOnePool 根据单个采集池生成Prometheus的预聚合规则配置YAML
	GeneratePrometheusRecordRuleConfigYamlOnePool(ctx context.Context, pool *model.MonitorScrapePool) map[string]string
}

type recordConfigCache struct {
	mu sync.RWMutex // 读写锁，保护缓存数据
	logx.Logger
	RecordRuleMap  map[string]string // 存储预聚合规则
	localYamlDir   string            // 本地YAML目录
	scrapePoolRepo repo.MonitorScrapePoolRepo
	recordRuleRepo repo.RecordRuleRepo
}

// RecordGroup 构造Prometheus record 结构体
type RecordGroup struct {
	Name  string         `yaml:"name"`
	Rules []rulefmt.Rule `yaml:"rules"`
}

// RecordGroups 生成Prometheus record yaml
type RecordGroups struct {
	Groups []RecordGroup `yaml:"groups"`
}

func NewRecordConfigCache(ctx context.Context, db *gorm.DB, config *config.Config) RecordConfigCache {
	return &recordConfigCache{
		Logger:         logx.WithContext(ctx),
		mu:             sync.RWMutex{},
		localYamlDir:   config.PrometheusConfig.LocalYamlDir,
		RecordRuleMap:  make(map[string]string),
		scrapePoolRepo: dao.NewMonitorScrapePoolDAO(db),
		recordRuleRepo: dao.NewMonitorRecordRuleDAO(db),
	}
}

func (r *recordConfigCache) GetPrometheusRecordRuleConfigYamlByIp(ip string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.RecordRuleMap[ip]
}

func (r *recordConfigCache) GenerateRecordRuleConfigYaml(ctx context.Context) error {
	// 获取支持预聚合配置的所有采集池
	pools, err := r.scrapePoolRepo.GetMonitorScrapePoolList(ctx)
	if err != nil {
		r.Logger.Error("[监控模块] 获取支持预聚合的采集池失败: %v", err)
		return err
	}

	if len(pools) == 0 {
		r.Logger.Info("没有找到采集池")
		return nil
	}

	recordConfigMap := r.RecordRuleMap

	// 遍历每个采集池生成对应的预聚合规则配置
	for _, pool := range pools {
		if pool.SupportRecord == 1 {
			oneMap := r.GeneratePrometheusRecordRuleConfigYamlOnePool(ctx, pool)
			for ip, out := range oneMap {
				recordConfigMap[ip] = out
			}
		}
	}

	r.mu.Lock()
	r.RecordRuleMap = recordConfigMap
	r.mu.Unlock()

	return nil
}

// GeneratePrometheusRecordRuleConfigYamlOnePool 根据单个采集池生成Prometheus的预聚合规则配置YAML
func (r *recordConfigCache) GeneratePrometheusRecordRuleConfigYamlOnePool(ctx context.Context, pool *model.MonitorScrapePool) map[string]string {
	rules, err := r.recordRuleRepo.GetMonitorRecordRuleByPoolId(ctx, pool.ID)
	if err != nil {
		r.Logger.Error("[监控模块] 根据采集池ID获取预聚合规则失败: %v", err)
		return nil
	}

	if len(rules) == 0 {
		return nil
	}

	var recordGroups RecordGroups

	// 构建规则组
	for _, rule := range rules {
		forD, err := pm.ParseDuration(rule.ForTime)
		if err != nil {
			r.Logger.Infof("[监控模块] 解析预聚合规则持续时间失败，使用默认值: %v", err)
			forD, _ = pm.ParseDuration("5s")
		}
		oneRule := rulefmt.Rule{
			Alert: rule.Name, // 告警名称
			Expr:  rule.Expr, // 预聚合表达式
			For:   forD,      // 持续时间
		}

		recordGroup := RecordGroup{
			Name:  rule.Name,
			Rules: []rulefmt.Rule{oneRule},
		}
		recordGroups.Groups = append(recordGroups.Groups, recordGroup)
	}

	numInstances := len(pool.PrometheusInstances)
	if numInstances == 0 {
		r.Logger.Infof("[监控模块] 采集池中没有Prometheus实例: %s", pool.Name)
		return nil
	}

	ruleMap := make(map[string]string)

	// 分片逻辑，将规则分配给不同的Prometheus实例
	for i, ip := range pool.PrometheusInstances {
		var myRecordGroups RecordGroups
		for j, group := range recordGroups.Groups {
			if j%numInstances == i { // 按顺序平均分片
				myRecordGroups.Groups = append(myRecordGroups.Groups, group)
			}
		}

		yamlData, err := yaml.Marshal(&myRecordGroups)
		if err != nil {
			r.Logger.Errorf("[监控模块] 序列化预聚合规则YAML失败: %v", err)
			continue
		}
		fileName := fmt.Sprintf("%s/prometheus_record_%s_%s.yml",
			r.localYamlDir,
			pool.Name,
			ip,
		)

		if err := os.WriteFile(fileName, yamlData, 0644); err != nil {
			r.Logger.Errorf("[监控模块] 写入预聚合规则文件失败: %v", err)
			continue
		}

		ruleMap[ip] = string(yamlData)
	}

	return ruleMap
}
