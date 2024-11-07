package es_service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"server/services/redis_service"
	"strings"
)

func CommList(option Option) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	if option.Key != "" {
		boolSearch.Must(
			//elastic.NewMatchQuery("title", key),
			//elastic.NewMultiMatchQuery(option.Key, "title", "abstract", "content"),
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}
	// 根据标签搜
	if option.Tag != "" {
		boolSearch.Must(
			//elastic.NewMultiMatchQuery(option.Tag, option.Fields...))
			elastic.NewMultiMatchQuery(option.Tag, "tags"))
	}

	type SortField struct {
		Field     string
		Ascending bool
	}

	sortField := SortField{
		Field:     "created_at", // 设置一个默认字段
		Ascending: false,        // 从小到大 从大到小
	}
	if option.Sort != "" {
		_list := strings.Split(option.Sort, " ")                          // 截取
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") { // 逻辑短路
			sortField.Field = _list[0]
			if _list[1] == "desc" {
				sortField.Ascending = false
			}
			if _list[1] == "asc" {
				sortField.Ascending = true
			}
		}
	}

	//fmt.Println(sortField)

	res, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Highlight(elastic.NewHighlight().Field("title")). // 高亮标题字段
		From(option.GetForm()).
		Sort(sortField.Field, sortField.Ascending).
		Size(option.Limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []models.ArticleModel{}

	diggInfo := redis_service.GetDiggInfo()
	lookInfo := redis_service.GetLookInfo()

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
		title, ok := hit.Highlight["title"]
		if ok {
			model.Title = title[0]
		}
		model.ID = hit.Id
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]

		model.DiggCount = model.DiggCount + digg
		model.LookCount = model.LookCount + look

		demoList = append(demoList, model)
	}
	return demoList, count, err
}

// CommDetail 通过id查询
func CommDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.EsClient.
		Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	// 一般来说是在函数内部处理这个error，要么就是抛给上层函数
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
		return
	}
	model.ID = res.Id
	model.LookCount = model.LookCount + redis_service.GetLook(res.Id)
	return
}

// CommDetailByKeyword 通过关键字查询
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

// ArticleUpdate 更新文章
func ArticleUpdate(id string, data map[string]any) error {
	_, err := global.EsClient.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).
		Do(context.Background())
	return err
}
