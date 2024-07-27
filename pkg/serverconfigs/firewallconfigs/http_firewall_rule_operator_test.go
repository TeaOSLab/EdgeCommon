package firewallconfigs_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
)

func TestRuleOperator_Markdown(t *testing.T) {
	var result = []string{}
	for _, def := range firewallconfigs.AllRuleOperators {
		def.Description = strings.ReplaceAll(def.Description, "<code-label>", "`")
		def.Description = strings.ReplaceAll(def.Description, "</code-label>", "`")

		var row = "## " + def.Name + "\n"
		row += "* 名称：" + def.Name + "\n"
		row += "* 代号：`" + def.Code + "`\n"
		row += "* 描述：" + def.Description + "\n"
		result = append(result, row)
	}

	fmt.Print(strings.Join(result, "\n") + "\n")
}
