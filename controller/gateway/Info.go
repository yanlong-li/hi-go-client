package gateway

import (
	"github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-server/packet_model/gateway"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(gateway.Info{}, Info)
}

func Info(info gateway.Info, conn connect.Connector) {

	logger.Debug("收到网关信息", 0, info)
}
