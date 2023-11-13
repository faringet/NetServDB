package middleware

import (
	"NetServDB/config"
	"NetServDB/logging"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(cfg *config.Config) gin.HandlerFunc {
	logger := logging.GetLogger()
	expectedUsername := cfg.Auth.Username
	expectedPassword := cfg.Auth.Password

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Проверяем заголовок Authorization на пустоту
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			logger.Error("empty authHeader - authorization denied")
			c.Abort()
			return
		}

		// Проверяем что начинается с "Basic "
		if !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			logger.Error("wrong authHeader - authorization denied")
			c.Abort()
			return
		}

		// Извлекаем данные
		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")

		// Раскодируем все
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			logger.Error("decode error - authorization denied")
			c.Abort()
			return
		}

		// В строку
		credentials := string(decodedCredentials)

		// Юзер пароль
		parts := strings.Split(credentials, ":")
		if len(parts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			logger.Error("decode error - authorization denied")
			c.Abort()
			return
		}

		// Сравниваем креденшалз
		username := parts[0]
		password := parts[1]
		if username != expectedUsername || password != expectedPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			logger.Error("wrong password or username - authorization denied")
			c.Abort()
			return
		}

		// Если аутентификация успешна, продолжайте выполнение запроса
		logger.Info("successful authorization !")
		c.Next()
	}
}
