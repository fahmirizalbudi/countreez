package configs

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
)

func GetRedisConnection() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
	})

	 _, err := Redis.Ping(Ctx).Result()
    if err != nil {
        panic(err)
    }

	fmt.Println("Redis connection established successfully.")

}