package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

/**
模拟客户端
 */
func main() {
	fmt.Println("hello client start...")
	time.Sleep(time.Second)
	// 1. 直接连接服务器，得到conn连接
	conn, e := net.Dial("tcp", "127.0.0.1:8999")
	if e!= nil{
		fmt.Println("client start err,exit!")
		return
	}
	for{
		// 不停的写数据
		/********************************* 客户端发送信息过程********************************/
		// 发送封包的消息，消息结构： msgLen/msgId/msgData
		dp := znet.NewDataPack()
		binaryMsg, e := dp.Pack(znet.NewMsgPackage(1, []byte("Zinx hello client Test Message")))
		if e!=nil{
			fmt.Println("hello client Pack error:",e)
			break
		}
		if _, e = conn.Write(binaryMsg); e!=nil{
			fmt.Println("hello client write error:", e)
			break
		}

		/********************************* 客户端接收信息过程********************************/
		// 服务器就应该给我们回复一个message数据，此时需要解包
		// 1. 先读取流中head部分，得到ID和dataLen
		binaryHead:= make([]byte, dp.GetHeadLen())
		if _, e = io.ReadFull(conn, binaryHead); e!=nil{
			fmt.Println("hello client read head error:", e)
			break
		}
		// 2. 将二进制的head拆包到msg结构体中
		message, e := dp.Unpack(binaryHead)
		if e!=nil{
			fmt.Println("hello client unpack head error:", e)
			break
		}
		// 2. 再根据DataLen进行第二次读取，将data读出来
		if message.GetMsgLen()>0{
			// 此时msg是有数据的
			msg := message.(*znet.Message)
			msg.Data = make([]byte, message.GetMsgLen())
			if _,err := io.ReadFull(conn, msg.Data); err!=nil{
				fmt.Println("hello client read msg error:", err)
				break
			}
			fmt.Println("-->hello client recv msg: ID=", msg.GetMsgId(), " len=", msg.GetMsgLen(), " data=", string(msg.GetData()))
		}

		time.Sleep(time.Second*2)
	}
}
