package jwts

import (
	"github.com/golang-jwt/jwt/v5"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	//Username string `json:"username"`  // 用户名
	NickName string `json:"nick_name"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint   `json:"user_id"`   // 用户id
}

// CustomClaims 必须要有的
type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

//var MySecret = []byte(global.Config.Jwt.Secret)
