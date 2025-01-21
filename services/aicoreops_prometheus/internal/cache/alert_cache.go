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
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"

	alertconfig "github.com/prometheus/alertmanager/config"
	al "github.com/prometheus/alertmanager/pkg/labels"
	pm "github.com/prometheus/common/model"
)

const alertSendGroupKey = "alert_send_group"
const defaultConfigTimeout = "5s"

type AlertConfigCache interface {
	// GetAlertManagerMainConfigYamlByIP 根据IP地址获取AlertManager的主配置内容
	GetAlertManagerMainConfigYamlByIP(ip string) string
	// GenerateAlertManagerMainConfig 生成所有AlertManager主配置文件
	GenerateAlertManagerMainConfig(ctx context.Context) error
	// GenerateAlertManagerMainConfigOnePool 生成单个AlertManager池的主配置
	// GenerateAlertManagerMainConfigOnePool(pool *model.MonitorAlertManagerPool) *altconfig.Config
	// GenerateAlertManagerRouteConfigOnePool 生成单个AlertManager池的routes和receivers配置
	// GenerateAlertManagerRouteConfigOnePool(ctx context.Context, pool *model.MonitorAlertManagerPool) ([]*altconfig.Route, []altconfig.Receiver)
}

type alertConfigCache struct {
	logx.Logger
	mu                        sync.RWMutex
	localYamlDir              string
	alertWebhookAddr          string
	alertManagerMainConfigMap map[string]string
	alertPoolRepo             repo.MonitorAlterManagerPoolRepo
	alertSendRepo             repo.SendGroupRepo
}

func NewAlertConfigCache(ctx context.Context, db *gorm.DB, config *config.Config) AlertConfigCache {
	return &alertConfigCache{
		Logger:                    logx.WithContext(ctx),
		mu:                        sync.RWMutex{},
		localYamlDir:              config.AlertManagerConfig.LocalYamlDir,
		alertWebhookAddr:          config.AlertManagerConfig.AlertWebhookAddr,
		alertManagerMainConfigMap: make(map[string]string),
		alertPoolRepo:             dao.NewAlertManagerPoolDao(db),
		alertSendRepo:             dao.NewSendGroupDao(db),
	}
}

func (a *alertConfigCache) GetAlertManagerMainConfigYamlByIP(ip string) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.alertManagerMainConfigMap[ip]
}

func (a *alertConfigCache) GenerateAlertManagerMainConfig(ctx context.Context) error {
	// 从数据库中获取所有AlertManager采集池
	pools, err := a.alertPoolRepo.GetMonitorAlertManagerPoolList(ctx)
	if err != nil {
		a.Logger.Errorf("[更新AlertManager]扫描数据库中的AlertManager集群失败: %v", err)
		return err
	}

	if len(pools) == 0 {
		a.Logger.Info("[更新AlertManager]没有找到任何AlertManager采集池")
		return nil
	}

	mainConfigMap := make(map[string]string)

	for _, pool := range pools {
		// 生成单个AlertManager池的主配置
		oneConfig := a.generateMainConfigOnePool(pool)

		// 生成对应的routes和receivers配置
		routes, receivers := a.generateRouteConfigOnePool(ctx, pool)
		if len(routes) > 0 {
			oneConfig.Route.Routes = routes
		}

		if len(receivers) > 0 {
			if oneConfig.Receivers == nil {
				oneConfig.Receivers = receivers
			} else {
				oneConfig.Receivers = append(oneConfig.Receivers, receivers...)
			}
		}
		// 序列化配置为YAML格式
		config, err := yaml.Marshal(oneConfig)
		if err != nil {
			a.Logger.Error("[定时任务更新AlertManager配置]根据alert配置生成AlertManager主配置文件错误",
				zap.Error(err),
				zap.String("池子", pool.Name),
			)
			continue
		}
		a.Logger.Debug("[定时任务更新AlertManager配置]根据alert配置生成AlertManager主配置文件成功",
			zap.String("池子", pool.Name),
			zap.ByteString("配置", config),
		)

		// 写入配置文件并更新缓存
		for index, ip := range pool.AlertManagerInstances {
			fileName := fmt.Sprintf("%s/alertmanager_pool_%s_%d.yaml",
				a.localYamlDir,
				pool.Name,
				index,
			)

			if err := os.WriteFile(fileName, config, 0644); err != nil {
				a.Logger.Error("[定时任务更新AlertManager配置]写入AlertManager配置文件失败",
					zap.Error(err),
					zap.String("文件路径", fileName),
				)
				continue
			}

			// 配置存入map中
			mainConfigMap[ip] = string(config)
		}
	}

	// 更新缓存
	a.mu.Lock()
	a.alertManagerMainConfigMap = mainConfigMap
	a.mu.Unlock()

	return nil
}

