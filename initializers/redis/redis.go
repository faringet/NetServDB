package redis

import (
	"NetServDB/config"
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func NewRedis(cfg *config.Config) (client *redis.Client, cleanup func() error, err error) {
	// Получаем адрес редиса через параметры `-host` и `-port`, если их нет - тогда дефолт
	host := flag.String("host", cfg.Redis.Host, "Redis host")
	port := flag.String("port", cfg.Redis.Port, "Redis port")
	flag.Parse()
	redisAddr := *host + ":" + *port

	// Новый клиент Redis
	client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	// Проверить подключение к Redis
	ctx := context.Background()
	result, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to connect to Redis: %v", err)
	}
	if result == "PONG" {
		log.Print("Redis connection successful")
	}

	cleanup = func() error {
		fmt.Println("cleanup from redis")
		return client.Close()
	}

	return client, cleanup, nil
}
