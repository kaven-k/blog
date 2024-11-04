package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// TagUpdateView 修改标签
// @Tags 标签管理
// @Summary 修改标签
// @Description 修改标签
// @Param data body AdvertisementRequest  true  "标签的一些参数"
// @Router /api/ad/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	// 参数校验
	var cr TagRequest
	// 通过ShouldBindJSON方法尝试将客户端传来的 JSON 数据绑定到AdvertisementRequest结构体上
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil { // 如果err为空就代表查到了
		res.FailWithMessage("标签不存在", c)
		return
	}
	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error
	// 结构体转map的第三方包
	if err != nil {
		global.Log.Error(err) // 错误就对外提出
		res.FailWithMessage("修改标签失败", c)
		return
	}
	res.OkWithMessage("修改标签成功", c)
}
