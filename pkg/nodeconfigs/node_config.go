package nodeconfigs

import (
	"encoding/json"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/Tea"
	"io/ioutil"
)

var sharedNodeConfig *NodeConfig = nil

type NodeConfig struct {
	Id      int64                         `yaml:"id" json:"id"`
	NodeId  string                        `yaml:"nodeId" json:"nodeId"`
	IsOn    bool                          `yaml:"isOn" json:"isOn"`
	Servers []*serverconfigs.ServerConfig `yaml:"servers" json:"servers"`
	Version int64                         `yaml:"version" json:"version"`
	Name    string                        `yaml:"name" json:"name"`
	MaxCPU  int32                         `yaml:"maxCPU" json:"maxCPU"`

	// 全局配置
	GlobalConfig *serverconfigs.GlobalConfig `yaml:"globalConfig" json:"globalConfig"` // 全局配置

	// TOA配置
	TOA *TOAConfig `yaml:"toa" json:"toa"`

	paddedId         string
	cachePolicies    []*serverconfigs.HTTPCachePolicy
	firewallPolicies []*firewallconfigs.HTTPFirewallPolicy
}

// 取得当前节点配置单例
func SharedNodeConfig() (*NodeConfig, error) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	if sharedNodeConfig != nil {
		return sharedNodeConfig, nil
	}

	data, err := ioutil.ReadFile(Tea.ConfigFile("node.json"))
	if err != nil {
		return &NodeConfig{}, err
	}

	config := &NodeConfig{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return &NodeConfig{}, err
	}

	sharedNodeConfig = config
	return config, nil
}

// 重置节点配置
func ResetNodeConfig(nodeConfig *NodeConfig) {
	shared.Locker.Lock()
	sharedNodeConfig = nodeConfig
	shared.Locker.Unlock()
}

// 初始化
func (this *NodeConfig) Init() error {
	this.paddedId = fmt.Sprintf("%08d", this.Id)

	// servers
	for _, server := range this.Servers {
		err := server.Init()
		if err != nil {
			return err
		}
	}

	// global config
	if this.GlobalConfig != nil {
		err := this.GlobalConfig.Init()
		if err != nil {
			return err
		}
	}

	// cache policies
	this.cachePolicies = []*serverconfigs.HTTPCachePolicy{}
	for _, server := range this.Servers {
		if server.Web != nil {
			this.lookupWeb(server.Web)
		}
	}

	// firewall policies
	this.firewallPolicies = []*firewallconfigs.HTTPFirewallPolicy{}
	for _, server := range this.Servers {
		if server.Web != nil {
			this.lookupWeb(server.Web)
		}
	}

	// TOA
	if this.TOA != nil {
		err := this.TOA.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

// 根据网络地址和协议分组
func (this *NodeConfig) AvailableGroups() []*serverconfigs.ServerGroup {
	groupMapping := map[string]*serverconfigs.ServerGroup{} // protocol://addr => Server Group
	for _, server := range this.Servers {
		if !server.IsOn {
			continue
		}
		for _, addr := range server.FullAddresses() {
			group, ok := groupMapping[addr]
			if ok {
				group.Add(server)
			} else {
				group = serverconfigs.NewServerGroup(addr)
				group.Add(server)
			}
			groupMapping[addr] = group
		}
	}
	result := []*serverconfigs.ServerGroup{}
	for _, group := range groupMapping {
		result = append(result, group)
	}
	return result
}

// 获取使用的所有的缓存策略
func (this *NodeConfig) AllCachePolicies() []*serverconfigs.HTTPCachePolicy {
	return this.cachePolicies
}

// 获取使用的所有的WAF策略
func (this *NodeConfig) AllHTTPFirewallPolicies() []*firewallconfigs.HTTPFirewallPolicy {
	return this.firewallPolicies
}

// 写入到文件
func (this *NodeConfig) Save() error {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	data, err := json.Marshal(this)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(Tea.ConfigFile("node.json"), data, 0777)
}

// 获取填充后的ID
func (this *NodeConfig) PaddedId() string {
	return this.paddedId
}

// 查找Web中的缓存策略、防火墙策略等
func (this *NodeConfig) lookupWeb(web *serverconfigs.HTTPWebConfig) {
	if web == nil {
		return
	}

	// cache
	if web.Cache != nil && len(web.Cache.CacheRefs) > 0 {
		for _, cacheRef := range web.Cache.CacheRefs {
			if cacheRef.CachePolicy != nil && !this.hasCachePolicy(cacheRef.CachePolicyId) {
				this.cachePolicies = append(this.cachePolicies, cacheRef.CachePolicy)
			}
		}
	}

	// firewall
	if web.FirewallPolicy != nil && !this.hasHTTPFirewallPolicy(web.FirewallPolicy.Id) {
		this.firewallPolicies = append(this.firewallPolicies, web.FirewallPolicy)
	}

	for _, location := range web.Locations {
		this.lookupWeb(location.Web)
	}
}

// 检查缓存策略是否已收集
func (this *NodeConfig) hasCachePolicy(cachePolicyId int64) bool {
	for _, cachePolicy := range this.cachePolicies {
		if cachePolicy.Id == cachePolicyId {
			return true
		}
	}
	return false
}

// 检查防火墙策略是否已收集
func (this *NodeConfig) hasHTTPFirewallPolicy(firewallPolicyId int64) bool {
	for _, p := range this.firewallPolicies {
		if p.Id == firewallPolicyId {
			return true
		}
	}
	return false
}
