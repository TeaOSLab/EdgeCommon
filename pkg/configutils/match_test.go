package configutils

import (
	"testing"

	"github.com/iwind/TeaGo/assert"
)

func TestMatchKeyword(t *testing.T) {
	a := assert.NewAssertion(t)
	a.IsTrue(MatchKeyword("a b c", "a"))
	a.IsFalse(MatchKeyword("a b c", ""))
	a.IsTrue(MatchKeyword("abc", "BC"))
}
