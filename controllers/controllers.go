package controllers

import (
	"NetServDB/initializers"
	"NetServDB/logging"
	"NetServDB/models"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedisIncr(c *gin.Context) {
	logger := logging.GetLogger()
	var request models.IncrRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("invalid input format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Логгируем принятые значения
	logger.Info(fmt.Sprintf("Received request - Key:%s Value:%d", request.Key, request.Value))

	// Инкрементируем значение в Redis
	ctx := context.Background()
	updatedValue, err := initializers.RedisClient.IncrBy(ctx, request.Key, int64(request.Value)).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем и логгируем обновленное значение в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"value": updatedValue})
	logger.Info(fmt.Sprintf("UpdatedValue:%d", updatedValue))
}

func RedisRefresh(c *gin.Context) {
	logger := logging.GetLogger()

	ctx := context.Background()
	initializers.RedisClient.Del(ctx, "age")
	initializers.SetRedisKey()

	// Отвечаем и логгируем что редис обновлен
	c.JSON(http.StatusOK, gin.H{"REDIS successfully refreshed": "OK"})
	logger.Info("REDIS successfully updated")

}

func SignHMACSHA512(c *gin.Context) {
	logger := logging.GetLogger()
	var request models.Ihmacsha512Request

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("invalid input format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создаем новый HMAC-SHA512 хэш с ключом
	h := hmac.New(sha512.New, []byte(request.Key))

	// Записываем данные для подписи
	h.Write([]byte(request.Text))

	// Вычисляем хэш
	signature := h.Sum(nil)

	// Преобразуем хэш в строку в шестнадцатеричном формате
	signatureHex := hex.EncodeToString(signature)

	// Возвращаем и логгируем HMAC-SHA512 подпись в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"signature": signatureHex})
	logger.Info(fmt.Sprintf("signature:%s", signatureHex))
}

func AddUser(c *gin.Context) {
	logger := logging.GetLogger()
	var request models.Users

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("invalid input format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Пишем юзера в БД
	user := models.Users{
		Name: request.Name,
		Age:  request.Age,
	}

	// Логгируем юзера
	logger.Info(fmt.Sprintf("new user - Name:%s Age:%d", request.Name, request.Age))

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Возвращаем и логгируем id нового юзера в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"id": user.ID})
	logger.Info(fmt.Sprintf("user's id:%d", user.ID))
}

func TableRefresh(c *gin.Context) {
	logger := logging.GetLogger()
	db := initializers.DB

	// Так как GORM не умеет дропать таблицы придется выполнить SQL-запрос руками
	if err := db.Exec("DROP TABLE IF EXISTS users;").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table 'users' refreshed successfully"})
	initializers.DB.AutoMigrate(&models.Users{})
	logger.Info("Table 'users' refreshed successfully")
}
