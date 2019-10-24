package controller

import (
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Disconnect{}, Disconnect)
}

func Disconnect(ID uint32) {
	fmt.Println("一个连接断开:", ID)
}
