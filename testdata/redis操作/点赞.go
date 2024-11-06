package main

import (
	"fmt"
	"server/core"
	"server/global"
	"server/services/redis_service"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()

	global.Redis = core.ConnectRedis()

	redis_service.Digg("M9Rd_JIBD9wpuaCLvksx")
	fmt.Println(redis_service.GetDigg("M9Rd_JIBD9wpuaCLvksx"))

	fmt.Println(redis_service.GetDiggInfo())
}
