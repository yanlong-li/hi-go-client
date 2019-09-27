package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/connect"
	"fmt"
)

func init() {
	route.Register(packet.HelloWorld{}, HelloWorld)
	route.Register(packet.Token{}, LoginSuccess)
	route.Register(packet.LoginFail{}, LoginFail)
}

func HelloWorld(world packet.HelloWorld, conn *connect.Connector) {

	conn.Send(packet.Login{Username: "张三", Password: "123456"})
}

func LoginSuccess(token packet.Token, conn *connect.Connector) {
	fmt.Println("登录成功：", token.Token)
}
func LoginFail(fail packet.LoginFail, conn *connect.Connector) {
	fmt.Println("登录失败：", fail.Message)
}
