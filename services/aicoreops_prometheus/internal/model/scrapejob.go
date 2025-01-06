package model

import (
	"time"

	"gorm.io/gorm"
)

// MonitorScrapeJob Prometheus 抓取任务配置
type MonitorScrapeJob struct {
	ID                       int64          `gorm:"column:id;primaryKey;autoIncrement"`
	Name                     string         `gorm:"column:name;type:varchar(255);not null;comment:任务名称"`
	UserID                   int64          `gorm:"column:user_id;not null;comment:用户ID"`
	Enable                   int32          `gorm:"column:enable;type:tinyint(1);not null;default:1;comment:是否启用 1:启用 2:禁用"`
	ServiceDiscoveryType     string         `gorm:"column:service_discovery_type;type:varchar(50);not null;comment:服务发现类型"`
	MetricsPath              string         `gorm:"column:metrics_path;type:varchar(255);not null;default:'/metrics';comment:指标路径"`
	Scheme                   string         `gorm:"column:scheme;type:varchar(10);not null;default:'http';comment:协议"`
	ScrapeInterval           int32          `gorm:"column:scrape_interval;not null;default:15;comment:抓取间隔(秒)"`
	ScrapeTimeout            int32          `gorm:"column:scrape_timeout;not null;default:10;comment:抓取超时(秒)"`
	PoolID                   int64          `gorm:"column:pool_id;not null;comment:采集池ID"`
	RelabelConfigsYamlString string         `gorm:"column:relabel_configs_yaml_string;type:text;comment:重新标记配置"`
	RefreshInterval          int32          `gorm:"column:refresh_interval;not null;default:300;comment:刷新间隔(秒)"`
	Port                     int32          `gorm:"column:port;not null;default:9090;comment:端口"`
	TreeNodeIDs              string         `gorm:"column:tree_node_ids;type:text;comment:服务树节点ID列表,json格式"`
	KubeConfigFilePath       string         `gorm:"column:kube_config_file_path;type:varchar(255);comment:k8s配置文件路径"`
	TlsCaFilePath            string         `gorm:"column:tls_ca_file_path;type:varchar(255);comment:TLS CA证书路径"`
	TlsCaContent             string         `gorm:"column:tls_ca_content;type:text;comment:TLS CA证书内容"`
	BearerToken              string         `gorm:"column:bearer_token;type:text;comment:Bearer Token"`
	BearerTokenFile          string         `gorm:"column:bearer_token_file;type:varchar(255);comment:Bearer Token文件路径"`
	KubernetesSdRole         string         `gorm:"column:kubernetes_sd_role;type:varchar(50);comment:k8s服务发现角色"`
	CreatedAt                time.Time      `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt                time.Time      `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt                gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

// TableName 表名
func (s *MonitorScrapeJob) TableName() string {
	return "monitor_scrape_job"
}
