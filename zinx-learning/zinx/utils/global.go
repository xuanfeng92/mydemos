package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go/zinx/ziface"
)

/**
存储一些有关zinx框架的全局参数，共其他模块使用
一些参数是可以通过zinx.json由用户进行配置
*/
type Global struct {
	/**
	server相关
	*/
	TcpServer ziface.IServer // 当前zinx全局的server对象
	Host      string         // 当前服务器主机监听IP
	TcpPort   int            // 当前服务器主机监听Port
	Name      string         // 当前服务器主机名

	/**
	zinx相关
	*/
	Version           string // 当前zinx版本
	MaxConn           int    // 当前服务器主机允许最大连接数
	MaxPackageSize    uint32 // 当前zinx框架数据包最大值
	WorkerPoolSize    uint32 // 当前工作Worker池的Goroutines数量
	MaxWorkerTaskSize uint32 //允许用户最多开辟多少个Worker（限定条件）
}

/**
定义一个全局对外的global
*/
var GlobalObject *Global

/**
提供一个init方法，初始化当前GlobalObj
*/
func init() {
	// 如果配置文件没有加载，则用默认值
	GlobalObject = &Global{
		Name:              "ZinxServerApp",
		Version:           "v0.9",
		TcpPort:           8999,
		Host:              "0.0.0.0",
		MaxConn:           2,
		MaxPackageSize:    4096,
		WorkerPoolSize:    10,   // Worker工作池队列的数量
		MaxWorkerTaskSize: 1024, // 每个worker对应的消息队列的任务数量最大值
	}
	// 应该尝试从con/zinx.json去加载一些用户自定义的参数
	//GlobalObject.Reload()
}
func (g *Global) Reload() {
	data, e := ioutil.ReadFile("conf/zinx.json")

	if e != nil {
		//panic(e)
		fmt.Println("load conf/zinx.json failed，reason:", e)
	} else {
		// 将json文件数据解析到struct中
		error := json.Unmarshal(data, &GlobalObject)
		if error != nil {
			//panic(error)
			fmt.Println("unmarshal conf/zinx.json failed，reason:", error)
		}
	}

}
