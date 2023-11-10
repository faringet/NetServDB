package http

import (
	"NetServDB/controllers"
	"NetServDB/middleware"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string
	//	...
}

type Router struct {
	r *gin.Engine
}

/*
 можно сделать через interface -> rc *RedisController
*/

// NewRouter dsa
func NewRouter(cfg Config, rc *RedisController, uc *UserController) (router *Router, cleanup func() error) {

	r := gin.Default()

	r.POST("/redis/incr", func(c *gin.Context) {
		rc.RedisIncr(c)
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

	return &Router{r: r}
}

func (r Router) Run() {
	r.Run()
}
