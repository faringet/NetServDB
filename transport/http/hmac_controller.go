package http

import (
	"NetServDB/logging"
	"NetServDB/transport/http/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HMAC interface {
	SignHMACSHA512(text string, key string) (string, error)
}

type HMACController struct {
	HMAC   HMAC
	logger *logging.Logger
}

func NewHMACController(logger *logging.Logger, hmacService HMAC) *HMACController {
	return &HMACController{
		HMAC:   hmacService,
		logger: logger,
	}
}

func (hc *HMACController) SignHMACSHA512(c *gin.Context) {
	var request model.Ihmacsha512Request

	if err := c.ShouldBindJSON(&request); err != nil {
		hc.logger.Error("invalid input format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := request.ValidationHmac()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Вызываем метод сервиса для создания подписи HMAC-SHA512
	signature, err := hc.HMAC.SignHMACSHA512(request.Text, request.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем и логгируем HMAC-SHA512 подпись в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"signature": signature})
	hc.logger.Info("HMAC-SHA512 signature generated")
}
