package controller

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
)

func init() {
	route.Register(packetModel.Disconnect{}, Disconnect)
}

func Disconnect(ID uint64) {
	fmt.Println("一个连接断开:", ID)
}
