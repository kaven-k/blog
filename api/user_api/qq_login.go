package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/models/res"
	"server/plugins/qq"
	"server/utils/jwts"
	"server/utils/pwd"
	"server/utils/random"
)

func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	fmt.Println(code)
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	// 打印是为了测试的时候可以看到数据
	//fmt.Println(qqInfo)
	//res.OkWithData(qqInfo, c)
	openID := qqInfo.OpenID
	//根据openID判断用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "token = ?", openID).Error
	if err != nil {
		// 不存在，就注册
		hashPwd := pwd.HashPwd(random.RandString(16))
		//res.FailWithMessage("用户不存在，请注册", c)
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,  // 直接用openID作为用户名，这样就可以直接使用qq登录，绑定邮箱之后可以使用邮箱+密码登录
			Password:   hashPwd, // 随机生成一个16位的密码
			Avatar:     qqInfo.Avatar,
			Addr:       "内网", // 根据ip算地址
			Token:      openID,
			IP:         c.ClientIP(),
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("注册失败", c)
			return
		}

	}
	// 登录操作
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)               //将具体的错误打印在控制台
		res.FailWithMessage("token生成失败", c) // 返回给前端的
		return
	}
	res.OkWithData(token, c)
}
