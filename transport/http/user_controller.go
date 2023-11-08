package http

import (
	"NetServDB/domain"
	"NetServDB/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	dbClient *gorm.DB
	logger   *logging.Logger
}

func NewUserController(logger *logging.Logger, dbClient *gorm.DB) *UserController {
	return &UserController{
		dbClient: dbClient,
		logger:   logger,
	}
}

func (uc *UserController) AddUser(c *gin.Context) {
	var request UserRequestAdd

	if err := c.ShouldBindJSON(&request); err != nil {
		uc.logger.Error("invalid input format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.mapToDomain()

	// Пишем юзера в БД
	user := domain.Users{
		Name: request.Name,
		Age:  request.Age,
	}

	// Логгируем юзера
	uc.logger.Info(fmt.Sprintf("new user - Name:%s Age:%d", request.Name, request.Age))

	result := uc.dbClient.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Возвращаем и логгируем id нового юзера в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"id": user.ID})
	uc.logger.Info(fmt.Sprintf("user's id:%d", user.ID))
}

func (uc *UserController) TableRefresh(c *gin.Context) {

	// Так как GORM не умеет дропать таблицы придется выполнить SQL-запрос руками
	if err := uc.dbClient.Exec("DROP TABLE IF EXISTS users;").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table 'users' refreshed successfully"})
	uc.dbClient.AutoMigrate(&domain.Users{})
	uc.logger.Info("Table 'users' refreshed successfully")
}
