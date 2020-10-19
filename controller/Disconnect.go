package controller

import (
	"fmt"
	"github.com/yanlong-li/hi-go-server/packet_model"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(packet_model.Disconnect{}, Disconnect)
}

func Disconnect(conn connect.Connector) {
	fmt.Println("一个连接断开:", conn.GetId())
}
