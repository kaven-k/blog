package user_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
	"server/services"
	"server/utils/jwts"
)

// UserLogoutView 用户注销接口 这个接口必须要登录之后才能使用
func (UserApi) UserLogoutView(c *gin.Context) {
	//ctx := context.Background()
	// 如果登录的话，就能拿到这两个信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 同时还要拿一个信息 claims.ExpiresAt 过期时间
	//fmt.Println(claims.ExpiresAt)

	token := c.Request.Header.Get("token")
	err := services.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		// 出错了
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c) // 返回给前端的信息
		return
	}
	res.OkWithMessage("注销成功", c)

}
