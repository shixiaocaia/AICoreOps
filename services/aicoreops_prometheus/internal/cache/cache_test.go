package cache

import (
	"context"
	"testing"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestPromConfigCache_GeneratePrometheusMainConfig(t *testing.T) {
	ctx := context.Background()
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/AICoreOps?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	promConfigCache := NewPromConfigCache(ctx, db, &config.Config{
		PrometheusConfig: config.PrometheusConfig{
			LocalYamlDir: "./local_yaml",
			HttpSdAPI:    "http://localhost:8888/api/not_auth/getTreeNodeBindIps",
		},
	})
	promConfigCache.GeneratePrometheusMainConfig(ctx)
}

func TestAlertConfigCache_generateRouteConfigOnePool(t *testing.T) {
	ctx := context.Background()
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/AICoreOps?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	alertConfigCache := NewAlertConfigCache(ctx, db, &config.Config{
		AlertManagerConfig: config.AlertManagerConfig{
			LocalYamlDir:     "./local_yaml",
			AlertWebhookAddr: "http://localhost:8888/api/not_auth/getTreeNodeBindIps",
		},
	})
	alertConfigCache.GenerateAlertManagerMainConfig(ctx)
}
