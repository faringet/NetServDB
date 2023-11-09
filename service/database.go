package service

import (
	"NetServDB/domain"
	"NetServDB/transport/http"
	"github.com/gin-gonic/gin"
)

type DataBaseRepository interface {
	AddUser(c *gin.Context, user *domain.Users) (uint64, error)
	TableRefresh(c *gin.Context) error
}

type DataBaseWorker struct {
	repo DataBaseRepository
}

func NewDataBaseWorker(repo DataBaseRepository) *DataBaseWorker {
	return &DataBaseWorker{
		repo: repo,
	}
}

func (dw *DataBaseWorker) Add(c *gin.Context, request *http.UserRequestAdd) (int64, error) { //
	user := request.MapToDomain()

	// Вызываем метод репозитория для добавления пользователя
	userID, err := dw.repo.AddUser(c, &user)
	if err != nil {
		return 0, err
	}

	return int64(userID), nil
}

func (dw *DataBaseWorker) Refresh(c *gin.Context) error { //
	if err := dw.repo.TableRefresh(c); err != nil {
		return err
	}

	return nil
}
