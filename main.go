package main

import (
	"NetServDB/config"
	"NetServDB/initializers/postgre"
	"NetServDB/initializers/redis"
	"NetServDB/logging"
	"NetServDB/service"
	"NetServDB/storage/dbpostgre"
	"NetServDB/storage/dbredis"
	"NetServDB/transport/http"
	"fmt"
)

const configPath = "config/conf.yaml"

func main() {
	logger := logging.GetLogger()
	logger.Info("Start app")

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("can't panic: %v", err))
	}

	redisClient, redisCleanup, err := redis.NewRedis(cfg)
	if err != nil {
		panic("can't panic")
	}

	db, postgCleanup, err := postgre.NewDB(cfg)
	if err != nil {
		panic("can't panic")
	}

	redisRepo := dbredis.NewRedisRepositoryImpl(redisClient)
	cacheWorker := service.NewCacheWorker(redisRepo)

	dataBaseRepo := dbpostgre.NewDataBaseRepositoryImpl(db)
	dataBaseWorker := service.NewDataBaseWorker(dataBaseRepo)

	redController := http.NewRedisController(logger, cacheWorker)
	userController := http.NewUserController(logger, dataBaseWorker)

	hmacService := service.NewHMACService()
	hmacController := http.NewHMACController(logger, hmacService)

	router := http.NewRouter(redController, userController, hmacController, logger, cfg)
	router.RegisterRoutes()

	//TODO: при создании gin использовать cleanup()

	// TODO:  сделать grasfullshutdown
	// чтобы можно было завершить программу из терминала по Ctrl + C когда запускаем через параметры
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала завершения
	//<-signals

	// Закрываем коннекты
	//err = postgCleanup()
	//err = redisCleanup()
	//logger.Error(err)

	defer func() {
		postgErr := postgCleanup()
		fmt.Print("defer from main")
		redisErr := redisCleanup()

		if postgErr != nil {
			logger.Error("Error during PostgreSQL cleanup:", postgErr)
		}

		if redisErr != nil {
			logger.Error("Error during Redis cleanup:", redisErr)
		}
	}()

}
