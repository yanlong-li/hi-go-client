package controller

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet"
	"fmt"
	"log"
)

func init() {
	route.Register(packet.HelloWorld{}, HelloWorld)
	route.Register(packet.Token{}, LoginSuccess)
	route.Register(packet.LoginFail{}, LoginFail)
	route.Register(packet.UserList{}, UserList)
	route.Register(packet.GlobalMessage{}, GlobalMessage)
	route.Register(packet.Logout{}, UserLogout)
	route.Register(packet.NewUser{}, NewUserOnline)
}

func HelloWorld(conn connect.Connector) {
	fmt.Println("请输入昵称")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal("输入失败")
	}
	//name = "张三2"
	conn.Send(packet.Login{Username: name, Password: "123456"})
}

func LoginSuccess(token packet.Token, conn connect.Connector) {
	fmt.Println("登录成功：", token.Token)
	//for {
	fmt.Println("请求用户列表")
	conn.Send(packet.GetUserList{})
	//}

}
func LoginFail(fail packet.LoginFail, conn connect.Connector) {
	fmt.Println("登录失败：", fail.Message, fail.Code)
	fmt.Println("请输入昵称:")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal("输入失败")
	}
	conn.Send(packet.Login{Username: name, Password: "123456"})
}

// 用户列表
func UserList(list packet.UserList, conn connect.Connector) {

	fmt.Println("打印用户列表：")
	fmt.Println(list)

	fmt.Println("输入消息：")
	for {
		var message string
		_, err := fmt.Scanln(&message)
		if err != nil {
			fmt.Println("输入失败")
			continue
		}
		conn.Send(packet.Message{Content: message})
	}

}

func GlobalMessage(message packet.GlobalMessage, conn connect.Connector) {
	fmt.Println(message.Nickname, ":", message.Content)
	//fmt.Println(message.User, time.Unix(message.User.LoginTime, 0))
}

func UserLogout(logout packet.Logout, conn connect.Connector) {
	fmt.Printf("用户%s离开了服务器\n", logout.Nickname)
}

func NewUserOnline(user packet.NewUser, conn connect.Connector) {
	fmt.Println("掌声欢迎 “", user.Nickname, "” 上线")
}
