// 程序的入口
package main

import (
	"server/core"
	_ "server/docs"
	"server/flag"
	"server/global"
	"server/routers"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// 连接Redis
	global.Redis = core.ConnectRedis()
	// 连接es
	global.EsClient = core.EsConnect()

	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("server运行在：%s", addr)
	if err := router.Run(addr); err != nil {
		global.Log.Fatal(err)
	}
}
