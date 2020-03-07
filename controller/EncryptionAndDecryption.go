package controller

import (
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/stream"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/encrypt"
)

//加密和解密

func init() {
	route.Register(packetModel.BeforeSending{}, Encryption)
	route.Register(encrypt.BytesData{}, Decryption)
}

// 加密动作
func Encryption(ps stream.Interface) []byte {

	//简单做了一层封装
	model := encrypt.BytesData{}
	model.Data = ps.ToData()

	ps.Marshal(model)

	return ps.ToData()
}
func Decryption(encryptData encrypt.BytesData, conn connect.Connector) {

	logger.Debug("收到一个加密数据", 0)
	conn.HandleData(encryptData.Data)

}
