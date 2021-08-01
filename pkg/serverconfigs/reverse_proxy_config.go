package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/schedulingconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/lists"
	"sync"
)

type RequestHostType = int8

const (
	RequestHostTypeProxyServer RequestHostType = 0
	RequestHostTypeOrigin      RequestHostType = 1
	RequestHostTypeCustomized  RequestHostType = 2
)

// ReverseProxyConfig 反向代理设置
type ReverseProxyConfig struct {
	Id                int64             `yaml:"id" json:"id"`                               // ID
	IsOn              bool              `yaml:"isOn" json:"isOn"`                           // 是否启用
	PrimaryOrigins    []*OriginConfig   `yaml:"primaryOrigins" json:"primaryOrigins"`       // 主要源站列表
	PrimaryOriginRefs []*OriginRef      `yaml:"primaryOriginRefs" json:"primaryOriginRefs"` // 主要源站引用
	BackupOrigins     []*OriginConfig   `yaml:"backupOrigins" json:"backupOrigins"`         // 备用源站列表
	BackupOriginRefs  []*OriginRef      `yaml:"backupOriginRefs" json:"backupOriginRefs"`   // 备用源站引用
	Scheduling        *SchedulingConfig `yaml:"scheduling" json:"scheduling"`               // 调度算法选项

	ConnTimeout  *shared.TimeDuration `yaml:"connTimeout" json:"connTimeout"`   // 连接失败超时 TODO
	ReadTimeout  *shared.TimeDuration `yaml:"readTimeout" json:"readTimeout"`   // 读取超时时间 TODO
	IdleTimeout  *shared.TimeDuration `yaml:"idleTimeout" json:"idleTimeout"`   // 空闲连接超时时间 TODO
	MaxFails     int                  `yaml:"maxFails" json:"maxFails"`         // 最多失败次数 TODO
	MaxConns     int                  `yaml:"maxConns" json:"maxConns"`         // 最大并发连接数 TODO
	MaxIdleConns int                  `yaml:"maxIdleConns" json:"maxIdleConns"` // 最大空闲连接数 TODO

	StripPrefix     string          `yaml:"stripPrefix" json:"stripPrefix"`         // 去除URL前缀
	RequestHostType RequestHostType `yaml:"requestHostType" json:"requestHostType"` // 请求Host类型
	RequestHost     string          `yaml:"requestHost" json:"requestHost"`         // 请求Host，支持变量
	RequestURI      string          `yaml:"requestURI" json:"requestURI"`           // 请求URI，支持变量，如果同时定义了StripPrefix，则先执行StripPrefix

	AddHeaders []string `yaml:"addHeaders" json:"addHeaders"` // 自动添加的Header

	AutoFlush bool `yaml:"autoFlush" json:"autoFlush"` // 是否自动刷新缓冲区，在比如SSE（server-sent events）场景下很有用

	requestHostHasVariables bool
	requestURIHasVariables  bool

	hasPrimaryOrigins  bool
	hasBackupOrigins   bool
	schedulingIsBackup bool
	schedulingObject   schedulingconfigs.SchedulingInterface
	schedulingLocker   sync.Mutex

	addXRealIPHeader         bool
	addXForwardedForHeader   bool
	addForwardedHeader       bool
	addXForwardedByHeader    bool
	addXForwardedHostHeader  bool
	addXForwardedProtoHeader bool
}

// Init 初始化
func (this *ReverseProxyConfig) Init() error {
	this.requestHostHasVariables = configutils.HasVariables(this.RequestHost)
	this.requestURIHasVariables = configutils.HasVariables(this.RequestURI)

	this.hasPrimaryOrigins = len(this.PrimaryOrigins) > 0
	this.hasBackupOrigins = len(this.BackupOrigins) > 0

	for _, origins := range [][]*OriginConfig{this.PrimaryOrigins, this.BackupOrigins} {
		for _, origin := range origins {
			// 覆盖参数设置
			if origin.MaxFails <= 0 && this.MaxFails > 0 {
				origin.MaxFails = this.MaxFails
			}
			if origin.MaxConns <= 0 && this.MaxConns > 0 {
				origin.MaxConns = this.MaxConns
			}
			if origin.MaxIdleConns <= 0 && this.MaxIdleConns > 0 {
				origin.MaxIdleConns = this.MaxIdleConns
			}
			if (origin.ConnTimeout == nil || origin.ConnTimeout.Count <= 0) && this.ConnTimeout != nil && this.ConnTimeout.Count > 0 {
				origin.ConnTimeout = this.ConnTimeout
			}
			if (origin.ReadTimeout == nil || origin.ReadTimeout.Count <= 0) && this.ReadTimeout != nil && this.ReadTimeout.Count > 0 {
				origin.ReadTimeout = this.ReadTimeout
			}
			if (origin.IdleTimeout == nil || origin.IdleTimeout.Count <= 0) && this.IdleTimeout != nil && this.IdleTimeout.Count > 0 {
				origin.IdleTimeout = this.IdleTimeout
			}

			// 初始化
			err := origin.Init()
			if err != nil {
				return err
			}
		}
	}

	// scheduling
	this.SetupScheduling(false, false, true)

	// Header
	if len(this.AddHeaders) == 0 {
		// 默认加入两项
		this.addXRealIPHeader = true
		this.addXForwardedForHeader = true
		this.addXForwardedByHeader = true
		this.addXForwardedHostHeader = true
		this.addXForwardedProtoHeader = true
	} else {
		this.addXRealIPHeader = lists.ContainsString(this.AddHeaders, "X-Real-IP")
		this.addXForwardedForHeader = lists.ContainsString(this.AddHeaders, "X-Forwarded-For")
		this.addXForwardedByHeader = lists.ContainsString(this.AddHeaders, "X-Forwarded-By")
		this.addXForwardedHostHeader = lists.ContainsString(this.AddHeaders, "X-Forwarded-Host")
		this.addXForwardedProtoHeader = lists.ContainsString(this.AddHeaders, "X-Forwarded-Proto")
	}

	return nil
}

