package http

import (
	"NetServDB/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Database interface {
	Add(c *gin.Context, request *UserRequestAdd) (int64, error)
	Refresh(c *gin.Context) error
}

type UserController struct {
	database Database
	logger   *logging.Logger
}

func NewUserController(logger *logging.Logger, database Database) *UserController {
	return &UserController{
		database: database,
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

	err := request.Validation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uc.database.Add(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем и логгируем id нового юзера в виде JSON-ответа
	c.JSON(http.StatusOK, gin.H{"id": userID})
	uc.logger.Info(fmt.Sprintf("user's id:%d", userID))
}

func (uc *UserController) TableRefresh(c *gin.Context) {
	// Вызываем метод сервиса для обновления таблицы, передавая контекст
	if err := uc.database.Refresh(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table 'users' refreshed successfully"})
	uc.logger.Info("Table 'users' refreshed successfully")
}
