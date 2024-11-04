package settings_api

import (
	"github.com/gin-gonic/gin"
	"server/config"
	"server/core"
	"server/global"
	"server/models/res"
)

// SettingsInfoUpdateView 修改某一项的配置信息
// 使用这个方法得到的是接口数量的减少，失去的是接口没法统一
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info

	case "email":
		var email config.Email
		err = c.ShouldBindJSON(&email)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = email
	case "qq":
		var qq config.QQ
		err = c.ShouldBindJSON(&qq)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = qq
	case "qiniu":
		var qiniu config.QiNiu
		err = c.ShouldBindJSON(&qiniu)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = qiniu
	case "jwt":
		var jwt config.Jwt
		err = c.ShouldBindJSON(&jwt)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = jwt
	default:
		res.FailWithMessage("没有对应的配置信息！", c)
		return
	}
	err = core.SetYaml()
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	res.OkWith(c)
}
