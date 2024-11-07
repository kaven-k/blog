package main

import (
	"server/core"
	"server/global"
	"server/services/es_service"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接es
	global.EsClient = core.EsConnect()
	es_service.DeleteFullTextByArticleID("M9Rd_JIBD9wpuaCLvksx")
}
