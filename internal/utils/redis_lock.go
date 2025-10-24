package utils

import (
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

var Rs *redsync.Redsync

func InitRedisLock() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	pool := redsyncredis.NewPool(client)
	Rs = redsync.New(pool)
}
