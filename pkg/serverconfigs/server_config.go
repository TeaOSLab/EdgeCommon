package serverconfigs

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
)

type ServerConfig struct {
	Id               int64               `yaml:"id" json:"id"`                             // ID
	Type             string              `yaml:"type" json:"type"`                         // 类型
	IsOn             bool                `yaml:"isOn" json:"isOn"`                         // 是否开启
	Name             string              `yaml:"name" json:"name"`                         // 名称
	Description      string              `yaml:"description" json:"description"`           // 描述
	AliasServerNames []string            `yaml:"aliasServerNames" json:"aliasServerNames"` // 关联的域名，比如CNAME之类的
	ServerNames      []*ServerNameConfig `yaml:"serverNames" json:"serverNames"`           // 域名

	// 前端协议
	HTTP  *HTTPProtocolConfig  `yaml:"http" json:"http"`   // HTTP配置
	HTTPS *HTTPSProtocolConfig `yaml:"https" json:"https"` // HTTPS配置
	TCP   *TCPProtocolConfig   `yaml:"tcp" json:"tcp"`     // TCP配置
	TLS   *TLSProtocolConfig   `yaml:"tls" json:"tls"`     // TLS配置
	Unix  *UnixProtocolConfig  `yaml:"unix" json:"unix"`   // Unix配置
	UDP   *UDPProtocolConfig   `yaml:"udp" json:"udp"`     // UDP配置

	// Web配置
	Web *HTTPWebConfig `yaml:"web" json:"web"`

	// 反向代理配置
	ReverseProxyRef *ReverseProxyRef    `yaml:"reverseProxyRef" json:"reverseProxyRef"`
	ReverseProxy    *ReverseProxyConfig `yaml:"reverseProxy" json:"reverseProxy"`

	isOk bool
}

// 从JSON中解析Server配置
func NewServerConfigFromJSON(data []byte) (*ServerConfig, error) {
	config := &ServerConfig{}
	err := json.Unmarshal(data, config)
	return config, err
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
}

func (this *ServerConfig) Init() error {
	if this.HTTP != nil {
		err := this.HTTP.Init()
		if err != nil {
			return err
		}
	}

	if this.HTTPS != nil {
		err := this.HTTPS.Init()
		if err != nil {
			return err
		}
	}

	if this.TCP != nil {
		err := this.TCP.Init()
		if err != nil {
			return err
		}
	}

	if this.TLS != nil {
		err := this.TLS.Init()
		if err != nil {
			return err
		}
	}

	if this.Unix != nil {
		err := this.Unix.Init()
		if err != nil {
			return err
		}
	}

	if this.UDP != nil {
		err := this.UDP.Init()
		if err != nil {
			return err
		}
	}

	if this.ReverseProxyRef != nil {
		err := this.ReverseProxyRef.Init()
		if err != nil {
			return err
		}
	}

	if this.ReverseProxy != nil {
		err := this.ReverseProxy.Init()
		if err != nil {
			return err
		}
	}

	if this.Web != nil {
		err := this.Web.Init()
		if err != nil {
			return err
		}
	}

	this.isOk = true

	return nil
}

// 配置是否正确
func (this *ServerConfig) IsOk() bool {
	return this.isOk
}

func (this *ServerConfig) FullAddresses() []string {
	result := []string{}
	if this.HTTP != nil && this.HTTP.IsOn {
		result = append(result, this.HTTP.FullAddresses()...)
	}
	if this.HTTPS != nil && this.HTTPS.IsOn {
		result = append(result, this.HTTPS.FullAddresses()...)
	}
	if this.TCP != nil && this.TCP.IsOn {
		result = append(result, this.TCP.FullAddresses()...)
	}
	if this.TLS != nil && this.TLS.IsOn {
		result = append(result, this.TLS.FullAddresses()...)
	}
	if this.Unix != nil && this.Unix.IsOn {
		result = append(result, this.Unix.FullAddresses()...)
	}
	if this.UDP != nil && this.UDP.IsOn {
		result = append(result, this.UDP.FullAddresses()...)
	}

	return result
}

func (this *ServerConfig) Listen() []*NetworkAddressConfig {
	result := []*NetworkAddressConfig{}
	if this.HTTP != nil {
		result = append(result, this.HTTP.Listen...)
	}
	if this.HTTPS != nil {
		result = append(result, this.HTTPS.Listen...)
	}
	if this.TCP != nil {
		result = append(result, this.TCP.Listen...)
	}
	if this.TLS != nil {
		result = append(result, this.TLS.Listen...)
	}
	if this.Unix != nil {
		result = append(result, this.Unix.Listen...)
	}
	if this.UDP != nil {
		result = append(result, this.UDP.Listen...)
	}
	return result
}

func (this *ServerConfig) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}

func (this *ServerConfig) IsHTTPFamily() bool {
	return this.HTTP != nil || this.HTTPS != nil
}

func (this *ServerConfig) IsTCPFamily() bool {
	return this.TCP != nil || this.TLS != nil
}

func (this *ServerConfig) IsUnixFamily() bool {
	return this.Unix != nil
}

func (this *ServerConfig) IsUDPFamily() bool {
	return this.UDP != nil
}

// 判断是否和域名匹配
func (this *ServerConfig) MatchName(name string) bool {
	if len(name) == 0 {
		return false
	}
	if len(this.AliasServerNames) > 0 && configutils.MatchDomains(this.AliasServerNames, name) {
		return true
	}
	for _, serverName := range this.ServerNames {
		if serverName.Match(name) {
			return true
		}
	}
	return false
}

// 判断是否严格匹配
func (this *ServerConfig) MatchNameStrictly(name string) bool {
	for _, serverName := range this.AliasServerNames {
		if serverName == name {
			return true
		}
	}
	for _, serverName := range this.ServerNames {
		if serverName.Name == name {
			return true
		}
	}
	return false
}

// SSL信息
func (this *ServerConfig) SSLPolicy() *sslconfigs.SSLPolicy {
	if this.HTTPS != nil {
		return this.HTTPS.SSLPolicy
	}
	if this.TLS != nil {
		return this.TLS.SSLPolicy
	}
	return nil
}

// 根据条件查找ReverseProxy
func (this *ServerConfig) FindAndCheckReverseProxy(dataType string) (*ReverseProxyConfig, error) {
	switch dataType {
	case "server":
		if this.ReverseProxy == nil {
			return nil, errors.New("reverse proxy not been configured")
		}
		return this.ReverseProxy, nil
	default:
		return nil, errors.New("invalid data type:'" + dataType + "'")
	}
}
