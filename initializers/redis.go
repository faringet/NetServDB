package initializers

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client

func ConnectToRedis() {

	// Получаем адрес редиса через параметры `-host` и `-port`, если их нет - тогда дефолт
	host := flag.String("host", "localhost", "Redis host")
	port := flag.String("port", "6379", "Redis port")
	flag.Parse()
	redisAddr := *host + ":" + *port

	// Новый клиент Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
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
