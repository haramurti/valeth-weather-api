package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client          //remote tv universal
var Ctx = context.Background() // giving context

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: " ",
		DB:       0,
	})
}
