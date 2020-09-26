package ziface

/**
路由的抽象接口： 路由李的数据都是IRequest
 */

 type IRouter interface {
 	// 处理conn业务之前的钩子方法
 	PreHandle(request IRequest)
 	// 处理conn业务的主方法hook
 	Handle(request IRequest)
 	// 处理conn业务之后的钩子方法
 	PostHandle(request IRequest)
 }
