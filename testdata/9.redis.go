package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

var rdb *redis.Client

func init() {
	ctx := context.Background()
	// 连接redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 地址
		Password: "",               // no password set
		DB:       0,                // use default DB
		PoolSize: 100,              // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond) // 500*time.Millisecond 超时的时间
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Error(err)
		return
	}

}
func main() {
	ctx := context.Background()
	err := rdb.Set(ctx, "xxx1", "value1", 10*time.Second).Err()
	fmt.Println(err)
	cmd := rdb.Keys(ctx, "*")
	keys, err := cmd.Result()
	fmt.Println(keys, err)
}
