package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/jwts"
	"server/utils/pwd"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入原来的密码"` // 旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`       // 新密码
}

// UserUpdatePasswordView 修改登录人的id
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 参数绑定
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	// 这里可能查不到,按照道理UserID是JWT里面的，
	// 但是JWT如果返回回来了，你把这个用户删掉了，那这个UserID就是被删掉了，那就不能继续往下走
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	// 判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("密码错误", c)
		return
	}
	// 密码一致
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
	return
}
