package znet

import "github.com/go/zinx/ziface"

// 实现router时，先潜入这个BaseRouter基类，然后更具需要对这个基类的方法进行重写就好了
type BaseRouter struct{}

// 以下方法是空实现，因此，后续只需要继承这个baseRouter，然后根据需要实现具体的方法，而不需要实现全部方法来实现IRouter接口
// 处理conn业务之前的钩子方法
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

// 处理conn业务的主方法hook
func (br *BaseRouter) Handle(request ziface.IRequest) {}

// 处理conn业务之后的钩子方法
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
