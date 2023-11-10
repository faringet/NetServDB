package main

import (
	"NetServDB/config"
	"NetServDB/controllers"
	"NetServDB/initializers"
	"NetServDB/initializers/postgre"
	"NetServDB/initializers/redis"
	"NetServDB/logging"
	"NetServDB/middleware"
	"NetServDB/service"
	"NetServDB/storage/dbpostgre"
	"NetServDB/storage/dbredis"
	"NetServDB/transport/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	initializers.LoadEnvVariables()

	logger := logging.GetLogger()
	logger.Info("Start app")
}

const configPath = "config/conf.yaml"

func main() {
	r := gin.Default()
	logger := logging.GetLogger()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("can't panic: %v", err))
	}

	redisClient, err := redis.NewRedis(cfg)
	if err != nil {
		panic("can't panic")
	}

	db, err := postgre.NewDB(cfg)
	if err != nil {
		panic("can't panic")
	}

	redisRepo := dbredis.NewRedisRepositoryImpl(redisClient)
	cacheWorker := service.NewCacheWorker(redisRepo)

	dataBaseRepo := dbpostgre.NewDataBaseRepositoryImpl(db)
	dataBaseWorker := service.NewDataBaseWorker(dataBaseRepo)

	redController := http.NewRedisController(logger, cacheWorker)
	userController := http.NewUserController(logger, dataBaseWorker)

	r.POST("/redis/incr", func(c *gin.Context) {
		redController.RedisIncr(c)
	})

	r.POST("/sign/hmacsha512", func(c *gin.Context) {
		controllers.SignHMACSHA512(c, logger)
	})

	r.POST("/postgres/users", func(c *gin.Context) {
		userController.AddUser(c)
	})

	r.DELETE("/redis/del", middleware.Authenticate(), func(c *gin.Context) {
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
