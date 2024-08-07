package serverconfigs

import (
	"testing"

	"github.com/iwind/TeaGo/assert"
)

func TestProtocol_IsHTTPFamily(t *testing.T) {
	a := assert.NewAssertion(t)
	t.Log(ProtocolHTTP.String(), ProtocolHTTPS.String(), ProtocolTCP)
	a.IsTrue(ProtocolHTTP.IsHTTPFamily())
	a.IsTrue(ProtocolHTTP4.IsHTTPFamily())
	a.IsTrue(ProtocolHTTP6.IsHTTPFamily())
	a.IsTrue(ProtocolHTTPS.IsHTTPSFamily())
	a.IsTrue(ProtocolHTTPS4.IsHTTPSFamily())
	a.IsTrue(ProtocolHTTPS6.IsHTTPSFamily())
	a.IsTrue(ProtocolTCP.IsTCPFamily())
	a.IsTrue(ProtocolTCP.IsTCPFamily())
	a.IsTrue(ProtocolTCP6.IsTCPFamily())
	a.IsTrue(ProtocolUDP.IsUDPFamily())
}
