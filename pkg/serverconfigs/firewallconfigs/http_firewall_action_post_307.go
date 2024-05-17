// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package firewallconfigs

type HTTPFirewallPost307Action struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	Life  int32         `yaml:"life" json:"life"`
	Scope FirewallScope `yaml:"scope" json:"scope"`
}
