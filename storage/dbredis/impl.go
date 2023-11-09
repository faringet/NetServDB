package dbredis

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type RedisRepositoryImpl struct {
	redisClient *redis.Client
}

func NewRedisRepositoryImpl(redisClient *redis.Client) *RedisRepositoryImpl {
	return &RedisRepositoryImpl{
		redisClient: redisClient,
	}
}

func (r *RedisRepositoryImpl) IncrBy(c *gin.Context, key string, value int64) (int64, error) {
	updatedValue, err := r.redisClient.IncrBy(c, key, value).Result()
	if err != nil {
		return 0, err
	}

	return updatedValue, nil
}

func (r *RedisRepositoryImpl) Delete(c *gin.Context, key string) error {

	ctx := c.Request.Context()

	err := r.redisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepositoryImpl) SetKeyDefault(key string, defaultValue int64) {
	ctx := context.Background()

	// Проверить существование ключа "age"
	exists, err := r.redisClient.Exists(ctx, key).Result()
	if err != nil {
		fmt.Println("!!!!")
		return
	}

	if exists == 0 {
		// Ключ "age" не существует, устанавливаем его значение в 1
		err := r.redisClient.Set(ctx, key, defaultValue, 0).Err()
		if err != nil {
			fmt.Println("??????")
			return
		}
	}

}
