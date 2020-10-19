package main

import (
	_ "github.com/yanlong-li/hi-go-client/controller"
	_ "github.com/yanlong-li/hi-go-server/packet_model"
	"github.com/yanlong-li/hi-go-socket/socket"
)

func main() {

	socket.Client("127.0.0.1:3000")
}