// AddPrimaryOrigin 添加主源站配置
func (this *ReverseProxyConfig) AddPrimaryOrigin(origin *OriginConfig) {
	this.PrimaryOrigins = append(this.PrimaryOrigins, origin)
}

// AddBackupOrigin 添加备用源站配置
func (this *ReverseProxyConfig) AddBackupOrigin(origin *OriginConfig) {
	this.BackupOrigins = append(this.BackupOrigins, origin)
}

// NextOrigin 取得下一个可用的后端服务
func (this *ReverseProxyConfig) NextOrigin(call *shared.RequestCall) *OriginConfig {
	this.schedulingLocker.Lock()
	defer this.schedulingLocker.Unlock()

	if this.schedulingObject == nil {
		return nil
	}

	if this.Scheduling != nil && call != nil && call.Options != nil {
		for k, v := range this.Scheduling.Options {
			call.Options[k] = v
		}
	}

	candidate := this.schedulingObject.Next(call)

	// 末了重置状态
	defer func() {
		if candidate == nil {
			this.schedulingIsBackup = false
		}
	}()

	if candidate == nil {
		// 启用备用服务器
		if !this.schedulingIsBackup {
			this.SetupScheduling(true, true, false)
			candidate = this.schedulingObject.Next(call)
			if candidate == nil {
				// 不检查主要源站
				this.SetupScheduling(false, false, false)
				candidate = this.schedulingObject.Next(call)
				if candidate == nil {
					// 不检查备用源站
					this.SetupScheduling(true, false, false)
					candidate = this.schedulingObject.Next(call)
					if candidate == nil {
						return nil
					}
				}
			}
		}

		if candidate == nil {
			return nil
		}
	}

	return candidate.(*OriginConfig)
}

// SetupScheduling 设置调度算法
func (this *ReverseProxyConfig) SetupScheduling(isBackup bool, checkOk bool, lock bool) {
	if lock {
		this.schedulingLocker.Lock()
		defer this.schedulingLocker.Unlock()
	}

	this.schedulingIsBackup = isBackup

	if this.Scheduling == nil {
		this.schedulingObject = &schedulingconfigs.RandomScheduling{}
	} else {
		typeCode := this.Scheduling.Code
		s := schedulingconfigs.FindSchedulingType(typeCode)
		if s == nil {
			this.Scheduling = nil
			this.schedulingObject = &schedulingconfigs.RandomScheduling{}
		} else {
			this.schedulingObject = s["instance"].(schedulingconfigs.SchedulingInterface)
		}
	}

	if !isBackup {
		for _, origin := range this.PrimaryOrigins {
			if origin.IsOn && (origin.IsOk || !checkOk) {
				this.schedulingObject.Add(origin)
			}
		}
	} else {
		for _, origin := range this.BackupOrigins {
			if origin.IsOn && (origin.IsOk || !checkOk) {
				this.schedulingObject.Add(origin)
			}
		}
	}

	if !this.schedulingObject.HasCandidates() {
		return
	}

	this.schedulingObject.Start()
}

// FindSchedulingConfig 获取调度配置对象
func (this *ReverseProxyConfig) FindSchedulingConfig() *SchedulingConfig {
	if this.Scheduling == nil {
		this.Scheduling = &SchedulingConfig{Code: "random"}
	}
	return this.Scheduling
}

// RequestHostHasVariables 判断RequestHost是否有变量
func (this *ReverseProxyConfig) RequestHostHasVariables() bool {
	return this.requestHostHasVariables
}

// RequestURIHasVariables 判断RequestURI是否有变量
func (this *ReverseProxyConfig) RequestURIHasVariables() bool {
	return this.requestURIHasVariables
}

// ShouldAddXRealIPHeader 是否添加X-Real-IP
func (this *ReverseProxyConfig) ShouldAddXRealIPHeader() bool {
	return this.addXRealIPHeader
}

// ShouldAddXForwardedForHeader 是否添加X-Forwarded-For
func (this *ReverseProxyConfig) ShouldAddXForwardedForHeader() bool {
	return this.addXForwardedForHeader
}

// ShouldAddXForwardedByHeader 是否添加X-Forwarded-By
func (this *ReverseProxyConfig) ShouldAddXForwardedByHeader() bool {
	return this.addXForwardedByHeader
}

// ShouldAddXForwardedHostHeader 是否添加X-Forwarded-Host
func (this *ReverseProxyConfig) ShouldAddXForwardedHostHeader() bool {
	return this.addXForwardedHostHeader
}

// ShouldAddXForwardedProtoHeader 是否添加X-Forwarded-Proto
func (this *ReverseProxyConfig) ShouldAddXForwardedProtoHeader() bool {
	return this.addXForwardedProtoHeader
}

// ResetScheduling 重置调度算法
func (this *ReverseProxyConfig) ResetScheduling() {
	this.SetupScheduling(false, true, true)
}
