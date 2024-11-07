package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"server/global"
	"server/models"
	"server/models/res"
	"server/services/common"
	"server/utils/jwts"
)

// CollResponse 只需要多定义一个收藏时间
type CollResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

func (ArticleApi) ArticleCollListView(c *gin.Context) {
	// 绑定分页的参数
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	// 需要登录
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 查表获取文章ID
	var articleIDList []interface{}
	// 这个接口需要分页
	list, count, err := common.CommList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
	})

	//
	var collMap = map[string]string{}
	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}

	boolSearch := elastic.NewTermsQuery("_id", articleIDList...) // 根据ID搜
	var collList = make([]CollResponse, 0)
	// 传id列表，查es
	result, err := global.EsClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(result.Hits.TotalHits.Value, articleIDList) // 将ID列表也打印出来

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		article.ID = hit.Id
		collList = append(collList, CollResponse{ //这个存
			ArticleModel: article,
			CreatedAt:    collMap[hit.Id],
		})
	}
	res.OkWithList(collList, count, c)
}
