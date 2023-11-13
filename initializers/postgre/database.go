package postgre

import (
	"NetServDB/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB(cfg *config.Config) (db *gorm.DB, cleanup func() error, err error) {
	dsn := cfg.Postgre.DbURL

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
		return nil, nil, err
	}

	log.Print("DB connection successful")

	cleanup = func() error {
		sqlDB, _ := db.DB()
		fmt.Println("cleanup from db")
		return sqlDB.Close()
	}

	return db, cleanup, nil
}
