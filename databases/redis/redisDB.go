package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb *redis.Client

func Connect() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})
}

func Set(k string, v string) error {
	return rdb.Set(ctx, k, v, 0).Err()
}

func Get(k string) (string, error) {
	return rdb.Get(ctx, k).Result()
}
