package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/jwts"
	"server/utils/pwd"
)

type EmailLoginRequest struct {
	Username string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	// 参数绑定
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 验证用户是否存在
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or user_name = ?", cr.Username, cr.Username).Error
	if err != nil {
		// 没找到
		global.Log.Warn("用户名不存在")           //打印在控制台
		res.FailWithMessage("用户名或者密码错误", c) // 返回给前端的
		return
	}
	// 校验密码
	idCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !idCheck {
		// 没找到
		global.Log.Warn("用户名密码错误")        //打印在控制台
		res.FailWithMessage("用户名密码错误", c) // 返回给前端的
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)               //将具体的错误打印在控制台
		res.FailWithMessage("token生成失败", c) // 返回给前端的
		return
	}
	res.OkWithData(token, c)
}
