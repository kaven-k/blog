package article_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/services/es_service"
	"server/services/redis_service"
)

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.EsIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 用户浏览量是在用户查看文章详情的时候，浏览量自动加的
	redis_service.Look(cr.ID)

	model, err := es_service.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	model, err := es_service.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
