package settings_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
)

type SettingUri struct {
	Name string `uri:"name"`
}

//var SettingsMap map[string]*SettingUri

// SettingsInfoView 该函数属于一个名为 SettingsApi 的类型，
// 接受一个 *gin.Context 类型的参数 c ，表示当前的 HTTP 请求上下文。
// 显示某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息！", c)
	}
}
