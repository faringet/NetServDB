package main

import (
	"NetServDB/controllers"
	"NetServDB/initializers"
	"NetServDB/logging"
	"NetServDB/middleware"
	"NetServDB/models"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Users{})
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

	r.POST("/redis/incr", func(c *gin.Context) {
		controllers.RedisIncr(c, logger, redisClient)
	})

	r.POST("/sign/hmacsha512", func(c *gin.Context) {
		controllers.SignHMACSHA512(c, logger)
	})

	r.POST("/postgres/users", func(c *gin.Context) {
		controllers.AddUser(c, logger, db)
	})

	r.DELETE("/redis/del", middleware.Authenticate(), func(c *gin.Context) {
		controllers.RedisRefresh(c, logger, redisClient)
	})

	r.DELETE("/postgres/users", middleware.Authenticate(), func(c *gin.Context) {
		controllers.TableRefresh(c, logger, db)
	})

	r.Run()

	// чтобы можно было завершить программу из терминала по Ctrl + C когда запускаем через параметры
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала завершения
	<-signals

}
