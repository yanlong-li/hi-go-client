package controller

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel/encryp"
)

//加密和解密

func init() {
	route.Register(encryp.EncryptData{}, Encryption)
}

// 加密动作
func Encryption(data encryp.EncryptData, connector connect.Connector) {
	fmt.Println(data)
	connector.HandleData(data.Data)
}
