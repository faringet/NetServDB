package myRedis

import (
	"github.com/gin-gonic/gin"
)

type RedisService struct {
	repo RedisRepository
}

func NewRedisService(repo RedisRepository) *RedisService {
	return &RedisService{
		repo: repo,
	}
}

func (rs *RedisService) IncrementByKey(c *gin.Context, key string, value int64) (int64, error) {
	return rs.repo.IncrBy(c, key, value)
}

func (rs *RedisService) RefreshRedis(c *gin.Context) error {
	key := "age"
	if err := rs.repo.Refresh(c, key); err != nil {
		return err
	}
	return nil
}
