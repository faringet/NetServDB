package service

import (
	http2 "NetServDB/transport/http/model"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

type HMACService interface {
	SignHMACSHA512(c *gin.Context, request *http2.Ihmacsha512Request) (string, error)
}

type HMACServiceImpl struct{}

func NewHMACService() HMACService {
	return &HMACServiceImpl{}
}

func (hs *HMACServiceImpl) SignHMACSHA512(c *gin.Context, request *http2.Ihmacsha512Request) (string, error) {
	// Создаем новый HMAC-SHA512 хэш с ключом
	h := hmac.New(sha512.New, []byte(request.Key))

	// Записываем данные для подписи
	h.Write([]byte(request.Text))

	// Вычисляем хэш
	signature := h.Sum(nil)

	// Преобразуем хэш в строку в шестнадцатеричном формате
	signatureHex := hex.EncodeToString(signature)

	return signatureHex, nil
}
