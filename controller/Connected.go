package controller

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packet"
)

func init() {
	route.Register(packet.Connected{}, Connected)
}

func Connected(connector connect.Connector) {
	fmt.Println("新连接", connector.GetId())
}
