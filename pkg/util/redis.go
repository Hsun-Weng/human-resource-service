package util

import (
	"context"
	"fmt"
	"github.com/Hsun-Weng/human-resource-service/internal/config"
	"github.com/redis/go-redis/v9"
	"log"
)

func NewRedisClient() *redis.Client {
	redisHost := config.GetEnv(config.RedisHost, "localhost")
	redisPort := config.GetEnv(config.RedisPort, "6379")
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		DB:   0,
	})

	// Perform basic diagnostic to check if the connection is working
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}

	return client
}
