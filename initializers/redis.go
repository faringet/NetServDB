package initializers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client

func ConnectToRedis() {

	// Новый клиент Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Проверить подключение к Redis
	ctx := context.Background()
	result, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	if result == "PONG" {
		log.Print("Redis connection successful")
	}

}

func SetRedisKey() {
	ctx := context.Background()

	// Проверить существование ключа "age"
	exists, err := RedisClient.Exists(ctx, "age").Result()
	if err != nil {
		fmt.Println("!!!!")
		return
	}

	if exists == 0 {
		// Ключ "age" не существует, устанавливаем его значение в 1
		err := RedisClient.Set(ctx, "age", 1, 0).Err()
		if err != nil {
			fmt.Println("??????")
			return
		}
	}

}
