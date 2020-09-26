package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/go/zinx/utils"
	"github.com/go/zinx/ziface"
)

/**
连接模块
*/

type Connection struct {

	// 当前conn隶属于那个Server
	TcpServer ziface.IServer

	// 当前连接的socket TCP
	Conn *net.TCPConn

	// 连接的ID
	ConnID uint32

	// 当前的连接状态
	isClosed bool

	// 当前连接锁绑定的处理业务方法API
	handleAPI ziface.HandleFunc

	//告知当前连接已推出/停止 channel
	ExitChan chan bool

	// 用于读和写之间的通信
	msgChan chan []byte

	// 消息的管理MsgID 和对应处理业务API关系
	MsgHandler ziface.IMsgHandle

	// 连接属性集合
	property map[string]interface{}
	// 保护连接属性的锁
	propertyLock sync.RWMutex
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	connection := &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClosed:   false,
		msgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
		TcpServer:  server,
		property:   make(map[string]interface{}),
	}
	// 将conn加入到ConnMgr
	connection.TcpServer.GetConnMgr().Add(connection)

	return connection
}

// 启动连接，让当前的连接准备开始工作
func (conn *Connection) Start() {
	fmt.Println("Conn Start()... ConnID=", conn.ConnID)
	// 启动从当前连接读的数据业务
	go conn.StartReader()
	// TODO 启动从当前业务连接写数据的业务
	go conn.StartWriter()

	// 按照开发者传递进来的 创建连接后需要调用的业务，执行对应的hook函数
	conn.TcpServer.CallOnConnStart(conn)
}

// 停止连接，结束当前连接工作
func (conn *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID=", conn.ConnID)
	if conn.isClosed {
		return
	}
	conn.isClosed = true
	// 关闭socket连接
	conn.TcpServer.CallOnConnStop(conn) // 关闭conn之前，调用需要执行的业务
	conn.Conn.Close()

	// 告知writer关闭
	conn.ExitChan <- true

	// 将当前连接从ConnMgr中摘除掉
	conn.TcpServer.GetConnMgr().Remove(conn)

	// 回收资源，出发退出管道
	close(conn.ExitChan)
	close(conn.msgChan)
}

// 获取当前连接的绑定socket conn
func (conn *Connection) GetTCPConnection() *net.TCPConn {
	return conn.Conn
}

// 获取当前连接模块的连接ID
func (conn *Connection) GetConnID() uint32 {
	return conn.ConnID
}

// 获取远程客户端的TCP状态 IP port
func (conn *Connection) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

// 发送数据，将数据发送给远程客户端
// 提供一个sendMsg方法，将我们要发送给客户端的数据，先进性封包，再发送
func (conn *Connection) SendMsg(msgId uint32, data []byte) error {
	if conn.isClosed == true {
		return errors.New("Connection  closed when send msg")
	}
	// 将data封包 msgDataLen-MsgID-Data
	dp := NewDataPack()

	binaryMsg, e := dp.Pack(NewMsgPackage(msgId, data))
	if e != nil {
		fmt.Println("Pack error msg is=", msgId)
		return errors.New("Pack  error msg")
	}
	// 将数据发送给MsgChannel，触发StartWriter操作
	conn.msgChan <- binaryMsg

	return nil
}

// 设置连接属性
func (conn *Connection) SetProperty(key string, value interface{}) {
	// 设置属性，加写锁
	conn.propertyLock.Lock()
	defer conn.propertyLock.Unlock()

	// 添加一个连接属性
	conn.property[key] = value
}

// 获取连接属性
func (conn *Connection) GetProperty(key string) (interface{}, error) {
	// 获取属性，加读锁
	conn.propertyLock.RLock()
	defer conn.propertyLock.RUnlock()

	// 读取属性
	if value, ok := conn.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New(fmt.Sprintf("property [%s] not found!", key))
	}
}

// 移除连接属性
func (conn *Connection) RemoveProperty(key string) {
	// 移除属性，加写锁
	conn.propertyLock.Lock()
	defer conn.propertyLock.Unlock()

	delete(conn.property, key)
}

// 连接的读业务方法
func (conn *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println("[Reader is exit] connID=", conn.ConnID, ", remote addr is ", conn.RemoteAddr().String())
	defer conn.Stop()
	for {
		// 开始组件一个消息包
		dp := NewDataPack()
		// 读取客户端的msg head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error:", err)
			break
		}
		// 拆包，得到msgID 和 msgDataLen 放在msg消息中
		message, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error:", err)
			break
		}
		// 根据dataLen，再次读取Data， 放在msg.Data中
		var data []byte
		if message.GetMsgLen() > 0 {
			data = make([]byte, message.GetMsgLen())
			_, err := io.ReadFull(conn.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("read msg data error:", err)
				break
			}
		}
		message.SetMsgData(data)

		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: conn,
			msg:  message,
		}
		if utils.GlobalObject.WorkerPoolSize > 0 {
			// 表示已经开启了工作池机制，将消息发送给worker工作池处理即可
			conn.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			// 从路由中，找到绑定的Conn对应的router调用
			// 根据绑定好的msgID，找到对应处理api业务
			go conn.MsgHandler.DoMsgHandler(&req)
		}

	}
}

/**
写消息的Goroutine,专门发送给客户端消息的模块
*/
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println("[Writer is exit] connID= ", c.ConnID, ", remote addr is ", c.RemoteAddr().String())

	// 不断地阻塞的等待channel的消息，进行写给客户端
	for {
		select {
		case data := <-c.msgChan:
			// 有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error:", err)
				return
			}
		case <-c.ExitChan:
			// 代表Reader已经退出，此时Writer也要退出
			return
		}
	}
}
