package controller

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packet"
)

func init() {
	route.Register(packet.Disconnect{}, Disconnect)
}

func Disconnect(ID uint32) {
	fmt.Println("一个连接断开:", ID)
}
