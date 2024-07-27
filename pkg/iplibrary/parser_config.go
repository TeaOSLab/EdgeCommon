// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package iplibrary

type ParserConfig struct {
	Template    *Template
	EmptyValues []string
	Iterator    func(values map[string]string) error
}
