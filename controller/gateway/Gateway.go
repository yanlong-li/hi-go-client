package gateway

import (
	"fmt"
	common2 "github.com/yanlong-li/hi-go-client/common"
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-gateway/model/server"
	"github.com/yanlong-li/hi-go-gateway/packet_model"
	"github.com/yanlong-li/hi-go-gateway/packet_model/client"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"github.com/yanlong-li/hi-go-socket/socket"
	"strconv"
	"time"
)

func init() {
	route.Register(common.GatewayAndClientGroup, packet_model.Connected{}, func(conn connect.Connector) {
		fmt.Println("Successfully connected to the gateway server")
		common2.Gateway = conn
		_ = conn.Send(client.GetServerList{})
	})
	route.Register(common.GatewayAndClientGroup, client.ServerList{}, func(serverList client.ServerList, conn connect.Connector) {
		fmt.Println("Get server list successfully")
		if len(serverList.List) >= 1 {
			var server1, server2 server.Server
			var w, w2 uint8 = 0, 0
			for _, v := range serverList.List {
				if v.CurrentLoad >= v.PeakLoad {
					continue
				} else if v.CurrentLoad > v.OptimumLoad {
					if w > 0 {
						continue
					}
					// server2
					_w2 := uint8((v.PeakLoad - v.CurrentLoad) / v.PeakLoad * uint32(v.Weight/100))
					if _w2 > w2 {
						w2 = _w2
						server2 = v
					}
				} else {
					// server2
					_w := uint8((v.OptimumLoad - v.CurrentLoad) / v.OptimumLoad * uint32(v.Weight/100))
					if _w > w {
						w = _w
						server1 = v
					}
				}
			}
			if w == 0 && w2 != 0 {
				server1 = server2
				w = w2
				fmt.Println("Enable standby server")
			}
			if w != 0 {
				server1 = serverList.List[0]
				common2.CurrentConnectServer = server1
				ip := server1.IP
				ipString := ip.String()
				port := server1.Port
				portStr := strconv.Itoa(int(port))

				fmt.Println(ipString + ":" + portStr)

				go socket.Client(common.ServerAndClientGroup, ipString+":"+portStr)
			} else {
				fmt.Println("No servers available")
			}
			return
		}
		time.Sleep(time.Second * 5)
		// 再次请求服务器信息
		_ = conn.Send(client.GetServerList{})

	})
}
