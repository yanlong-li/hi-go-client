package main

import (
	"HelloWorld/io/network/socket"
	_ "HelloWorldClient/controller"
	_ "HelloWorldServer/packet"
)

func main() {

	socket.Client()
}
