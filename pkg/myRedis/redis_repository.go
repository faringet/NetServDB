package myRedis

import (
	"github.com/gin-gonic/gin"
)

type RedisRepository interface {
	IncrBy(c *gin.Context, key string, value int64) (int64, error)
	Refresh(c *gin.Context, key string) error
}
