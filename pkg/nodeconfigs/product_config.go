// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package nodeconfigs

const DefaultProductName = "GoEdge"

// ProductConfig 产品相关设置
type ProductConfig struct {
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}
