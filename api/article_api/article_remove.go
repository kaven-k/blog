package article_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"server/global"
	"server/models"
	"server/models/res"
	"server/services/es_service"
)

// IDListRequest 请求的参数，通过ID查询并删除
type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	// 参数绑定
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 如果文章删除了，用户收藏这篇文章怎么办
	// 解决方案：
	// 1. 删除之后，顺带与这个文章关联的收藏也删除了
	// 2. 用户收藏表，新增一个字段，表示文章是否删除，用户可以删除收藏记录，但是找不到文章去修改收藏数
	// 批量删除
	bulkService := global.EsClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	// 根据ID进行删除
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)

		go es_service.DeleteFullTextByArticleID(id) // 删除全文搜索中的ID
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
	return
}
