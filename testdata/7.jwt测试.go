package main

import (
	"fmt"
	"server/core"
	"server/global"
	"server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID: 1,
		Role:   1,
		//Username: "yky",
		NickName: "xxx",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InlreSIsIm5pY2tfbmFtZSI6Inh4eCIsInJvbGUiOjEsInVzZXJfaWQiOjEsImlzcyI6Inh4IiwiZXhwIjoxNzMwMTIxNjIxfQ.j2jpXvT4Lffxd4YuHLWGB_pD051_g-qHZJKfZc0LL3E")
	fmt.Println(claims, err)
}
