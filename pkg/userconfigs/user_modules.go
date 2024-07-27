// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package userconfigs

type UserModule = string

const (
	UserModuleCDN      UserModule = "cdn"
	UserModuleAntiDDoS UserModule = "antiDDoS"
	UserModuleNS       UserModule = "ns"
)

var DefaultUserModules = []UserModule{UserModuleCDN}
