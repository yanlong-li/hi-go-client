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
	logger.SetLevel(logger.ALL)

	go socket.Client(common.GatewayAndClientGroup, "127.0.0.1:3001")

	for {
		time.Sleep(time.Second * 5)
		fmt.Println(connect.Count())
	}
}
