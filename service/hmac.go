package service

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

type HMACService interface {
	SignHMACSHA512(text string, key string) (string, error)
}

type HMACServiceImpl struct {
	repo HMACService
}

func NewHMACService() HMACService {
	return &HMACServiceImpl{}
}

func (hs *HMACServiceImpl) SignHMACSHA512(text string, key string) (string, error) {
	// Создаем новый HMAC-SHA512 хэш с ключом
	h := hmac.New(sha512.New, []byte(key))

	// Записываем данные для подписи
	h.Write([]byte(text))

	// Вычисляем хэш
	signature := h.Sum(nil)

	// Преобразуем хэш в строку в шестнадцатеричном формате
	signatureHex := hex.EncodeToString(signature)

	return signatureHex, nil
}
