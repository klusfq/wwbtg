package main

import (
	"wwbtg/com/frame"
	"wwbtg/uibook/controller"
)

/**
 * 启动框架
 *	1. 加载基础服务配置
 *  2. 注册对外接口
 *  3. 启动服务
 */
func main() {
	sv := &frame.Serv{}

	sv.Init("uibook", "", controller.ApiList)

	sv.Run()
}
