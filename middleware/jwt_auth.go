package middleware

import (
	"github.com/gin-gonic/gin"
	"server/models/ctype"
	"server/models/res"
	"server/services/redis_service"
	"server/utils/jwts"
)

// JwtAuth 判断是否登录，登录之后才可以调用某些接口
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			res.FailWithMessage("未携带token", c)
			c.Abort() // 拦截
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
		}
		// 判断是否在redis中
		if redis_service.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}
}

// JwtAdmin 管理员才能使用的中间件
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			res.FailWithMessage("未携带token", c)
			c.Abort() // 拦截
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		if redis_service.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		// 写之前需要判断是不是超级管理员
		if claims.Role != int(ctype.PermissionAdmin) {
			if err != nil {
				res.FailWithMessage("权限错误", c)
				c.Abort()
				return
			}
		}
		c.Set("claims", claims)
	}
}
