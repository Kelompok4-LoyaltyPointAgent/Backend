package redisclient

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/kelompok4-loyaltypointagent/backend/config"
)

func Init() *redis.Client {
	redisConfig := config.LoadRedisConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr + ":" + redisConfig.Port,
		Username: redisConfig.Username,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	pong, err := client.Ping(client.Context()).Result()
	fmt.Println(pong, err)

	return client
}
