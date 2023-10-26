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

	r.POST("/redis/incr", controllers.RedisIncr)
	r.POST("/sign/hmacsha512", controllers.SignHMACSHA512)
	r.POST("/postgres/users", controllers.AddUser)

	r.DELETE("/redis/del", middleware.Authenticate(), controllers.RedisRefresh)
	r.DELETE("/postgres/users", middleware.Authenticate(), controllers.TableRefresh)
	r.Run()

	// чтобы можно было завершить программу из терминала по Ctrl + C когда запускаем через параметры
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала завершения
	<-signals

}
