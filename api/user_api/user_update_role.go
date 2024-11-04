package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/models/res"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
	UserID   uint       `json:"user_id" binding:"required" msg:"用户ID错误"`
	//UserEmail string     `json:"user_email"`
	//UserTel   string     `json:"user_tel"`
}

// UserUpdateRoleView 用户权限变更
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	// 1.参数绑定
	var cr UserRole
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 2.查找用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		// 没找都用户
		res.FailWithMessage("用户ID错误,用户不存在", c)
		return
	}
	// 3.查到了，进行修改
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
		//"user_email": cr.UserEmail,
		//"user_tel":   cr.UserTel,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("修改权限失败", c)
		return
	}
	res.OkWithMessage("修改权限成功", c)
}
