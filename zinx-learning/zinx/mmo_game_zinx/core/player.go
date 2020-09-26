package core

import (
	"fmt"
	"math/rand"
	"sync"

	pb "github.com/go/zinx/mmo_game_zinx/pd"
	"github.com/go/zinx/ziface"
	"github.com/golang/protobuf/proto"
)

// 玩家对象
type Player struct {
	Pid  int32              // 玩家ID
	Conn ziface.IConnection // 玩家当前的连接（用于和客户端的连接）
	X    float32            // 平面x坐标
	Y    float32            // 高度
	Z    float32            // 平面y坐标（注意，不是Y）
	V    float32            // 旋转角度0-360
}

// Player ID 生成器
var PidGen int32 = 1  //用来生成玩家ID的计数器
var IdLock sync.Mutex // 保护PidGen的锁

// 创建玩家
func NewPlayer(conn ziface.IConnection) *Player {
	// 生成一个玩家ID
	IdLock.Lock()

	id := PidGen
	PidGen++

	IdLock.Unlock()

	// 创建一个玩家对象
	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)), // 随机在160坐标点，基于X轴若干偏移
		Y:    0,
		Z:    float32(140 + rand.Intn(20)), // 随机在140坐标点，给予Y轴若干偏移
		V:    0,
	}
	return p
}

/**
提供一个发送给客户端消息的方法
主要是将pd的protobuf数据序列化，再调用zinx的sendMsg方法
*/
func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	// 将proto Message 结构序列化 转换成二进制
	msg, e := proto.Marshal(data)
	if e != nil {
		fmt.Println("marshal msg err:", e)
		return
	}

	// 将二进制文件，通过zinx框架的sendMsg数据发送给客户端
	if p.Conn == nil {
		fmt.Println("connection in player is nil")
		return
	}
	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("Player SendMsg error!")
		return
	}

	return
}

// 告知客户端玩家Pid,同步已生成玩家ID给客户端
func (p *Player) SynPid() {
	// 组件MsgID:1 的proto数据
	data := &pb.SyncPid{
		Pid: p.Pid,
	}

	// 将消息发送给客户端
	p.SendMsg(1, data)
}

// 广播玩家自己的出生地
func (p *Player) BrodCastStartPosition() {
	// 组件MsgID:200 的proto数据
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, // 代表广播的是位置坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	// 将消息发送给客户端
	p.SendMsg(200, msg)
}

// 玩家广播世界聊天消息
func (p *Player) Talk(content string) {
	// 1. 组件MsgID:200 proto数据
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1, // tp-1:代表聊天广播
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	// 2. 得到当前世界所有玩家
	players := WorldMgr.GetAllPlayers()

	// 3. 向所有的玩家（包括自己）
	for _, player := range players {
		// 给所有的玩家发送消息
		player.SendMsg(200, proto_msg)
	}
}

// 同步玩家上线的位置信息
func (p *Player) SyncSurrounding() {
	//  1.	获取当前玩家周围的玩家有哪些（九宫格）
	pids := WorldMgr.AoiMgr.GetPidsByPos(p.X, p.Z)
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldMgr.GetPlayerByPid(int32(pid)))
	}
	//  2.	将当前玩家的位置信息通过MsgID:200 发送给周围玩家
	//  2.1 组件MsgID：200 proto数据
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, // tp2 代表广播坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	//  2.2 向周围玩家的格子客户端发送 200消息
	for _, player := range players {
		player.SendMsg(200, proto_msg)
	}

	//	3.  需要将周围的全部玩家的信息发送给当前玩家客户端
	// 3.1 组件MsgID:202 proto数据
	players_proto_msg := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		p := &pb.Player{
			Pid: player.Pid,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		players_proto_msg = append(players_proto_msg, p)
	}
	SyncPlayer_proto_msg := &pb.SyncPlayers{
		Ps: players_proto_msg[:],
	}
	// 3.2 将组建好的数据发送给当前玩家
	p.SendMsg(202, SyncPlayer_proto_msg)
}

// 广播并更新当前玩家的坐标
func (p *Player) UpdatePos(x float32, y float32, z float32, v float32) {
	//更新当前玩家player对象的坐标
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v

	//组建广播proto协议，MsgID:200 Tp=4
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4, // tp-4 代表移动的消息
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	//获取当前玩家的周边玩家AOI九宫格内的玩家
	players := p.GetSurroundingPlayer()

	//一次性给每个玩家对应的客户端发送当前玩家位置更新的消息
	for _, player := range players {
		player.SendMsg(200, proto_msg)
	}

}

// 根据当前玩家的位置，获取周围所有玩家
func (p *Player) GetSurroundingPlayer() []*Player {
	// 得到当前AOI九宫格内所有玩家PID
	pids := WorldMgr.AoiMgr.GetPidsByPos(p.X, p.Z)

	// 将所有的pid对应的Player放到player切片中
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldMgr.GetPlayerByPid(int32(pid)))
	}
	return players
}

// 处理玩家下线
func (p *Player) Offline() {
	//	得到当前玩家周边的九宫格内都有哪些玩家
	players := p.GetSurroundingPlayer()

	//	给周围的玩家广播MsgID:201消息
	proto_msg := &pb.SyncPid{
		Pid: p.Pid,
	}

	for _, player := range players {
		// 如果是自己下线，不用管
		if !(player.Pid == p.Pid){
			player.SendMsg(201, proto_msg)
		}
	}

	//将当前玩家从AOI管理器删除
	WorldMgr.AoiMgr.RemoveFromGridByPos(int(p.Pid), p.X, p.Z)

	//将当前玩家从世界管理器删除
	WorldMgr.RemovePlayerByPid(p.Pid)

}
