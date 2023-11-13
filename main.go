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
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

	// создаем канал ошибок errChain
	errChain := make(chan error, 1)

	/*
		Запускаем горутину, которая содержит код для запуска роутера
		Если происходит ошибка при запуске, она отправляется в errChain
	*/
	go func() {
		err = router.Start()
		if err != nil {
			fmt.Print("exit router start with error:", err)
		}

		errChain <- err
	}()

	/*
		Еще одна асинхронная горутина, которая слушает сигналы прерывания (Ctrl+C) или завершения программы (SIGTERM)
		При получении сигнала она отправляет ошибку в errChain
	*/
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		fmt.Print("\t<-signals")
		//Ожидаем сигнала завершения
		s := <-signals

		errChain <- errors.New("get os signal" + s.String())
	}()

	// инфу из канала errChain сохраняем в errRun
	errRun := <-errChain
	logger.Error(errRun)

	//Закрываем коннекты
	err = postgCleanup()
	logger.Error("postgres cleanup: ", err)
	err = redisCleanup()
	logger.Error("redis cleanup: ", err)
}
