package main

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket"
	_ "github.com/yanlong-li/HelloWorldClient/controller"
	_ "github.com/yanlong-li/HelloWorldServer/packetModel"
)

func main() {

	socket.Client("127.0.0.1:3000")
}
