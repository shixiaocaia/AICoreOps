package model

// MonitorScrapePool 采集池的配置
type MonitorScrapePool struct {
	ID                    int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:采集池ID"`
	Name                  string     `json:"name" binding:"required,min=1,max=50" gorm:"uniqueIndex;size:100;comment:采集池名称，支持使用通配符*进行模糊搜索"`
	PrometheusInstances   StringList `json:"prometheusInstances,omitempty" gorm:"type:text;comment:选择多个Prometheus实例"`
	AlertManagerInstances StringList `json:"alertManagerInstances,omitempty" gorm:"type:text;comment:选择多个AlertManager实例"`
	UserID                int64      `json:"userId" gorm:"comment:创建该采集池的用户ID"`
	ScrapeInterval        int32      `json:"scrapeInterval,omitempty" gorm:"default:30;type:int;comment:采集间隔（秒）"`
	ScrapeTimeout         int32      `json:"scrapeTimeout,omitempty" gorm:"default:10;type:int;comment:采集超时时间（秒）"`
	ExternalLabels        StringList `json:"externalLabels,omitempty" gorm:"type:text;comment:remote_write时添加的标签组，格式为 key=v，例如 scrape_ip=1.1.1.1"`
	SupportAlert          int32      `json:"supportAlert" gorm:"type:int;comment:是否支持告警：1支持，2不支持"`
	SupportRecord         int32      `json:"supportRecord" gorm:"type:int;comment:是否支持预聚合：1支持，2不支持"`
	RemoteReadUrl         string     `json:"remoteReadUrl,omitempty" gorm:"size:255;comment:远程读取的地址"`
	AlertManagerUrl       string     `json:"alertManagerUrl,omitempty" gorm:"size:255;comment:AlertManager的地址"`
	RuleFilePath          string     `json:"ruleFilePath,omitempty" gorm:"size:255;comment:规则文件路径"`
	RecordFilePath        string     `json:"recordFilePath,omitempty" gorm:"size:255;comment:记录文件路径"`
	RemoteWriteUrl        string     `json:"remoteWriteUrl,omitempty" gorm:"size:255;comment:远程写入的地址"`
	RemoteTimeoutSeconds  int32      `json:"remoteTimeoutSeconds,omitempty" gorm:"default:5;type:int;comment:远程写入的超时时间（秒）"`
	CreateTime            int64      `gorm:"column:create_time;type:int;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime            int64      `gorm:"column:update_time;type:int;autoUpdateTime" json:"update_time"` // 更新时间
	IsDeleted             int        `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`    // 软删除标志（0:否, 1:是）

	// 前端使用字段
	ExternalLabelsFront string `json:"externalLabelsFront,omitempty" gorm:"-"`
	Key                 string `json:"key" gorm:"-"`
	CreateUserName      string `json:"createUserName,omitempty" gorm:"-"`
}

func (MonitorScrapePool) TableName() string {
	return "monitor_scrape_pool"
}
