package server

import (
	common2 "github.com/yanlong-li/hi-go-client/common"
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-gateway/packet_model/client"
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-server/packet_model"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"time"
)

func init() {
	route.Register(common.ServerAndClientGroup, packet_model.Disconnect{}, Disconnect)
	route.Register(common.ServerAndClientGroup, packet_model.Connected{}, func(conn connect.Connector) {
		common2.Server = conn
		logger.Debug("The connection to the server is successful. It's time to disconnect the gateway", 0)

		_ = common2.Gateway.Send(client.ConnectServerSuccess{Id: common2.CurrentConnectServer.Id})

		time.Sleep(time.Second)
		common2.Gateway.Disconnect()
	})
}

func Disconnect(conn connect.Connector) {
	//fmt.Println("一个连接断开:", conn.GetId())
}
