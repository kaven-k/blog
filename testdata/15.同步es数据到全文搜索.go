package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"server/core"
	"server/global"
	"server/models"
	"server/services/es_service"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接es
	global.EsClient = core.EsConnect()

	// 先查所有的文章
	boolSearch := elastic.NewMatchAllQuery()
	res, _ := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())

	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)

		indexList := es_service.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)

		// 创建一个桶
		bulk := global.EsClient.Bulk()
		for _, IndexData := range indexList {
			req := elastic.NewBulkIndexRequest().
				Index(models.FullTextModel{}.Index()).
				Doc(IndexData)
			bulk.Add(req)
		}
		result, err := bulk.Do(context.Background())
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
		fmt.Println(article.Title, "添加成功", "共", len(result.Succeeded()), "条!")
	}
}
