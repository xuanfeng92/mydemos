package znet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/go/zinx/utils"
	"github.com/go/zinx/ziface"
)

/**
封包，拆包，模块
直接面向tcp连接中的数据流，用于处理tcp粘包问题
*/
type DataPack struct{}

// 拆包封包实例的一个初始化方法
func NewDataPack() *DataPack {
	return new(DataPack)
}

// 获取包的头的长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen uint32(4字节) + ID uint32(4字节)
	return 8
}

// 封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	buffer := bytes.NewBuffer([]byte{})

	// 格式： dataLen(4字节)-dataid(4字节)-data
	// 将dataLen写进dataBuff中
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	// 将mesgId写进dataBuff中
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 将data数据写进dataBuff中
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// 拆包方法
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	// 读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// 读dataId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	// 读data内容
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Data); err != nil {
		return nil, err
	}

	// 判断datalen是否已经超出我们的运行的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New(" too large msg data recv!")
	}
	return msg, nil
}
