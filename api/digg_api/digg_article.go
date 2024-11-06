package digg_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/services/redis_service"
)

func (DiggApi) DiggArticleView(c *gin.Context) {
	// 参数绑定
	var cr models.EsIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 对长度校验
	// 查es
	redis_service.Digg(cr.ID)
	res.OkWithMessage("文章点赞成功", c)
}
