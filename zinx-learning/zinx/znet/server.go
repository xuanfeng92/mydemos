package znet

import (
	"fmt"
	"net"

	"github.com/go/zinx/ziface"

	"github.com/go/zinx/utils"
)

// IServer接口的实现。定义一个Server的服务模块

type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定版本
	IPVersion string
	// 服务器监听IP
	IP string
	// 服务器监听端口
	Port int
	// 当前server的消息管理模块，用来绑定MsgID与对应的处理业务API关系
	MsgHandler ziface.IMsgHandle

	// 该server的连接管理器
	ConnMgr ziface.IConnManager

	// 该Server创建连接后自动调用Hook函数--OnConnStart
	OnConnStart func(conn ziface.IConnection)
	// 该Server销毁连接之前自动调用的Hook函数--OnConnStop
	OnConnStop func(conn ziface.IConnection)
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP : %s. Port %d, starting\n", s.IP, s.Port)
	go func() {
		// 0. 开启消息队列及worker工作池
		s.MsgHandler.StartWorkerPool()

		// 1.获取一个tcp的addr
		addr, e := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if e != nil {
			fmt.Println("Resolve tcp addr err:", e)
			return
		}
		// 2. 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err ", err)
			return
		}
		fmt.Println("start Zinx server success, ", s.Name, " success, Listening...")
		// 3. 阻塞等待客户端连接，处理客户端连接业务
		var cid uint32 = 0
		for {
			// 4. 如果有客户端连接过来，阻塞会返回！！！！
			conn, e := listener.AcceptTCP()
			if e != nil {
				fmt.Println("Accept err", e)
				continue
			}
			// 将连接加入连接管理中ConnManager
			// 1.设置最大连接个数，如果超过最大连接，那么关闭此链接
			if s.ConnMgr.Len() > utils.GlobalObject.MaxConn {
				fmt.Println("Too Many Connections,MaxConn=", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
				// TODO 优化：可以给客户端相应的一个超出最大连接的错误包
			}

			fmt.Printf("remote IP 【%s】 coming in \n", conn.RemoteAddr().String())
			// 将处理新连接的业务方法，和conn进行绑定，得到我们的连接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			dealConn.Start()
		}
	}()
}

// 停止服务器
func (s *Server) Stop() {
	// TODO 将一些服务的资源，状态或者一些已经开启的连接信息，进行停止或回收
	fmt.Println("[STOP] zinx server name:", s.Name)
	s.ConnMgr.ClearConn()
}

// 运行服务器
func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO 做一些启动服务之后的额外业务
	// 阻塞一下，防止server启动马上done掉goroutine
	select {}

}

// 初始化server模块方法
func NewServer(name string) *Server {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

// 路由功能：给当前的服务注册一个路由方法，共客户端的连接处理使用
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router success!!,msgID=", msgID)
}

// 获取当前server对应的连接管理器connMgr
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

// 注册OnConnStart钩子函数的方法
func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// 注册OnConnStop钩子函数的方法
func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用OnConnStart钩子函数的方法
func (s *Server) CallOnConnStart(connection ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("--> Call OnConnStart()....")
		s.OnConnStart(connection)
	}
}

// 调用OnConnStop钩子函数的方法
func (s *Server) CallOnConnStop(connection ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("--> Call OnConnStop()....")
		s.OnConnStop(connection)
	}
}