func (a *alertConfigCache) generateMainConfigOnePool(pool *model.MonitorAlertManagerPool) *alertconfig.Config {
	// 解析默认恢复时间
	resolveTimeout, err := pm.ParseDuration(pool.ResolveTimeout)
	if err != nil {
		a.Logger.Errorf("alertPool [%v] 解析ResolveTimeout失败：%v，使用默认值: %v", pool.Name, err, defaultConfigTimeout)
		resolveTimeout, _ = pm.ParseDuration(defaultConfigTimeout)
	}

	// 解析分组第一次等待时间
	groupWait, err := pm.ParseDuration(pool.GroupWait)
	if err != nil {
		a.Logger.Errorf("alertPool [%v] 解析GroupWait失败：%v，使用默认值: %v", pool.Name, err, defaultConfigTimeout)
		groupWait, _ = pm.ParseDuration(defaultConfigTimeout)
	}

	// 解析分组等待间隔时间
	groupInterval, err := pm.ParseDuration(pool.GroupInterval)
	if err != nil {
		a.Logger.Errorf("alertPool [%v] 解析GroupInterval失败：%v，使用默认值: %v", pool.Name, err, defaultConfigTimeout)
		groupInterval, _ = pm.ParseDuration(defaultConfigTimeout)
	}

	// 解析重复发送时间
	repeatInterval, err := pm.ParseDuration(pool.RepeatInterval)
	if err != nil {
		a.Logger.Errorf("alertPool [%v] 解析RepeatInterval失败：%v，使用默认值: %v", pool.Name, err, defaultConfigTimeout)
		repeatInterval, _ = pm.ParseDuration(defaultConfigTimeout)
	}

	// 生成 Alertmanager 默认配置
	config := &alertconfig.Config{
		Global: &alertconfig.GlobalConfig{
			ResolveTimeout: resolveTimeout, // 设置恢复超时时间
		},
		Route: &alertconfig.Route{ // 设置默认路由
			Receiver:       pool.Receiver,   // 设置默认接收者
			GroupWait:      &groupWait,      // 设置分组等待时间
			GroupInterval:  &groupInterval,  // 设置分组等待间隔
			RepeatInterval: &repeatInterval, // 设置重复发送时间
			GroupByStr:     pool.GroupBy,    // 设置分组分组标签
		},
	}

	return config
}

