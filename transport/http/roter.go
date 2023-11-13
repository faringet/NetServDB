package http

import (
	"NetServDB/config"
	"NetServDB/logging"
	"NetServDB/middleware"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Start()
	RegisterRoutes()
}

type RouterImpl struct {
	redisController *RedisController
	userController  *UserController
	hmacController  *HMACController
	logger          *logging.Logger
	config          *config.Config
	server          *gin.Engine
}

func NewRouter(redisController *RedisController, userController *UserController, hmacController *HMACController, logger *logging.Logger, cfg *config.Config) *RouterImpl {
	return &RouterImpl{
		redisController: redisController,
		userController:  userController,
		hmacController:  hmacController,
		logger:          logger,
		config:          cfg,
	}
}

func (r *RouterImpl) RegisterRoutes() {
	r.logger.Info("Registering routes")

	router := gin.Default()

	router.POST("/redis/incr", func(c *gin.Context) {
		r.redisController.RedisIncr(c)
	})

	router.POST("/sign/hmacsha512", func(c *gin.Context) {
		r.hmacController.SignHMACSHA512(c)
	})

	router.POST("/postgres/users", func(c *gin.Context) {
		r.userController.AddUser(c)
	})

	router.DELETE("/redis/del", middleware.Authenticate(r.config), func(c *gin.Context) {
		r.redisController.RedisRefresh(c)
	})

	router.DELETE("/postgres/users", middleware.Authenticate(r.config), func(c *gin.Context) {
		r.userController.TableRefresh(c)
	})

	router.Run(r.config.Port)

}

func (r *RouterImpl) Start() error {
	port := r.config.Port
	return r.server.Run(port)
}
