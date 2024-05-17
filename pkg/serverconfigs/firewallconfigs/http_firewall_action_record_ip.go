// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package firewallconfigs

type HTTPFirewallRecordIPAction struct {
	Type     string        `yaml:"type" json:"type"`
	IPListId int64         `yaml:"ipListId" json:"ipListId"`
	Level    string        `yaml:"level" json:"level"`
	Timeout  int32         `yaml:"timeout" json:"timeout"`
	Scope    FirewallScope `yaml:"scope" json:"scope"`
}
