package gateway

import (
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
)

func init() {
	route.Register(gateway.Info{}, Info)
}

func Info(info gateway.Info, conn connect.Connector) {

	logger.Debug("收到网关信息", 0, info)
}
