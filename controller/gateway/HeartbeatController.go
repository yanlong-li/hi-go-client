package gateway

import (
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
	"time"
)

func init() {
	route.Register(gateway.Heartbeat{}, Heartbeat)
}

func Heartbeat(heartbeat gateway.Heartbeat, conn connect.Connector) {

	logger.Debug("收到心跳，进入休眠", 0, heartbeat)
	time.Sleep(time.Minute)
	logger.Debug("休眠结束,发送心跳", 0)
	_ = conn.Send(gateway.Heartbeat{Sn: heartbeat.Sn + 1})
}
