package http

import (
	"NetServDB/domain"
	"NetServDB/initializers"
	"NetServDB/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
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

func (rc *RedisController) RedisRefresh(c *gin.Context, logger *logging.Logger, redisClient *redis.Client) {

	redisClient.Del(c, "age")
	initializers.SetRedisKey() //todo придумать как уменьшить связанность

	// Отвечаем и логгируем что редис обновлен
	c.JSON(http.StatusOK, gin.H{"REDIS successfully refreshed": "OK"})
	logger.Info("REDIS successfully updated")

}

func AddUser(c *gin.Context, logger *logging.Logger, db *gorm.DB) {
	var request UserRequestAdd

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("invalid input format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.mapToDomain()

	// Пишем юзера в БД
	user := domain.Users{
		Name: request.Name,
		Age:  request.Age,
	}

	// Логгируем юзера
	logger.Info(fmt.Sprintf("new user - Name:%s Age:%d", request.Name, request.Age))

	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Возвращаем и логгируем id нового юзера в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"id": user.ID})
	logger.Info(fmt.Sprintf("user's id:%d", user.ID))
}

func TableRefresh(c *gin.Context, logger *logging.Logger, db *gorm.DB) {

	// Так как GORM не умеет дропать таблицы придется выполнить SQL-запрос руками
	if err := db.Exec("DROP TABLE IF EXISTS users;").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table 'users' refreshed successfully"})
	db.AutoMigrate(&domain.Users{})
	logger.Info("Table 'users' refreshed successfully")
}
