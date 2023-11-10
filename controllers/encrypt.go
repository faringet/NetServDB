package controllers

import (
	"NetServDB/logging"
	http2 "NetServDB/transport/http"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignHMACSHA512(c *gin.Context, logger *logging.Logger) {
	var request http2.Ihmacsha512Request

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("invalid input format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: Это выделить на слой сервиса
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
