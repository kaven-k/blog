package advertisement_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/services/common"
	"strings"
)

// AdvertisementListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo   false  "查询参数"
// @Router /api/ad [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertisementApi) AdvertisementListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 列表页就可以拿admin的参数
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		// admin来的
		isShow = false
	}
	// 判断 Referer 是否包换admin，如果是 就全部返回，不是 就返回is_show = true
	list, count, _ := common.CommList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		//Debug:    true,
	})
	res.OkWithList(list, count, c)
}
