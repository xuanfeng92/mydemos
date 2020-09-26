package ziface

import "net"

// 定义连接模块的抽象层

type IConnection interface {
	// 启动连接，让当前的连接准备开始工作
	Start()
	// 停止连接，结束当前连接工作
	Stop()
	// 获取当前连接的绑定socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接模块的连接ID
	GetConnID() uint32
	// 获取远程哭护短的TCP状态 IP port
	RemoteAddr() net.Addr

	// 发送数据，将数据发送给远程客户端
	SendMsg(uint32, []byte) error

	// 设置属性
	SetProperty(string, interface{})

	// 获取连接属性
	GetProperty(key string) (interface{}, error)

	// 移除连接属性
	RemoveProperty(key string)
}

// 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error