package redis_service

import (
	"context"
	"server/global"
	"strconv"
)

const diggPrefix = "digg_api"

// Digg 点赞某一篇文章
func Digg(id string) error {
	ctx := context.Background()
	num, _ := global.Redis.HGet(ctx, diggPrefix, id).Int()
	num++
	err := global.Redis.HSet(ctx, diggPrefix, id, num).Err()
	return err
}

// 如果想知道某人在何时点的赞

// GetDigg 获取某一篇文章下的点赞数
func GetDigg(id string) int {
	ctx := context.Background()
	num, _ := global.Redis.HGet(ctx, diggPrefix, id).Int()
	return num
}

// GetDiggInfo 取出点赞数据 // 每隔一段时间同步点赞数据到es
func GetDiggInfo() map[string]int {
	ctx := context.Background()
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(ctx, diggPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val) // 转成int类型
		DiggInfo[id] = num
	}
	return DiggInfo
}

// DiggClear 文章同步完毕之后需要清除掉缓存里面的内容
func DiggClear() {
	ctx := context.Background()
	global.Redis.Del(ctx, diggPrefix) // 清空的方法，直接将索引删掉
}
