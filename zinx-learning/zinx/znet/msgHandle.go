package znet

import (
	"fmt"

	"github.com/go/zinx/utils"
	"github.com/go/zinx/ziface"
)

/**
消息处理模块的实现
*/
type MsgHandle struct {
	// 存放每个msgID 对应的处理方法
	Apis map[uint32]ziface.IRouter

	// 负责Worker取任务的消息队列
	TaskQueue []chan ziface.IRequest // 设置了IRequest 的chan集合
	// 负责工作Worker池的worker数量
	WorkerPoolSize uint32
}

// 初始化/创建MsgHandle方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,                               // 从全局配置中获取
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize), // 开辟worker工作池
	}
}

// 调度/执行对应的Router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	// 1. 从request 中找到msgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID=", request.GetMsgID(), " is NOT FOUND! Need to registry!")
		return
	}
	// 根据MsgID调度对应router业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	// 1.  判断当前msg绑定的api处理方法是否存在
	if _, ok := mh.Apis[msgID]; ok {
		// id已经注册
		fmt.Println("repeat api,msgID=", msgID)
	}
	// 2. 添加msg与api的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api MsgID=", msgID, " success!")
}

// 启动一个Worker工作池(开启工作池的操作，只能发生一次，一个zinx只有一个工作池)
func (mh *MsgHandle) StartWorkerPool() {
	// 根据workerPoolSize 分别开启Worker， 每个Worker用一个go来承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 一个worker被启动
		// 1. 当前的worker对应的channel消息队列，开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskSize)
		// 2. 启动当前的Worker，阻塞等待消息从channel传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动一个Worker工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID=", workerID, " is started...")
	for {
		select {
		// 如果有消息进来，出列的就是一个客户端的Request,执行当前Request所绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

// 将消息交给TaskQueue，有worker进行处理(在StartOneWorker方法里面定义的阻塞循环)
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// 1. 分配策略： 平均分配  （扩展：可以考虑分布式的算法进行负载均衡）
	// 根据客户端建立的ConnID进行分配
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(),
		" request MsgID=", request.GetMsgID(), " to WorkerID=", workerID)

	//2. 将消息发送给对应的worker的TaskQueue即可
	mh.TaskQueue[workerID] <- request
}
