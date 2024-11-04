package redis_service

import (
	"context"
	"server/global"
	"server/utils"
	"time"
)

const prefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	ctx := context.Background()
	err := global.Redis.Set(ctx, prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	ctx := context.Background()
	keys := global.Redis.Keys(ctx, prefix+"*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
