package cache

import (
	"context"
	"sync"
	"time"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type MonitorCache interface {
	// MonitorCacheManager 更新缓存
	MonitorCacheManager(ctx context.Context) error
}

type monitorCache struct {
	logx.Logger
	PrometheusMainConfig PromConfigCache
}

func NewMonitorCache(ctx context.Context, db *gorm.DB, config *config.Config) MonitorCache {
	return &monitorCache{
		Logger:               logx.WithContext(ctx),
		PrometheusMainConfig: NewPromConfigCache(ctx, db, config),
	}
}

// MonitorCacheManager 定期更新缓存并监听退出信号
func (mc *monitorCache) MonitorCacheManager(ctx context.Context) error {
	mc.Logger.Info("开始更新所有监控缓存配置")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(4)

	// 创建一个通道来收集错误
	errChan := make(chan error, 1)

	// 定义一个辅助函数来执行任务
	executeTask := func(taskName string, taskFunc func(context.Context) error) {
		defer wg.Done()
		mc.Logger.Info("开始执行任务：%s", taskName)
		err := taskFunc(ctx)
		if err != nil {
			mc.Logger.Error("任务执行失败：%s, error: %v", taskName, err)
			errChan <- err
			return
		}
		mc.Logger.Info("任务执行成功：%s", taskName)
	}

	// 并发执行任务
	go executeTask("生成 Prometheus 配置", mc.PrometheusMainConfig.GeneratePrometheusMainConfig)
	

	mc.Logger.Info("更新所有监控缓存配置完成")
	return nil
}
