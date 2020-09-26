package main

import (
	"github.com/go/zinx/mmo_game_zinx/apis"
	"github.com/go/zinx/mmo_game_zinx/core"
	"github.com/go/zinx/ziface"
	"github.com/go/zinx/znet"
)

func main() {
	// 创建zinx server句柄
	server := znet.NewServer("MMO Game Zinx")

	// 连接创建和销毁的hook钩子函数
	server.SetOnConnStart(OnConnectionAdd) // 注册连接创建后的操作
	server.SetOnConnStop(OnConnectionLost) // 注册连接断开之前的操作

	// 注册一些路由业务
	server.AddRouter(2, new(apis.WorldChatApi)) // 聊天路由处理
	server.AddRouter(3, new(apis.MoveApi))      // 移动路由处理

	// 启动业务
	server.Serve()
}

/**
当前客户端建立连接之后的hook函数
*/
func OnConnectionAdd(conn ziface.IConnection) {
	// 创建一个Player对象
	player := core.NewPlayer(conn)
	// 给客户端发送MsgID:1 的消息 登陆确认
	player.SynPid()

	// 给客户端发送MsgID:200 的消息 发送的内容为位置信息
	player.BrodCastStartPosition()

	// 将当前新上线的玩家添加到WorldManager中
	core.WorldMgr.AddPlayer(player)

	// 将该玩家的信息绑定到当前连接
	conn.SetProperty("pid", player.Pid)

	// 同步周边玩家，告知他们当前玩家已经上线广播当前玩家的位置
	player.SyncSurrounding()

	//fmt.Println("====> Player pid=", player.Pid, " is arrived <=====")
}

// 给当前连接断开之前触发的hook钩子函数
func OnConnectionLost(conn ziface.IConnection) {
	//  通过连接属性pid获取当前玩家
	pid, _ := conn.GetProperty("pid")
	player := core.WorldMgr.GetPlayerByPid(pid.(int32))

	// 下线操作
	player.Offline()
}
