package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/socket/connect"
	"fmt"
)

func HelloWorld2(world packet.HelloWorld, conn *connect.Connector) {

	fmt.Println(world)
	model := packet.HelloWorld{
		Message: "Hello World 我是客户端 2",
	}
	conn.Send(model)
}
