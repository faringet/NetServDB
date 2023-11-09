package redis

import (
	"context"
	"flag"
	"github.com/go-redis/redis/v8"
	"log"
)

type Config struct {
	//TODO: реализовать потом
}

func NewRedis(cfg Config) (*redis.Client, error) {
	// Получаем адрес редиса через параметры `-host` и `-port`, если их нет - тогда дефолт
	host := flag.String("host", "localhost", "Redis host")
	port := flag.String("port", "6379", "Redis port")
	flag.Parse()
	redisAddr := *host + ":" + *port

	// Новый клиент Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	// Проверить подключение к Redis
	ctx := context.Background()
	result, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	if result == "PONG" {
		log.Print("Redis connection successful")
	}

	return client, nil
}
