package main

import (
	"HelloWorld/io/network/socket"
	_ "HelloWorldClient/controller"
	_ "HelloWorldServer/packet"
)

func main() {

	socket.Client("127.0.0.1:3000")
}
