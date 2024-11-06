package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"server/core"
	"server/global"
	"server/models"
	"server/services/redis_service"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()

	global.Redis = core.ConnectRedis()
	global.EsClient = core.EsConnect()

	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	diggInfo := redis_service.GetDiggInfo()
	lookInfo := redis_service.GetLookInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article) // 进行json解析

		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]

		newDigg := article.DiggCount + digg //
		newLook := article.LookCount + look

		if article.DiggCount == newDigg && article.LookCount == newLook { // 如果点赞数没变化，就不用变
			logrus.Info(article.Title, "点赞数和浏览量都无变化")
			continue
		}
		_, err := global.EsClient. // 调更新的语法
						Update().
						Index(models.ArticleModel{}.Index()).
						Id(hit.Id).
						Doc(map[string]int{
				"digg_count": newDigg,
				"look_count": newLook,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Infof("%s,点赞数据同步成功，点赞数 %d 浏览数 %d", article.Title, newDigg, newLook) // 有变化，打印出来
	}
	redis_service.DiggClear() // 变化完之后，需要将缓存中的数据清除掉
	redis_service.LookClear() // 变化完之后，需要将缓存中的数据清除掉
}
