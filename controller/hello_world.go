package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/socket/connect"
)

func HelloWorld(world packet.HelloWorld, conn *connect.Connector) {

	//fmt.Println(world)
	model := packet.HelloWorld{
		Message: "Hello World 我是客户端",
	}
	conn.Send(model)

}
