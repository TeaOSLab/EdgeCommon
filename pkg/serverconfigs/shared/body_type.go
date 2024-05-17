// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package shared

type BodyType = string

const (
	BodyTypeURL  BodyType = "url"
	BodyTypeHTML BodyType = "html"
)

func FindAllBodyTypes() []*Definition {
	return []*Definition{
		{
			Name: "HTML",
			Code: BodyTypeHTML,
		},
		{
			Name: "读取URL",
			Code: BodyTypeURL,
		},
	}
}
