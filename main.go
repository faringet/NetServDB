package main

import (
	"NetServDB/controllers"
	"NetServDB/domain"
	"NetServDB/initializers"
	"NetServDB/initializers/redis"
	"NetServDB/logging"
	"NetServDB/middleware"
	"NetServDB/service"
	"NetServDB/storage/dbredis"
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

	logger := logging.GetLogger()
	logger.Info("Start app")
}

func main() {
	r := gin.Default()
	logger := logging.GetLogger()

	rCfg := redis.Config{}

	redisClient, err := redis.NewRedis(rCfg)
	if err != nil {
		panic("can't panic")
	}

	db := initializers.DB

	redisRepo := dbredis.NewRedisRepositoryImpl(redisClient)
	cacheWorker := service.NewCacheWorker(redisRepo)

	redController := http.NewRedisController(logger, cacheWorker)
	userController := http.NewUserController(logger, db)

	//TODO: ошибка в пути
	r.POST("/myRedis/incr", func(c *gin.Context) {
		redController.RedisIncr(c)
	})

	r.POST("/sign/hmacsha512", func(c *gin.Context) {
		controllers.SignHMACSHA512(c, logger)
	})

	r.POST("/postgres/users", func(c *gin.Context) {
		userController.AddUser(c)
	})

	r.DELETE("/myRedis/del", middleware.Authenticate(), func(c *gin.Context) {
		redController.RedisRefresh(c)
	})

	r.DELETE("/postgres/users", middleware.Authenticate(), func(c *gin.Context) {
		userController.TableRefresh(c)
	})

	r.Run()

	// чтобы можно было завершить программу из терминала по Ctrl + C когда запускаем через параметры
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала завершения
	<-signals

}
