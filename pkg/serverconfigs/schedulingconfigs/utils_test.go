// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package schedulingconfigs

import "testing"

func TestFindSchedulingType(t *testing.T) {
	t.Logf("%p", FindSchedulingType("roundRobin"))
	t.Logf("%p", FindSchedulingType("roundRobin"))
	t.Logf("%p", FindSchedulingType("roundRobin"))
}
