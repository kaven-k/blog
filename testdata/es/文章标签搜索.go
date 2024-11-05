package main

import (
	"server/core"
	"server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接es
	global.EsClient = core.EsConnect()

	/*
		[{"tag": "go", "article_count": 2, "article_list": [] }]
	*/

}
