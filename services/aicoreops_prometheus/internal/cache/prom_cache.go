package cache

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/config"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/pkg"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/repo"
	pcc "github.com/prometheus/common/config"
	pm "github.com/prometheus/common/model"
	pc "github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/discovery/http"
	"github.com/prometheus/prometheus/discovery/kubernetes"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/relabel"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

const hashTmpKey = "__tmp_hash"

type PromConfigCache interface {
	// GetPrometheusMainConfigByIP 根据IP地址获取Prometheus的主配置内容
	GetPrometheusMainConfigByIP(ip string) string
	// GeneratePrometheusMainConfig 生成所有Prometheus主配置文件
	GeneratePrometheusMainConfig(ctx context.Context) error
	// CreateBasePrometheusConfig 创建基础Prometheus配置
	CreateBasePrometheusConfig(pool *model.MonitorScrapePool) (pc.Config, error)
	// GenerateScrapeConfigs 生成采集配置
	GenerateScrapeConfigs(ctx context.Context, pool *model.MonitorScrapePool) []*pc.ScrapeConfig
	// ApplyHashMod 应用HashMod和Keep Relabel配置进行分片
	ApplyHashMod(scrapeConfigs []*pc.ScrapeConfig, modNum, index int) []*pc.ScrapeConfig
}

type promConfigCache struct {
	logx.Logger
	mu                      sync.RWMutex      // 读写锁，保护缓存数据
	localYamlDir            string            // 本地 YAML 目录
	PrometheusMainConfigMap map[string]string // 存储 Prometheus 主配置
	httpSdAPI               string            // HTTP 服务发现 API 地址
	scrapePoolRepo          repo.MonitorScrapePoolRepo
	scrapeJobRepo           repo.MonitorScrapeJobRepo
}

func NewPromConfigCache(ctx context.Context, db *gorm.DB, config *config.Config) PromConfigCache {
	return &promConfigCache{
		Logger:                  logx.WithContext(ctx),
		mu:                      sync.RWMutex{},
		localYamlDir:            config.PrometheusConfig.LocalYamlDir,
		httpSdAPI:               config.PrometheusConfig.HttpSdAPI,
		PrometheusMainConfigMap: make(map[string]string),
		scrapePoolRepo:          dao.NewMonitorScrapePoolDAO(db),
		scrapeJobRepo:           dao.NewMonitorScrapeJobDAO(db),
	}
}

func (p *promConfigCache) GetPrometheusMainConfigByIP(ip string) string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.PrometheusMainConfigMap[ip]
}

func (p *promConfigCache) GeneratePrometheusMainConfig(ctx context.Context) error {
	// 获取所有采集池
	pools, err := p.scrapePoolRepo.GetMonitorScrapePoolList(ctx)
	if err != nil {
		p.Logger.Errorf("获取采集池失败: %v", err)
		return err
	}

	if len(pools) == 0 {
		p.Logger.Info("没有找到任何采集池")
		return nil
	}

	// 创建新的配置映射key为ip，val为配置
	newConfigMap := make(map[string]string)

	for _, pool := range pools {
		// 创建基础配置
		baseConfig, err := p.CreateBasePrometheusConfig(pool)
		if err != nil {
			p.Logger.Errorf("%v 创建基础 Prometheus 配置失败: %v", pool.Name, err)
			continue
		}

		// 生成采集配置
		scrapeConfigs := p.GenerateScrapeConfigs(ctx, pool)
		if len(scrapeConfigs) == 0 {
			p.Logger.Debugf("%v 没有找到任何采集任务", pool.Name)
			continue
		}
		baseConfig.ScrapeConfigs = scrapeConfigs

		for idx, ip := range pool.PrometheusInstances {
			configCopy := baseConfig // 浅拷贝
			// 如果有多个实例，应用哈希分片
			if len(pool.PrometheusInstances) > 1 {
				configCopy.ScrapeConfigs = p.ApplyHashMod(scrapeConfigs, len(pool.PrometheusInstances), idx)
			}

			// 序列化配置为 YAML
			yamlData, err := yaml.Marshal(configCopy)
			if err != nil {
				p.Logger.Errorf("%v 生成 Prometheus 配置失败: %v", pool.Name, err)
				continue
			}

			// 写入配置文件
			filePath := fmt.Sprintf("%s/prometheus_pool_%s.yaml", p.localYamlDir, ip)

			// 创建目录
			dir := filepath.Dir(filePath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				p.Logger.Errorf("%v 创建目录失败: %v", pool.Name, err)
				continue
			}

			if err := os.WriteFile(filePath, yamlData, 0644); err != nil {
				p.Logger.Errorf("%v 写入 Prometheus 配置文件失败: %v", pool.Name, err)
				continue
			}

			newConfigMap[ip] = string(yamlData)
			p.Logger.Debugf("%v 成功生成 Prometheus 配置", pool.Name)
		}
	}

	// 更新缓存
	p.mu.Lock()
	p.PrometheusMainConfigMap = newConfigMap
	p.mu.Unlock()

	return nil
}

