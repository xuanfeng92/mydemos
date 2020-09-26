package apis

import (
	"fmt"

	"github.com/go/zinx/mmo_game_zinx/core"
	pb "github.com/go/zinx/mmo_game_zinx/pd"
	"github.com/go/zinx/ziface"
	"github.com/go/zinx/znet"
	"github.com/golang/protobuf/proto"
)

type WorldChatApi struct {
	znet.BaseRouter
}

// 处理conn业务的主方法hook
func (wc *WorldChatApi) Handle(request ziface.IRequest) {
	// 1. 解析客户端传递进来的proto协议
	proto_msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("Talk Unmarshal error:", err)
		return
	}
	// 2. 从属性中，获取当前的聊天数据是属于哪个玩家发送的
	pid, err := request.GetConnection().GetProperty("pid")

	// 3. 根据pid得到对应的player对象
	player := core.WorldMgr.GetPlayerByPid(pid.(int32))

	// 4. 将这个消息广播给其他全部在线玩家
	player.Talk(proto_msg.Content)
}
