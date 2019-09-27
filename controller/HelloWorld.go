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
	route.Register(packet.UserList{}, UserList)
}

func HelloWorld(world packet.HelloWorld, conn *connect.Connector) {
	fmt.Println("请输入姓名")
	var name string
	//_, err := fmt.Scanln(&name)
	//if err != nil {
	//	log.Fatal("输入失败")
	//}
	name = "张三"
	conn.Send(packet.Login{Username: name, Password: "123456"})
}

func LoginSuccess(token packet.Token, conn *connect.Connector) {
	fmt.Println("登录成功：", token.Token)
	//for {
	fmt.Println("请求用户列表")
	conn.Send(packet.GetUserList{})
	//}

}
func LoginFail(fail packet.LoginFail, conn *connect.Connector) {
	fmt.Println("登录失败：", fail.Message)
	fmt.Println("请输入姓名")
	var name string
	//_, err := fmt.Scanln(&name)
	//if err != nil {
	//	log.Fatal("输入失败")
	//}
	name = "张三"
	conn.Send(packet.Login{Username: name, Password: "123456"})
}

func UserList(list packet.UserList, conn *connect.Connector) {

	fmt.Println(list)

}
