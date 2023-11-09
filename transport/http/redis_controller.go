package http

import (
	"NetServDB/initializers"
	"NetServDB/logging"
	"NetServDB/pkg/myRedis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RedisController struct {
	redisService myRedis.RedisService
	logger       *logging.Logger
}

func NewRedisController(logger *logging.Logger, redisService myRedis.RedisService) *RedisController {
	return &RedisController{
		redisService: redisService,
		logger:       logger,
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
	updatedValue, err := rc.redisService.IncrementByKey(c, request.Key, int64(request.Value))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем и логгируем обновленное значение в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"value": updatedValue})
	rc.logger.Info(fmt.Sprintf("UpdatedValue:%d", updatedValue))
}

func (rc *RedisController) RedisRefresh(c *gin.Context) {

	err := rc.redisService.RefreshRedis(c)
	initializers.SetRedisKey() //todo придумать как уменьшить связанность
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отвечаем и логгируем что редис обновлен
	c.JSON(http.StatusOK, gin.H{"REDIS successfully refreshed": "OK"})
	rc.logger.Info("REDIS successfully updated")

}
