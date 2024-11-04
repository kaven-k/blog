package user_service

import (
	"server/services/redis_service"
	"server/utils/jwts"
	"time"
)

// Logout 过期token
func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	//ctx := context.Background()
	// 需要计算距离现在的过期时间
	// redis设置的时间是 time.Duration() 距离现在多少多少秒
	exp := claims.ExpiresAt   // ExpiresAt 这个时间是截止时间
	now := time.Now()         // 当前时间
	diff := exp.Time.Sub(now) // 这个时间是到现在的截止时间
	//fmt.Println(diff)

	//err := global.Redis.Set(ctx, fmt.Sprintf("logout_%s", token), "", diff).Err()
	return redis_service.Logout(token, diff)
}
