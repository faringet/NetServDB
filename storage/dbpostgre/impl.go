package dbpostgre

import (
	"NetServDB/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DataBaseRepositoryImpl struct {
	postgreClient *gorm.DB
}

func NewDataBaseRepositoryImpl(postgreClient *gorm.DB) *DataBaseRepositoryImpl {
	return &DataBaseRepositoryImpl{
		postgreClient: postgreClient,
	}
}

func (dr *DataBaseRepositoryImpl) AddUser(c *gin.Context, user *domain.Users) (uint64, error) {
	result := dr.postgreClient.Create(user)

	if result.Error != nil {
		return 0, result.Error
	}

	return uint64(user.ID), nil
}

func (dr *DataBaseRepositoryImpl) TableRefresh(c *gin.Context) error {

	// Выполняем SQL-запрос для дропа и создания таблицы
	if err := dr.postgreClient.Exec("DROP TABLE IF EXISTS users;").Error; err != nil {
		return err
	}

	// Мигрируем таблицу
	dr.postgreClient.AutoMigrate(&domain.Users{})
	return nil
}
