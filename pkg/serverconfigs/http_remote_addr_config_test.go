// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package serverconfigs

import (
	"testing"

	"github.com/iwind/TeaGo/assert"
)

func TestHTTPRemoteAddrConfig_IsEmpty(t *testing.T) {
	var a = assert.NewAssertion(t)

	{
		var config = &HTTPRemoteAddrConfig{}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsEmpty())
	}

	{
		var config = &HTTPRemoteAddrConfig{
			Value: "${remoteAddr}",
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsEmpty())
	}

	{
		var config = &HTTPRemoteAddrConfig{
			Value: "${ remoteAddr }",
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(config.IsEmpty())
	}

	{
		var config = &HTTPRemoteAddrConfig{
			Value: "[${remoteAddr}]",
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(config.IsEmpty())
	}
}

func TestHTTPRemoteAddrConfig_Values(t *testing.T) {
	for _, value := range []string{"${remoteAddr}", "${header.x-real-ip}", "${header.x-client-ip,x-real-ip,x-forwarded-for}"} {
		var config = &HTTPRemoteAddrConfig{Value: value}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(value, "=>", config.Values())
	}
}
