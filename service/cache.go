package service

import (
	"github.com/gin-gonic/gin"
)

const (
	defaultAge = 1
	keyAge     = "age"
)

type RedisRepository interface {
	IncrBy(c *gin.Context, key string, value int64) (int64, error)
	Delete(c *gin.Context, key string) error
	SetKeyDefault(key string, defaultValue int64)
}

type CacheWorker struct {
	repo RedisRepository
}

func NewCacheWorker(repo RedisRepository) *CacheWorker {
	return &CacheWorker{
		repo: repo,
	}
}

func (rs *CacheWorker) IncrementByKey(c *gin.Context, key string, value int64) (int64, error) {
	rs.repo.SetKeyDefault(key, defaultAge)

	return rs.repo.IncrBy(c, key, value)
}

func (rs *CacheWorker) Refresh(c *gin.Context) error {
	if err := rs.repo.Delete(c, keyAge); err != nil {
		return err
	}

	return nil
}
