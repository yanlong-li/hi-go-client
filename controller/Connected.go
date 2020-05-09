package controller

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user/Login"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user/contacts"
)

func init() {
	route.Register(packetModel.Connected{}, Connected)
	route.Register(Login.Success{}, loginSuccess)
	route.Register(contacts.List{}, func(list contacts.List, conn connect.Connector) {
		logger.Debug("获取朋友列表成功", 0, list.List)
	})
	route.Register(contacts.RequestList{}, func(list contacts.RequestList, conn connect.Connector) {
		logger.Debug("获取新朋友列表成功", 0, list.List)
	})
	route.Register(contacts.Blacklist{}, func(list contacts.Blacklist, conn connect.Connector) {
		logger.Debug("获取黑名单列表成功", 0, list.List)
	})

	route.Register(contacts.SearchUserSuccess{}, func(info contacts.SearchUserSuccess, conn connect.Connector) {
		logger.Debug("搜索用户成功", 0, info)

		//申请加为好友
		conn.Send(contacts.AddContact{Id: info.Id, Remark: "Dev Test Add Contact"})
	})
}

func loginSuccess(success Login.Success, conn connect.Connector) {

	logger.Debug("登陆成功", 0, success.Token)

	// 获取联系人列表
	conn.Send(contacts.GetList{})
	//获取新朋友列表
	conn.Send(contacts.GetRequestList{})
	//获取黑名单列表
	conn.Send(contacts.GetBlacklist{})
	// 搜索这个账户
	conn.Send(contacts.SearchUser{Account: "ahlyl94@gmail.com"})
}

func LoginFail(fail Login.Fail, conn connect.Connector) {

	logger.Debug("登陆失败", 0)
}

func Connected(connector connect.Connector) {
	fmt.Println("新连接", connector.GetId())
	// 发送心跳
	_ = connector.Send(gateway.Heartbeat{Sn: 0})

	// 发送认证请求
	_ = connector.Send(Login.ForEmail{Email: "895185921@qq.com", Password: "123456"})
}
