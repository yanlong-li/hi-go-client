package main

import (
	"HelloWorld/io/network/handle"
	"HelloWorld/io/network/socket"
	"HelloWorldClient/controller"
)

func main() {

	handle.Register(0, controller.HelloWorld)
	handle.Register(1, controller.HelloWorld2)

	socket.Client()
}
