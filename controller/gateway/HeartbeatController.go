package gateway

import (
	"github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-server/packet_model/gateway"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
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
