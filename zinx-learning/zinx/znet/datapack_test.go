package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// 只负责测试datapack拆包，封包单元测试
func TestDataPack(t *testing.T) {
	/**
	模拟服务器
	*/
	// 1.创建socketTcp
	listener, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}
	// 创建一个go ,负责从客户端处理业务
	go func() {
		//2. 从客户端读取数据，拆包处理
		for {
			conn, e := listener.Accept()
			if e != nil {
				fmt.Println("accept err:", e)
				break
			}
			go func(conn2 net.Conn) {
				// 处理客户端请求
				// 拆包过程
				// 1. 定义一个拆包对象dp
				dp := NewDataPack()
				for {
					// 2. 第一次从conn读，把包的head读出来
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error")
						return
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err:", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						// msg是有数据的。需要进行第二次读取
						// 3. 第二次从conn读，根据head中的dataLen，再读取data内容
						msg := msgHead.(*Message) // 这里转换成Message，注意这里的类型转换是指针
						msg.Data = make([]byte, msg.GetMsgLen())
						// 根据dataLen的长度再次从io中读取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack err:", err)
							return
						}
						// 完整的一个消息已经读取完毕
						fmt.Println("-->Recv MsgID:", msg.Id, ", dataLen=", msg.GetMsgLen(), " data =", string(msg.GetData()))
					}
				}
			}(conn)
		}
	}()

	/**
	模拟客户端
	*/

	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}
	// 创建粘包过程
	dp := NewDataPack()

	// 封装第一个消息包
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'm'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error:", err)
		return
	}
	// 封装第二个消息包
	msg2 := &Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte{'n', 'i', 'h', 'a', 'o', '!', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 error:", err)
		return
	}
	// 测试粘包：将两个包粘在一起
	sendData := append(sendData1, sendData2...)
	// 一次性发送给服务端
	_, err = conn.Write(sendData)

	// 客户端阻塞
	select {}

}
