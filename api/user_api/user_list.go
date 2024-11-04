package user_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/ctype"
	"server/models/res"
	"server/services/common"
	"server/utils/desensitiza"
	"server/utils/jwts"
)

//type UserResponse struct {
//	models.UserModel
//}

// 权限

func (UserApi) UserListView(c *gin.Context) {
	// 如何判断是管理员
	//token := c.Request.Header.Get("token")
	//if len(token) == 0 {
	//	res.FailWithMessage("未携带token", c)
	//	return
	//}
	//claims, err := jwts.ParseToken(token)
	//if err != nil {
	//	res.FailWithMessage("token错误", c)
	//	return
	//}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 这里会用到分页操作
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c) // ArgumentError 参数错误
		return
	}
	var users []models.UserModel
	list, count, _ := common.CommList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 非管理员
			user.UserName = ""
		}
		// 脱敏
		user.Tel = desensitiza.DesensitizationTel(user.Tel)
		user.Email = desensitiza.DesensitizationEmail(user.Email)
		users = append(users, user)
	}
	res.OkWithList(users, count, c)
}