func (p *promConfigCache) CreateBasePrometheusConfig(pool *model.MonitorScrapePool) (pc.Config, error) {
	var config pc.Config

	// 创建prometheus global全局配置
	if pool.ScrapeInterval <= 0 || pool.ScrapeTimeout <= 0 || pool.ScrapeTimeout > pool.ScrapeInterval {
		return pc.Config{}, fmt.Errorf("采集间隔和采集超时时间不能小于等于0，且采集超时时间不能大于采集间隔")
	}
	config.GlobalConfig = pc.GlobalConfig{
		ScrapeInterval: pkg.GenPromDuration(int(pool.ScrapeInterval)), // 采集间隔
		ScrapeTimeout:  pkg.GenPromDuration(int(pool.ScrapeTimeout)),  // 采集超时时间
	}

	// 解析外部标签
	externalLabels := pkg.ParseExternalLabels(pool.ExternalLabels)
	if len(externalLabels) > 0 {
		config.GlobalConfig.ExternalLabels = labels.FromStrings(externalLabels...)
	}

	// 解析 RemoteWrite URL
	if pool.RemoteWriteUrl != "" {
		remoteWriteURL, err := pkg.ParseURL(pool.RemoteWriteUrl)
		if err != nil {
			p.Logger.Errorf("%v 解析 RemoteWriteUrl 失败: %v", pool.Name, err)
			return pc.Config{}, fmt.Errorf("解析 RemoteWriteUrl 失败: %w", err)
		}

		config.RemoteWriteConfigs = []*pc.RemoteWriteConfig{
			{
				URL:           remoteWriteURL,
				RemoteTimeout: pkg.GenPromDuration(int(pool.RemoteTimeoutSeconds)),
			},
		}
	}

	// 解析 RemoteRead URL
	if pool.RemoteReadUrl != "" {
		remoteReadURL, err := pkg.ParseURL(pool.RemoteReadUrl)
		if err != nil {
			p.Logger.Errorf("解析 RemoteReadUrl 失败: %v", err)
			return pc.Config{}, fmt.Errorf("解析 RemoteReadUrl 失败: %w", err)
		}

		config.RemoteReadConfigs = []*pc.RemoteReadConfig{
			{
				URL:           remoteReadURL,
				RemoteTimeout: pkg.GenPromDuration(int(pool.RemoteTimeoutSeconds)),
			},
		}
	}

	// 启用告警，配置 Alertmanager
	if pool.SupportAlert == 1 {
		alertConfig := &pc.AlertmanagerConfig{
			APIVersion: "v2",
			ServiceDiscoveryConfigs: []discovery.Config{ // 服务发现配置
				discovery.StaticConfig{
					{
						Targets: []pm.LabelSet{
							{
								pm.AddressLabel: pm.LabelValue(pool.AlertManagerUrl), // 配置抓取目标地址
							},
						},
					},
				},
			},
		}

		// 组装Alertmanager基础配置
		config.AlertingConfig = pc.AlertingConfig{
			AlertmanagerConfigs: []*pc.AlertmanagerConfig{alertConfig},
		}
	}

	// 启用预聚合，添加规则文件
	if pool.SupportRecord == 1 && pool.RecordFilePath != "" {
		config.RuleFiles = append(config.RuleFiles, pool.RecordFilePath)
	}

	return config, nil
}

