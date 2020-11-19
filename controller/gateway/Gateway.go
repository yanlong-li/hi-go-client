package gateway

import (
	"fmt"
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-gateway/packet_model"
	"github.com/yanlong-li/hi-go-gateway/packet_model/client"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"github.com/yanlong-li/hi-go-socket/socket"
	"strconv"
)

func init() {
	route.Register(common.GatewayAndClientGroup, packet_model.Connected{}, func(conn connect.Connector) {
		fmt.Println("成功连接到网关服务器")
		_ = conn.Send(client.GetServerList{})
	})
	route.Register(common.GatewayAndClientGroup, client.ServerList{}, func(serverList client.ServerList, conn connect.Connector) {
		fmt.Println("获取服务器列表成功")

		if len(serverList.List) >= 1 {
			server := serverList.List[0]
			ip := server.IP
			ipString := ip.String()
			port := server.Port
			portStr := strconv.Itoa(int(port))
			go socket.Client(common.ServerAndClientGroup, ipString+":"+portStr)

		} else {

		}

	})
}
