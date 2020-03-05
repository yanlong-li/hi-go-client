package controller

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
)

func init() {
	route.Register(packetModel.Connected{}, Connected)
}

func Connected(connector connect.Connector) {
	fmt.Println("新连接", connector.GetId())
}
