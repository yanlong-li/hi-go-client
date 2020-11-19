package server

import (
	"fmt"
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-server/packet_model"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(common.ServerAndClientGroup, packet_model.Disconnect{}, Disconnect)
	route.Register(common.ServerAndClientGroup, packet_model.Connected{}, func(conn connect.Connector) {

		fmt.Println("连接服务端成功，该断开网关了")
	})
}

func Disconnect(conn connect.Connector) {
	//fmt.Println("一个连接断开:", conn.GetId())
}
