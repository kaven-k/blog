package tag_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// AdvertisementRequest 设置了binding:"required"验证规则，表示该字段在接收到的请求数据中是必需的，
// 如果缺失则会触发验证错误。并且还定义了一个自定义的错误提示消息msg:"请输入标题"，
// 当验证不通过时（即该字段缺失时），可以返回这个消息给客户端以明确告知错误原因
type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

// TagCreateView 添加标签
// @Tags 标签管理
// @Summary 创建标签
// @Description 创建标签
// @Param data body AdvertisementRequest   true  "表示多个参数"
// @Router /api/ad [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagCreateView(c *gin.Context) {
	// 参数校验
	var cr TagRequest
	// 通过ShouldBindJSON方法尝试将客户端传来的 JSON 数据绑定到AdvertisementRequest结构体上
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复的判断（如果广告列表里面已经有了）
	// 有两种办法：
	// 1、在数据库中给标题（或者其他）设置一个唯一索引（使用唯一索引会报错，如果使用就需要去判断这个错误是什么）
	// 2、每次添加之前去查一下数据库
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil { // 如果err为空就代表查到了
		res.FailWithMessage("该标签已存在", c)
		return
	}
	// 添加广告
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err) // 错误就对外提出
		res.FailWithMessage("添加标签失败", c)
		return
	}
	res.OkWithMessage("添加标签成功", c)
}
