package http

import (
	"NetServDB/initializers"
	"NetServDB/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type RedisController struct {
	redisClient *redis.Client
	logger      *logging.Logger
}

func NewRedisController(logger *logging.Logger, redisClient *redis.Client) *RedisController {
	return &RedisController{
		redisClient: redisClient,
		logger:      logger,
	}
}

func (rc *RedisController) RedisIncr(c *gin.Context) {
	var request IncrRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		rc.logger.Error("invalid input format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Логгируем принятые значения
	rc.logger.Info(fmt.Sprintf("Received request - Key:%s Value:%d", request.Key, request.Value))

	// Инкрементируем значение в Redis
	updatedValue, err := rc.redisClient.IncrBy(c, request.Key, int64(request.Value)).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем и логгируем обновленное значение в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"value": updatedValue})
	rc.logger.Info(fmt.Sprintf("UpdatedValue:%d", updatedValue))
}

func (rc *RedisController) RedisRefresh(c *gin.Context) {

	rc.redisClient.Del(c, "age")
	initializers.SetRedisKey() //todo придумать как уменьшить связанность

	// Отвечаем и логгируем что редис обновлен
	c.JSON(http.StatusOK, gin.H{"REDIS successfully refreshed": "OK"})
	rc.logger.Info("REDIS successfully updated")

}