func (a *alertConfigCache) generateRouteConfigOnePool(ctx context.Context, pool *model.MonitorAlertManagerPool) ([]*alertconfig.Route, []alertconfig.Receiver) {
	// 从数据库中查找该AlertManager池的所有发送组
	sendGroups, err := a.alertSendRepo.GetMonitorSendGroupByPoolId(ctx, pool.ID)
	if err != nil {
		a.Logger.Errorf("alertPool [%v] 查找所有发送组错误：%v", pool.Name, err)
		return nil, nil
	}
	if len(sendGroups) == 0 {
		a.Logger.Info("alertPool [%v] 没有找到发送组", pool.Name)
		return nil, nil
	}

	var routes []*alertconfig.Route
	var receivers []alertconfig.Receiver

	var default_receiver bool

	for _, sendGroup := range sendGroups {
		// 去重 default_receiver
		if sendGroup.Name == pool.Receiver {
			default_receiver = true
		}

		// 解析RepeatInterval
		repeatInterval, err := pm.ParseDuration(sendGroup.RepeatInterval)
		if err != nil {
			a.Logger.Errorf("alertPool [%v] 解析 RepeatInterval 失败: %v，使用默认值: %v", pool.Name, err, defaultConfigTimeout)
			repeatInterval, _ = pm.ParseDuration(defaultConfigTimeout)
		}

		// 创建 Matcher 并设置匹配条件
		// 默认匹配条件为: alert_send_group=sendGroup.ID
		matcher, err := al.NewMatcher(al.MatchEqual, alertSendGroupKey, fmt.Sprintf("%d", sendGroup.ID))
		if err != nil {
			a.Logger.Errorf("alertPool [%v] 创建 Matcher 失败: %v", pool.Name, err)
			continue
		}

		// 创建Route
		route := &alertconfig.Route{
			Receiver:       sendGroup.Name,         // 设置接收者
			Continue:       true,                   // 继续匹配下一个路由
			Matchers:       []*al.Matcher{matcher}, // 设置匹配条件
			RepeatInterval: &repeatInterval,        // 设置重复发送时间
		}

		// 拼接Webhook URL
		webHookURL := fmt.Sprintf("%s?%s=%d",
			a.alertWebhookAddr,
			alertSendGroupKey,
			sendGroup.ID,
		)

		// 将 URL 写入到 .txt 文件
		urlFilePath := fmt.Sprintf("%s/webhook_url_%d.txt", a.localYamlDir, sendGroup.ID)
		err = os.WriteFile(urlFilePath, []byte(webHookURL), 0644)
		if err != nil {
			a.Logger.Errorf("alertPool [%v] 写入Webhook URL文件失败: %v", pool.Name, err)
			continue
		}

		// 创建Receiver
		receiver := alertconfig.Receiver{
			Name: sendGroup.Name, // 接收者名称
			WebhookConfigs: []*alertconfig.WebhookConfig{ // Webhook配置
				{
					NotifierConfig: alertconfig.NotifierConfig{ // Notifier配置 用于告警通知
						VSendResolved: sendGroup.SendResolved == 1, // 在告警解决时是否发送通知
					},
					URLFile: urlFilePath,
				},
			},
		}
		// 添加到routes和receivers中
		routes = append(routes, route)
		receivers = append(receivers, receiver)
	}

	// 如果有默认rs列表中Receiver，则添加到Receivers
	// 如果default_receiver为true，包含在sendGroups中
	if pool.Receiver != "" && !default_receiver {
		defaultSendGroup, err := a.alertSendRepo.GetMonitorSendGroupByName(ctx, pool.Receiver)
		if err != nil {
			a.Logger.Errorf("alertPool [%v] 获取默认Receiver失败：%v", pool.Name, err)
			return nil, nil
		}

		webHookURL := fmt.Sprintf("%s?%s=%d",
			a.alertWebhookAddr,
			alertSendGroupKey,
			defaultSendGroup.ID,
		)

		// 将 URL 写入到 .txt 文件
		urlFilePath := fmt.Sprintf("%s/webhook_url_%d.txt", a.localYamlDir, defaultSendGroup.ID)
		err = os.WriteFile(urlFilePath, []byte(webHookURL), 0644)
		if err != nil {
			a.Logger.Errorf("alertPool [%v] 写入Webhook URL文件失败: %v", pool.Name, err)
			return nil, nil
		}

		receivers = append(receivers, alertconfig.Receiver{
			Name: defaultSendGroup.Name, // 接收者名称
			WebhookConfigs: []*alertconfig.WebhookConfig{ // Webhook配置
				{
					NotifierConfig: alertconfig.NotifierConfig{ // Notifier配置 用于告警通知
						VSendResolved: defaultSendGroup.SendResolved == 1, // 在告警解决时是否发送通知
					},
					URLFile: urlFilePath,
				},
			},
		})

	}

	return routes, receivers
}
