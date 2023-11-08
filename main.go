package main

import (
	"NetServDB/controllers"
	"NetServDB/domain"
	"NetServDB/initializers"
	"NetServDB/logging"
	"NetServDB/middleware"
	"NetServDB/transport/http"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&domain.Users{})
	initializers.ConnectToRedis()
	initializers.SetRedisKey()

	logger := logging.GetLogger()
	logger.Info("Start app")
}

func main() {
	r := gin.Default()
	logger := logging.GetLogger()
	redisClient := initializers.RedisClient
	db := initializers.DB

	redController := http.NewRedisController(logger, redisClient)

	r.POST("/redis/incr", func(c *gin.Context) {
		redController.RedisIncr(c)
	})

	r.POST("/sign/hmacsha512", func(c *gin.Context) {
		controllers.SignHMACSHA512(c, logger)
	})

	r.POST("/postgres/users", func(c *gin.Context) {
		http.AddUser(c, logger, db)
	})

	r.DELETE("/redis/del", middleware.Authenticate(), func(c *gin.Context) {
		redController.RedisRefresh(c, logger, redisClient)
	})

	r.DELETE("/postgres/users", middleware.Authenticate(), func(c *gin.Context) {
		http.TableRefresh(c, logger, db)
	})

	r.Run()

	// чтобы можно было завершить программу из терминала по Ctrl + C когда запускаем через параметры
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала завершения
	<-signals

}
