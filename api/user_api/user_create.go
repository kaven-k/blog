package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/ctype"
	"server/models/res"
	"server/services/user_service"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   // 密码,密码不展示的话，就将json里面的password改为-
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
	Email    string     `json:"email" msg:"请输入邮箱"`
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err := user_service.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, cr.Email, c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	global.Log.Infof("用户%s创建成功！", cr.UserName)                  //打印在控制台
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功！", cr.UserName), c) // 返回给前端
	return
}
