package main

import (
	"fmt"
	"github.com/go/zinx/ziface"
	"github.com/go/zinx/znet"
)

/**
基于zinx框架开发的，服务器端应用程序
 */
func main() {
	// 1.创建一个server句柄，使用zinx的api
	s:=znet.NewServer("[zinx v0.1]")

	// 2. 创建用于连接前后的hook函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	// 给当前zinx框架添加一个自定义的router
	s.AddRouter(0,new(PingRouter))
	s.AddRouter(1,new(HelloRouter))

	// 3. 启动server
	s.Serve()
}

type PingRouter struct {
	znet.BaseRouter // 匿名类型，表示继承了该类型所有的方法
}

type HelloRouter struct {
	znet.BaseRouter
}
// 创建连接之后执行的钩子函数
func DoConnectionBegin(conn ziface.IConnection){
	fmt.Println("==> DoConnectionBegin is called")
	if err := conn.SendMsg(202, []byte("DoConnection Begin")); err!=nil{
		fmt.Println(err)
	}

	// 测试给当前的连接设置一些属性
	conn.SetProperty("name", "测试内置属性")
}
// 创建连接断开之前需要执行的钩子函数
func DoConnectionLost(conn ziface.IConnection)  {
	fmt.Println("==> DoConnectionLost is called")
	fmt.Println("connID =",conn.GetConnID(), " is Lost")

	// 测试连接属性
	if name,err := conn.GetProperty("name"); err == nil{
		fmt.Println("属性[name] 的值为：",name)
	}
}

// 处理业务前的钩子
/*func (br *PingRouter) PreHandle(request ziface.IRequest){
	fmt.Println("Call Router Prehandle..")
	_, e := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if e!=nil{
		fmt.Println("call back before ping error:",e)
	}

}*/
// 处理conn业务的主方法hook
func (br *PingRouter) Handle(request ziface.IRequest){
	fmt.Println("Call Router handle..")
	/*_, e := request.GetConnection().GetTCPConnection().Write([]byte("ping... ping...\n"))
	if e!=nil{
		fmt.Println("call back ping error:",e)
	}*/
	// 先读取客户端数据，再回写ping
	fmt.Println("recv from client:msgID=", request.GetMsgID(), " ,data=",string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping....ping...ping..."))
	if err!=nil{
		fmt.Println("Server send ping msg error:",err)
	}
}

// 处理conn业务的主方法hook
func (br *HelloRouter) Handle(request ziface.IRequest){
	fmt.Println("Call Hello Router handle..")
	/*_, e := request.GetConnection().GetTCPConnection().Write([]byte("ping... ping...\n"))
	if e!=nil{
		fmt.Println("call back ping error:",e)
	}*/
	// 先读取客户端数据，再回写内容
	fmt.Println("recv from client:msgID=", request.GetMsgID(), " ,data=",string(request.GetData()))
	err := request.GetConnection().SendMsg(200, []byte("hello....hello...hello..."))
	if err!=nil{
		fmt.Println("Server send hello msg error:",err)
	}
}
// 处理conn业务之后的钩子方法
/*func (br *PingRouter) PostHandle(request ziface.IRequest){
	fmt.Println("Call Router Posthandle..")
	_, e := request.GetConnection().GetTCPConnection().Write([]byte("post ping...\n"))
	if e!=nil{
		fmt.Println("call back after ping error:",e)
	}
}*/