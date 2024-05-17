// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package reporterconfigs

type NodeConfig struct {
	Id int64 `json:"id"`
}

func (this *NodeConfig) Init() error {
	return nil
}
