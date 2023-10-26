package tests

import (
	"NetServDB/controllers"
	"NetServDB/initializers"
	"NetServDB/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedisIncrJsonOk(t *testing.T) {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Users{})
	initializers.ConnectToRedis()
	initializers.SetRedisKey()

	router := gin.New()

	router.POST("/redis/incr", controllers.RedisIncr)

	requestBody := map[string]interface{}{
		"key":   "age",
		"value": 19,
	}
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/redis/incr", bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRedisIncrJsonNotOk(t *testing.T) {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Users{})
	initializers.ConnectToRedis()
	initializers.SetRedisKey()

	router := gin.New()
	router.POST("/redis/incr", controllers.RedisIncr)

	emptyReq, err := http.NewRequest(http.MethodPost, "/redis/incr", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}
	emptyRes := httptest.NewRecorder()
	router.ServeHTTP(emptyRes, emptyReq)
	assert.Equal(t, http.StatusBadRequest, emptyRes.Code)
}

func TestSignHMACSHA512JsonOk(t *testing.T) {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Users{})
	initializers.ConnectToRedis()
	initializers.SetRedisKey()

	router := gin.New()

	router.POST("/sign/hmacsha512", controllers.SignHMACSHA512)

	requestBody := map[string]interface{}{
		"text": "test",
		"key":  "test123",
	}
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/sign/hmacsha512", bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSignHMACSHA512JsonNotOk(t *testing.T) {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Users{})
	initializers.ConnectToRedis()
	initializers.SetRedisKey()

	router := gin.New()
	router.POST("/sign/hmacsha512", controllers.SignHMACSHA512)

	emptyReq, err := http.NewRequest(http.MethodPost, "/sign/hmacsha512", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}
	emptyRes := httptest.NewRecorder()
	router.ServeHTTP(emptyRes, emptyReq)
	assert.Equal(t, http.StatusBadRequest, emptyRes.Code)
}

func TestAddUserJsonOk(t *testing.T) {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Users{})
	initializers.ConnectToRedis()
	initializers.SetRedisKey()

	router := gin.New()
	router.POST("/postgres/users", controllers.AddUser)

	requestBody := map[string]interface{}{
		"name": "Alex",
		"age":  21,
	}
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/postgres/users", bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestAddUserJsonNotOk(t *testing.T) {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Users{})
	initializers.ConnectToRedis()
	initializers.SetRedisKey()

	router := gin.New()
	router.POST("/postgres/users", controllers.AddUser)

	emptyReq, err := http.NewRequest(http.MethodPost, "/postgres/users", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}
	emptyRes := httptest.NewRecorder()
	router.ServeHTTP(emptyRes, emptyReq)
	assert.Equal(t, http.StatusBadRequest, emptyRes.Code)
}
