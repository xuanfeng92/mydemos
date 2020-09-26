package apis

import (
	"fmt"

	"github.com/go/zinx/mmo_game_zinx/core"
	pb "github.com/go/zinx/mmo_game_zinx/pd"
	"github.com/go/zinx/ziface"
	"github.com/go/zinx/znet"
	"github.com/golang/protobuf/proto"
)

// 玩家移动
type MoveApi struct {
	znet.BaseRouter
}

func (m *MoveApi) Handle(request ziface.IRequest) {
	// 解析客户端传递进来的proto协议
	proto_msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("Move:Position Unmarshal error:", err)
		return
	}
	//得到当前发送位置的是哪个玩家
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("GetProperty pid error:", err)
		return
	}
	fmt.Printf("Player pid = [%d], move [x=%f, y=%f, z=%f, v=%f]\n", pid, proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)

	//给其他玩家广播当前客户端玩家的位置信息
	player := core.WorldMgr.GetPlayerByPid(pid.(int32))
	// 广播并更新当前玩家的坐标
	player.UpdatePos(proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)
}
