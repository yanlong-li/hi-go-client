package gateway

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
)

func init() {
	route.Register(gateway.Heartbeat{}, Heartbeat)
}

func Heartbeat(heartbeat gateway.Heartbeat, conn connect.Connector) {

	conn.Send(gateway.Heartbeat{Sn: heartbeat.Sn + 1})
}
