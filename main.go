package main

import (
	"fmt"
	_ "github.com/yanlong-li/hi-go-client/controller"
	"github.com/yanlong-li/hi-go-gateway/common"
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/socket"
	"time"
)

func main() {
	logger.SetLevel(logger.INFO)
	test()
	for {
		time.Sleep(time.Second * 5)
		fmt.Println(connect.Count())
	}
}

func test() {
	for i := 0; i < 20000; i++ {
		go conn()
	}
}

func conn() {
	defer func() {
		socket.Client(common.GatewayAndClientGroup, "gateway:3001")
	}()

	socket.Client(common.GatewayAndClientGroup, "gateway:3001")
}
