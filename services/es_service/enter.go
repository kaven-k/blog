package es_service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
)

func CommList(key string, page, limit int) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []models.ArticleModel{}
	for _, hit := range res.Hits.Hits {
		var model models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		model.ID = hit.Id
		demoList = append(demoList, model)
	}
	return demoList, count, err
}

func CommDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.EsClient.
		Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	// 一般来说是在函数内部处理这个error，要么就是抛给上层函数
	if err != nil {
		//logrus.Error(err.Error())
		return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
		//logrus.Error(err)
		return
	}
	model.ID = res.Id
	return
}

func CommDetailByKeyword(key string) (model models.ArticleModel, err error) {
	res, err := global.EsClient.
		Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	// 一般来说是在函数内部处理这个error，要么就是抛给上层函数
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &model)
	if err != nil {

		return
	}
	model.ID = hit.Id
	return
}
