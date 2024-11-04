package tag_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/services/common"
)

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description 标签列表
// @Param data query models.PageInfo   false  "查询参数"
// @Router /api/ad [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.CommList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	// 需要展示这个标签下的文章的数量
	res.OkWithList(list, count, c)
}
