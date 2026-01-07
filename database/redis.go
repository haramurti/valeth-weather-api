package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: " ",
		DB:       0,
	})
}
