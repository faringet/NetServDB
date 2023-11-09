package myRedis

import (
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

	ctx := c.Request.Context()

	updatedValue, err := r.redisClient.IncrBy(ctx, key, value).Result()
	if err != nil {
		return 0, err
	}

	return updatedValue, nil
}

func (r *RedisRepositoryImpl) Refresh(c *gin.Context, key string) error {

	ctx := c.Request.Context()

	err := r.redisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