func (p *promConfigCache) GenerateScrapeConfigs(ctx context.Context, pool *model.MonitorScrapePool) []*pc.ScrapeConfig {
	// 获取与指定池相关的采集任务
	scrapeJobs, err := p.scrapeJobRepo.SearchMonitorScrapeJobByID(ctx, pool.ID)
	if err != nil {
		p.Logger.Errorf("获取采集任务失败: %v", err)
		return nil
	}
	if len(scrapeJobs) == 0 {
		p.Logger.Infof("%v 没有找到任何采集任务", pool.Name)
		return nil
	}

	var scrapeConfigs []*pc.ScrapeConfig

	for _, job := range scrapeJobs {
		sc := &pc.ScrapeConfig{
			JobName:        job.Name,
			Scheme:         job.Scheme,
			MetricsPath:    job.MetricsPath,
			ScrapeInterval: pkg.GenPromDuration(int(job.ScrapeInterval)),
			ScrapeTimeout:  pkg.GenPromDuration(int(job.ScrapeTimeout)),
		}

		// 解析 Relabel 配置
		if job.RelabelConfigsYamlString != "" {
			if err := yaml.Unmarshal([]byte(job.RelabelConfigsYamlString), &sc.RelabelConfigs); err != nil {
				p.Logger.Errorf("scrapeJob [%v] 解析 Relabel 配置失败: %v", job.Name, err)
				continue
			}
		}

		// 根据服务发现类型配置 ServiceDiscoveryConfigs
		switch job.ServiceDiscoveryType {
		case "http":
			if p.httpSdAPI == "" {
				p.Logger.Errorf("scrapeJob [%v] 获取 HTTP SD API 失败: %v", job.Name, err)
				continue
			}

			// 拼接 SD API URL
			sdURL := fmt.Sprintf("%s?port=%d&leafNodeIds=%s", p.httpSdAPI, job.Port, strings.Join(job.TreeNodeIDs, ","))

			sc.ServiceDiscoveryConfigs = discovery.Configs{
				&http.SDConfig{
					URL:             sdURL,
					RefreshInterval: pkg.GenPromDuration(int(job.RefreshInterval)),
				},
			}
		case "k8s":
			sc.HTTPClientConfig = pcc.HTTPClientConfig{ // 配置 HTTP 客户端配置
				BearerTokenFile: job.BearerTokenFile, // 设置鉴权文件路径
				TLSConfig: pcc.TLSConfig{ // 配置 TLS 配置
					CAFile:             job.TlsCaFilePath, // 设置 CA 证书文件路径
					InsecureSkipVerify: true,              // 跳过证书验证
				},
			}

			sc.ServiceDiscoveryConfigs = discovery.Configs{
				&kubernetes.SDConfig{
					Role:             kubernetes.Role(job.KubernetesSdRole), // 设置k8s服务发现角色
					KubeConfig:       job.KubeConfigFilePath,                // kubeconfig文件路径
					HTTPClientConfig: pcc.DefaultHTTPClientConfig,           // 使用默认的HTTP客户端配置
				},
			}
		default:
			p.Logger.Errorf("scrapeJob [%v] 未知的服务发现类型: %v", job.Name, job.ServiceDiscoveryType)
			continue
		}

		scrapeConfigs = append(scrapeConfigs, sc)
	}

	return scrapeConfigs
}

func (p *promConfigCache) ApplyHashMod(scrapeConfigs []*pc.ScrapeConfig, modNum, index int) []*pc.ScrapeConfig {
	var modified []*pc.ScrapeConfig

	for _, sc := range scrapeConfigs {
		// 深度拷贝 ScrapeConfig
		copySc := pkg.DeepCopyScrapeConfig(sc)
		// 添加新的 Relabel 配置
		newRelabelConfigs := []*relabel.Config{
			{
				Action:       relabel.HashMod,                // 使用哈希取模操作
				SourceLabels: pm.LabelNames{pm.AddressLabel}, // 使用抓取目标地址作为源标签
				Regex:        relabel.MustNewRegexp("(.*)"),  // 匹配所有字符
				Replacement:  "$1",                           // 将匹配的整个值作为替换结果
				Modulus:      uint64(modNum),                 // 设置模数
				TargetLabel:  hashTmpKey,                     // 目标标签 用于存储哈希取模后的结果
			},
			{
				Action:       relabel.Keep,                                      // 保留符合条件的目标 丢弃不符合条件的目标
				SourceLabels: pm.LabelNames{hashTmpKey},                         // 使用上一步计算出的哈希结果作为源标签
				Regex:        relabel.MustNewRegexp(fmt.Sprintf("^%d$", index)), // 只保留哈希结果等于当前实例索引 (index) 的目标
			},
		}

		copySc.RelabelConfigs = append(copySc.RelabelConfigs, newRelabelConfigs...)
		modified = append(modified, copySc)
	}

	return modified
}
