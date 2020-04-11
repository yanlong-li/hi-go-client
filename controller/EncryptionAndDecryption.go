package controller

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/stream"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/encrypt"
	"io/ioutil"
)

//加密和解密

func init() {
	route.Register(packetModel.BeforeSending{}, Encryption)
	route.Register(encrypt.BytesData{}, Decryption)
}

// 加密动作
func Encryption(ps stream.Interface, conn connect.Connector) []byte {

	//简单做了一层封装
	model := encrypt.BytesData{}
	model.Data = ps.ToData()

	ps.Marshal(model)

	return ps.ToData()
}
func Decryption(encryptData encrypt.BytesData, conn connect.Connector) {

	logger.Debug("收到一个加密数据", 0)

	privateKeyBytes, err := ioutil.ReadFile("config/8FF1A48D9B97C67F9B8C4F633708AAA986797D5D.asc")

	if err != nil {
		logger.Warning("读取ecdsa密钥失败", 0, err)
	}
	c, e := helper.DecryptMessageArmored(string(privateKeyBytes), []byte("123456"), string(encryptData.Data))
	if e != nil {
		logger.Warning("解密失败", 0, e)
	}

	conn.HandleData([]byte(c))

}
