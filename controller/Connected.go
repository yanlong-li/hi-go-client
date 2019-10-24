package controller

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Connected{}, Connected)
}

func Connected(connector connect.Connector) {
	fmt.Println("新连接", connector.GetId())
}
