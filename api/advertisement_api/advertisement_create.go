package advertisement_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// AdvertisementRequest 设置了binding:"required"验证规则，表示该字段在接收到的请求数据中是必需的，
// 如果缺失则会触发验证错误。并且还定义了一个自定义的错误提示消息msg:"请输入标题"，
// 当验证不通过时（即该字段缺失时），可以返回这个消息给客户端以明确告知错误原因
type AdvertisementRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`     // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` // 图片
	IsShow bool   `json:"is_show" msg:"请选择是否展示" structs:"is_show"`                     // 是否展示
}

// AdvertisementCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertisementRequest   true  "表示多个参数"
// @Router /api/ad [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertisementApi) AdvertisementCreateView(c *gin.Context) {
	// 参数校验
	var cr AdvertisementRequest
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
	var ad models.AdvertModel
	err = global.DB.Take(&ad, "title = ?", cr.Title).Error
	if err == nil { // 如果err为空就代表查到了
		res.FailWithMessage("该广告已存在", c)
		return
	}
	// 添加广告
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err) // 错误就对外提出
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)
}
