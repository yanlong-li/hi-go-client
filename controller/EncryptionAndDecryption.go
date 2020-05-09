package controller

import (
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorld-GO/io/network/stream"
	"github.com/yanlong-li/HelloWorldServer/packetModel/encrypt"
)

//加密和解密

func init() {

	//route.Register(packetModel.BeforeSending{}, Encryption)
	//route.Register(encrypt.BytesData{}, Decryption)
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

	if len(encryptData.Data) < packet.OpCodeLen+packet.BufLenLen {
		logger.Warning("解密后的数据无效", 0, encryptData.Data)
		return
	}
	//自行处理粘包和拆包，单个包处理后丢给 conn.HandleData()

	//bufLen := buf[:packet.BufLenLen]
	bufData := encryptData.Data[packet.BufLenLen:]
	conn.HandleData(bufData)

}
