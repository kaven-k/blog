package message_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/services/common"
)

func (MessageApi) MessageListAdminView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.CommList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)
}
