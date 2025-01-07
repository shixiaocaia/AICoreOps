package model

// MonitorScrapePool 采集池的配置
type MonitorScrapePool struct {
	ID                    int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:采集池ID"`
	Name                  string     `json:"name" binding:"required,min=1,max=50" gorm:"uniqueIndex;size:100;comment:采集池名称，支持使用通配符*进行模糊搜索"`
	PrometheusInstances   StringList `json:"prometheusInstances,omitempty" gorm:"type:text;comment:选择多个Prometheus实例"`
	AlertManagerInstances StringList `json:"alertManagerInstances,omitempty" gorm:"type:text;comment:选择多个AlertManager实例"`
	UserID                int        `json:"userId" gorm:"comment:创建该采集池的用户ID"`
	ScrapeInterval        int        `json:"scrapeInterval,omitempty" gorm:"default:30;type:int;comment:采集间隔（秒）"`
	ScrapeTimeout         int        `json:"scrapeTimeout,omitempty" gorm:"default:10;type:int;comment:采集超时时间（秒）"`
	ExternalLabels        StringList `json:"externalLabels,omitempty" gorm:"type:text;comment:remote_write时添加的标签组，格式为 key=v，例如 scrape_ip=1.1.1.1"`
	SupportAlert          int        `json:"supportAlert" gorm:"type:int;comment:是否支持告警：1支持，2不支持"`
	SupportRecord         int        `json:"supportRecord" gorm:"type:int;comment:是否支持预聚合：1支持，2不支持"`
	RemoteReadUrl         string     `json:"remoteReadUrl,omitempty" gorm:"size:255;comment:远程读取的地址"`
	AlertManagerUrl       string     `json:"alertManagerUrl,omitempty" gorm:"size:255;comment:AlertManager的地址"`
	RuleFilePath          string     `json:"ruleFilePath,omitempty" gorm:"size:255;comment:规则文件路径"`
	RecordFilePath        string     `json:"recordFilePath,omitempty" gorm:"size:255;comment:记录文件路径"`
	RemoteWriteUrl        string     `json:"remoteWriteUrl,omitempty" gorm:"size:255;comment:远程写入的地址"`
	RemoteTimeoutSeconds  int        `json:"remoteTimeoutSeconds,omitempty" gorm:"default:5;type:int;comment:远程写入的超时时间（秒）"`

	// 前端使用字段
	ExternalLabelsFront string `json:"externalLabelsFront,omitempty" gorm:"-"`
	Key                 string `json:"key" gorm:"-"`
	CreateUserName      string `json:"createUserName,omitempty" gorm:"-"`
}

func (MonitorScrapePool) TableName() string {
	return "monitor_scrape_pool"
}
