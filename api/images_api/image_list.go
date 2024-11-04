package images_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/services/common"
)

// ImageListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.PageInfo   false  "查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.CommList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false, // 如果不想显示debug，就将true改为false
	})
	res.OkWithList(list, count, c)
	return
}
